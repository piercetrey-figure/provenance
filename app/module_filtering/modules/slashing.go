//go:build (include && slashing) || (exclude && !slashing)

package modules

import (
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

func init() {

	ActiveModules[slashingtypes.ModuleName] = true
}
