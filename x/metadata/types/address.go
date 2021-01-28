package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

const (
	// PrefixScope is the address human readable prefix used with bech32 encoding of Scope IDs
	PrefixScope = "scope"
	// PrefixGroup is the address human readable prefix used with bech32 encoding of Group IDs
	PrefixGroup = "group"
	// PrefixRecord is the address human readable prefix used with bech32 encoding of Record IDs
	PrefixRecord = "record"
	// PrefixScopeSpecification is the address human readable prefix used with bech32 encoding of ScopeSpecification IDs
	PrefixScopeSpecification = "scopespec"
	// PrefixGroupSpecification is the address human readable prefix used with bech32 encoding of GroupSpecification IDs
	PrefixGroupSpecification = "groupspec"
)

var (
	// Ensure MetadataAddress implements the sdk.Address interface
	_ sdk.Address = MetadataAddress{}
)

// MetadataAddress is a blockchain compliant address based on UUIDs
type MetadataAddress []byte

// VerifyMetadataAddressFormat checks a sequence of bytes for proper format as a MetadataAddress instance
// returns the associated bech32 hrp/type name or any errors encountered during verification
func VerifyMetadataAddressFormat(bz []byte) (string, error) {
	hrp := ""
	requiredLength := 1 + 16 // type byte plus size of one uuid
	if len(bz) < requiredLength {
		return hrp, fmt.Errorf("incorrect address length (must be at least 17, actual: %d)", len(bz))
	}
	switch bz[0] {
	case ScopeKeyPrefix[0]:
		hrp = PrefixScope
		requiredLength = 1 + 16 // type byte plus size of one uuid
	case GroupKeyPrefix[0]:
		hrp = PrefixGroup
		requiredLength = 1 + 16 + 16 // type byte plus size of two uuids
		// TODO -- check the format of the second uuid in this type here
	case RecordKeyPrefix[0]:
		hrp = PrefixRecord
		requiredLength = 1 + 16 + 32 // type byte plus size of one uuid and one sha256 hash
	case ScopeSpecificationPrefix[0]:
		hrp = PrefixScopeSpecification
		requiredLength = 1 + 16 // type byte plus size of one uuid
	case GroupSpecificationPrefix[0]:
		hrp = PrefixGroupSpecification
		requiredLength = 1 + 16 // type byte plus size of one uuid

	default:
		return hrp, fmt.Errorf("invalid metadata address type (must be 0-4, actual: %d)", bz[0])
	}
	if len(bz) != requiredLength {
		return hrp, fmt.Errorf("incorrect address length (expected: %d, actual: %d)", requiredLength, len(bz))
	}
	// all valid metdata address have at least one uuid
	if _, err := uuid.FromBytes(bz[1:17]); err != nil {
		return hrp, fmt.Errorf("invalid address bytes, expected uuid compliant: %w", err)
	}
	return hrp, nil
}

// ConvertHashToAddress constructs a MetadataAddress using the provided type code and the first 16 bytes of the
// base64 decoded hash.  Resulting Address is not guaranteed to contain a valid V4 UUID (random only)
func ConvertHashToAddress(typeCode []byte, hash string) (addr MetadataAddress, err error) {
	var raw []byte
	raw, err = base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return addr, err
	}
	if len(raw) < 16 {
		return addr, fmt.Errorf("invalid specification identifier, expected at least 16 bytes, found %d", len(raw))
	}
	// The codes 0,3,4 all start with a code followed by uuid bytes meaning we can create a valid address with them
	if len(typeCode) > 1 || !(typeCode[0] == 0x00 || typeCode[0] == 0x03 || typeCode[0] == 0x04) {
		return addr, fmt.Errorf("invalid address type code 0x%X, expected 0x00, 0x03, or 0x04", typeCode)
	}
	err = addr.Unmarshal(append(typeCode, raw[0:16]...))
	return
}

// MetadataAddressFromHex creates a MetadataAddress from a hex string.  NOTE: Does not perform validation on address,
// only performs basic HEX decoding checks.  This method matches the sdk.AccAddress approach
func MetadataAddressFromHex(address string) (MetadataAddress, error) {
	if len(address) == 0 {
		return MetadataAddress{}, errors.New("address decode failed: must provide an address")
	}
	bz, err := hex.DecodeString(address)
	return MetadataAddress(bz), err
}

// MetadataAddressFromBech32 creates a MetadataAddress from a Bech32 string.  The encoded data is checked against the
// provided bech32 hrp along with an overall verification of the byte format.
func MetadataAddressFromBech32(address string) (addr MetadataAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return MetadataAddress{}, errors.New("empty address string is not allowed")
	}

	hrp, bz, err := bech32.DecodeAndConvert(address)
	if err != nil {
		return nil, err
	}
	expectedHrp, err := VerifyMetadataAddressFormat(bz)
	if err != nil {
		return nil, err
	}
	if expectedHrp != hrp {
		return MetadataAddress{}, fmt.Errorf("invalid bech32 prefix; expected %s, got %s", expectedHrp, hrp)
	}

	return MetadataAddress(bz), nil
}

// ScopeMetadataAddress creates a MetadataAddress instance for the given scope by its uuid
func ScopeMetadataAddress(scopeUUID uuid.UUID) MetadataAddress {
	bz, err := scopeUUID.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return append(ScopeKeyPrefix, bz...)
}

// GroupMetadataAddress creates a MetadataAddress instance for a group within a scope by uuids
func GroupMetadataAddress(scopeUUID uuid.UUID, groupUUID uuid.UUID) MetadataAddress {
	bz, err := scopeUUID.MarshalBinary()
	if err != nil {
		panic(err)
	}
	addr := append(GroupKeyPrefix, bz...)
	bz, err = groupUUID.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return append(addr, bz...)
}

// RecordMetadataAddress creates a MetadataAddress instance for a record within a scope by scope uuid/record name
func RecordMetadataAddress(scopeUUID uuid.UUID, name string) MetadataAddress {
	bz, err := scopeUUID.MarshalBinary()
	if err != nil {
		panic(err)
	}
	addr := append(RecordKeyPrefix, bz...)
	name = strings.ToLower(strings.TrimSpace(name))
	nameBytes := sha256.Sum256([]byte(name))
	return append(addr, nameBytes[:]...)
}

// ScopeSpecMetadataAddress creates a MetadataAddress instance for a scope specification
func ScopeSpecMetadataAddress(specUUID uuid.UUID) MetadataAddress {
	bz, err := specUUID.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return append(ScopeSpecificationPrefix, bz...)
}

// GroupSpecMetadataAddress creates a MetadataAddress instance for a group specification
func GroupSpecMetadataAddress(specUUID uuid.UUID) MetadataAddress {
	bz, err := specUUID.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return append(GroupSpecificationPrefix, bz...)
}

// Equals determines if the current MetadataAddress is equal to another sdk.Address
func (ma MetadataAddress) Equals(ma2 sdk.Address) bool {
	if ma.Empty() && ma2.Empty() {
		return true
	}
	return bytes.Equal(ma.Bytes(), ma2.Bytes())
}

// Empty returns true if the MetadataAddress is uninitialized
func (ma MetadataAddress) Empty() bool {
	if ma == nil {
		return true
	}

	ma2 := MetadataAddress{}
	return bytes.Equal(ma.Bytes(), ma2.Bytes())
}

// Validate determines if the contained bytes form a valid MetadataAddress according to its type
func (ma MetadataAddress) Validate() (err error) {
	_, err = VerifyMetadataAddressFormat(ma)
	return
}

// Marshal returns the bytes underlying the MetadataAddress instance
func (ma MetadataAddress) Marshal() ([]byte, error) {
	return ma, nil
}

// Unmarshal initializes a MetadataAddress instance using the given bytes.  An error will be returned if the
// given bytes do not form a valid Address
func (ma *MetadataAddress) Unmarshal(data []byte) error {
	*ma = data
	if len(data) == 0 {
		return nil
	}
	_, err := VerifyMetadataAddressFormat(data)
	return err
}

// MarshalJSON returns a JSON representation for the current address using a bech32 encoded string
func (ma MetadataAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(ma.String())
}

// MarshalYAML returns a YAML representation for the current address using a bech32 encoded string
func (ma MetadataAddress) MarshalYAML() (interface{}, error) {
	return ma.String(), nil
}

// UnmarshalJSON creates a MetadataAddress instance from the given JSON data
func (ma *MetadataAddress) UnmarshalJSON(data []byte) error {
	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*ma = MetadataAddress{}
		return nil
	}

	ma2, err := MetadataAddressFromBech32(s)
	if err != nil {
		return err
	}

	*ma = ma2
	return nil
}

// UnmarshalYAML creates a MetadataAddress instance from the given YAML data
func (ma *MetadataAddress) UnmarshalYAML(data []byte) error {
	var s string
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*ma = MetadataAddress{}
		return nil
	}

	ma2, err := MetadataAddressFromBech32(s)
	if err != nil {
		return err
	}

	*ma = ma2
	return nil
}

// Bytes implements Address interface, returns the raw bytes for this Address
func (ma MetadataAddress) Bytes() []byte {
	return ma
}

// String implements the stringer interface and encodes as a bech32
func (ma MetadataAddress) String() string {
	if ma.Empty() {
		return ""
	}

	hrp, err := VerifyMetadataAddressFormat(ma)
	if err != nil {
		panic(err)
	}

	bech32Addr, err := bech32.ConvertAndEncode(hrp, ma.Bytes())
	if err != nil {
		panic(err)
	}

	return bech32Addr
}

// Size implements gogoproto custom type interface and returns the number of bytes in this instance
func (ma MetadataAddress) Size() int {
	return len(ma)
}

// MarshallTo implements gogoproto custom type interface and writes the current bytes into the provided data structure
func (ma *MetadataAddress) MarshalTo(data []byte) (int, error) {
	if len(*ma) == 0 {
		return 0, nil
	}
	copy(data, *ma)
	return len(*ma), nil
}

// Compare exists to fit gogoprotobuf custom type interface.
func (ma MetadataAddress) Compare(other MetadataAddress) int {
	return bytes.Compare(ma[0:], other[0:])
}

// ScopeUUID returns the scope uuid component of a MetadataAddress (if appropriate)
func (ma MetadataAddress) ScopeUUID() (uuid.UUID, error) {
	if len(ma) < 1 {
		return uuid.UUID{}, fmt.Errorf("address empty")
	}
	// if we don't know this type
	if ma[0] > 0x04 {
		return uuid.UUID{}, fmt.Errorf("invalid address type out of valid range (got: %d)", ma[0])
	}
	// Scope, RecordGroup, Record all have a ScopeUUID
	if ma[0] == ScopeKeyPrefix[0] || ma[0] == GroupKeyPrefix[0] || ma[0] == RecordKeyPrefix[0] {
		return uuid.FromBytes(ma[1:17])
	}
	// otherwise there isn't a scope uuid so return empty.
	return uuid.UUID{}, fmt.Errorf("this metadata address does not contain a scope uuid")
}

// GroupUUID returns the group uuid component of a MetadataAddress (if appropriate)
func (ma MetadataAddress) GroupUUID() (uuid.UUID, error) {
	if len(ma) < 1 {
		return uuid.UUID{}, fmt.Errorf("address empty")
	}
	// if we don't know this type
	if ma[0] > 0x04 {
		return uuid.UUID{}, fmt.Errorf("invalid address type out of valid range (got: %d)", ma[0])
	}
	if ma[0] == GroupKeyPrefix[0] {
		return uuid.FromBytes(ma[17:33])
	}
	// otherwise there isn't a group uuid so return empty.
	return uuid.UUID{}, fmt.Errorf("this metadata address does not contain a group uuid")
}

// GetRecordAddress returns the MetadataAddress for a record with the given name within the current scope context
// panics if the current context is not associated with a scope.
func (ma MetadataAddress) GetRecordAddress(name string) MetadataAddress {
	scopeUUID, err := ma.ScopeUUID()
	if err != nil {
		panic(err)
	}
	return RecordMetadataAddress(scopeUUID, name)
}

// ScopeGroupIteratorPrefix returns an iterator prefix that finds all Groups assigned to the metadata address scope
// if the current address is empty then returns a prefix to iterate through all groups
func (ma MetadataAddress) ScopeGroupIteratorPrefix() ([]byte, error) {
	if len(ma) < 1 {
		return GroupKeyPrefix, nil
	}
	// if we don't know this type
	if ma[0] > 0x04 {
		return []byte{}, fmt.Errorf("invalid address type out of valid range (got: %d)", ma[0])
	}
	// Scope, RecordGroup, Record all have a ScopeUUID
	if ma[0] == ScopeKeyPrefix[0] || ma[0] == GroupKeyPrefix[0] || ma[0] == RecordKeyPrefix[0] {
		return append(GroupKeyPrefix, ma[1:17]...), nil
	}
	// otherwise there isn't a scope uuid so return empty.
	return []byte{}, fmt.Errorf("this metadata address does not contain a scope uuid")
}

// ScopeRecordIteratorPrefix returns an iterator prefix that finds all Records assigned to the metadata address scope
// if the current address is empty the returns a prefix to iterate through all records
func (ma MetadataAddress) ScopeRecordIteratorPrefix() ([]byte, error) {
	if len(ma) < 1 {
		return RecordKeyPrefix, nil
	}
	// if we don't know this type
	if ma[0] > 0x04 {
		return []byte{}, fmt.Errorf("invalid address type out of valid range (got: %d)", ma[0])
	}
	// Scope, RecordGroup, Record all have a ScopeUUID
	if ma[0] == ScopeKeyPrefix[0] || ma[0] == GroupKeyPrefix[0] || ma[0] == RecordKeyPrefix[0] {
		return append(RecordKeyPrefix, ma[1:17]...), nil
	}
	// otherwise there isn't a scope uuid so return empty.
	return []byte{}, fmt.Errorf("this metadata address does not contain a scope uuid")
}

// Format implements fmt.Format interface
func (ma MetadataAddress) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(ma.String()))
	case 'p':
		s.Write([]byte(fmt.Sprintf("%p", ma)))
	default:
		s.Write([]byte(fmt.Sprintf("%X", []byte(ma))))
	}
}

// IsScopeAddress returns true is the address is valid and matches this type
func (ma MetadataAddress) IsScopeAddress() bool {
	hrp, err := VerifyMetadataAddressFormat(ma)
	return (err == nil && hrp == PrefixScope)
}

// IsGroupAddress returns true is the address is valid and matches this type
func (ma MetadataAddress) IsGroupAddress() bool {
	hrp, err := VerifyMetadataAddressFormat(ma)
	return (err == nil && hrp == PrefixGroup)
}

// IsRecordAddress returns true is the address is valid and matches this type
func (ma MetadataAddress) IsRecordAddress() bool {
	hrp, err := VerifyMetadataAddressFormat(ma)
	return (err == nil && hrp == PrefixRecord)
}

// IsScopeSpecificationAddress returns true is the address is valid and matches this type
func (ma MetadataAddress) IsScopeSpecificationAddress() bool {
	hrp, err := VerifyMetadataAddressFormat(ma)
	return (err == nil && hrp == PrefixScopeSpecification)
}

// IsGroupSpecificationAddress returns true is the address is valid and matches this type
func (ma MetadataAddress) IsGroupSpecificationAddress() bool {
	hrp, err := VerifyMetadataAddressFormat(ma)
	return (err == nil && hrp == PrefixGroupSpecification)
}