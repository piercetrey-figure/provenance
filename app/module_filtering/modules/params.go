//go:build (include && params) || (exclude && !params)

package modules

import (
	"fmt"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func init() {
	fmt.Println("Including Module: ", paramstypes.ModuleName)
	ActiveModules[paramstypes.ModuleName] = true
}
