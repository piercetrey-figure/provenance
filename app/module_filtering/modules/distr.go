//go:build (include && distr) || (exclude && !distr)

package modules

import (
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

func init() {

	ActiveModules[distrtypes.ModuleName] = true
}
