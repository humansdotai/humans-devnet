package observer

import (
	"context"
	"fmt"
	"time"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	// "github.com/tendermint/tendermint/rpc/coretypes"
)

// Fetches the USDC balance of Ethereum pool account
func (o *Observer) FetchBalanceOfHumanPool() bool {
	accBalance, err := o.HumanChainBridge.GetBalance(Humanchain_Pool_Address)
	if err != nil {
		return false
	}

	_, voter := o.HumanChainBridge.GetVoterInfo()

	msg := types.NewMsgUpdateBalance(voter, types.CHAIN_HUMAN, fmt.Sprintf("%f", accBalance.Balances[0].Amount/1e6), fmt.Sprintf("%v", 6))
	o.ArrMsgUpdateBalance = append(o.ArrMsgUpdateBalance, msg)

	return true
}

// Transfer token on Human
func (o *Observer) HumanTransferTokenToTarget(txdata *types.TransactionData, moniker string) bool {
	// Semaphore for transfer
	if moniker != types.MAIN_VALIDATOR_MONIKER {
		return true
	}

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

func (o *Observer) HumanParseLog(txs map[string][]string) {
	msgActions := txs["message.action"]
	if len(msgActions) < 1 {
		return
	}
	msgAction := msgActions[0]
	if msgAction != "/cosmos.bank.v1beta1.MsgSend" {
		return
	}

	sender := txs["coin_spent.spender"][0]
	receiver := txs["coin_received.receiver"][0]

	if sender == Humanchain_Pool_Address {
		// Send true to HmPoolchange channel
		o.HmPoolChanged <- true
		return
	}

	if receiver != Humanchain_Pool_Address {
		return
	}

	txHash := txs["tx.hash"][0]
	amt := txs["transfer.amount"][0]

	_, voter := o.HumanChainBridge.GetVoterInfo()
	msg := types.NewMsgObservationVote(voter, txHash, types.CHAIN_HUMAN, sender, receiver, amt)
	o.ArrMsgObservationVote = append(o.ArrMsgObservationVote, msg)

	// Send true to HmPoolchange channel
	o.HmPoolChanged <- true

	fmt.Println("Coin deposited into pool")
}
