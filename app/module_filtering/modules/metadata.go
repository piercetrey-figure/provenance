//go:build (include && metadata) || (exclude && !metadata)

package modules

import (
	"fmt"
	metadatatypes "github.com/provenance-io/provenance/x/metadata/types"
)

func init() {
	fmt.Println("Including Module: ", metadatatypes.ModuleName)
	ActiveModules[metadatatypes.ModuleName] = true
}
