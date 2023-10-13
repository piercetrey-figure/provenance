//go:build (include && sanction) || (exclude && !sanction)

package modules

import (
	"github.com/cosmos/cosmos-sdk/x/sanction"
)

func init() {

	ActiveModules[sanction.ModuleName] = true
}
