package observer

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/VigorousDeveloper/poc-human/processor/diverclient"
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cenkalti/backoff"
	stypes "github.com/cosmos/cosmos-sdk/types"
)

// Observer observer service
type Observer struct {
	lock                 *sync.Mutex
	stopChan             chan struct{}
	diversifiChainBridge *diverclient.DiversifiChainBridge
	storage              *ObserverStorage

	CurEthHeight  uint64
	CurSolHeight  uint64
	CurKimaHeight uint64

	approve_voted []string
	keysign_voted []string

	EthPoolChanged chan bool
	SolPoolChanged chan bool
	PolPoolChanged chan bool

	ArrMsgUpdateBalance      []*types.MsgUpdateBalance
	ArrMsgKeysignVote        []*types.MsgKeysignVote
	ArrMsgObservationVote    []*types.MsgObservationVote
	ArrMsgApproveTransaction []*types.MsgApproveTransaction

	ArrMsgsToSend []*stypes.Msg
}

const (
	//----------ETHEREUM---------
	//-------------------------
	// Ethereum RPC Node Provider URL from Pokt
	URL_Ethereum_RPC_Node_Provider = "https://eth-rinkeby.alchemyapi.io/v2/JiLlXSz2HgdpuutVqt-irguqqDBxPV4D"

	// Ethereum RPC Node Provider WSS URL from Infura, rinkeby
	URL_Ethereum_RPC_Node_Provider_WSS = "wss://eth-rinkeby.alchemyapi.io/v2/JiLlXSz2HgdpuutVqt-irguqqDBxPV4D"

	// Ethereum Rinkeby USDK Contract Address
	Ethereum_USDK_Token_Address = "0x7Ba1E70BF249eEF06de34af3E2695eFccFc4a0d2"

	// Ethereum Pool Account Address
	Ethereum_Pool_Address = "0x369b28f227C0188478cb05F8467bdd52002EcC4E"

	// Ethereum Pool Account Private Key
	Ethereum_Pool_Account_Private_Key = "4b11634f979c262e33def94f52a0a82e57d0db5d7f94efd2844a1892623e063c"

	//----------POLYGON---------
	//--------------------------
	// Ethereum RPC Node Provider URL from Pokt
	URL_Polygon_RPC_Node_Provider = "https://polygon-mumbai.g.alchemy.com/v2/gm02AeSIBsrJVs1cKpQRVEf383Qlhh5s"

	// Ethereum RPC Node Provider WSS URL from Infura, rinkeby
	URL_Polygon_RPC_Node_Provider_WSS = "wss://polygon-mumbai.g.alchemy.com/v2/gm02AeSIBsrJVs1cKpQRVEf383Qlhh5s"

	// Polygon Mumbai USDK Contract Address
	Polygon_USDK_Token_Address = "0x5bd4865a6dEd507dA08ed1aBE3cd971a7e0405D7"

	// Polygon Pool Account Address
	Polygon_Pool_Address = "0x369b28f227C0188478cb05F8467bdd52002EcC4E"

	// Polygon Pool Account Private Key
	Polygon_Pool_Account_Private_Key = "4b11634f979c262e33def94f52a0a82e57d0db5d7f94efd2844a1892623e063c"

	//----------SOLANA---------
	//-------------------------
	// Solana RPC Node Provider URL from ChainStack
	URL_Solana_RPC_Node_Provider = "https://nd-013-629-002.p2pify.com/b2367a95592abda85ae5802581b8880f"

	// Solana WSS Node provider
	ChainStack_Solana_WSS_Provider = "wss://ws-nd-013-629-002.p2pify.com/b2367a95592abda85ae5802581b8880f"

	// Solana Pool Devnet USDC token account address
	Solana_Pool_USDC_Token_Account_Address = "AnrnjMrYHhdLGsKKXi37VxK2PGDpp5LrXvdXFMs8J8yA"

	// Solana Pool Account Private Key
	Solana_USDC_Pool_Account_Private_Key = "3UN7JaU8WFphHMnijRSvWChh8GQgJnVm6eGu88xqmgJecfvGwNLTmPUmKZpqWX8pXX73sNheRVcZky8kD8jHSjR5"

	// Solana USDC token address
	Solana_USDK_Token_Address = "GkbnUDkymDTF4U6Z5wM5kKJn3GmGndMn2rN5typmyUHY"

	// Solana Pool Public Key
	Solana_Pool_Pub_key = "DRmLANN1qXBELs69gW5upY4qH4iWc23MTcRPjDuzZYuH"
)

// NewObserver create a new instance of Observer for chain
func NewObserver(chainBridge *diverclient.DiversifiChainBridge, dataPath string) (*Observer, error) {
	storage, err := NewObserverStorage(dataPath)
	if err != nil {
		return nil, fmt.Errorf("fail to create observer storage: %w", err)
	}

	return &Observer{
		lock:                  &sync.Mutex{},
		stopChan:              make(chan struct{}),
		diversifiChainBridge:  chainBridge,
		storage:               storage,
		CurEthHeight:          0,
		CurSolHeight:          0,
		CurKimaHeight:         0,
		EthPoolChanged:        make(chan bool),
		SolPoolChanged:        make(chan bool),
		PolPoolChanged:        make(chan bool),
		approve_voted:         make([]string, 0),
		keysign_voted:         make([]string, 0),
		ArrMsgUpdateBalance:   make([]*types.MsgUpdateBalance, 0),
		ArrMsgKeysignVote:     make([]*types.MsgKeysignVote, 0),
		ArrMsgObservationVote: make([]*types.MsgObservationVote, 0),
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
	go o.ProcessUpdateEthPoolBalance()
	go o.ProcessKeysignTx()
	go o.ProcessSendTxToDiversifiChain()

	o.EthPoolChanged <- true
	o.SolPoolChanged <- true
	o.PolPoolChanged <- true

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
		}
	}
}

// Ethereum Checking
func (o *Observer) ProcessSendTxToDiversifiChain() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-time.After(time.Second):
			o.SendTxToDiversifiChain()
		}
	}
}

// Send msgs to diversifi chain
func (o *Observer) SendTxToDiversifiChain() error {
	msgs := make([]stypes.Msg, 0)
	for _, m := range o.ArrMsgKeysignVote {
		msgs = append(msgs, m)
	}

	for _, m := range o.ArrMsgObservationVote {
		msgs = append(msgs, m)
	}

	for _, m := range o.ArrMsgUpdateBalance {
		msgs = append(msgs, m)
	}

	for _, m := range o.ArrMsgApproveTransaction {
		msgs = append(msgs, m)
	}

	if len(msgs) < 1 {
		return nil
	}

	err := o.SendBroadcast(msgs...)
	if err == nil {
		o.ArrMsgKeysignVote = o.ArrMsgKeysignVote[:0]
		o.ArrMsgObservationVote = o.ArrMsgObservationVote[:0]
		o.ArrMsgUpdateBalance = o.ArrMsgUpdateBalance[:0]
		o.ArrMsgApproveTransaction = o.ArrMsgApproveTransaction[:0]
	}

	return err
}

// Send try b
func (o *Observer) SendBroadcast(msgs ...stypes.Msg) error {
	bf := backoff.NewExponentialBackOff()
	bf.MaxElapsedTime = time.Second * 5

	return backoff.Retry(func() error {
		_, err := o.diversifiChainBridge.Broadcast(msgs...)
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

// Fetch DiversifiChain & Broadcast Keysign Transaction
func (o *Observer) FetchTransactionAndBroadcastKeysignTx() bool {
	// Get PubKey & Voter Address
	pubKey, voter := o.diversifiChainBridge.GetVoterInfo()

	// Get All Transaction Data from DiversifiChain
	txDataList, err := o.diversifiChainBridge.GetTxDataList("")

	//
	if err != nil {
		return false
	}

	// Looping
	for _, tx := range txDataList.TransactionData {
		// If it is not confirmed, continue
		if tx.Status != types.PAY_CONFIRMED && tx.Status != types.PAY_KEYSIGNED {
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

			if tx.OriginChain == types.CHAIN_SOLANA {
				o.SolPoolChanged <- true
			}

			if tx.OriginChain == types.CHAIN_POLYGON {
				o.PolPoolChanged <- true
			}

			return true
		}

		// It shouldn't be pay confirmed and voted hash.
		if tx.Status == types.PAY_KEYSIGNED && !o.continsHash(o.approve_voted, tx.ConfirmedBlockHash) {
			// observe voted list
			o.approve_voted = append(o.approve_voted, tx.ConfirmedBlockHash)

			moniker := o.diversifiChainBridge.GetMonikerName()
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

			bResult := false

			if tx.TargetChain == types.CHAIN_ETHEREUM {
				bResult = o.EthereumTransferTokenToTarget(data, moniker)
			} else if tx.TargetChain == types.CHAIN_POLYGON {
			}

			pubKey, _ := o.diversifiChainBridge.GetVoterInfo()
			securedKey := types.EncryptMsgSHA256(pubKey)

			// construct msg
			if bResult {
				msg := types.NewMsgApproveTransaction(voter, tx.ConfirmedBlockHash, types.PAY_PAID, securedKey)
				o.ArrMsgApproveTransaction = append(o.ArrMsgApproveTransaction, msg)
			} else {
				msg := types.NewMsgApproveTransaction(voter, tx.ConfirmedBlockHash, types.PAY_FAILED, securedKey)
				o.ArrMsgApproveTransaction = append(o.ArrMsgApproveTransaction, msg)
			}

			return true
		}
	}

	return true
}
