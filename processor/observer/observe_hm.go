package observer

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
	// for {
	// 	select {
	// 	case <-o.stopChan:
	// 		return
	// 	case <-sub.Err():
	// 		time.Sleep(time.Second * 5)
	// 	case vLog := <-logs:
	// 		o.EthereumParseLog(vLog)
	// 	}
	// }
}

func (o *Observer) HumanParseLog(vLog etherTypes.Log) {
	contractAbi, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	switch vLog.Topics[0].Hex() {
	case logTransferSigHash.Hex():
		var transferEvent LogTransfer

		i, err := contractAbi.Unpack("Transfer", vLog.Data)
		if err != nil {
			return
		}

		amt := i[0].(*big.Int)
		famt := new(big.Float).SetInt(amt)
		ff, _ := famt.Float64()
		tokenAmount := ff / 1e18

		transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
		transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

		if transferEvent.From.String() == Ethereum_Pool_Address {
			// Send true to SolPoolchange channel
			o.EthPoolChanged <- true
			return
		}

		if transferEvent.To.String() != Ethereum_Pool_Address || tokenAmount == 0.0 {
			return
		}

		if transferEvent.From.String() == Ethereum_USDK_Token_Address {
			return
		}

		_, voter := o.HumanChainBridge.GetVoterInfo()
		msg := types.NewMsgObservationVote(voter, vLog.TxHash.String(), types.CHAIN_HUMAN, transferEvent.From.Hex(), transferEvent.To.Hex(), fmt.Sprintf("%f", tokenAmount))
		o.ArrMsgObservationVote = append(o.ArrMsgObservationVote, msg)

		// Send true to HmPoolchange channel
		o.HmPoolChanged <- true
	}
}
