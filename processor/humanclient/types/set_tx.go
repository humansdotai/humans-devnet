package types

import (
	legacytypes "github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	cosmos "github.com/humansdotai/humans/common/cosmos"
)

type SetTx struct {
	Mode string `json:"mode"`
	Tx   struct {
		Msg        []cosmos.Msg               `json:"msg"`
		Fee        legacytypes.StdFee         `json:"fee"`
		Signatures []legacytypes.StdSignature `json:"signatures"` // nolint
		Memo       string                     `json:"memo"`
	} `json:"tx"`
}
