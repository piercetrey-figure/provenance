//go:build (include && ica) || (exclude && !ica)

package modules

import (
    icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", icatypes.ModuleName)
    ActiveModules[icatypes.ModuleName] = true
}
