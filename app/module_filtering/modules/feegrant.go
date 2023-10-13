//go:build (include && feegrant) || (exclude && !feegrant)

package modules

import (
	"github.com/cosmos/cosmos-sdk/x/feegrant"
)

func init() {

	ActiveModules[feegrant.ModuleName] = true
}
