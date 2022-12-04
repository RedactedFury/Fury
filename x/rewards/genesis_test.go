package rewards_test

import (
	"testing"

	"github.com/redactedfury/fury/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/redactedfury/fury/x/rewards"
	"github.com/redactedfury/fury/x/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	furyApp := app.Setup(false)
	ctx := furyApp.BaseApp.NewContext(false, tmproto.Header{})

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	rewards.InitGenesis(ctx, furyApp.Rewardskeeper, &genesisState)
	got := rewards.ExportGenesis(ctx, furyApp.Rewardskeeper)
	require.NotNil(t, got)
}
