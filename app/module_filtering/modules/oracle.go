//go:build (include && oracle) || (exclude && !oracle)

package modules

import (
	"fmt"
	oracletypes "github.com/provenance-io/provenance/x/oracle/types"
)

func init() {
	fmt.Println("Including Module: ", oracletypes.ModuleName)
	ActiveModules[oracletypes.ModuleName] = true
}
