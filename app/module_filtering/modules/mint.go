//go:build (include && mint) || (exclude && !mint)

package modules

import (
	"fmt"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func init() {
	fmt.Println("Including Module: ", minttypes.ModuleName)
	ActiveModules[minttypes.ModuleName] = true
}
