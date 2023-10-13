//go:build (include && authz) || (exclude && !authz)

package modules

import (
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func init() {

	ActiveModules[authz.ModuleName] = true
}
