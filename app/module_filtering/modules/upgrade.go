//go:build (include && upgrade) || (exclude && !upgrade)

package modules

import (
	"fmt"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func init() {
	fmt.Println("Including Module: ", upgradetypes.ModuleName)
	ActiveModules[upgradetypes.ModuleName] = true
}
