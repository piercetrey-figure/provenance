//go:build (include && distr) || (exclude && !distr)

package modules

import (
	"fmt"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

func init() {
	fmt.Println("Including Module: ", distrtypes.ModuleName)
	ActiveModules[distrtypes.ModuleName] = true
}
