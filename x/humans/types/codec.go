package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestTransaction{}, "humans/RequestTransaction", nil)
	cdc.RegisterConcrete(&MsgObservationVote{}, "humans/ObservationVote", nil)
	cdc.RegisterConcrete(&MsgUpdateBalance{}, "humans/UpdateBalance", nil)
	cdc.RegisterConcrete(&MsgKeysignVote{}, "humans/KeysignVote", nil)
	cdc.RegisterConcrete(&MsgApproveTransaction{}, "humans/ApproveTransaction", nil)
	cdc.RegisterConcrete(&MsgTranfserPoolcoin{}, "humans/TranfserPoolcoin", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestTransaction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgObservationVote{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateBalance{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgKeysignVote{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveTransaction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTranfserPoolcoin{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
