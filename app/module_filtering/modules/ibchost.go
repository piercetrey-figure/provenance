//go:build (include && ibchost) || (exclude && !ibchost)

package modules

import (
    ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", ibchost.ModuleName)
    ActiveModules[ibchost.ModuleName] = true
}
