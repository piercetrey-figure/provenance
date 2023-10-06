//go:build (include && msgfees) || (exclude && !msgfees)

package modules

import (
	"fmt"
	msgfeestypes "github.com/provenance-io/provenance/x/msgfees/types"
)

func init() {
	fmt.Println("Including Module: ", msgfeestypes.ModuleName)
	ActiveModules[msgfeestypes.ModuleName] = true
}
