//go:build (include && params) || (exclude && !params)

package modules

import (
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func init() {

	ActiveModules[paramstypes.ModuleName] = true
}
