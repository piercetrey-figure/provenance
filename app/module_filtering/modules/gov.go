//go:build (include && gov) || (exclude && !gov)

package modules

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func init() {

	ActiveModules[govtypes.ModuleName] = true
}
