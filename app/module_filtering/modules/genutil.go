//go:build (include && genutil) || (exclude && !genutil)

package modules

import (
	"fmt"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

func init() {
	fmt.Println("Including Module: ", genutiltypes.ModuleName)
	ActiveModules[genutiltypes.ModuleName] = true
}
