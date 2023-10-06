package module_filtering

import "github.com/provenance-io/provenance/app/module_filtering/modules"

// var ActiveModules = map[string]bool{
// 	// upgradetypes.ModuleName:     true,
// 	// capabilitytypes.ModuleName:  true,
// 	// minttypes.ModuleName:        true,
// 	// distrtypes.ModuleName:       true,
// 	// slashingtypes.ModuleName:    true, // to disable
// 	// evidencetypes.ModuleName:    true, // to disable
// 	// stakingtypes.ModuleName:     true,
// 	// ibchost.ModuleName:          true,
// 	// markertypes.ModuleName:      true,
// 	// icatypes.ModuleName:         true,
// 	// attributetypes.ModuleName:   true, // to disable
// 	// rewardtypes.ModuleName:      true, // to disable
// 	// triggertypes.ModuleName:     true, // to disable
// 	// authtypes.ModuleName:        true,
// 	// banktypes.ModuleName:        true,
// 	// govtypes.ModuleName:         true,
// 	// crisistypes.ModuleName:      true,
// 	// genutiltypes.ModuleName:     true,
// 	// authz.ModuleName:            true,
// 	// group.ModuleName:            true, // to disable
// 	// feegrant.ModuleName:         true, // to disable
// 	// paramstypes.ModuleName:      true,
// 	// msgfeestypes.ModuleName:     true, // to disable
// 	// metadatatypes.ModuleName:    true, // to disable
// 	// oracletypes.ModuleName:      true,
// 	// wasm.ModuleName:             true,
// 	// ibchookstypes.ModuleName:    true,
// 	// ibctransfertypes.ModuleName: true,
// 	// icqtypes.ModuleName:         true,
// 	// nametypes.ModuleName:        true,
// 	// vestingtypes.ModuleName:     true,
// 	// quarantine.ModuleName:       true, // to disable
// 	// sanction.ModuleName:         true, // to disable
// 	// hold.ModuleName:             true, // to disable
// }

func FilterActiveModules(moduleList ...string) []string {
	var filteredModules = make([]string, 0, len(modules.ActiveModules))
	for _, v := range moduleList {
		if modules.ActiveModules[v] {
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
	var filteredModules = make([]T, 0, len(modules.ActiveModules))
	for _, v := range moduleList {
		if modules.ActiveModules[v.First] {
			filteredModules = append(filteredModules, v.Second())
		}
	}
	return filteredModules
}
