//go:build (include && ibchooks) || (exclude && !ibchooks)

package modules

import (
	ibchookstypes "github.com/provenance-io/provenance/x/ibchooks/types"
)

func init() {

	ActiveModules[ibchookstypes.ModuleName] = true
}
