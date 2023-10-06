//go:build (include && group) || (exclude && !group)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/group"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", group.ModuleName)
    ActiveModules[group.ModuleName] = true
}
