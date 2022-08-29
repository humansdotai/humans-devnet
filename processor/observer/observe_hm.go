package observer

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/humansdotai/humans/x/humans/types"
)

// WasmRawExecuteTxMsg
type Transfer struct {
	Message    string `json:"message,string"`
	Signature  string `json:"signature,string"`
	To_address string `json:"to_address,string"`
	Amount     string `json:"amount,string"`
}

type WasmRawExecuteTxMsg struct {
	Transfer Transfer `json:"transfer"`
}

// Fetches the USDC balance of Ethereum pool account
func (o *Observer) FetchBalanceOfHumanPool() bool {
	accBalance, err := o.HumanChainBridge.GetBalance(o.config.Humanchain_Pool_Address)
	if err != nil {
		return false
	}

	_, voter := o.HumanChainBridge.GetVoterInfo()

	famt, _ := strconv.ParseFloat(accBalance.Balances[0].Amount, 64)

	// Constrcut msg to be broadcasted
	msg := types.NewMsgUpdateBalance(voter, types.CHAIN_HUMAN, fmt.Sprintf("%f", famt/1e6), fmt.Sprintf("%v", 6))
	o.ArrMsgUpdateBalance = append(o.ArrMsgUpdateBalance, msg)

	return true
}

// Transfer token on Human
func (o *Observer) HumanTransferTokenToTarget(txdata *types.TransactionData, signature string, transMsg string, moniker string) bool {
	// Semaphore for transfer
	if moniker != types.MAIN_VALIDATOR_MONIKER {
		return true
	}

	// Construct uHEART, decimal 6
	famt, err := strconv.ParseFloat(txdata.Amount, 64)
	if err != nil {
		return false
	}

	// Fee
	amtFee, err := strconv.ParseFloat(txdata.Fee, 64)
	famt -= amtFee

	// String conv
	amt := fmt.Sprintf("%d", (int64)(famt*1e6))
	hexMsg := fmt.Sprintf("0x%s", hex.EncodeToString(([]byte)(transMsg)))

	// m := string(jsonData)
	m := `{"transfer": { "message": "` + hexMsg + `", "signature": "`
	m += signature + `", "to_address": "` + txdata.TargetAddress + `", "amount": "`
	m += amt + `"}}`

	msgExecContract := &wasmtypes.MsgExecuteContract{
		Sender:   o.config.Humanchain_Pool_Owner_Address,
		Contract: o.config.Humanchain_Pool_Address,
		Msg:      []byte(m),
		Funds:    nil,
	}

	txId, err := o.WasmTxBridge.Broadcast(msgExecContract)
	if err != nil {
		fmt.Println("Err:", err)
		return false
	}

	fmt.Println("Hash:", txId)

	// Send true to HumanPoolchange channel
	o.HmPoolChanged <- true

	return true
}

// Keep listening to WSS and fetch transaction deposited to the pool
func (o *Observer) ProcessTxInsHmExternal() {
	ctx := o.HumanChainBridge.GetContext()
	client := ctx.Client
	err := client.Start()
	if err != nil {
		return
	}
	defer client.Stop()

	ctx0, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	query := "tm.event = 'Tx'"
	txs, err := ctx.Client.Subscribe(ctx0, "test-client", query)
	if err != nil {
		time.Sleep(time.Second * 5)
		o.HmSocketErr <- true
		return
	}

	for {
		select {
		case <-o.stopChan:
			return
		case tx := <-txs:
			o.HumanParseLog(tx.Events)
		}
	}
}

// ---------------------------------- //
// --- Parse Human chain Transfer log //
// ---------------------------------- //
func (o *Observer) HumanParseLog(txs map[string][]string) {
	msgActions := txs["message.action"]
	if len(msgActions) < 1 {
		return
	}

	msgAction := msgActions[0]
	if msgAction != "/cosmos.bank.v1beta1.MsgSend" {
		return
	}

	// hash
	txHash := txs["tx.hash"][0]
	if o.continsHash(o.HumTxHasVoted, txHash) {
		return
	}

	sender := txs["coin_spent.spender"][0]
	receiver := txs["coin_received.receiver"][1]

	if sender == o.config.Humanchain_Pool_Address {
		// Send true to HmPoolchange channel
		o.HmPoolChanged <- true
		return
	}

	if receiver != o.config.Humanchain_Pool_Address {
		return
	}

	// amt
	amt := txs["transfer.amount"][1]
	amt = amt[:len(amt)-6]

	// convert uHEART to HEART
	famt, _ := strconv.ParseFloat(amt, 64)
	amount := fmt.Sprintf("%f", famt/1e6)

	o.HumTxHasVoted = append(o.HumTxHasVoted, txHash)

	_, voter := o.HumanChainBridge.GetVoterInfo()
	msg := types.NewMsgObservationVote(voter, txHash, types.CHAIN_HUMAN, sender, receiver, amount)
	o.ArrMsgObservationVote = append(o.ArrMsgObservationVote, msg)

	// Send true to HmPoolchange channel
	o.HmPoolChanged <- true
}
