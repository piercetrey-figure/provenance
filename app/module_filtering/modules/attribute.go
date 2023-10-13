//go:build (include && attribute) || (exclude && !attribute)

package modules

import (
	attributetypes "github.com/provenance-io/provenance/x/attribute/types"
)

func init() {

	ActiveModules[attributetypes.ModuleName] = true
}
