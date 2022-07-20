package keeper

import (
	"github.com/VigorousDeveloper/humans/x/humans/types"
)

var _ types.QueryServer = Keeper{}
