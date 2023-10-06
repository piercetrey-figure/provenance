//go:build (include && wasm) || (exclude && !wasm)

package modules

import (
    "github.com/CosmWasm/wasmd/x/wasm"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", wasm.ModuleName)
    ActiveModules[wasm.ModuleName] = true
}
