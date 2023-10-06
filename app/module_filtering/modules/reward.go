//go:build (include && reward) || (exclude && !reward)

package modules

import (
    rewardtypes "github.com/provenance-io/provenance/x/reward/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", rewardtypes.ModuleName)
    ActiveModules[rewardtypes.ModuleName] = true
}
