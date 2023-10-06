//go:build (include && reward) || (exclude && !reward)

package modules

import (
	"fmt"
	rewardtypes "github.com/provenance-io/provenance/x/reward/types"
)

func init() {
	fmt.Println("Including Module: ", rewardtypes.ModuleName)
	ActiveModules[rewardtypes.ModuleName] = true
}
