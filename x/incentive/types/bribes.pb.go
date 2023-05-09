// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hhand/incentive/bribes.proto

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

type Bribes struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Proposer string `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Block    uint64 `protobuf:"varint,4,opt,name=block,proto3" json:"block,omitempty"`
	Denom    string `protobuf:"bytes,5,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount   uint64 `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Chain    uint64 `protobuf:"varint,7,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *Bribes) Reset()         { *m = Bribes{} }
func (m *Bribes) String() string { return proto.CompactTextString(m) }
func (*Bribes) ProtoMessage()    {}
func (*Bribes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc92c88a36a3661d, []int{0}
}
func (m *Bribes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bribes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bribes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bribes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bribes.Merge(m, src)
}
func (m *Bribes) XXX_Size() int {
	return m.Size()
}
func (m *Bribes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bribes.DiscardUnknown(m)
}

var xxx_messageInfo_Bribes proto.InternalMessageInfo

func (m *Bribes) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Bribes) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *Bribes) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Bribes) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Bribes) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Bribes) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Bribes) GetChain() uint64 {
	if m != nil {
		return m.Chain
	}
	return 0
}

func init() {
	proto.RegisterType((*Bribes)(nil), "hhand.incentive.Bribes")
}

func init() { proto.RegisterFile("hhand/incentive/bribes.proto", fileDescriptor_fc92c88a36a3661d) }

var fileDescriptor_fc92c88a36a3661d = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xc8, 0x48, 0xcc,
	0x4b, 0xd1, 0xcf, 0xcc, 0x4b, 0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0xd5, 0x4f, 0x2a, 0xca, 0x4c,
	0x4a, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x07, 0xcb, 0xea, 0xc1, 0x65, 0x95,
	0x56, 0x31, 0x72, 0xb1, 0x39, 0x81, 0x55, 0x08, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5, 0xa4, 0x56,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x52, 0x5c, 0x1c, 0x05, 0x45, 0xf9,
	0x05, 0xf9, 0xc5, 0xa9, 0x45, 0x12, 0x4c, 0x60, 0x09, 0x38, 0x1f, 0xa4, 0xa3, 0x24, 0xb3, 0x24,
	0x27, 0x55, 0x82, 0x19, 0xa2, 0x03, 0xcc, 0x01, 0x89, 0x26, 0xe5, 0xe4, 0x27, 0x67, 0x4b, 0xb0,
	0x28, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x38, 0x20, 0xd1, 0x94, 0xd4, 0xbc, 0xfc, 0x5c, 0x09, 0x56,
	0x88, 0x5a, 0x30, 0x47, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x0d,
	0xac, 0x18, 0xca, 0x03, 0xa9, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0x93, 0x60, 0x87, 0x98, 0x01, 0xe6,
	0x38, 0x19, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x38, 0xc4, 0xd7,
	0x15, 0x48, 0xfe, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xdb, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x2d, 0x67, 0xfb, 0x2e, 0x17, 0x01, 0x00, 0x00,
}

func (m *Bribes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bribes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bribes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Chain != 0 {
		i = encodeVarintBribes(dAtA, i, uint64(m.Chain))
		i--
		dAtA[i] = 0x38
	}
	if m.Amount != 0 {
		i = encodeVarintBribes(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintBribes(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Block != 0 {
		i = encodeVarintBribes(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintBribes(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintBribes(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBribes(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBribes(dAtA []byte, offset int, v uint64) int {
	offset -= sovBribes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bribes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBribes(uint64(l))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovBribes(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovBribes(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovBribes(uint64(m.Block))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovBribes(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovBribes(uint64(m.Amount))
	}
	if m.Chain != 0 {
		n += 1 + sovBribes(uint64(m.Chain))
	}
	return n
}

func sovBribes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBribes(x uint64) (n int) {
	return sovBribes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bribes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBribes
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
			return fmt.Errorf("proto: Bribes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bribes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBribes
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
				return ErrInvalidLengthBribes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBribes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBribes
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
				return ErrInvalidLengthBribes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBribes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBribes
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
				return ErrInvalidLengthBribes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBribes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBribes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBribes
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
				return ErrInvalidLengthBribes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBribes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBribes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			m.Chain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBribes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chain |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBribes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBribes
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
func skipBribes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBribes
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
					return 0, ErrIntOverflowBribes
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
					return 0, ErrIntOverflowBribes
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
				return 0, ErrInvalidLengthBribes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBribes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBribes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBribes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBribes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBribes = fmt.Errorf("proto: unexpected end of group")
)
