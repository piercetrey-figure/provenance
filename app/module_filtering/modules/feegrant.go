//go:build (include && feegrant) || (exclude && !feegrant)

package modules

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
)

func init() {
	fmt.Println("Including Module: ", feegrant.ModuleName)
	ActiveModules[feegrant.ModuleName] = true
}
