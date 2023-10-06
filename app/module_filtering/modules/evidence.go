//go:build (include && evidence) || (exclude && !evidence)

package modules

import (
	"fmt"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
)

func init() {
	fmt.Println("Including Module: ", evidencetypes.ModuleName)
	ActiveModules[evidencetypes.ModuleName] = true
}
