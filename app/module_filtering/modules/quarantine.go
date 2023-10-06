//go:build (include && quarantine) || (exclude && !quarantine)

package modules

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/quarantine"
)

func init() {
	fmt.Println("Including Module: ", quarantine.ModuleName)
	ActiveModules[quarantine.ModuleName] = true
}
