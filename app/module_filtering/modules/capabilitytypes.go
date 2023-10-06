//go:build (include && capability) || (exclude && !capability)

package modules
    
import (
    capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", capabilitytypes.ModuleName)
    ActiveModules[capabilitytypes.ModuleName] = true
}
