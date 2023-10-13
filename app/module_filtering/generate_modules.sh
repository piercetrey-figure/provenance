cat << EOF > modules/modules.go
package modules
var ActiveModules = map[string]bool{}
EOF

cat << EOF > modules/upgrade.go
//go:build (include && upgrade) || (exclude && !upgrade)

package modules

import (
    upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
    "fmt"
)

func init() {

    ActiveModules[upgradetypes.ModuleName] = true
}
EOF
cat << EOF > modules/capability.go
//go:build (include && capability) || (exclude && !capability)

package modules
    
import (
    capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
    "fmt"
)

func init() {

    ActiveModules[capabilitytypes.ModuleName] = true
}
EOF
cat << EOF > modules/mint.go
//go:build (include && mint) || (exclude && !mint)

package modules
    
import (
    minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
    "fmt"
)

func init() {

    ActiveModules[minttypes.ModuleName] = true
}
EOF
cat << EOF > modules/distr.go
//go:build (include && distr) || (exclude && !distr)

package modules

import (
    distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
    "fmt"
)

func init() {

    ActiveModules[distrtypes.ModuleName] = true
}
EOF
cat << EOF > modules/slashing.go
//go:build (include && slashing) || (exclude && !slashing)

package modules

import (
    slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
    "fmt"
)

func init() {

    ActiveModules[slashingtypes.ModuleName] = true
}
EOF
cat << EOF > modules/evidence.go
//go:build (include && evidence) || (exclude && !evidence)

package modules

import (
    evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
    "fmt"
)

func init() {

    ActiveModules[evidencetypes.ModuleName] = true
}
EOF
cat << EOF > modules/staking.go
//go:build (include && staking) || (exclude && !staking)

package modules

import (
    stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
    "fmt"
)

func init() {

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

    ActiveModules[ibchost.ModuleName] = true
}
EOF
cat << EOF > modules/marker.go
//go:build (include && marker) || (exclude && !marker)

package modules

import (
    markertypes "github.com/provenance-io/provenance/x/marker/types"
    "fmt"
)

func init() {

    ActiveModules[markertypes.ModuleName] = true
}
EOF
cat << EOF > modules/ica.go
//go:build (include && ica) || (exclude && !ica)

package modules

import (
    icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
    "fmt"
)

func init() {

    ActiveModules[icatypes.ModuleName] = true
}
EOF
cat << EOF > modules/attribute.go
//go:build (include && attribute) || (exclude && !attribute)

package modules

import (
    attributetypes "github.com/provenance-io/provenance/x/attribute/types"
    "fmt"
)

func init() {

    ActiveModules[attributetypes.ModuleName] = true
}
EOF
cat << EOF > modules/reward.go
//go:build (include && reward) || (exclude && !reward)

package modules

import (
    rewardtypes "github.com/provenance-io/provenance/x/reward/types"
    "fmt"
)

func init() {

    ActiveModules[rewardtypes.ModuleName] = true
}
EOF
cat << EOF > modules/trigger.go
//go:build (include && trigger) || (exclude && !trigger)

package modules

import (
    triggertypes "github.com/provenance-io/provenance/x/trigger/types"
    "fmt"
)

func init() {

    ActiveModules[triggertypes.ModuleName] = true
}
EOF
cat << EOF > modules/auth.go
//go:build (include && auth) || (exclude && !auth)

package modules

import (
    authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
    "fmt"
)

func init() {

    ActiveModules[authtypes.ModuleName] = true
}
EOF
cat << EOF > modules/bank.go
//go:build (include && bank) || (exclude && !bank)

package modules

import (
    banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
    "fmt"
)

func init() {

    ActiveModules[banktypes.ModuleName] = true
}
EOF
cat << EOF > modules/gov.go
//go:build (include && gov) || (exclude && !gov)

package modules

import (
    govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
    "fmt"
)

func init() {

    ActiveModules[govtypes.ModuleName] = true
}
EOF
cat << EOF > modules/crisis.go
//go:build (include && crisis) || (exclude && !crisis)

package modules

import (
    crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
    "fmt"
)

func init() {

    ActiveModules[crisistypes.ModuleName] = true
}
EOF
cat << EOF > modules/genutil.go
//go:build (include && genutil) || (exclude && !genutil)

package modules

import (
    genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
    "fmt"
)

func init() {

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

    ActiveModules[feegrant.ModuleName] = true
}
EOF
cat << EOF > modules/params.go
//go:build (include && params) || (exclude && !params)

package modules

import (
    paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
    "fmt"
)

func init() {

    ActiveModules[paramstypes.ModuleName] = true
}
EOF
cat << EOF > modules/msgfees.go
//go:build (include && msgfees) || (exclude && !msgfees)

package modules

import (
    msgfeestypes "github.com/provenance-io/provenance/x/msgfees/types"
    "fmt"
)

func init() {

    ActiveModules[msgfeestypes.ModuleName] = true
}
EOF
cat << EOF > modules/metadata.go
//go:build (include && metadata) || (exclude && !metadata)

package modules

import (
    metadatatypes "github.com/provenance-io/provenance/x/metadata/types"
    "fmt"
)

func init() {

    ActiveModules[metadatatypes.ModuleName] = true
}
EOF
cat << EOF > modules/oracle.go
//go:build (include && oracle) || (exclude && !oracle)

package modules

import (
    oracletypes "github.com/provenance-io/provenance/x/oracle/types"
    "fmt"
)

func init() {

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

    ActiveModules[wasm.ModuleName] = true
}
EOF
cat << EOF > modules/ibchooks.go
//go:build (include && ibchooks) || (exclude && !ibchooks)

package modules

import (
    ibchookstypes "github.com/provenance-io/provenance/x/ibchooks/types"
    "fmt"
)

func init() {

    ActiveModules[ibchookstypes.ModuleName] = true
}
EOF
cat << EOF > modules/ibctransfer.go
//go:build (include && ibctransfer) || (exclude && !ibctransfer)

package modules

import (
    ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
    "fmt"
)

func init() {

    ActiveModules[ibctransfertypes.ModuleName] = true
}
EOF
cat << EOF > modules/icq.go
//go:build (include && icq) || (exclude && !icq)

package modules

import (
    icqtypes "github.com/strangelove-ventures/async-icq/v6/types"
    "fmt"
)

func init() {

    ActiveModules[icqtypes.ModuleName] = true
}
EOF
cat << EOF > modules/name.go
//go:build (include && name) || (exclude && !name)

package modules

import (
    nametypes "github.com/provenance-io/provenance/x/name/types"
    "fmt"
)

func init() {

    ActiveModules[nametypes.ModuleName] = true
}
EOF
cat << EOF > modules/vesting.go
//go:build (include && vesting) || (exclude && !vesting)

package modules

import (
    vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
    "fmt"
)

func init() {

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

    ActiveModules[hold.ModuleName] = true
}
EOF