//go:build (include && msgfees) || (exclude && !msgfees)

package modules

import (
	msgfeestypes "github.com/provenance-io/provenance/x/msgfees/types"
)

func init() {

	ActiveModules[msgfeestypes.ModuleName] = true
}
