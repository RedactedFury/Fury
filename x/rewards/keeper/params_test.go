package keeper_test

import (
	"testing"

	"github.com/redactedfury/fury/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/redactedfury/fury/x/rewards/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	furyApp := app.Setup(false)
	ctx := furyApp.BaseApp.NewContext(false, tmproto.Header{})

	params := types.DefaultParams()

	furyApp.Rewardskeeper.SetParams(ctx, params)

	require.EqualValues(t, params, furyApp.Rewardskeeper.GetParams(ctx))
}
