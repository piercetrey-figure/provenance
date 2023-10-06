//go:build (include && wasm) || (exclude && !wasm)

package modules

import (
	"fmt"
	"github.com/CosmWasm/wasmd/x/wasm"
)

func init() {
	fmt.Println("Including Module: ", wasm.ModuleName)
	ActiveModules[wasm.ModuleName] = true
}
