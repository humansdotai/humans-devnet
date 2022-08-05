package types

import (
	stypes "github.com/humansdotai/humans/x/humans/types"
)

type Msg struct {
	Type  string                    `json:"type"`
	Value stypes.MsgObservationVote `json:"value"`
}
