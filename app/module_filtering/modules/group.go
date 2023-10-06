//go:build (include && group) || (exclude && !group)

package modules

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/group"
)

func init() {
	fmt.Println("Including Module: ", group.ModuleName)
	ActiveModules[group.ModuleName] = true
}
