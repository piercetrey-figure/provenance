//go:build (include && hold) || (exclude && !hold)

package modules

import (
    "github.com/provenance-io/provenance/x/hold"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", hold.ModuleName)
    ActiveModules[hold.ModuleName] = true
}
