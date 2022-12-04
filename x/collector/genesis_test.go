package collector_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	app "github.com/redactedfury/fury/app"
	"github.com/redactedfury/fury/x/collector"
	"github.com/redactedfury/fury/x/collector/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	furyApp := app.Setup(false)
	ctx := furyApp.BaseApp.NewContext(false, tmproto.Header{})
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k := furyApp.CollectorKeeper
	collector.InitGenesis(ctx, k, &genesisState)
	got := collector.ExportGenesis(ctx, k)
	require.NotNil(t, got)
}
