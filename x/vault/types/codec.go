package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateRequest{}, "fury/vault/MsgCreateRequest", nil)
	cdc.RegisterConcrete(&MsgCloseRequest{}, "fury/vault/MsgCloseRequest", nil)
	cdc.RegisterConcrete(&MsgDepositRequest{}, "fury/vault/MsgDepositRequest", nil)
	cdc.RegisterConcrete(&MsgRepayRequest{}, "fury/vault/MsgRepayRequest", nil)
	cdc.RegisterConcrete(&MsgWithdrawRequest{}, "fury/vault/MsgWithdrawRequest", nil)
	cdc.RegisterConcrete(&MsgDrawRequest{}, "fury/vault/MsgDrawRequest", nil)
	cdc.RegisterConcrete(&MsgDepositAndDrawRequest{}, "fury/vault/MsgDepositAndDrawRequest", nil)
	cdc.RegisterConcrete(&MsgCreateStableMintRequest{}, "fury/vault/MsgCreateStableMintRequest", nil)
	cdc.RegisterConcrete(&MsgDepositStableMintRequest{}, "fury/vault/MsgDepositStableMintRequest", nil)
	cdc.RegisterConcrete(&MsgWithdrawStableMintRequest{}, "fury/vault/MsgWithdrawStableMintRequest", nil)
	cdc.RegisterConcrete(&MsgVaultInterestCalcRequest{}, "fury/vault/MsgVaultInterestCalcRequest", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateRequest{},
		&MsgDepositRequest{},
		&MsgWithdrawRequest{},
		&MsgDrawRequest{},
		&MsgRepayRequest{},
		&MsgCloseRequest{},
		&MsgDepositAndDrawRequest{},
		&MsgCreateStableMintRequest{},
		&MsgDepositStableMintRequest{},
		&MsgWithdrawStableMintRequest{},
		&MsgVaultInterestCalcRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
