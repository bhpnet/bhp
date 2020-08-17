package token

import (
	"github.com/bhpnet/bhp-dev/x/common/perf"
	"github.com/bhpnet/bhp-dev/x/token/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called when dapp handles with abci::BeginBlock
func beginBlocker(ctx sdk.Context, keeper Keeper) {
	seq := perf.GetPerf().OnBeginBlockEnter(ctx, types.ModuleName)
	defer perf.GetPerf().OnBeginBlockExit(ctx, types.ModuleName, seq)

	keeper.ResetCache(ctx)
}
