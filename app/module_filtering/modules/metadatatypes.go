//go:build (include && metadata) || (exclude && !metadata)

package modules

import (
    metadatatypes "github.com/provenance-io/provenance/x/metadata/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", metadatatypes.ModuleName)
    ActiveModules[metadatatypes.ModuleName] = true
}
