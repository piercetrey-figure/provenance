//go:build (include && ibchooks) || (exclude && !ibchooks)

package modules

import (
    ibchookstypes "github.com/provenance-io/provenance/x/ibchooks/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", ibchookstypes.ModuleName)
    ActiveModules[ibchookstypes.ModuleName] = true
}
