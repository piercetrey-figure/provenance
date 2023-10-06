//go:build (include && attribute) || (exclude && !attribute)

package modules

import (
    attributetypes "github.com/provenance-io/provenance/x/attribute/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", attributetypes.ModuleName)
    ActiveModules[attributetypes.ModuleName] = true
}
