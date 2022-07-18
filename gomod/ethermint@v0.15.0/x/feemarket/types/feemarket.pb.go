


package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3


type Params struct {

	NoBaseFee bool `protobuf:"varint,1,opt,name=no_base_fee,json=noBaseFee,proto3" json:"no_base_fee,omitempty"`


	BaseFeeChangeDenominator uint32 `protobuf:"varint,2,opt,name=base_fee_change_denominator,json=baseFeeChangeDenominator,proto3" json:"base_fee_change_denominator,omitempty"`


	ElasticityMultiplier uint32 `protobuf:"varint,3,opt,name=elasticity_multiplier,json=elasticityMultiplier,proto3" json:"elasticity_multiplier,omitempty"`

	EnableHeight int64 `protobuf:"varint,5,opt,name=enable_height,json=enableHeight,proto3" json:"enable_height,omitempty"`

	BaseFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=base_fee,json=baseFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"base_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4feb8b20cf98e6e1, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetNoBaseFee() bool {
	if m != nil {
		return m.NoBaseFee
	}
	return false
}

func (m *Params) GetBaseFeeChangeDenominator() uint32 {
	if m != nil {
		return m.BaseFeeChangeDenominator
	}
	return 0
}

func (m *Params) GetElasticityMultiplier() uint32 {
	if m != nil {
		return m.ElasticityMultiplier
	}
	return 0
}

func (m *Params) GetEnableHeight() int64 {
	if m != nil {
		return m.EnableHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "ethermint.feemarket.v1.Params")
}

func init() {
	proto.RegisterFile("ethermint/feemarket/v1/feemarket.proto", fileDescriptor_4feb8b20cf98e6e1)
}

var fileDescriptor_4feb8b20cf98e6e1 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0x87, 0x33, 0xfe, 0xbb, 0x9a, 0x7b, 0x05, 0x09, 0xde, 0x4b, 0xb8, 0x85, 0x18, 0x5a, 0x90,
	0x6c, 0x9a, 0x20, 0xae, 0xbb, 0xb1, 0xa5, 0x68, 0xa1, 0x50, 0xb2, 0xec, 0x26, 0x4c, 0xe2, 0x31,
	0x19, 0x4c, 0x66, 0x64, 0xe6, 0x28, 0xf5, 0x2d, 0xfa, 0x10, 0x7d, 0x18, 0x97, 0x2e, 0x4b, 0x17,
	0x52, 0xf4, 0x45, 0x4a, 0xa3, 0x4d, 0x5c, 0xcd, 0xcc, 0xf9, 0xbe, 0xe1, 0x1c, 0xce, 0x4f, 0xef,
	0x03, 0x26, 0x20, 0x33, 0xc6, 0xd1, 0x9b, 0x01, 0x64, 0x54, 0xce, 0x01, 0xbd, 0xd5, 0xa0, 0x7c,
	0xb8, 0x0b, 0x29, 0x50, 0x18, 0xff, 0x0a, 0xcf, 0x2d, 0xd1, 0x6a, 0xf0, 0xbf, 0x1b, 0x8b, 0x58,
	0xe4, 0x8a, 0xf7, 0x7d, 0x3b, 0xda, 0x97, 0x6f, 0x15, 0xbd, 0xf1, 0x44, 0x25, 0xcd, 0x94, 0x61,
	0xe9, 0xbf, 0xb9, 0x08, 0x42, 0xaa, 0x20, 0x98, 0x01, 0x98, 0xc4, 0x26, 0x4e, 0xd3, 0x6f, 0x71,
	0x31, 0xa2, 0x0a, 0xee, 0x01, 0x8c, 0x1b, 0xfd, 0xe2, 0x07, 0x06, 0x51, 0x42, 0x79, 0x0c, 0xc1,
	0x14, 0xb8, 0xc8, 0x18, 0xa7, 0x28, 0xa4, 0x59, 0xb1, 0x89, 0xd3, 0xf6, 0xcd, 0xf0, 0x68, 0xdf,
	0xe6, 0xc2, 0x5d, 0xc9, 0x8d, 0xa1, 0xfe, 0x17, 0x52, 0xaa, 0x90, 0x45, 0x0c, 0xd7, 0x41, 0xb6,
	0x4c, 0x91, 0x2d, 0x52, 0x06, 0xd2, 0xac, 0xe6, 0x1f, 0xbb, 0x25, 0x7c, 0x2c, 0x98, 0x71, 0xa5,
	0xb7, 0x81, 0xd3, 0x30, 0x85, 0x20, 0x01, 0x16, 0x27, 0x68, 0xd6, 0x6d, 0xe2, 0x54, 0xfd, 0x3f,
	0xc7, 0xe2, 0x38, 0xaf, 0x19, 0x13, 0xbd, 0x59, 0x4c, 0xdd, 0xb0, 0x89, 0xd3, 0x1a, 0xb9, 0x9b,
	0x5d, 0x4f, 0xfb, 0xd8, 0xf5, 0xfa, 0x31, 0xc3, 0x64, 0x19, 0xba, 0x91, 0xc8, 0xbc, 0x48, 0xa8,
	0x4c, 0xa8, 0xd3, 0x71, 0xad, 0xa6, 0x73, 0x0f, 0xd7, 0x0b, 0x50, 0xee, 0x84, 0xa3, 0xff, 0xeb,
	0x34, 0xf5, 0x43, 0xad, 0x59, 0xeb, 0xd4, 0xfd, 0x0e, 0xe3, 0x0c, 0x19, 0x4d, 0x8b, 0x65, 0x8c,
	0xc6, 0x9b, 0xbd, 0x45, 0xb6, 0x7b, 0x8b, 0x7c, 0xee, 0x2d, 0xf2, 0x7a, 0xb0, 0xb4, 0xed, 0xc1,
	0xd2, 0xde, 0x0f, 0x96, 0xf6, 0xec, 0x9e, 0xb5, 0xc0, 0x84, 0x4a, 0xc5, 0x94, 0x57, 0x26, 0xf5,
	0x72, 0x96, 0x55, 0xde, 0x2e, 0x6c, 0xe4, 0x7b, 0x1f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xee,
	0xfc, 0x5c, 0x06, 0xcf, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BaseFee.Size()
		i -= size
		if _, err := m.BaseFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFeemarket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.EnableHeight != 0 {
		i = encodeVarintFeemarket(dAtA, i, uint64(m.EnableHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.ElasticityMultiplier != 0 {
		i = encodeVarintFeemarket(dAtA, i, uint64(m.ElasticityMultiplier))
		i--
		dAtA[i] = 0x18
	}
	if m.BaseFeeChangeDenominator != 0 {
		i = encodeVarintFeemarket(dAtA, i, uint64(m.BaseFeeChangeDenominator))
		i--
		dAtA[i] = 0x10
	}
	if m.NoBaseFee {
		i--
		if m.NoBaseFee {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeemarket(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeemarket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoBaseFee {
		n += 2
	}
	if m.BaseFeeChangeDenominator != 0 {
		n += 1 + sovFeemarket(uint64(m.BaseFeeChangeDenominator))
	}
	if m.ElasticityMultiplier != 0 {
		n += 1 + sovFeemarket(uint64(m.ElasticityMultiplier))
	}
	if m.EnableHeight != 0 {
		n += 1 + sovFeemarket(uint64(m.EnableHeight))
	}
	l = m.BaseFee.Size()
	n += 1 + l + sovFeemarket(uint64(l))
	return n
}

func sovFeemarket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeemarket(x uint64) (n int) {
	return sovFeemarket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeemarket
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoBaseFee", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeemarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoBaseFee = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFeeChangeDenominator", wireType)
			}
			m.BaseFeeChangeDenominator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeemarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseFeeChangeDenominator |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ElasticityMultiplier", wireType)
			}
			m.ElasticityMultiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeemarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ElasticityMultiplier |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableHeight", wireType)
			}
			m.EnableHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeemarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EnableHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeemarket
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
				return ErrInvalidLengthFeemarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeemarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeemarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeemarket
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
func skipFeemarket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeemarket
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
					return 0, ErrIntOverflowFeemarket
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
					return 0, ErrIntOverflowFeemarket
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
				return 0, ErrInvalidLengthFeemarket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeemarket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeemarket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeemarket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeemarket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeemarket = fmt.Errorf("proto: unexpected end of group")
)
