package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/VigorousDeveloper/humans/testutil/keeper"
	"github.com/VigorousDeveloper/humans/x/humans/keeper"
	"github.com/VigorousDeveloper/humans/x/humans/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HumansKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
