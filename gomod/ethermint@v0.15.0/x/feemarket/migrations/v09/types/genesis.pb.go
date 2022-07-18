


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


type GenesisState struct {

	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`


	BaseFee github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=base_fee,json=baseFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"base_fee"`


	BlockGas uint64 `protobuf:"varint,3,opt,name=block_gas,json=blockGas,proto3" json:"block_gas,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6241c21661288629, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBlockGas() uint64 {
	if m != nil {
		return m.BlockGas
	}
	return 0
}

var fileDescriptor_6241c21661288629 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x73, 0x5a, 0x6a, 0x1b, 0x9d, 0x82, 0x48, 0xa9, 0x70, 0x0d, 0x22, 0x25, 0x8b, 0x77,
	0x54, 0x57, 0xa7, 0x0c, 0xd6, 0x6e, 0x12, 0x37, 0x97, 0x72, 0x89, 0xaf, 0x49, 0x88, 0xc9, 0x85,
	0x7b, 0x67, 0xd1, 0x6f, 0xe1, 0x87, 0xf1, 0x43, 0x74, 0xec, 0x28, 0x0e, 0x45, 0x92, 0x2f, 0x22,
	0xb9, 0x96, 0xb6, 0x83, 0x4e, 0x77, 0xbc, 0xf7, 0xfb, 0xbf, 0x1f, 0xfc, 0xed, 0x4b, 0xd0, 0x09,
	0xa8, 0x3c, 0x2d, 0x34, 0x9f, 0x01, 0xe4, 0x42, 0x65, 0xa0, 0xf9, 0x7c, 0xc4, 0x63, 0x28, 0x00,
	0x53, 0x64, 0xa5, 0x92, 0x5a, 0x3a, 0x67, 0x5b, 0x8a, 0x6d, 0x29, 0x36, 0x1f, 0xf5, 0x4f, 0x63,
	0x19, 0x4b, 0x83, 0xf0, 0xe6, 0xb7, 0xa6, 0xfb, 0xc3, 0x7f, 0x6e, 0xee, 0xa2, 0x86, 0xbb, 0xf8,
	0x24, 0xf6, 0xc9, 0x78, 0xed, 0x79, 0xd4, 0x42, 0x83, 0x73, 0x6b, 0xb7, 0x4b, 0xa1, 0x44, 0x8e,
	0x3d, 0xe2, 0x12, 0xef, 0xf8, 0x9a, 0xb2, 0xbf, 0xbd, 0xec, 0xc1, 0x50, 0x7e, 0x6b, 0xb1, 0x1a,
	0x58, 0xc1, 0x26, 0xe3, 0x4c, 0xec, 0x4e, 0x28, 0x10, 0xa6, 0x33, 0x80, 0xde, 0x81, 0x4b, 0xbc,
	0xae, 0xcf, 0x9a, 0xfd, 0xf7, 0x6a, 0x30, 0x8c, 0x53, 0x9d, 0xbc, 0x86, 0x2c, 0x92, 0x39, 0x8f,
	0x24, 0xe6, 0x12, 0x37, 0xcf, 0x15, 0x3e, 0x67, 0x5c, 0xbf, 0x97, 0x80, 0x6c, 0x52, 0xe8, 0xe0,
	0xa8, 0xc9, 0xdf, 0x01, 0x38, 0xe7, 0x76, 0x37, 0x7c, 0x91, 0x51, 0x36, 0x8d, 0x05, 0xf6, 0x0e,
	0x5d, 0xe2, 0xb5, 0x82, 0x8e, 0x19, 0x8c, 0x05, 0xfa, 0xf7, 0x8b, 0x8a, 0x92, 0x65, 0x45, 0xc9,
	0x4f, 0x45, 0xc9, 0x47, 0x4d, 0xad, 0x65, 0x4d, 0xad, 0xaf, 0x9a, 0x5a, 0x4f, 0x6c, 0xcf, 0xa3,
	0x13, 0xa1, 0x30, 0x45, 0xbe, 0xeb, 0xe2, 0x6d, 0xaf, 0x0d, 0xe3, 0x0c, 0xdb, 0xa6, 0x87, 0x9b,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xbf, 0x7e, 0x16, 0x85, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockGas != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.BlockGas))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.BaseFee.Size()
		i -= size
		if _, err := m.BaseFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.BaseFee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.BlockGas != 0 {
		n += 1 + sovGenesis(uint64(m.BlockGas))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockGas", wireType)
			}
			m.BlockGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
