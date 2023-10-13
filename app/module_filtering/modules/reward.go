//go:build (include && reward) || (exclude && !reward)

package modules

import (
	rewardtypes "github.com/provenance-io/provenance/x/reward/types"
)

func init() {

	ActiveModules[rewardtypes.ModuleName] = true
}
