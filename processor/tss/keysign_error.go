package tss

import (
	"fmt"

	"github.com/humansdotai/humans/x/humans/types"
)

// KeysignError is a custom error create to include which party to blame
type KeysignError struct {
	Blame types.Blame
}

// NewKeysignError create a new instance of KeysignError
func NewKeysignError(blame types.Blame) KeysignError {
	return KeysignError{
		Blame: blame,
	}
}

// Error implement error interface
func (k KeysignError) Error() string {
	return fmt.Sprintf("fail to complete TSS keysign,reason:%s, culprit:%+v", k.Blame.FailReason, k.Blame.BlameNodes)
}
