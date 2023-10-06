//go:build (include && trigger) || (exclude && !trigger)

package modules

import (
    triggertypes "github.com/provenance-io/provenance/x/trigger/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", triggertypes.ModuleName)
    ActiveModules[triggertypes.ModuleName] = true
}
