package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/glottologist/checkers/testutil/keeper"
	"github.com/glottologist/checkers/x/checkers/keeper"
	"github.com/glottologist/checkers/x/checkers/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
