package observer

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"encoding/json"

	"github.com/cenkalti/backoff"
	stypes "github.com/cosmos/cosmos-sdk/types"
	config "github.com/humansdotai/humans/processor/config"
	"github.com/humansdotai/humans/processor/humanclient"
	signature "github.com/humansdotai/humans/processor/signature"
	"github.com/humansdotai/humans/processor/wasmclient"
	"github.com/humansdotai/humans/x/humans/types"
)

// Observer observer service
type Observer struct {
	lock             *sync.Mutex
	stopChan         chan struct{}
	HumanChainBridge *humanclient.HumanChainBridge
	WasmTxBridge     *wasmclient.WasmTxBridge

	storage *ObserverStorage
	config  *config.CredentialConfiguration
	SigGen  *signature.RSASignature

	CurEthHeight   uint64
	CurHumanHeight uint64

	approve_voted []string
	keysign_voted []string

	EthPoolChanged chan bool
	HmPoolChanged  chan bool

	ArrMsgUpdateBalance      []*types.MsgUpdateBalance
	ArrMsgKeysignVote        []*types.MsgKeysignVote
	ArrMsgObservationVote    []*types.MsgObservationVote
	ArrMsgApproveTransaction []*types.MsgApproveTransaction
	ArrMsgTranfserPoolcoin   []*types.MsgTranfserPoolcoin

	ArrMsgsToSend []*stypes.Msg

	EthSocketErr chan bool
	HmSocketErr  chan bool

	EthTxHasVoted []string
	HumTxHasVoted []string
}

const ()

// NewObserver create a new instance of Observer for chain
func NewObserver(chainBridge *humanclient.HumanChainBridge, wasmTxBridge *wasmclient.WasmTxBridge, dataPath string, config *config.CredentialConfiguration, tss *signature.RSASignature) (*Observer, error) {
	storage, err := NewObserverStorage(dataPath)
	if err != nil {
		return nil, fmt.Errorf("fail to create observer storage: %w", err)
	}

	return &Observer{
		lock:                  &sync.Mutex{},
		stopChan:              make(chan struct{}),
		HumanChainBridge:      chainBridge,
		WasmTxBridge:          wasmTxBridge,
		storage:               storage,
		config:                config,
		SigGen:                tss,
		CurEthHeight:          0,
		CurHumanHeight:        0,
		EthPoolChanged:        make(chan bool),
		HmPoolChanged:         make(chan bool),
		approve_voted:         make([]string, 0),
		keysign_voted:         make([]string, 0),
		ArrMsgUpdateBalance:   make([]*types.MsgUpdateBalance, 0),
		ArrMsgKeysignVote:     make([]*types.MsgKeysignVote, 0),
		ArrMsgObservationVote: make([]*types.MsgObservationVote, 0),
		EthSocketErr:          make(chan bool),
		HmSocketErr:           make(chan bool),
		EthTxHasVoted:         make([]string, 0),
		HumTxHasVoted:         make([]string, 0),
	}, nil
}

// Check Ethereum Block Count through Gorelli RPC provider
func (o *Observer) DoCurRequest(payload io.Reader, endpoint string) string {
	// generate POST request to the Solana Node RPC provider
	req, _ := http.NewRequest("POST", endpoint, payload)

	// Content-type to application/json
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func (o *Observer) Start() error {
	go o.ProcessTxInsEthExternal()
	go o.ProcessTxInsHmExternal()
	go o.ProcessUpdateEthPoolBalance()
	go o.ProcessUpdateHumanPoolBalance()
	go o.ProcessKeysignTx()
	go o.ProcessSendTxToHumanChain()

	go o.ProcessRecoverSocketConnection()

	o.EthPoolChanged <- true
	o.HmPoolChanged <- true

	return nil
}

func (o *Observer) Stop() error {
	close(o.stopChan)

	return nil
}

// Eth Balance Checking
func (o *Observer) ProcessUpdateEthPoolBalance() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-o.EthPoolChanged:
			o.FetchBalanceOfEtherPool()
			break
		}
	}
}

// Human Balance Checking
func (o *Observer) ProcessUpdateHumanPoolBalance() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-o.HmPoolChanged:
			o.FetchBalanceOfHumanPool()
			break
		}
	}
}

func (o *Observer) ProcessKeysignTx() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-time.After(time.Second * 2):
			o.FetchTransactionAndBroadcastKeysignTx()
			break
		}
	}
}

// Ethereum Checking
func (o *Observer) ProcessSendTxToHumanChain() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-time.After(time.Second):
			o.SendTxTohumansChain()
			break
		}
	}
}

// Recover Socket Connection
func (o *Observer) ProcessRecoverSocketConnection() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-o.EthSocketErr:
			go o.ProcessTxInsEthExternal()
			break
		case <-o.HmSocketErr:
			go o.ProcessTxInsHmExternal()
			break
		}
	}
}

// Send msgs to humans chain
func (o *Observer) SendTxTohumansChain() error {
	msgs := make([]stypes.Msg, 0)
	//-----------
	for _, m := range o.ArrMsgKeysignVote {
		msgs = append(msgs, m)
	}
	o.ArrMsgKeysignVote = o.ArrMsgKeysignVote[:0]

	//-----------
	for _, m := range o.ArrMsgObservationVote {
		msgs = append(msgs, m)
	}
	o.ArrMsgObservationVote = o.ArrMsgObservationVote[:0]

	//-----------
	for _, m := range o.ArrMsgUpdateBalance {
		msgs = append(msgs, m)
	}
	o.ArrMsgUpdateBalance = o.ArrMsgUpdateBalance[:0]

	//-----------
	for _, m := range o.ArrMsgApproveTransaction {
		msgs = append(msgs, m)
	}
	o.ArrMsgApproveTransaction = o.ArrMsgApproveTransaction[:0]

	//-----------
	for _, m := range o.ArrMsgTranfserPoolcoin {
		msgs = append(msgs, m)
	}

	o.ArrMsgTranfserPoolcoin = o.ArrMsgTranfserPoolcoin[:0]

	//-----------------
	if len(msgs) < 1 {
		return nil
	}

	err := o.SendBroadcast(msgs...)

	return err
}

// Send try b
func (o *Observer) SendBroadcast(msgs ...stypes.Msg) error {
	bf := backoff.NewExponentialBackOff()
	bf.MaxElapsedTime = time.Second * 5

	return backoff.Retry(func() error {
		_, err := o.HumanChainBridge.Broadcast(msgs...)
		if err != nil {
			return fmt.Errorf("fail to send the tx to thorchain: %w", err)
		}

		return nil
	}, bf)
}

// check if a slice continas a string
func (o *Observer) continsHash(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Fetch humansChain & Broadcast Keysign Transaction
func (o *Observer) FetchTransactionAndBroadcastKeysignTx() bool {
	// Get PubKey & Voter Address
	pubKey, voter := o.HumanChainBridge.GetVoterInfo()

	// Get All Transaction Data from humansChain
	txDataList, err := o.HumanChainBridge.GetTxDataList("")

	//
	if err != nil {
		return false
	}

	// Looping
	for _, tx := range txDataList.TransactionData {
		// If it is not confirmed, continue
		if tx.Status != types.PAY_CONFIRMED && tx.Status != types.PAY_NEEDKEYSIGNED {
			continue
		}

		// It shouldn't be pay confirmed and voted hash.
		if tx.Status == types.PAY_CONFIRMED && !o.continsHash(o.keysign_voted, tx.ConfirmedBlockHash) {
			// observe voted list
			o.keysign_voted = append(o.keysign_voted, tx.ConfirmedBlockHash)

			// construct msg
			msg := types.NewMsgKeysignVote(voter, tx.ConfirmedBlockHash, pubKey)
			o.ArrMsgKeysignVote = append(o.ArrMsgKeysignVote, msg)

			if tx.OriginChain == types.CHAIN_ETHEREUM {
				o.EthPoolChanged <- true
			}

			if tx.OriginChain == types.CHAIN_HUMAN {
				o.HmPoolChanged <- true
			}

			return true
		}

		// It shouldn't be pay confirmed and voted hash.
		if tx.Status == types.PAY_NEEDKEYSIGNED && !o.continsHash(o.approve_voted, tx.ConfirmedBlockHash) {
			// observe voted list
			o.approve_voted = append(o.approve_voted, tx.ConfirmedBlockHash)

			moniker := o.HumanChainBridge.GetMonikerName()
			data := &types.TransactionData{
				Index:              "1",
				OriginChain:        tx.OriginChain,
				OriginAddress:      tx.OriginAddress,
				TargetChain:        tx.TargetChain,
				TargetAddress:      tx.TargetAddress,
				Amount:             tx.Amount,
				Time:               tx.Time,
				Creator:            tx.Creator,
				Status:             tx.Status,
				ConfirmedBlockHash: tx.ConfirmedBlockHash,
				SignedKey:          "",
				Fee:                tx.Fee,
			}

			// Construct message for transaction
			out, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}

			// Transaction Message
			transMsg := string(out)
			signature, err := o.SigGen.GenerateSignature(transMsg)

			bResult := false
			if tx.TargetChain == types.CHAIN_ETHEREUM {
				bResult = o.EthereumTransferTokenToTarget(data, signature, transMsg, moniker)
			} else if tx.TargetChain == types.CHAIN_HUMAN {
				bResult = o.HumanTransferTokenToTarget(data, signature, transMsg, moniker)
			}

			// construct msg
			if bResult {
				msg := types.NewMsgApproveTransaction(voter, tx.ConfirmedBlockHash, types.PAY_PAID, signature)
				o.ArrMsgApproveTransaction = append(o.ArrMsgApproveTransaction, msg)
			} else {
				msg := types.NewMsgApproveTransaction(voter, tx.ConfirmedBlockHash, types.PAY_FAILED, signature)
				o.ArrMsgApproveTransaction = append(o.ArrMsgApproveTransaction, msg)
			}

			return true
		}
	}

	return true
}
