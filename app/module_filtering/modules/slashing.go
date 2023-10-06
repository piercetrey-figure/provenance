//go:build (include && slashing) || (exclude && !slashing)

package modules

import (
	"fmt"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

func init() {
	fmt.Println("Including Module: ", slashingtypes.ModuleName)
	ActiveModules[slashingtypes.ModuleName] = true
}
