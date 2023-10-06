//go:build (include && ica) || (exclude && !ica)

package modules

import (
	"fmt"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
)

func init() {
	fmt.Println("Including Module: ", icatypes.ModuleName)
	ActiveModules[icatypes.ModuleName] = true
}
