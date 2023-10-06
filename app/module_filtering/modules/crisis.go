//go:build (include && crisis) || (exclude && !crisis)

package modules

import (
	"fmt"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
)

func init() {
	fmt.Println("Including Module: ", crisistypes.ModuleName)
	ActiveModules[crisistypes.ModuleName] = true
}
