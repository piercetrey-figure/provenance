//go:build (include && evidence) || (exclude && !evidence)

package modules

import (
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
)

func init() {

	ActiveModules[evidencetypes.ModuleName] = true
}
