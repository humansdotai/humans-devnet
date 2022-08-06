// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rdnet/pubkeys.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Pubkeys struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Moniker  string `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Pubkey   string `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Issigner string `protobuf:"bytes,4,opt,name=issigner,proto3" json:"issigner,omitempty"`
	Timeat   string `protobuf:"bytes,5,opt,name=timeat,proto3" json:"timeat,omitempty"`
}

func (m *Pubkeys) Reset()         { *m = Pubkeys{} }
func (m *Pubkeys) String() string { return proto.CompactTextString(m) }
func (*Pubkeys) ProtoMessage()    {}
func (*Pubkeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_e52cc9cb3ba70215, []int{0}
}
func (m *Pubkeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pubkeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pubkeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pubkeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pubkeys.Merge(m, src)
}
func (m *Pubkeys) XXX_Size() int {
	return m.Size()
}
func (m *Pubkeys) XXX_DiscardUnknown() {
	xxx_messageInfo_Pubkeys.DiscardUnknown(m)
}

var xxx_messageInfo_Pubkeys proto.InternalMessageInfo

func (m *Pubkeys) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Pubkeys) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Pubkeys) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *Pubkeys) GetIssigner() string {
	if m != nil {
		return m.Issigner
	}
	return ""
}

func (m *Pubkeys) GetTimeat() string {
	if m != nil {
		return m.Timeat
	}
	return ""
}

func init() {
	proto.RegisterType((*Pubkeys)(nil), "humansdotai.humans.humans.Pubkeys")
}

func init() { proto.RegisterFile("humans/pubkeys.proto", fileDescriptor_e52cc9cb3ba70215) }

var fileDescriptor_e52cc9cb3ba70215 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0xc9, 0x4b,
	0x2d, 0xd1, 0x2f, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0x4a, 0x89, 0xcf,
	0x4b, 0x2d, 0xd1, 0x03, 0x2b, 0x53, 0x6a, 0x65, 0xe4, 0x62, 0x0f, 0x80, 0x28, 0x15, 0x12, 0xe1,
	0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84,
	0x24, 0xb8, 0xd8, 0x73, 0xf3, 0xf3, 0x32, 0xb3, 0x53, 0x8b, 0x24, 0x98, 0xc0, 0xe2, 0x30, 0xae,
	0x90, 0x18, 0x17, 0x1b, 0xc4, 0x16, 0x09, 0x66, 0xb0, 0x04, 0x94, 0x27, 0x24, 0xc5, 0xc5, 0x91,
	0x59, 0x5c, 0x9c, 0x99, 0x9e, 0x97, 0x5a, 0x24, 0xc1, 0x02, 0x96, 0x81, 0xf3, 0x41, 0x7a, 0x4a,
	0x32, 0x73, 0x53, 0x13, 0x4b, 0x24, 0x58, 0x21, 0x7a, 0x20, 0x3c, 0x27, 0x97, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x0f, 0x00, 0xf9, 0xc2, 0x19, 0xe4, 0x0b, 0x7d, 0x88, 0x2f, 0xf4, 0x2b, 0xf4,
	0x21, 0xde, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xd6, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0x12, 0x4e, 0x48, 0x04, 0x01, 0x00, 0x00,
}

func (m *Pubkeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pubkeys) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pubkeys) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Timeat) > 0 {
		i -= len(m.Timeat)
		copy(dAtA[i:], m.Timeat)
		i = encodeVarintPubkeys(dAtA, i, uint64(len(m.Timeat)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Issigner) > 0 {
		i -= len(m.Issigner)
		copy(dAtA[i:], m.Issigner)
		i = encodeVarintPubkeys(dAtA, i, uint64(len(m.Issigner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintPubkeys(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintPubkeys(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintPubkeys(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPubkeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovPubkeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pubkeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovPubkeys(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovPubkeys(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovPubkeys(uint64(l))
	}
	l = len(m.Issigner)
	if l > 0 {
		n += 1 + l + sovPubkeys(uint64(l))
	}
	l = len(m.Timeat)
	if l > 0 {
		n += 1 + l + sovPubkeys(uint64(l))
	}
	return n
}

func sovPubkeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPubkeys(x uint64) (n int) {
	return sovPubkeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pubkeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPubkeys
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
			return fmt.Errorf("proto: Pubkeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pubkeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkeys
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
				return ErrInvalidLengthPubkeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubkeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkeys
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
				return ErrInvalidLengthPubkeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubkeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkeys
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
				return ErrInvalidLengthPubkeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubkeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issigner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkeys
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
				return ErrInvalidLengthPubkeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubkeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issigner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeat", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkeys
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
				return ErrInvalidLengthPubkeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPubkeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timeat = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPubkeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPubkeys
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
func skipPubkeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPubkeys
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
					return 0, ErrIntOverflowPubkeys
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
					return 0, ErrIntOverflowPubkeys
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
				return 0, ErrInvalidLengthPubkeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPubkeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPubkeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPubkeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPubkeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPubkeys = fmt.Errorf("proto: unexpected end of group")
)
