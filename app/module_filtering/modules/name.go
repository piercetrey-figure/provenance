//go:build (include && name) || (exclude && !name)

package modules

import (
	"fmt"
	nametypes "github.com/provenance-io/provenance/x/name/types"
)

func init() {
	fmt.Println("Including Module: ", nametypes.ModuleName)
	ActiveModules[nametypes.ModuleName] = true
}
