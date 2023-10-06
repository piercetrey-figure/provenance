//go:build (include && bank) || (exclude && !bank)

package modules

import (
	"fmt"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func init() {
	fmt.Println("Including Module: ", banktypes.ModuleName)
	ActiveModules[banktypes.ModuleName] = true
}
