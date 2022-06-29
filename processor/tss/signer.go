package tss

import (
	"fmt"
	"sync"
	"time"

	"github.com/VigorousDeveloper/poc-human/processor/diverclient"
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
)

type TssSigner struct {
	lock                 *sync.Mutex
	stopChan             chan struct{}
	diversifiChainBridge *diverclient.DiversifiChainBridge
	observe_voted        []string
	keysign_voted        []string
}

// NewObserver create a new instance of Observer for chain
func NewTssSigner(chainBridge *diverclient.DiversifiChainBridge) (*TssSigner, error) {
	return &TssSigner{
		lock:                 &sync.Mutex{},
		stopChan:             make(chan struct{}),
		diversifiChainBridge: chainBridge,
		observe_voted:        make([]string, 0),
		keysign_voted:        make([]string, 0),
	}, nil
}

func (tss *TssSigner) Start() error {
	go tss.processKeysignTx()

	return nil
}

func (tss *TssSigner) Stop() error {
	close(tss.stopChan)

	return nil
}

// Keysign Tx
func (tss *TssSigner) processKeysignTx() {
	for {
		select {
		case <-tss.stopChan:
			return
		case <-time.After(time.Millisecond * 100):
			tss.lock.Lock()
			// Check & Process Tx
			tss.fetchTransactionAndBroadcastKeysignTx()
			tss.lock.Unlock()
		}
	}
}

// check if a slice continas a string
func (tss *TssSigner) continsHash(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Fetch DiversifiChain & Broadcast Keysign Transaction
func (tss *TssSigner) fetchTransactionAndBroadcastKeysignTx() bool {
	// Get PubKey & Voter Address
	pubKey, voter := tss.diversifiChainBridge.GetVoterInfo()

	// Get All Transaction Data from DiversifiChain
	txDataList, err := tss.diversifiChainBridge.GetTxDataList("")

	//
	if err != nil {
		return false
	}

	// Looping
	for _, tx := range txDataList.TransactionData {
		// If it is not confirmed, continue
		if tx.Status != types.PAY_CONFIRMED {
			continue
		}

		// It shouldn't be pay confirmed and voted hash.
		if tx.Status == types.PAY_CONFIRMED && !tss.continsHash(tss.observe_voted, tx.ConfirmedBlockHash) {
			// observe voted list
			tss.observe_voted = append(tss.observe_voted, tx.ConfirmedBlockHash)

			// construct msg
			msg := types.NewMsgKeysignVote(voter, tx.ConfirmedBlockHash, pubKey)

			fmt.Println("!!!!!!!KeySign Voted!!!!!!!!")
			// Broadcast Tx to diversifiChain
			_, err := tss.diversifiChainBridge.Broadcast(msg)
			if err != nil {
				fmt.Println(err)

				return false
			}
		}

		// // If the transaction gets ready for transfer
		// // It shouldn't be keysigned && voted
		// if tx.Status == types.PAY_KEYSIGNED && !tss.continsHash(tss.keysign_voted, tx.ConfirmedBlockHash) {
		// 	// add keysigned voted hash to avoid multiple voting
		// 	tss.keysign_voted = append(tss.keysign_voted, tx.ConfirmedBlockHash)

		// 	fmt.Println("!!!!!!!Approve Transaction!!!!!")
		// 	// construct msg
		// 	msg := types.NewMsgApproveTransaction(voter, tx.ConfirmedBlockHash)

		// 	// Broadcast Tx to diversifiChain
		// 	_, err := tss.diversifiChainBridge.Broadcast(msg)
		// 	if err != nil {
		// 		fmt.Println(err)

		// 		return false
		// 	}
		// }

		return true
	}

	return true
}
