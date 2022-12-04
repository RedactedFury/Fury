package keeper_test

import (
	"testing"

	"github.com/redactedfury/fury/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/redactedfury/fury/x/collector/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	furyApp := app.Setup(false)
	ctx := furyApp.BaseApp.NewContext(false, tmproto.Header{})

	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	furyApp.CollectorKeeper.SetParams(ctx, params)

	response, err := furyApp.CollectorKeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
