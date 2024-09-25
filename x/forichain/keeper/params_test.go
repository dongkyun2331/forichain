package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/dongkyun2331/forichain/testutil/keeper"
	"github.com/dongkyun2331/forichain/x/forichain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ForichainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
