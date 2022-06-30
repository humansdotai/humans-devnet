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
	_, voter := o.HumanChainBridge.GetVoterInfo()

	msg := types.NewMsgUpdateBalance(voter, types.CHAIN_HUMAN, fmt.Sprintf("%f", 5.0), fmt.Sprintf("%v", 5))
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
			fmt.Println(tx)
			// 	o.HumanParseLog(tx)
		}
	}
}

// func (o *Observer) HumanParseLog(txs coretypes.ResultEvent) {
// 	// for e := range txs {
// 	// 	fmt.Println("got", e.Data.(ttypes.EventDataTx))
// 	// }

// 	_, voter := o.HumanChainBridge.GetVoterInfo()
// 	msg := types.NewMsgObservationVote(voter, "vLog.TxHash.String()", types.CHAIN_HUMAN, "transferEvent.From.Hex()", "transferEvent.To.Hex()", fmt.Sprintf("%f", 100.0))
// 	o.ArrMsgObservationVote = append(o.ArrMsgObservationVote, msg)

// 	// Send true to HmPoolchange channel
// 	o.HmPoolChanged <- true
// }
