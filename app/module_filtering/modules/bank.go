//go:build (include && bank) || (exclude && !bank)

package modules

import (
    banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", banktypes.ModuleName)
    ActiveModules[banktypes.ModuleName] = true
}
