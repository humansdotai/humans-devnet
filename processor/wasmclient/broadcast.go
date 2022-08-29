package wasmclient

import (
	"fmt"
	"sync/atomic"

	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	flag "github.com/spf13/pflag"

	stypes "github.com/cosmos/cosmos-sdk/types"
)

// Broadcast Broadcasts tx to humans
func (b *WasmTxBridge) Broadcast(msgs ...stypes.Msg) (TxID, error) {
	b.broadcastLock.Lock()
	defer b.broadcastLock.Unlock()

	noTxID := TxID("")

	blockHeight, err := b.GetBlockHeight()
	if err != nil {
		fmt.Println(err)
		return noTxID, err
	}
	if blockHeight > b.blockHeight {
		var seqNum uint64
		b.accountNumber, seqNum, err = b.getAccountNumberAndSequenceNumber()
		if err != nil {
			return noTxID, fmt.Errorf("fail to get account number and sequence number from humans : %w", err)
		}
		b.blockHeight = blockHeight
		if seqNum > b.seqNumber {
			b.seqNumber = seqNum
		}
	}

	flags := flag.NewFlagSet("humans", 0)

	ctx := b.GetContext()
	factory := clienttx.NewFactoryCLI(ctx, flags)
	factory = factory.WithAccountNumber(b.accountNumber)
	factory = factory.WithSequence(b.seqNumber)
	factory = factory.WithSignMode(signing.SignMode_SIGN_MODE_DIRECT)

	builder, err := clienttx.BuildUnsignedTx(factory, msgs...)
	if err != nil {
		return noTxID, err
	}
	builder.SetGasLimit(4000000000)
	err = clienttx.Sign(factory, ctx.GetFromName(), builder, true)
	if err != nil {
		return noTxID, err
	}

	txBytes, err := ctx.TxConfig.TxEncoder()(builder.GetTx())
	if err != nil {
		return noTxID, err
	}

	// broadcast to a Tendermint node
	commit, err := ctx.BroadcastTx(txBytes)
	if err != nil {
		return noTxID, fmt.Errorf("fail to broadcast tx: %w", err)
	}

	txHash, err := NewTxID(commit.TxHash)
	if err != nil {
		return BlankTxID, fmt.Errorf("fail to convert txhash: %w", err)
	}

	// Code will be the tendermint ABICode , it start at 1 , so if it is an error , code will not be zero
	if commit.Code > 0 {
		if commit.Code == 32 {
			// bad sequence number, fetch new one
			_, seqNum, _ := b.getAccountNumberAndSequenceNumber()
			if seqNum > 0 {
				b.seqNumber = seqNum
			}
		}
		// b.logger.Info().Msgf("messages: %+v", msgs)
		// commit code 6 means `unknown request` , which means the tx can't be accepted by humans
		// if that's the case, let's just ignore it and move on
		if commit.Code != 6 {
			return txHash, fmt.Errorf("fail to broadcast to HumanChain,code:%d, log:%s", commit.Code, commit.RawLog)
		}
	}

	// increment seqNum
	atomic.AddUint64(&b.seqNumber, 1)

	return txHash, nil
}
