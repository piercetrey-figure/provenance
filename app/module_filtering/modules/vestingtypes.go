//go:build (include && vesting) || (exclude && !vesting)

package modules

import (
    vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", vestingtypes.ModuleName)
    ActiveModules[vestingtypes.ModuleName] = true
}
