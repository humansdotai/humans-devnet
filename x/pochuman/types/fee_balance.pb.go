// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pochuman/fee_balance.proto

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

type FeeBalance struct {
	Index     string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ChainName string `protobuf:"bytes,2,opt,name=chainName,proto3" json:"chainName,omitempty"`
	Balance   string `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Decimal   string `protobuf:"bytes,4,opt,name=decimal,proto3" json:"decimal,omitempty"`
}

func (m *FeeBalance) Reset()         { *m = FeeBalance{} }
func (m *FeeBalance) String() string { return proto.CompactTextString(m) }
func (*FeeBalance) ProtoMessage()    {}
func (*FeeBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5dff84e1c37a6d0, []int{0}
}
func (m *FeeBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeBalance.Merge(m, src)
}
func (m *FeeBalance) XXX_Size() int {
	return m.Size()
}
func (m *FeeBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeBalance.DiscardUnknown(m)
}

var xxx_messageInfo_FeeBalance proto.InternalMessageInfo

func (m *FeeBalance) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *FeeBalance) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *FeeBalance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *FeeBalance) GetDecimal() string {
	if m != nil {
		return m.Decimal
	}
	return ""
}

func init() {
	proto.RegisterType((*FeeBalance)(nil), "vigorousdeveloper.pochuman.pochuman.FeeBalance")
}

func init() { proto.RegisterFile("pochuman/fee_balance.proto", fileDescriptor_e5dff84e1c37a6d0) }

var fileDescriptor_e5dff84e1c37a6d0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xc8, 0x4f, 0xce,
	0x28, 0xcd, 0x4d, 0xcc, 0xd3, 0x4f, 0x4b, 0x4d, 0x8d, 0x4f, 0x4a, 0xcc, 0x49, 0xcc, 0x4b, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2e, 0xcb, 0x4c, 0xcf, 0x2f, 0xca, 0x2f, 0x2d,
	0x4e, 0x49, 0x2d, 0x4b, 0xcd, 0xc9, 0x2f, 0x48, 0x2d, 0xd2, 0x83, 0xa9, 0x86, 0x33, 0x94, 0x4a,
	0xb8, 0xb8, 0xdc, 0x52, 0x53, 0x9d, 0x20, 0x1a, 0x85, 0x44, 0xb8, 0x58, 0x33, 0xf3, 0x52, 0x52,
	0x2b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x19, 0x2e, 0xce, 0xe4, 0x8c,
	0xc4, 0xcc, 0x3c, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x0c, 0x42, 0x40, 0x48, 0x82, 0x8b,
	0x1d, 0x6a, 0xaf, 0x04, 0x33, 0x58, 0x0e, 0xc6, 0x05, 0xc9, 0xa4, 0xa4, 0x26, 0x67, 0xe6, 0x26,
	0xe6, 0x48, 0xb0, 0x40, 0x64, 0xa0, 0x5c, 0xa7, 0xc0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0x32, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f,
	0x83, 0xba, 0xdf, 0x05, 0xe6, 0x7e, 0xfd, 0x82, 0xfc, 0x64, 0x5d, 0x88, 0x77, 0x2b, 0xf4, 0xe1,
	0x3e, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x7b, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x20, 0x9a, 0xd8, 0xec, 0x12, 0x01, 0x00, 0x00,
}

func (m *FeeBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Decimal) > 0 {
		i -= len(m.Decimal)
		copy(dAtA[i:], m.Decimal)
		i = encodeVarintFeeBalance(dAtA, i, uint64(len(m.Decimal)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintFeeBalance(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintFeeBalance(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFeeBalance(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeeBalance(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeeBalance(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFeeBalance(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovFeeBalance(uint64(l))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovFeeBalance(uint64(l))
	}
	l = len(m.Decimal)
	if l > 0 {
		n += 1 + l + sovFeeBalance(uint64(l))
	}
	return n
}

func sovFeeBalance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeeBalance(x uint64) (n int) {
	return sovFeeBalance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeeBalance
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
			return fmt.Errorf("proto: FeeBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeBalance
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
				return ErrInvalidLengthFeeBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeeBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeBalance
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
				return ErrInvalidLengthFeeBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeeBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeBalance
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
				return ErrInvalidLengthFeeBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeeBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeBalance
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
				return ErrInvalidLengthFeeBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeeBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Decimal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeeBalance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeeBalance
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
func skipFeeBalance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeeBalance
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
					return 0, ErrIntOverflowFeeBalance
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
					return 0, ErrIntOverflowFeeBalance
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
				return 0, ErrInvalidLengthFeeBalance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeeBalance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeeBalance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeeBalance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeeBalance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeeBalance = fmt.Errorf("proto: unexpected end of group")
)
