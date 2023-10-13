//go:build (include && bank) || (exclude && !bank)

package modules

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func init() {

	ActiveModules[banktypes.ModuleName] = true
}
