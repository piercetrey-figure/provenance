//go:build (include && upgrade) || (exclude && !upgrade)

package modules

import (
    upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", upgradetypes.ModuleName)
    ActiveModules[upgradetypes.ModuleName] = true
}
