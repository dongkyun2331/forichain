package forichain_test

import (
	"testing"

	keepertest "github.com/dongkyun2331/forichain/testutil/keeper"
	"github.com/dongkyun2331/forichain/testutil/nullify"
	forichain "github.com/dongkyun2331/forichain/x/forichain/module"
	"github.com/dongkyun2331/forichain/x/forichain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ForichainKeeper(t)
	forichain.InitGenesis(ctx, k, genesisState)
	got := forichain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
