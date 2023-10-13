//go:build (include && genutil) || (exclude && !genutil)

package modules

import (
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

func init() {

	ActiveModules[genutiltypes.ModuleName] = true
}
