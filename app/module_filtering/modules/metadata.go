//go:build (include && metadata) || (exclude && !metadata)

package modules

import (
	metadatatypes "github.com/provenance-io/provenance/x/metadata/types"
)

func init() {

	ActiveModules[metadatatypes.ModuleName] = true
}
