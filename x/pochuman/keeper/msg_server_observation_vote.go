package keeper

import (
	"context"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ObservationVote(goCtx context.Context, msg *types.MsgObservationVote) (*types.MsgObservationVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgObservationVoteResponse{}, nil
}
