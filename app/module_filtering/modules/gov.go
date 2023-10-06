//go:build (include && gov) || (exclude && !gov)

package modules

import (
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func init() {
	fmt.Println("Including Module: ", govtypes.ModuleName)
	ActiveModules[govtypes.ModuleName] = true
}
