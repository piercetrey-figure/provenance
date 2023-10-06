//go:build (include && oracle) || (exclude && !oracle)

package modules

import (
    oracletypes "github.com/provenance-io/provenance/x/oracle/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", oracletypes.ModuleName)
    ActiveModules[oracletypes.ModuleName] = true
}
