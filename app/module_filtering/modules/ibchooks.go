//go:build (include && ibchooks) || (exclude && !ibchooks)

package modules

import (
	"fmt"
	ibchookstypes "github.com/provenance-io/provenance/x/ibchooks/types"
)

func init() {
	fmt.Println("Including Module: ", ibchookstypes.ModuleName)
	ActiveModules[ibchookstypes.ModuleName] = true
}
