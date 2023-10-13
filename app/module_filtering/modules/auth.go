//go:build (include && auth) || (exclude && !auth)

package modules

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func init() {

	ActiveModules[authtypes.ModuleName] = true
}
