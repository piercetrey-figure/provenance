// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/marker/v1/accessgrant.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Access defines the different types of permissions that a marker supports granting to an address.
type Access int32

const (
	// ACCESS_UNSPECIFIED defines a no-op vote option.
	Access_Unknown Access = 0
	// ACCESS_MINT is the ability to increase the supply of a marker
	Access_Mint Access = 1
	// ACCESS_BURN is the ability to decrease the supply of the marker using coin held by the marker.
	Access_Burn Access = 2
	// ACCESS_DEPOSIT is the ability to set a marker reference to this marker in the metadata/scopes module
	Access_Deposit Access = 3
	// ACCESS_WITHDRAW is the ability to remove marker references to this marker in from metadata/scopes or
	// transfer coin from this marker account to another account.
	Access_Withdraw Access = 4
	// ACCESS_DELETE is the ability to move a proposed, finalized or active marker into the cancelled state. This
	// access also allows cancelled markers to be marked for deletion
	Access_Delete Access = 5
	// ACCESS_ADMIN is the ability to add access grants for accounts to the list of marker permissions.
	Access_Admin Access = 6
	// ACCESS_TRANSFER is the ability to invoke a send operation using the marker module to facilitate exchange.
	// This capability is useful when the marker denomination has "send enabled = false" preventing normal bank transfer
	Access_Transfer Access = 7
)

var Access_name = map[int32]string{
	0: "ACCESS_UNSPECIFIED",
	1: "ACCESS_MINT",
	2: "ACCESS_BURN",
	3: "ACCESS_DEPOSIT",
	4: "ACCESS_WITHDRAW",
	5: "ACCESS_DELETE",
	6: "ACCESS_ADMIN",
	7: "ACCESS_TRANSFER",
}

var Access_value = map[string]int32{
	"ACCESS_UNSPECIFIED": 0,
	"ACCESS_MINT":        1,
	"ACCESS_BURN":        2,
	"ACCESS_DEPOSIT":     3,
	"ACCESS_WITHDRAW":    4,
	"ACCESS_DELETE":      5,
	"ACCESS_ADMIN":       6,
	"ACCESS_TRANSFER":    7,
}

func (x Access) String() string {
	return proto.EnumName(Access_name, int32(x))
}

func (Access) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7242c30a84644575, []int{0}
}

// AccessGrant associates a colelction of permisssions with an address for delegated marker account control.
type AccessGrant struct {
	Address     string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Permissions AccessList `protobuf:"varint,2,rep,packed,name=permissions,proto3,enum=provenance.marker.v1.Access,castrepeated=AccessList" json:"permissions,omitempty"`
}

func (m *AccessGrant) Reset()      { *m = AccessGrant{} }
func (*AccessGrant) ProtoMessage() {}
func (*AccessGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_7242c30a84644575, []int{0}
}
func (m *AccessGrant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessGrant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessGrant.Merge(m, src)
}
func (m *AccessGrant) XXX_Size() int {
	return m.Size()
}
func (m *AccessGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessGrant.DiscardUnknown(m)
}

var xxx_messageInfo_AccessGrant proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("provenance.marker.v1.Access", Access_name, Access_value)
	proto.RegisterType((*AccessGrant)(nil), "provenance.marker.v1.AccessGrant")
}

func init() {
	proto.RegisterFile("provenance/marker/v1/accessgrant.proto", fileDescriptor_7242c30a84644575)
}

var fileDescriptor_7242c30a84644575 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x31, 0x6f, 0xd3, 0x4c,
	0x18, 0xc7, 0xed, 0xb4, 0x4d, 0xda, 0x4b, 0xde, 0xbc, 0xd6, 0xa9, 0x12, 0xa9, 0x29, 0x8e, 0x01,
	0x09, 0x55, 0x48, 0xb5, 0xd5, 0xb2, 0xb1, 0x39, 0xb1, 0x0b, 0x96, 0x1a, 0x13, 0xd9, 0x8e, 0x22,
	0xb1, 0x54, 0xae, 0x73, 0xb8, 0xa7, 0xe2, 0x3b, 0xeb, 0xce, 0x4d, 0xe1, 0x0b, 0x20, 0xe4, 0x89,
	0x05, 0x89, 0xc5, 0x52, 0x66, 0x66, 0x3e, 0x04, 0x63, 0xc5, 0xc4, 0x06, 0x4a, 0x16, 0x3e, 0x06,
	0x6a, 0x2e, 0xa5, 0x1e, 0xba, 0x3d, 0xff, 0xfb, 0xff, 0xf4, 0xd3, 0xa3, 0xd3, 0x03, 0x9e, 0x64,
	0x8c, 0x4e, 0x11, 0x89, 0x48, 0x8c, 0xcc, 0x34, 0x62, 0xe7, 0x88, 0x99, 0xd3, 0x03, 0x33, 0x8a,
	0x63, 0xc4, 0x79, 0xc2, 0x22, 0x92, 0x1b, 0x19, 0xa3, 0x39, 0x85, 0xdb, 0xb7, 0x9c, 0x21, 0x38,
	0x63, 0x7a, 0xa0, 0x6e, 0x27, 0x34, 0xa1, 0x4b, 0xc0, 0xbc, 0x9e, 0x04, 0xab, 0xee, 0xc4, 0x94,
	0xa7, 0x94, 0x9f, 0x88, 0x42, 0x04, 0x51, 0x3d, 0xfa, 0x2c, 0x83, 0xa6, 0xb5, 0x94, 0xbf, 0xb8,
	0x96, 0xc3, 0x0e, 0x68, 0x44, 0x93, 0x09, 0x43, 0x9c, 0x77, 0x64, 0x5d, 0xde, 0xdb, 0xf2, 0x6f,
	0x22, 0xf4, 0x40, 0x33, 0x43, 0x2c, 0xc5, 0x9c, 0x63, 0x4a, 0x78, 0xa7, 0xa6, 0xaf, 0xed, 0xb5,
	0x0f, 0x77, 0x8d, 0xbb, 0xd6, 0x30, 0x84, 0xb1, 0xd7, 0xfe, 0xfa, 0xab, 0x0b, 0xc4, 0x7c, 0x8c,
	0x79, 0xee, 0x57, 0x05, 0xcf, 0x77, 0x3f, 0xce, 0xba, 0xd2, 0x97, 0x59, 0x57, 0xfa, 0x33, 0xeb,
	0xca, 0x3f, 0xbe, 0xed, 0xb7, 0x2a, 0x6b, 0xb8, 0x4f, 0x3f, 0xd4, 0x40, 0x5d, 0x3c, 0xc0, 0xc7,
	0x00, 0x5a, 0xfd, 0xbe, 0x13, 0x04, 0x27, 0x23, 0x2f, 0x18, 0x3a, 0x7d, 0xf7, 0xc8, 0x75, 0x6c,
	0x45, 0x52, 0x9b, 0x45, 0xa9, 0x37, 0x46, 0xe4, 0x9c, 0xd0, 0x4b, 0x02, 0x77, 0x40, 0x73, 0x05,
	0x0d, 0x5c, 0x2f, 0x54, 0x64, 0x75, 0xb3, 0x28, 0xf5, 0xf5, 0x01, 0x26, 0x79, 0xa5, 0xea, 0x8d,
	0x7c, 0x4f, 0xa9, 0x89, 0xaa, 0x77, 0xc1, 0x08, 0xec, 0x82, 0xf6, 0xaa, 0xb2, 0x9d, 0xe1, 0xab,
	0xc0, 0x0d, 0x95, 0x35, 0xa1, 0xb5, 0x51, 0x46, 0x39, 0xce, 0xe1, 0x43, 0xf0, 0xff, 0x0a, 0x18,
	0xbb, 0xe1, 0x4b, 0xdb, 0xb7, 0xc6, 0xca, 0xba, 0xda, 0x2a, 0x4a, 0x7d, 0x73, 0x8c, 0xf3, 0xb3,
	0x09, 0x8b, 0x2e, 0xe1, 0x03, 0xf0, 0xdf, 0x3f, 0xc7, 0xb1, 0x13, 0x3a, 0xca, 0x86, 0x0a, 0x8a,
	0x52, 0xaf, 0xdb, 0xe8, 0x2d, 0xca, 0x11, 0xbc, 0x0f, 0x5a, 0xab, 0xda, 0xb2, 0x07, 0xae, 0xa7,
	0xd4, 0xd5, 0xad, 0xa2, 0xd4, 0x37, 0xac, 0x49, 0x8a, 0x49, 0x45, 0x1f, 0xfa, 0x96, 0x17, 0x1c,
	0x39, 0xbe, 0xd2, 0x10, 0xfa, 0x90, 0x45, 0x84, 0xbf, 0x41, 0xac, 0x97, 0x7c, 0x9f, 0x6b, 0xf2,
	0xd5, 0x5c, 0x93, 0x7f, 0xcf, 0x35, 0xf9, 0xd3, 0x42, 0x93, 0xae, 0x16, 0x9a, 0xf4, 0x73, 0xa1,
	0x49, 0xe0, 0x1e, 0xa6, 0x77, 0xfe, 0xfe, 0x50, 0x7e, 0x7d, 0x98, 0xe0, 0xfc, 0xec, 0xe2, 0xd4,
	0x88, 0x69, 0x6a, 0xde, 0x22, 0xfb, 0x98, 0x56, 0x92, 0xf9, 0xee, 0xe6, 0xbe, 0xf2, 0xf7, 0x19,
	0xe2, 0xa7, 0xf5, 0xe5, 0x41, 0x3c, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x68, 0xb2, 0x7c, 0x78,
	0x81, 0x02, 0x00, 0x00,
}

func (this *AccessGrant) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessGrant)
	if !ok {
		that2, ok := that.(AccessGrant)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if len(this.Permissions) != len(that1.Permissions) {
		return false
	}
	for i := range this.Permissions {
		if this.Permissions[i] != that1.Permissions[i] {
			return false
		}
	}
	return true
}
func (m *AccessGrant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessGrant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessGrant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		dAtA2 := make([]byte, len(m.Permissions)*10)
		var j1 int
		for _, num := range m.Permissions {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintAccessgrant(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccessgrant(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccessgrant(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccessgrant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessGrant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccessgrant(uint64(l))
	}
	if len(m.Permissions) > 0 {
		l = 0
		for _, e := range m.Permissions {
			l += sovAccessgrant(uint64(e))
		}
		n += 1 + sovAccessgrant(uint64(l)) + l
	}
	return n
}

func sovAccessgrant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccessgrant(x uint64) (n int) {
	return sovAccessgrant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessGrant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessgrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AccessGrant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessGrant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessgrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccessgrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessgrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Access
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAccessgrant
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Access(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Permissions = append(m.Permissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAccessgrant
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthAccessgrant
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthAccessgrant
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Permissions) == 0 {
					m.Permissions = make([]Access, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Access
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAccessgrant
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Access(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Permissions = append(m.Permissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccessgrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccessgrant
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccessgrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAccessgrant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccessgrant
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccessgrant
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccessgrant
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAccessgrant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccessgrant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccessgrant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccessgrant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccessgrant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccessgrant = fmt.Errorf("proto: unexpected end of group")
)