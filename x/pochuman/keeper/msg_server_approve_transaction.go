package keeper

import (
	"context"

    "github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) ApproveTransaction(goCtx context.Context,  msg *types.MsgApproveTransaction) (*types.MsgApproveTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgApproveTransactionResponse{}, nil
}
