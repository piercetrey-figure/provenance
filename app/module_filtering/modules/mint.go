//go:build (include && mint) || (exclude && !mint)

package modules

import (
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func init() {

	ActiveModules[minttypes.ModuleName] = true
}
