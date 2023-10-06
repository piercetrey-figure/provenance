//go:build (include && msgfees) || (exclude && !msgfees)

package modules

import (
    msgfeestypes "github.com/provenance-io/provenance/x/msgfees/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", msgfeestypes.ModuleName)
    ActiveModules[msgfeestypes.ModuleName] = true
}
