package simulation

import (
	"math/rand"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/keeper"
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRequestTransaction(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRequestTransaction{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RequestTransaction simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RequestTransaction simulation not implemented"), nil, nil
	}
}
