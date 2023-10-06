//go:build (include && evidence) || (exclude && !evidence)

package modules

import (
    evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", evidencetypes.ModuleName)
    ActiveModules[evidencetypes.ModuleName] = true
}
