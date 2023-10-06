//go:build (include && vesting) || (exclude && !vesting)

package modules

import (
	"fmt"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func init() {
	fmt.Println("Including Module: ", vestingtypes.ModuleName)
	ActiveModules[vestingtypes.ModuleName] = true
}
