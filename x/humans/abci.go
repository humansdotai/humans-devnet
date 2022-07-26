package humans

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/humansdotai/humans/x/humans/keeper"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// Keysign tx request
	k.KeysignTxRequest(ctx)

	// Update tx request
	k.UpdateTxRequestByObservationVote(ctx)
}
