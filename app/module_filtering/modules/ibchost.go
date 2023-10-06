//go:build (include && ibchost) || (exclude && !ibchost)

package modules

import (
	"fmt"
	ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

func init() {
	fmt.Println("Including Module: ", ibchost.ModuleName)
	ActiveModules[ibchost.ModuleName] = true
}
