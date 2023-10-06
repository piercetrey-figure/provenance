//go:build (include && auth) || (exclude && !auth)

package modules

import (
    authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", authtypes.ModuleName)
    ActiveModules[authtypes.ModuleName] = true
}
