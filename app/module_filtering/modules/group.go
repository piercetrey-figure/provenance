//go:build (include && group) || (exclude && !group)

package modules

import (
	"github.com/cosmos/cosmos-sdk/x/group"
)

func init() {

	ActiveModules[group.ModuleName] = true
}
