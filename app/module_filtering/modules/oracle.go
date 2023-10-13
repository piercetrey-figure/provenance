//go:build (include && oracle) || (exclude && !oracle)

package modules

import (
	oracletypes "github.com/provenance-io/provenance/x/oracle/types"
)

func init() {

	ActiveModules[oracletypes.ModuleName] = true
}
