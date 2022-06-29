package observer

import (
	"context"
	"fmt"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
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
	sub, err := ctx.Client.Subscribe(context.Background(), "subscriber", "tm.even='Tx'")

	if err != nil {
		return
	}

	for {
		select {
		case <-o.stopChan:
			return
		case vLog := <-sub:
			o.HumanParseLog(vLog)
		}
	}
}

func (o *Observer) HumanParseLog(vLog coretypes.ResultEvent) {

	fmt.Println(vLog)

	_, voter := o.HumanChainBridge.GetVoterInfo()
	msg := types.NewMsgObservationVote(voter, "vLog.TxHash.String()", types.CHAIN_HUMAN, "transferEvent.From.Hex()", "transferEvent.To.Hex()", fmt.Sprintf("%f", 100.0))
	o.ArrMsgObservationVote = append(o.ArrMsgObservationVote, msg)

	// Send true to HmPoolchange channel
	o.HmPoolChanged <- true
}
