package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgLend{}, "fury/lend/MsgLend", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "fury/lend/MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "fury/lend/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgCloseLend{}, "fury/lend/MsgCloseLend", nil)
	cdc.RegisterConcrete(&MsgBorrow{}, "fury/lend/MsgBorrow", nil)
	cdc.RegisterConcrete(&MsgDepositBorrow{}, "fury/lend/MsgDepositBorrow", nil)
	cdc.RegisterConcrete(&MsgDraw{}, "fury/lend/MsgDraw", nil)
	cdc.RegisterConcrete(&MsgCloseBorrow{}, "fury/lend/MsgCloseBorrow", nil)
	cdc.RegisterConcrete(&MsgRepay{}, "fury/lend/MsgRepay", nil)
	cdc.RegisterConcrete(&MsgBorrowAlternate{}, "fury/lend/MsgBorrowAlternate", nil)
	cdc.RegisterConcrete(&MsgFundModuleAccounts{}, "fury/lend/MsgFundModuleAccounts", nil)
	cdc.RegisterConcrete(&LendPairsProposal{}, "fury/lend/LendPairsProposal", nil)
	cdc.RegisterConcrete(&AddPoolsProposal{}, "fury/lend/AddPoolsProposal", nil)
	cdc.RegisterConcrete(&AddAssetToPairProposal{}, "fury/lend/AddAssetToPairProposal", nil)
	cdc.RegisterConcrete(&AddAssetRatesParams{}, "fury/lend/AddAssetRatesParams", nil)
	cdc.RegisterConcrete(&AddAuctionParamsProposal{}, "fury/lend/AddAuctionParamsProposal", nil)
	cdc.RegisterConcrete(&MsgCalculateInterestAndRewards{}, "fury/lend/MsgCalculateInterestAndRewards", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&LendPairsProposal{},
		&AddPoolsProposal{},
		&AddAssetToPairProposal{},
		&AddAssetRatesParams{},
		&AddAuctionParamsProposal{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgLend{},
		&MsgWithdraw{},
		&MsgDeposit{},
		&MsgCloseLend{},
		&MsgBorrow{},
		&MsgDepositBorrow{},
		&MsgDraw{},
		&MsgCloseBorrow{},
		&MsgRepay{},
		&MsgBorrowAlternate{},
		&MsgFundModuleAccounts{},
		&MsgCalculateInterestAndRewards{},
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
