//go:build (include && staking) || (exclude && !staking)

package modules

import (
    stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", stakingtypes.ModuleName)
    ActiveModules[stakingtypes.ModuleName] = true
}
