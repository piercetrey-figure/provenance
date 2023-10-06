//go:build (include && icq) || (exclude && !icq)

package modules

import (
    icqtypes "github.com/strangelove-ventures/async-icq/v6/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", icqtypes.ModuleName)
    ActiveModules[icqtypes.ModuleName] = true
}
