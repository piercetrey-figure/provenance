//go:build (include && staking) || (exclude && !staking)

package modules

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func init() {

	ActiveModules[stakingtypes.ModuleName] = true
}
