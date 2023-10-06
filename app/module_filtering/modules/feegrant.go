//go:build (include && feegrant) || (exclude && !feegrant)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/feegrant"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", feegrant.ModuleName)
    ActiveModules[feegrant.ModuleName] = true
}
