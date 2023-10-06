//go:build (include && quarantine) || (exclude && !quarantine)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/quarantine"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", quarantine.ModuleName)
    ActiveModules[quarantine.ModuleName] = true
}
