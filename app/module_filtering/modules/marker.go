//go:build (include && marker) || (exclude && !marker)

package modules

import (
	"fmt"
	markertypes "github.com/provenance-io/provenance/x/marker/types"
)

func init() {
	fmt.Println("Including Module: ", markertypes.ModuleName)
	ActiveModules[markertypes.ModuleName] = true
}
