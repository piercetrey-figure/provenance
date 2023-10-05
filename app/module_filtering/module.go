package module_filtering

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	icqtypes "github.com/strangelove-ventures/async-icq/v6/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"

	ibchookstypes "github.com/provenance-io/provenance/x/ibchooks/types"
	markertypes "github.com/provenance-io/provenance/x/marker/types"
	nametypes "github.com/provenance-io/provenance/x/name/types"
	oracletypes "github.com/provenance-io/provenance/x/oracle/types"
)

var activeModules = map[string]bool{
	upgradetypes.ModuleName:    true,
	capabilitytypes.ModuleName: true,
	minttypes.ModuleName:       true,
	distrtypes.ModuleName:      true,
	// slashingtypes.ModuleName:    true,
	// evidencetypes.ModuleName:    true,
	stakingtypes.ModuleName: true,
	ibchost.ModuleName:      true,
	markertypes.ModuleName:  true,
	icatypes.ModuleName:     true,
	// attributetypes.ModuleName:   true,
	// rewardtypes.ModuleName:      true,
	// triggertypes.ModuleName:     true,
	authtypes.ModuleName:    true,
	banktypes.ModuleName:    true,
	govtypes.ModuleName:     true,
	crisistypes.ModuleName:  true,
	genutiltypes.ModuleName: true,
	authz.ModuleName:        true,
	// group.ModuleName:            true,
	// feegrant.ModuleName:         true,
	paramstypes.ModuleName: true,
	// msgfeestypes.ModuleName:     true,
	// metadatatypes.ModuleName:    true,
	oracletypes.ModuleName:      true,
	wasm.ModuleName:             true,
	ibchookstypes.ModuleName:    true,
	ibctransfertypes.ModuleName: true,
	icqtypes.ModuleName:         true,
	nametypes.ModuleName:        true,
	vestingtypes.ModuleName:     true,
	// quarantine.ModuleName:       true,
	// sanction.ModuleName:         true,
	// hold.ModuleName:             true,
}

func FilterActiveModules(moduleList ...string) []string {
	var filteredModules = make([]string, 0, len(activeModules))
	for _, v := range moduleList {
		if activeModules[v] {
			filteredModules = append(filteredModules, v)
		}
	}
	return filteredModules
}

type pair[T, U any] struct {
	First  T
	Second U
}

func Pair[T, U any](first T, second U) pair[T, U] {
	return pair[T, U]{first, second}
}

func ApplyActiveModules[T interface{}](moduleList ...pair[string, func() T]) []T {
	var filteredModules = make([]T, 0, len(activeModules))
	for _, v := range moduleList {
		if activeModules[v.First] {
			filteredModules = append(filteredModules, v.Second())
		}
	}
	return filteredModules
}
