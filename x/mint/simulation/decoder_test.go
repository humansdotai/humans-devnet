package simulation_test

import (
	"fmt"
	"testing"

	"github.com/ignite-hq/cli/ignite/pkg/cosmoscmd"
	"github.com/stretchr/testify/require"

	"github.com/VigorousDeveloper/poc-human/app"
	"github.com/VigorousDeveloper/poc-human/x/mint/simulation"
	"github.com/VigorousDeveloper/poc-human/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

func TestDecodeStore(t *testing.T) {
	cdc := cosmoscmd.MakeEncodingConfig(app.ModuleBasics).Marshaler
	dec := simulation.NewDecodeStore(cdc)

	minter := types.NewMinter(sdk.OneDec(), sdk.NewDec(15), 1, 1)

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.MinterKey, Value: cdc.MustMarshal(&minter)},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Minter", fmt.Sprintf("%v\n%v", minter, minter)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
