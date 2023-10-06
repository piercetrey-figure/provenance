//go:build (include && marker) || (exclude && !marker)

package modules

import (
    markertypes "github.com/provenance-io/provenance/x/marker/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", markertypes.ModuleName)
    ActiveModules[markertypes.ModuleName] = true
}
