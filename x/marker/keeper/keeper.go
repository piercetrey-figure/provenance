package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibctypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"

	"github.com/provenance-io/provenance/x/marker/types"
)

// Handler is a handler function for use with IterateRecords.
type Handler func(record types.MarkerAccountI) error

// MarkerKeeperI provides a read/write iterate interface to marker acccounts in the auth account keeper store
type MarkerKeeperI interface {
	// Returns a new marker instance with the address and baseaccount assigned.  Does not save to auth store
	NewMarker(sdk.Context, types.MarkerAccountI) types.MarkerAccountI

	// GetMarker looks up a marker by a given address
	GetMarker(sdk.Context, sdk.AccAddress) (types.MarkerAccountI, error)
	// Set a marker in the auth account store
	SetMarker(sdk.Context, types.MarkerAccountI)
	// Remove a marker from the auth account store
	RemoveMarker(sdk.Context, types.MarkerAccountI)

	GetEscrow(sdk.Context, types.MarkerAccountI) sdk.Coins

	// IterateMarker processes all markers with the given handler function.
	IterateMarkers(sdk.Context, func(types.MarkerAccountI) bool)

	// GetAuthority returns the signing authority
	GetAuthority() string
}

// Keeper defines the name module Keeper
type Keeper struct {
	// The reference to the Paramstore to get and set account specific params
	paramSpace paramtypes.Subspace

	// To check whether accounts exist for addresses.
	authKeeper types.AccountKeeper

	// To check whether accounts exist for addresses.
	authzKeeper types.AuthzKeeper

	// To handle movement of coin between accounts and check total supply
	bankKeeper types.BankKeeper

	// To pass through grant creation for callers with admin access on a marker.
	feegrantKeeper types.FeeGrantKeeper

	// To access attributes for addresses
	attrKeeper types.AttrKeeper
	// To access names and normalize required attributes
	nameKeeper types.NameKeeper

	// Key to access the key-value store from sdk.Context.
	storeKey storetypes.StoreKey

	// The codec for binary encoding/decoding.
	cdc codec.BinaryCodec

	// the signing authority for the gov proposals
	authority string

	markerModuleAddr sdk.AccAddress

	ibcTransferModuleAddr sdk.AccAddress

	// Used to transfer the ibc marker
	ibcTransferServer types.IbcTransferMsgServer

	// reqAttrBypassAddrs is a set of addresses that are allowed to bypass the required attribute check.
	// When sending to one of these, if there are required attributes, it behaves as if the addr has them;
	// if there aren't required attributes, the sender still needs transfer permission.
	// When sending from one of these, if there are required attributes, the destination must have them;
	// if there aren't required attributes, it behaves as if the sender has transfer permission.
	reqAttrBypassAddrs types.ImmutableAccAddresses
}

// NewKeeper returns a marker keeper. It handles:
// - managing MarkerAccounts
// - enforcing permissions for marker creation/deletion/management
//
// CONTRACT: the parameter Subspace must have the param key table already initialized
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	authKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	authzKeeper types.AuthzKeeper,
	feegrantKeeper types.FeeGrantKeeper,
	attrKeeper types.AttrKeeper,
	nameKeeper types.NameKeeper,
	ibcTransferServer types.IbcTransferMsgServer,
	reqAttrBypassAddrs []sdk.AccAddress,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	rv := Keeper{
		paramSpace:            paramSpace,
		authKeeper:            authKeeper,
		authzKeeper:           authzKeeper,
		bankKeeper:            bankKeeper,
		feegrantKeeper:        feegrantKeeper,
		attrKeeper:            attrKeeper,
		nameKeeper:            nameKeeper,
		storeKey:              key,
		cdc:                   cdc,
		authority:             authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		markerModuleAddr:      authtypes.NewModuleAddress(types.CoinPoolName),
		ibcTransferModuleAddr: authtypes.NewModuleAddress(ibctypes.ModuleName),
		ibcTransferServer:     ibcTransferServer,
		reqAttrBypassAddrs:    types.NewImmutableAccAddresses(reqAttrBypassAddrs),
	}
	bankKeeper.AppendSendRestriction(rv.SendRestrictionFn)
	return rv
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

var _ MarkerKeeperI = &Keeper{}

// NewMarker returns a new marker instance with the address and baseaccount assigned.  Does not save to auth store
func (k Keeper) NewMarker(ctx sdk.Context, marker types.MarkerAccountI) types.MarkerAccountI {
	return k.authKeeper.NewAccount(ctx, marker).(types.MarkerAccountI)
}

// GetMarker looks up a marker by a given address
func (k Keeper) GetMarker(ctx sdk.Context, address sdk.AccAddress) (types.MarkerAccountI, error) {
	mac := k.authKeeper.GetAccount(ctx, address)
	if mac != nil {
		macc, ok := mac.(types.MarkerAccountI)
		if !ok {
			return nil, fmt.Errorf("account at %s is not a marker account", address.String())
		}
		return macc, nil
	}
	return nil, nil
}

// SetMarker sets a marker in the auth account store will panic if the marker account is not valid or
// if the auth module account keeper fails to marshall the account.
func (k Keeper) SetMarker(ctx sdk.Context, marker types.MarkerAccountI) {
	store := ctx.KVStore(k.storeKey)

	if err := marker.Validate(); err != nil {
		panic(err)
	}
	k.authKeeper.SetAccount(ctx, marker)
	store.Set(types.MarkerStoreKey(marker.GetAddress()), marker.GetAddress())
}

// RemoveMarker removes a marker from the auth account store. Note: if the account holds coins this will
// likely cause an invariant constraint violation for the coin supply
func (k Keeper) RemoveMarker(ctx sdk.Context, marker types.MarkerAccountI) {
	store := ctx.KVStore(k.storeKey)
	k.authKeeper.RemoveAccount(ctx, marker)

	k.RemoveNetAssetValues(ctx, marker.GetAddress())
	store.Delete(types.MarkerStoreKey(marker.GetAddress()))
}

// IterateMarkers iterates all markers with the given handler function.
func (k Keeper) IterateMarkers(ctx sdk.Context, cb func(marker types.MarkerAccountI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.MarkerStoreKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		account := k.authKeeper.GetAccount(ctx, iterator.Value())
		ma, ok := account.(types.MarkerAccountI)
		if !ok {
			panic(fmt.Errorf("invalid account type in marker account registry"))
		}
		if cb(ma) {
			break
		}
	}
}

// GetEscrow returns the balances of all coins held in escrow in the marker
func (k Keeper) GetEscrow(ctx sdk.Context, marker types.MarkerAccountI) sdk.Coins {
	return k.bankKeeper.GetAllBalances(ctx, marker.GetAddress())
}

// GetAuthority is signer of the proposal
func (k Keeper) GetAuthority() string {
	return k.authority
}

// IsSendDeny returns true if sender address is denied for marker
func (k Keeper) IsSendDeny(ctx sdk.Context, markerAddr, senderAddr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.DenySendKey(markerAddr, senderAddr))
}

// AddSendDeny set sender address to denied for marker
func (k Keeper) AddSendDeny(ctx sdk.Context, markerAddr, senderAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.DenySendKey(markerAddr, senderAddr), []byte{})
}

// RemoveSendDeny removes sender address from marker deny list
func (k Keeper) RemoveSendDeny(ctx sdk.Context, markerAddr, senderAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DenySendKey(markerAddr, senderAddr))
}

// IterateMarkers  iterates all markers with the given handler function.
func (k Keeper) IterateSendDeny(ctx sdk.Context, handler func(key []byte) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DenySendKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		if handler(iterator.Key()) {
			break
		}
	}
}

// AddSetNetAssetValues adds a set of net asset values to a marker
func (k Keeper) AddSetNetAssetValues(ctx sdk.Context, marker types.MarkerAccountI, netAssetValues []types.NetAssetValue, source string) error {
	for _, nav := range netAssetValues {
		if nav.Price.Denom == marker.GetDenom() {
			return fmt.Errorf("net asset value denom cannot match marker denom %q", marker.GetDenom())
		}
		if nav.Price.Denom != types.UsdDenom {
			_, err := k.GetMarkerByDenom(ctx, nav.Price.Denom)
			if err != nil {
				return fmt.Errorf("net asset value denom does not exist: %v", err.Error())
			}
		}

		if err := k.SetNetAssetValue(ctx, marker, nav, source); err != nil {
			return fmt.Errorf("cannot set net asset value : %v", err.Error())
		}
	}
	return nil
}

// SetNetAssetValue adds/updates a net asset value to marker
func (k Keeper) SetNetAssetValue(ctx sdk.Context, marker types.MarkerAccountI, netAssetValue types.NetAssetValue, source string) error {
	netAssetValue.UpdatedBlockHeight = uint64(ctx.BlockHeight())
	if err := netAssetValue.Validate(); err != nil {
		return err
	}

	setNetAssetValueEvent := types.NewEventSetNetAssetValue(marker.GetDenom(), netAssetValue.Price, netAssetValue.Volume, source)
	if err := ctx.EventManager().EmitTypedEvent(setNetAssetValueEvent); err != nil {
		return err
	}

	key := types.NetAssetValueKey(marker.GetAddress(), netAssetValue.Price.Denom)
	store := ctx.KVStore(k.storeKey)
	if math.NewIntFromUint64(netAssetValue.Volume).GT(marker.GetSupply().Amount) {
		return fmt.Errorf("volume(%v) cannot exceed marker %q supply(%v) ", netAssetValue.Volume, marker.GetDenom(), marker.GetSupply())
	}

	bz, err := k.cdc.Marshal(&netAssetValue)
	if err != nil {
		return err
	}
	store.Set(key, bz)

	return nil
}

// IterateNetAssetValues iterates net asset values for marker
func (k Keeper) IterateNetAssetValues(ctx sdk.Context, markerAddr sdk.AccAddress, handler func(state types.NetAssetValue) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, types.NetAssetValueKeyPrefix(markerAddr))
	defer it.Close()
	for ; it.Valid(); it.Next() {
		var markerNav types.NetAssetValue
		err := k.cdc.Unmarshal(it.Value(), &markerNav)
		if err != nil {
			return err
		} else if handler(markerNav) {
			break
		}
	}
	return nil
}

// RemoveNetAssetValues removes all net asset values for a marker
func (k Keeper) RemoveNetAssetValues(ctx sdk.Context, markerAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, types.NetAssetValueKeyPrefix(markerAddr))
	var keys [][]byte
	for ; it.Valid(); it.Next() {
		keys = append(keys, it.Key())
	}
	it.Close()

	for _, key := range keys {
		store.Delete(key)
	}
}

// GetReqAttrBypassAddrs returns a deep copy of the addresses that bypass the required attributes checking.
func (k Keeper) GetReqAttrBypassAddrs() []sdk.AccAddress {
	return k.reqAttrBypassAddrs.GetSlice()
}

// IsReqAttrBypassAddr returns true if the provided addr can bypass the required attributes checking.
func (k Keeper) IsReqAttrBypassAddr(addr sdk.AccAddress) bool {
	return k.reqAttrBypassAddrs.Has(addr)
}
