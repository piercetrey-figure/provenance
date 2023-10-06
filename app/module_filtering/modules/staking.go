//go:build (include && staking) || (exclude && !staking)

package modules

import (
	"fmt"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func init() {
	fmt.Println("Including Module: ", stakingtypes.ModuleName)
	ActiveModules[stakingtypes.ModuleName] = true
}
