//go:build (include && auth) || (exclude && !auth)

package modules

import (
	"fmt"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func init() {
	fmt.Println("Including Module: ", authtypes.ModuleName)
	ActiveModules[authtypes.ModuleName] = true
}
