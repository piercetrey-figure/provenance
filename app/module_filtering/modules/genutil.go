//go:build (include && genutil) || (exclude && !genutil)

package modules

import (
    genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", genutiltypes.ModuleName)
    ActiveModules[genutiltypes.ModuleName] = true
}
