//go:build (include && capability) || (exclude && !capability)

package modules

import (
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
)

func init() {

	ActiveModules[capabilitytypes.ModuleName] = true
}
