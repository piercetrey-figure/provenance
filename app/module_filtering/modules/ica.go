//go:build (include && ica) || (exclude && !ica)

package modules

import (
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
)

func init() {

	ActiveModules[icatypes.ModuleName] = true
}
