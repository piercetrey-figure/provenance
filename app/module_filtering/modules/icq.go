//go:build (include && icq) || (exclude && !icq)

package modules

import (
	icqtypes "github.com/strangelove-ventures/async-icq/v6/types"
)

func init() {

	ActiveModules[icqtypes.ModuleName] = true
}
