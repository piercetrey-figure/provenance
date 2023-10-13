//go:build (include && ibchost) || (exclude && !ibchost)

package modules

import (
	ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

func init() {

	ActiveModules[ibchost.ModuleName] = true
}
