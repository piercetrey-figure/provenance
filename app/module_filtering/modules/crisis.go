//go:build (include && crisis) || (exclude && !crisis)

package modules

import (
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
)

func init() {

	ActiveModules[crisistypes.ModuleName] = true
}
