//go:build (include && hold) || (exclude && !hold)

package modules

import (
	"fmt"
	"github.com/provenance-io/provenance/x/hold"
)

func init() {
	fmt.Println("Including Module: ", hold.ModuleName)
	ActiveModules[hold.ModuleName] = true
}
