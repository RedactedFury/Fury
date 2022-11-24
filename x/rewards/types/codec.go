package type    s

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGauge{}, "fury/rewards/MsgCreateGauge", nil)
	cdc.RegisterConcrete(&ActivateExternalRewardsLockers{}, "fury/rewards/activateExternalRewardsLockers", nil)
	cdc.RegisterConcrete(&ActivateExternalRewardsVault{}, "fury/rewards/activateExternalRewardsVault", nil)
	cdc.RegisterConcrete(&ActivateExternalRewardsLend{}, "fury/rewards/activateExternalRewardsLend", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateGauge{},
		&ActivateExternalRewardsLockers{},
		&ActivateExternalRewardsVault{},
		&ActivateExternalRewardsLend{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
