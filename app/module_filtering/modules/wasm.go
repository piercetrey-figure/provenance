//go:build (include && wasm) || (exclude && !wasm)

package modules

import (
	"github.com/CosmWasm/wasmd/x/wasm"
)

func init() {

	ActiveModules[wasm.ModuleName] = true
}
