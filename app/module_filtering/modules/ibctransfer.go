//go:build (include && ibctransfer) || (exclude && !ibctransfer)

package modules

import (
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
)

func init() {

	ActiveModules[ibctransfertypes.ModuleName] = true
}
