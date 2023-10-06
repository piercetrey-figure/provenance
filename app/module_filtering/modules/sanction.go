//go:build (include && sanction) || (exclude && !sanction)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/sanction"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", sanction.ModuleName)
    ActiveModules[sanction.ModuleName] = true
}
