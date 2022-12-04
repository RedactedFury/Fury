package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	assetkeeper "github.com/redactedfury/fury/x/asset/keeper"
	auctionKeeper "github.com/redactedfury/fury/x/auction/keeper"
	collectorKeeper "github.com/redactedfury/fury/x/collector/keeper"
	esmKeeper "github.com/redactedfury/fury/x/esm/keeper"
	lendKeeper "github.com/redactedfury/fury/x/lend/keeper"
	liquidationKeeper "github.com/redactedfury/fury/x/liquidation/keeper"
	liquidityKeeper "github.com/redactedfury/fury/x/liquidity/keeper"
	lockerkeeper "github.com/redactedfury/fury/x/locker/keeper"
	rewardsKeeper "github.com/redactedfury/fury/x/rewards/keeper"
	tokenMintkeeper "github.com/redactedfury/fury/x/tokenmint/keeper"
	vaultKeeper "github.com/redactedfury/fury/x/vault/keeper"
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
	furyQueryPlugin := NewQueryPlugin(asset, locker, tokenMint, rewards, collector, liquidation, esm, vault, lend, liquidity)

	appDataQueryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(furyQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(*locker, *rewards, *asset, *collector, *liquidation, *auction, *tokenMint, *esm, *vault),
	)

	return []wasm.Option{
		appDataQueryPluginOpt,
		messengerDecoratorOpt,
	}
}
