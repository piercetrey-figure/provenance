cat << EOF > modules/modules.go
package modules
var ActiveModules = map[string]bool{}
EOF

cat << EOF > modules/upgradetypes.go
//go:build (include && upgrade) || (exclude && !upgrade)

package modules

import (
    upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", upgradetypes.ModuleName)
    ActiveModules[upgradetypes.ModuleName] = true
}
EOF
cat << EOF > modules/capabilitytypes.go
//go:build (include && capability) || (exclude && !capability)

package modules
    
import (
    capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", capabilitytypes.ModuleName)
    ActiveModules[capabilitytypes.ModuleName] = true
}
EOF
cat << EOF > modules/minttypes.go
//go:build (include && mint) || (exclude && !mint)

package modules
    
import (
    minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", minttypes.ModuleName)
    ActiveModules[minttypes.ModuleName] = true
}
EOF
cat << EOF > modules/distrtypes.go
//go:build (include && distr) || (exclude && !distr)

package modules

import (
    distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", distrtypes.ModuleName)
    ActiveModules[distrtypes.ModuleName] = true
}
EOF
cat << EOF > modules/slashingtypes.go
//go:build (include && slashing) || (exclude && !slashing)

package modules

import (
    slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", slashingtypes.ModuleName)
    ActiveModules[slashingtypes.ModuleName] = true
}
EOF
cat << EOF > modules/evidencetypes.go
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
EOF
cat << EOF > modules/stakingtypes.go
//go:build (include && staking) || (exclude && !staking)

package modules

import (
    stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", stakingtypes.ModuleName)
    ActiveModules[stakingtypes.ModuleName] = true
}
EOF
cat << EOF > modules/ibchost.go
//go:build (include && ibchost) || (exclude && !ibchost)

package modules

import (
    ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", ibchost.ModuleName)
    ActiveModules[ibchost.ModuleName] = true
}
EOF
cat << EOF > modules/markertypes.go
//go:build (include && marker) || (exclude && !marker)

package modules

import (
    markertypes "github.com/provenance-io/provenance/x/marker/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", markertypes.ModuleName)
    ActiveModules[markertypes.ModuleName] = true
}
EOF
cat << EOF > modules/icatypes.go
//go:build (include && ica) || (exclude && !ica)

package modules

import (
    icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", icatypes.ModuleName)
    ActiveModules[icatypes.ModuleName] = true
}
EOF
cat << EOF > modules/attributetypes.go
//go:build (include && attribute) || (exclude && !attribute)

package modules

import (
    attributetypes "github.com/provenance-io/provenance/x/attribute/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", attributetypes.ModuleName)
    ActiveModules[attributetypes.ModuleName] = true
}
EOF
cat << EOF > modules/rewardtypes.go
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
EOF
cat << EOF > modules/triggertypes.go
//go:build (include && trigger) || (exclude && !trigger)

package modules

import (
    triggertypes "github.com/provenance-io/provenance/x/trigger/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", triggertypes.ModuleName)
    ActiveModules[triggertypes.ModuleName] = true
}
EOF
cat << EOF > modules/authtypes.go
//go:build (include && auth) || (exclude && !auth)

package modules

import (
    authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", authtypes.ModuleName)
    ActiveModules[authtypes.ModuleName] = true
}
EOF
cat << EOF > modules/banktypes.go
//go:build (include && bank) || (exclude && !bank)

package modules

import (
    banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", banktypes.ModuleName)
    ActiveModules[banktypes.ModuleName] = true
}
EOF
cat << EOF > modules/govtypes.go
//go:build (include && gov) || (exclude && !gov)

package modules

import (
    govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", govtypes.ModuleName)
    ActiveModules[govtypes.ModuleName] = true
}
EOF
cat << EOF > modules/crisistypes.go
//go:build (include && crisis) || (exclude && !crisis)

package modules

import (
    crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", crisistypes.ModuleName)
    ActiveModules[crisistypes.ModuleName] = true
}
EOF
cat << EOF > modules/genutiltypes.go
//go:build (include && genutil) || (exclude && !genutil)

package modules

import (
    genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", genutiltypes.ModuleName)
    ActiveModules[genutiltypes.ModuleName] = true
}
EOF
cat << EOF > modules/authz.go
//go:build (include && authz) || (exclude && !authz)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/authz"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", authz.ModuleName)
    ActiveModules[authz.ModuleName] = true
}
EOF
cat << EOF > modules/group.go
//go:build (include && group) || (exclude && !group)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/group"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", group.ModuleName)
    ActiveModules[group.ModuleName] = true
}
EOF
cat << EOF > modules/feegrant.go
//go:build (include && feegrant) || (exclude && !feegrant)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/feegrant"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", feegrant.ModuleName)
    ActiveModules[feegrant.ModuleName] = true
}
EOF
cat << EOF > modules/paramstypes.go
//go:build (include && params) || (exclude && !params)

package modules

import (
    paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", paramstypes.ModuleName)
    ActiveModules[paramstypes.ModuleName] = true
}
EOF
cat << EOF > modules/msgfeestypes.go
//go:build (include && msgfees) || (exclude && !msgfees)

package modules

import (
    msgfeestypes "github.com/provenance-io/provenance/x/msgfees/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", msgfeestypes.ModuleName)
    ActiveModules[msgfeestypes.ModuleName] = true
}
EOF
cat << EOF > modules/metadatatypes.go
//go:build (include && metadata) || (exclude && !metadata)

package modules

import (
    metadatatypes "github.com/provenance-io/provenance/x/metadata/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", metadatatypes.ModuleName)
    ActiveModules[metadatatypes.ModuleName] = true
}
EOF
cat << EOF > modules/oracletypes.go
//go:build (include && oracle) || (exclude && !oracle)

package modules

import (
    oracletypes "github.com/provenance-io/provenance/x/oracle/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", oracletypes.ModuleName)
    ActiveModules[oracletypes.ModuleName] = true
}
EOF
cat << EOF > modules/wasm.go
//go:build (include && wasm) || (exclude && !wasm)

package modules

import (
    "github.com/CosmWasm/wasmd/x/wasm"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", wasm.ModuleName)
    ActiveModules[wasm.ModuleName] = true
}
EOF
cat << EOF > modules/ibchookstypes.go
//go:build (include && ibchooks) || (exclude && !ibchooks)

package modules

import (
    ibchookstypes "github.com/provenance-io/provenance/x/ibchooks/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", ibchookstypes.ModuleName)
    ActiveModules[ibchookstypes.ModuleName] = true
}
EOF
cat << EOF > modules/ibctransfertypes.go
//go:build (include && ibctransfer) || (exclude && !ibctransfer)

package modules

import (
    ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", ibctransfertypes.ModuleName)
    ActiveModules[ibctransfertypes.ModuleName] = true
}
EOF
cat << EOF > modules/icqtypes.go
//go:build (include && icq) || (exclude && !icq)

package modules

import (
    icqtypes "github.com/strangelove-ventures/async-icq/v6/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", icqtypes.ModuleName)
    ActiveModules[icqtypes.ModuleName] = true
}
EOF
cat << EOF > modules/nametypes.go
//go:build (include && name) || (exclude && !name)

package modules

import (
    nametypes "github.com/provenance-io/provenance/x/name/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", nametypes.ModuleName)
    ActiveModules[nametypes.ModuleName] = true
}
EOF
cat << EOF > modules/vestingtypes.go
//go:build (include && vesting) || (exclude && !vesting)

package modules

import (
    vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", vestingtypes.ModuleName)
    ActiveModules[vestingtypes.ModuleName] = true
}
EOF
cat << EOF > modules/quarantine.go
//go:build (include && quarantine) || (exclude && !quarantine)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/quarantine"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", quarantine.ModuleName)
    ActiveModules[quarantine.ModuleName] = true
}
EOF
cat << EOF > modules/sanction.go
//go:build (include && sanction) || (exclude && !sanction)

package modules

import (
    "github.com/cosmos/cosmos-sdk/x/sanction"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", sanction.ModuleName)
    ActiveModules[sanction.ModuleName] = true
}
EOF
cat << EOF > modules/hold.go
//go:build (include && hold) || (exclude && !hold)

package modules

import (
    "github.com/provenance-io/provenance/x/hold"
    "fmt"
)

func init() {
    fmt.Println("Including Module: ", hold.ModuleName)
    ActiveModules[hold.ModuleName] = true
}
EOF