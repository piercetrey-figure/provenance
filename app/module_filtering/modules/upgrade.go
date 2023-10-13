//go:build (include && upgrade) || (exclude && !upgrade)

package modules

import (
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func init() {

	ActiveModules[upgradetypes.ModuleName] = true
}
