package keeper

import (
	"github.com/glottologist/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
