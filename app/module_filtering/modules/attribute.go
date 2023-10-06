//go:build (include && attribute) || (exclude && !attribute)

package modules

import (
	"fmt"
	attributetypes "github.com/provenance-io/provenance/x/attribute/types"
)

func init() {
	fmt.Println("Including Module: ", attributetypes.ModuleName)
	ActiveModules[attributetypes.ModuleName] = true
}
