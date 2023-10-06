//go:build (include && sanction) || (exclude && !sanction)

package modules

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/sanction"
)

func init() {
	fmt.Println("Including Module: ", sanction.ModuleName)
	ActiveModules[sanction.ModuleName] = true
}
