//go:build (include && authz) || (exclude && !authz)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/authz"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", authz.ModuleName)
    ActiveModules[authz.ModuleName] = true
}
