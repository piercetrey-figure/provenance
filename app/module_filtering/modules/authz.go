//go:build (include && authz) || (exclude && !authz)

package modules

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func init() {
	fmt.Println("Including Module: ", authz.ModuleName)
	ActiveModules[authz.ModuleName] = true
}
