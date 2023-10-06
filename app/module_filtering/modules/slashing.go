//go:build (include && slashing) || (exclude && !slashing)

package modules

import (
    slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", slashingtypes.ModuleName)
    ActiveModules[slashingtypes.ModuleName] = true
}
