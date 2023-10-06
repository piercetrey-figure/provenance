//go:build (include && params) || (exclude && !params)

package modules

import (
    paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", paramstypes.ModuleName)
    ActiveModules[paramstypes.ModuleName] = true
}
