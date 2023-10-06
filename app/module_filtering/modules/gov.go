//go:build (include && gov) || (exclude && !gov)

package modules

import (
    govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", govtypes.ModuleName)
    ActiveModules[govtypes.ModuleName] = true
}
