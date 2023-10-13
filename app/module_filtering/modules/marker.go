//go:build (include && marker) || (exclude && !marker)

package modules

import (
	markertypes "github.com/provenance-io/provenance/x/marker/types"
)

func init() {

	ActiveModules[markertypes.ModuleName] = true
}
