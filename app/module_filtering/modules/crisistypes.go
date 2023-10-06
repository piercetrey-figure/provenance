//go:build (include && crisis) || (exclude && !crisis)

package modules

import (
    crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", crisistypes.ModuleName)
    ActiveModules[crisistypes.ModuleName] = true
}
