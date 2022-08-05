package types

import (
	"github.com/humansdotai/humans/common"
)

// Solvency structure is to hold all the information necessary to report solvency to THORNode
type Solvency struct {
	Height int64
	Chain  common.Chain
	PubKey common.PubKey
	Coins  common.Coins
}
