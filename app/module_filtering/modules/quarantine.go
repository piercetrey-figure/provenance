//go:build (include && quarantine) || (exclude && !quarantine)

package modules

import (
	"github.com/cosmos/cosmos-sdk/x/quarantine"
)

func init() {

	ActiveModules[quarantine.ModuleName] = true
}
