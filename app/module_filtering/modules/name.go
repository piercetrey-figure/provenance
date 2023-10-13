//go:build (include && name) || (exclude && !name)

package modules

import (
	nametypes "github.com/provenance-io/provenance/x/name/types"
)

func init() {

	ActiveModules[nametypes.ModuleName] = true
}
