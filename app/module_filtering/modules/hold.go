//go:build (include && hold) || (exclude && !hold)

package modules

import (
	"github.com/provenance-io/provenance/x/hold"
)

func init() {

	ActiveModules[hold.ModuleName] = true
}
