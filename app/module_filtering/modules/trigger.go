//go:build (include && trigger) || (exclude && !trigger)

package modules

import (
	triggertypes "github.com/provenance-io/provenance/x/trigger/types"
)

func init() {

	ActiveModules[triggertypes.ModuleName] = true
}
