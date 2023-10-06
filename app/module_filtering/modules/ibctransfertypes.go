//go:build (include && ibctransfer) || (exclude && !ibctransfer)

package modules

import (
    ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", ibctransfertypes.ModuleName)
    ActiveModules[ibctransfertypes.ModuleName] = true
}
