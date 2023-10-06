//go:build (include && mint) || (exclude && !mint)

package modules
    
import (
    minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", minttypes.ModuleName)
    ActiveModules[minttypes.ModuleName] = true
}
