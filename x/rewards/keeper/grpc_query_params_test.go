package keeper_test

import (
	"testing"

	"github.com/Fury-Labs/fury/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/Fury-Labs/fury/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	furyApp := app.Setup(false)
	ctx := furyApp.BaseApp.NewContext(false, tmproto.Header{})

	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	furyApp.Rewardskeeper.SetParams(ctx, params)

	response, err := furyApp.Rewardskeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
