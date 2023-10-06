//go:build (include && distr) || (exclude && !distr)

package modules

import (
    distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", distrtypes.ModuleName)
    ActiveModules[distrtypes.ModuleName] = true
}
