package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/dongkyun2331/forichain/testutil/keeper"
	"github.com/dongkyun2331/forichain/x/forichain/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.ForichainKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
