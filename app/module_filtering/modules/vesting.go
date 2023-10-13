//go:build (include && vesting) || (exclude && !vesting)

package modules

import (
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

func init() {

	ActiveModules[vestingtypes.ModuleName] = true
}
