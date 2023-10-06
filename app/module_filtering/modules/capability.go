//go:build (include && capability) || (exclude && !capability)

package modules

import (
	"fmt"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
)

func init() {
	fmt.Println("Including Module: ", capabilitytypes.ModuleName)
	ActiveModules[capabilitytypes.ModuleName] = true
}
