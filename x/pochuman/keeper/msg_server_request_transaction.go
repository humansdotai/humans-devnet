package keeper

import (
	"context"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestTransaction(goCtx context.Context, msg *types.MsgRequestTransaction) (*types.MsgRequestTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestTransactionResponse{}, nil
}
