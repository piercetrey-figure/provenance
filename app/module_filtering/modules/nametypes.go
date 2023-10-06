//go:build (include && name) || (exclude && !name)

package modules

import (
    nametypes "github.com/provenance-io/provenance/x/name/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", nametypes.ModuleName)
    ActiveModules[nametypes.ModuleName] = true
}
