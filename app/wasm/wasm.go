package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	assetkeeper "github.com/Fury-Labs/fury/x/asset/keeper"
	auctionKeeper "github.com/Fury-Labs/fury/x/auction/keeper"
	collectorKeeper "github.com/Fury-Labs/fury/x/collector/keeper"
	esmKeeper "github.com/Fury-Labs/fury/x/esm/keeper"
	lendKeeper "github.com/Fury-Labs/fury/x/lend/keeper"
	liquidationKeeper "github.com/Fury-Labs/fury/x/liquidation/keeper"
	liquidityKeeper "github.com/Fury-Labs/fury/x/liquidity/keeper"
	lockerkeeper "github.com/Fury-Labs/fury/x/locker/keeper"
	rewardsKeeper "github.com/Fury-Labs/fury/x/rewards/keeper"
	tokenMintkeeper "github.com/Fury-Labs/fury/x/tokenmint/keeper"
	vaultKeeper "github.com/Fury-Labs/fury/x/vault/keeper"
)

func RegisterCustomPlugins(
	locker *lockerkeeper.Keeper,
	tokenMint *tokenMintkeeper.Keeper,
	asset *assetkeeper.Keeper,
	rewards *rewardsKeeper.Keeper,
	collector *collectorKeeper.Keeper,
	liquidation *liquidationKeeper.Keeper,
	auction *auctionKeeper.Keeper,
	esm *esmKeeper.Keeper,
	vault *vaultKeeper.Keeper,
	lend *lendKeeper.Keeper,
	liquidity *liquidityKeeper.Keeper,
) []wasmkeeper.Option {
	comdexQueryPlugin := NewQueryPlugin(asset, locker, tokenMint, rewards, collector, liquidation, esm, vault, lend, liquidity)

	appDataQueryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(comdexQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(*locker, *rewards, *asset, *collector, *liquidation, *auction, *tokenMint, *esm, *vault),
	)

	return []wasm.Option{
		appDataQueryPluginOpt,
		messengerDecoratorOpt,
	}
}
