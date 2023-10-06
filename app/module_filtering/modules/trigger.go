//go:build (include && trigger) || (exclude && !trigger)

package modules

import (
	"fmt"
	triggertypes "github.com/provenance-io/provenance/x/trigger/types"
)

func init() {
	fmt.Println("Including Module: ", triggertypes.ModuleName)
	ActiveModules[triggertypes.ModuleName] = true
}
