//go:build (include && icq) || (exclude && !icq)

package modules

import (
	"fmt"
	icqtypes "github.com/strangelove-ventures/async-icq/v6/types"
)

func init() {
	fmt.Println("Including Module: ", icqtypes.ModuleName)
	ActiveModules[icqtypes.ModuleName] = true
}
