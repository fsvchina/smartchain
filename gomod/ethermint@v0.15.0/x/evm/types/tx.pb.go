


package types

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3


type MsgEthereumTx struct {

	Data *types.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`

	Size_ float64 `protobuf:"fixed64,2,opt,name=size,proto3" json:"-"`

	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty" rlp:"-"`



	From string `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
}

func (m *MsgEthereumTx) Reset()         { *m = MsgEthereumTx{} }
func (m *MsgEthereumTx) String() string { return proto.CompactTextString(m) }
func (*MsgEthereumTx) ProtoMessage()    {}
func (*MsgEthereumTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f75ac0a12d075f21, []int{0}
}
func (m *MsgEthereumTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEthereumTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEthereumTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEthereumTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEthereumTx.Merge(m, src)
}
func (m *MsgEthereumTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgEthereumTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEthereumTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEthereumTx proto.InternalMessageInfo


type LegacyTx struct {

	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`

	GasPrice *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=gas_price,json=gasPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"gas_price,omitempty"`

	GasLimit uint64 `protobuf:"varint,3,opt,name=gas,proto3" json:"gas,omitempty"`

	To string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`

	Amount *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"value,omitempty"`

	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`

	V []byte `protobuf:"bytes,7,opt,name=v,proto3" json:"v,omitempty"`

	R []byte `protobuf:"bytes,8,opt,name=r,proto3" json:"r,omitempty"`

	S []byte `protobuf:"bytes,9,opt,name=s,proto3" json:"s,omitempty"`
}

func (m *LegacyTx) Reset()         { *m = LegacyTx{} }
func (m *LegacyTx) String() string { return proto.CompactTextString(m) }
func (*LegacyTx) ProtoMessage()    {}
func (*LegacyTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f75ac0a12d075f21, []int{1}
}
func (m *LegacyTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyTx.Merge(m, src)
}
func (m *LegacyTx) XXX_Size() int {
	return m.Size()
}
func (m *LegacyTx) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyTx.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyTx proto.InternalMessageInfo


type AccessListTx struct {

	ChainID *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"chainID"`

	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`

	GasPrice *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=gas_price,json=gasPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"gas_price,omitempty"`

	GasLimit uint64 `protobuf:"varint,4,opt,name=gas,proto3" json:"gas,omitempty"`

	To string `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`

	Amount *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"value,omitempty"`

	Data     []byte     `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	Accesses AccessList `protobuf:"bytes,8,rep,name=accesses,proto3,castrepeated=AccessList" json:"accessList"`

	V []byte `protobuf:"bytes,9,opt,name=v,proto3" json:"v,omitempty"`

	R []byte `protobuf:"bytes,10,opt,name=r,proto3" json:"r,omitempty"`

	S []byte `protobuf:"bytes,11,opt,name=s,proto3" json:"s,omitempty"`
}

func (m *AccessListTx) Reset()         { *m = AccessListTx{} }
func (m *AccessListTx) String() string { return proto.CompactTextString(m) }
func (*AccessListTx) ProtoMessage()    {}
func (*AccessListTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f75ac0a12d075f21, []int{2}
}
func (m *AccessListTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessListTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessListTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessListTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessListTx.Merge(m, src)
}
func (m *AccessListTx) XXX_Size() int {
	return m.Size()
}
func (m *AccessListTx) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessListTx.DiscardUnknown(m)
}

var xxx_messageInfo_AccessListTx proto.InternalMessageInfo


type DynamicFeeTx struct {

	ChainID *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"chainID"`

	Nonce uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`

	GasTipCap *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=gas_tip_cap,json=gasTipCap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"gas_tip_cap,omitempty"`

	GasFeeCap *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=gas_fee_cap,json=gasFeeCap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"gas_fee_cap,omitempty"`

	GasLimit uint64 `protobuf:"varint,5,opt,name=gas,proto3" json:"gas,omitempty"`

	To string `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`

	Amount *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"value,omitempty"`

	Data     []byte     `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	Accesses AccessList `protobuf:"bytes,9,rep,name=accesses,proto3,castrepeated=AccessList" json:"accessList"`

	V []byte `protobuf:"bytes,10,opt,name=v,proto3" json:"v,omitempty"`

	R []byte `protobuf:"bytes,11,opt,name=r,proto3" json:"r,omitempty"`

	S []byte `protobuf:"bytes,12,opt,name=s,proto3" json:"s,omitempty"`
}

func (m *DynamicFeeTx) Reset()         { *m = DynamicFeeTx{} }
func (m *DynamicFeeTx) String() string { return proto.CompactTextString(m) }
func (*DynamicFeeTx) ProtoMessage()    {}
func (*DynamicFeeTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f75ac0a12d075f21, []int{3}
}
func (m *DynamicFeeTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DynamicFeeTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DynamicFeeTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DynamicFeeTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicFeeTx.Merge(m, src)
}
func (m *DynamicFeeTx) XXX_Size() int {
	return m.Size()
}
func (m *DynamicFeeTx) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicFeeTx.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicFeeTx proto.InternalMessageInfo

type ExtensionOptionsEthereumTx struct {
}

func (m *ExtensionOptionsEthereumTx) Reset()         { *m = ExtensionOptionsEthereumTx{} }
func (m *ExtensionOptionsEthereumTx) String() string { return proto.CompactTextString(m) }
func (*ExtensionOptionsEthereumTx) ProtoMessage()    {}
func (*ExtensionOptionsEthereumTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f75ac0a12d075f21, []int{4}
}
func (m *ExtensionOptionsEthereumTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtensionOptionsEthereumTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtensionOptionsEthereumTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtensionOptionsEthereumTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionOptionsEthereumTx.Merge(m, src)
}
func (m *ExtensionOptionsEthereumTx) XXX_Size() int {
	return m.Size()
}
func (m *ExtensionOptionsEthereumTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionOptionsEthereumTx.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionOptionsEthereumTx proto.InternalMessageInfo


type MsgEthereumTxResponse struct {



	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`


	Logs []*Log `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`


	Ret []byte `protobuf:"bytes,3,opt,name=ret,proto3" json:"ret,omitempty"`

	VmError string `protobuf:"bytes,4,opt,name=vm_error,json=vmError,proto3" json:"vm_error,omitempty"`

	GasUsed uint64 `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
}

func (m *MsgEthereumTxResponse) Reset()         { *m = MsgEthereumTxResponse{} }
func (m *MsgEthereumTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEthereumTxResponse) ProtoMessage()    {}
func (*MsgEthereumTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f75ac0a12d075f21, []int{5}
}
func (m *MsgEthereumTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEthereumTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEthereumTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEthereumTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEthereumTxResponse.Merge(m, src)
}
func (m *MsgEthereumTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEthereumTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEthereumTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEthereumTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgEthereumTx)(nil), "ethermint.evm.v1.MsgEthereumTx")
	proto.RegisterType((*LegacyTx)(nil), "ethermint.evm.v1.LegacyTx")
	proto.RegisterType((*AccessListTx)(nil), "ethermint.evm.v1.AccessListTx")
	proto.RegisterType((*DynamicFeeTx)(nil), "ethermint.evm.v1.DynamicFeeTx")
	proto.RegisterType((*ExtensionOptionsEthereumTx)(nil), "ethermint.evm.v1.ExtensionOptionsEthereumTx")
	proto.RegisterType((*MsgEthereumTxResponse)(nil), "ethermint.evm.v1.MsgEthereumTxResponse")
}

func init() { proto.RegisterFile("ethermint/evm/v1/tx.proto", fileDescriptor_f75ac0a12d075f21) }

var fileDescriptor_f75ac0a12d075f21 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0x31, 0x8f, 0xe3, 0x44,
	0x14, 0xce, 0x24, 0x4e, 0xec, 0x4c, 0xc2, 0xe9, 0x64, 0xed, 0x49, 0x4e, 0xc4, 0xc5, 0x91, 0x25,
	0x20, 0x20, 0xc5, 0xd6, 0x2d, 0x54, 0x5b, 0xb1, 0xbe, 0xdd, 0x3b, 0xdd, 0x29, 0x27, 0x90, 0x15,
	0x1a, 0xae, 0x88, 0x66, 0x9d, 0x59, 0x67, 0x44, 0xec, 0xb1, 0x3c, 0x13, 0xcb, 0x41, 0xa2, 0x41,
	0x14, 0x74, 0x20, 0xf1, 0x07, 0x28, 0xa8, 0x68, 0xe1, 0x07, 0x50, 0x5e, 0x79, 0x82, 0x06, 0x51,
	0x18, 0x94, 0xa5, 0xda, 0x0e, 0x7e, 0x01, 0x9a, 0xb1, 0xb3, 0x9b, 0x10, 0xe5, 0x80, 0xe3, 0x10,
	0x55, 0xe6, 0xf9, 0x7b, 0xf3, 0xe6, 0xbd, 0xf7, 0x7d, 0x79, 0x0f, 0x76, 0x30, 0x9f, 0xe1, 0x24,
	0x24, 0x11, 0x77, 0x70, 0x1a, 0x3a, 0xe9, 0x1d, 0x87, 0x67, 0x76, 0x9c, 0x50, 0x4e, 0xf5, 0x9b,
	0x57, 0x90, 0x8d, 0xd3, 0xd0, 0x4e, 0xef, 0x74, 0x0f, 0x02, 0x1a, 0x50, 0x09, 0x3a, 0xe2, 0x54,
	0xf8, 0x75, 0x5f, 0x0e, 0x28, 0x0d, 0xe6, 0xd8, 0x41, 0x31, 0x71, 0x50, 0x14, 0x51, 0x8e, 0x38,
	0xa1, 0x11, 0x2b, 0xd1, 0x4e, 0x89, 0x4a, 0xeb, 0x6c, 0x71, 0xee, 0xa0, 0x68, 0xb9, 0x86, 0x7c,
	0xca, 0x42, 0xca, 0x26, 0x45, 0xc4, 0xc2, 0x28, 0xa1, 0xee, 0x4e, 0x5a, 0x22, 0x05, 0x89, 0x59,
	0x9f, 0x01, 0xf8, 0xd2, 0x23, 0x16, 0x9c, 0x0a, 0x0f, 0xbc, 0x08, 0xc7, 0x99, 0x3e, 0x80, 0xca,
	0x14, 0x71, 0x64, 0x80, 0x3e, 0x18, 0xb4, 0x0e, 0x0f, 0xec, 0xe2, 0x49, 0x7b, 0xfd, 0xa4, 0x7d,
	0x1c, 0x2d, 0x3d, 0xe9, 0xa1, 0x77, 0xa0, 0xc2, 0xc8, 0x87, 0xd8, 0xa8, 0xf6, 0xc1, 0x00, 0xb8,
	0xf5, 0xcb, 0xdc, 0x04, 0x43, 0x4f, 0x7e, 0xd2, 0x4d, 0xa8, 0xcc, 0x10, 0x9b, 0x19, 0xb5, 0x3e,
	0x18, 0x34, 0xdd, 0xd6, 0xef, 0xb9, 0xa9, 0x26, 0xf3, 0xf8, 0xc8, 0x1a, 0x5a, 0x9e, 0x04, 0x74,
	0x1d, 0x2a, 0xe7, 0x09, 0x0d, 0x0d, 0x45, 0x38, 0x78, 0xf2, 0x7c, 0xa4, 0x7c, 0xfa, 0xa5, 0x59,
	0xb1, 0xbe, 0xa9, 0x42, 0x6d, 0x84, 0x03, 0xe4, 0x2f, 0xc7, 0x99, 0x7e, 0x00, 0xeb, 0x11, 0x8d,
	0x7c, 0x2c, 0xb3, 0x51, 0xbc, 0xc2, 0xd0, 0xef, 0xc3, 0x66, 0x80, 0x44, 0xa9, 0xc4, 0x2f, 0x5e,
	0x6f, 0xba, 0x6f, 0xfc, 0x94, 0x9b, 0xaf, 0x06, 0x84, 0xcf, 0x16, 0x67, 0xb6, 0x4f, 0xc3, 0xb2,
	0x01, 0xe5, 0xcf, 0x90, 0x4d, 0x3f, 0x70, 0xf8, 0x32, 0xc6, 0xcc, 0x7e, 0x10, 0x71, 0x4f, 0x0b,
	0x10, 0x7b, 0x57, 0xdc, 0xd5, 0x7b, 0xb0, 0x16, 0x20, 0x26, 0xb3, 0x54, 0xdc, 0xf6, 0x2a, 0x37,
	0xb5, 0xfb, 0x88, 0x8d, 0x48, 0x48, 0xb8, 0x27, 0x00, 0xfd, 0x06, 0xac, 0x72, 0x5a, 0xe6, 0x58,
	0xe5, 0x54, 0x7f, 0x08, 0xeb, 0x29, 0x9a, 0x2f, 0xb0, 0x51, 0x97, 0x8f, 0xbe, 0xf5, 0xf7, 0x1f,
	0x5d, 0xe5, 0x66, 0xe3, 0x38, 0xa4, 0x8b, 0x88, 0x7b, 0x45, 0x08, 0xd1, 0x01, 0xd9, 0xe7, 0x46,
	0x1f, 0x0c, 0xda, 0x65, 0x47, 0xdb, 0x10, 0xa4, 0x86, 0x2a, 0x3f, 0x80, 0x54, 0x58, 0x89, 0xa1,
	0x15, 0x56, 0x22, 0x2c, 0x66, 0x34, 0x0b, 0x8b, 0x1d, 0xdd, 0x10, 0xbd, 0xfa, 0xfe, 0xdb, 0x61,
	0x63, 0x9c, 0x9d, 0x20, 0x8e, 0xac, 0xdf, 0x6a, 0xb0, 0x7d, 0xec, 0xfb, 0x98, 0xb1, 0x11, 0x61,
	0x7c, 0x9c, 0xe9, 0x8f, 0xa1, 0xe6, 0xcf, 0x10, 0x89, 0x26, 0x64, 0x2a, 0x9b, 0xd7, 0x74, 0xdf,
	0xfe, 0x47, 0xd9, 0xaa, 0x77, 0xc5, 0xed, 0x07, 0x27, 0x97, 0xb9, 0xa9, 0xfa, 0xc5, 0xd1, 0x2b,
	0x0f, 0xd3, 0x6b, 0x5a, 0xaa, 0x7b, 0x69, 0xa9, 0xfd, 0x7b, 0x5a, 0x94, 0x67, 0xd3, 0x52, 0xdf,
	0xa5, 0xa5, 0xf1, 0xe2, 0x68, 0x51, 0x37, 0x68, 0x79, 0x0c, 0x35, 0x24, 0x7b, 0x8b, 0x99, 0xa1,
	0xf5, 0x6b, 0x83, 0xd6, 0xe1, 0x6d, 0xfb, 0xcf, 0xff, 0x67, 0xbb, 0xe8, 0xfe, 0x78, 0x11, 0xcf,
	0xb1, 0xdb, 0x7f, 0x92, 0x9b, 0x95, 0xcb, 0xdc, 0x84, 0xe8, 0x8a, 0x92, 0xaf, 0x7f, 0x36, 0xe1,
	0x35, 0x41, 0xde, 0x55, 0xc0, 0x82, 0xf3, 0xe6, 0x16, 0xe7, 0x70, 0x8b, 0xf3, 0xd6, 0x3e, 0xce,
	0xbf, 0x53, 0x60, 0xfb, 0x64, 0x19, 0xa1, 0x90, 0xf8, 0xf7, 0x30, 0xfe, 0x7f, 0x38, 0x7f, 0x08,
	0x5b, 0x82, 0x73, 0x4e, 0xe2, 0x89, 0x8f, 0xe2, 0xe7, 0x60, 0x5d, 0x48, 0x66, 0x4c, 0xe2, 0xbb,
	0x28, 0x5e, 0xc7, 0x3a, 0xc7, 0x58, 0xc6, 0x52, 0x9e, 0x2b, 0xd6, 0x3d, 0x8c, 0x45, 0xac, 0x52,
	0x42, 0xf5, 0x67, 0x4b, 0xa8, 0xb1, 0x2b, 0x21, 0xf5, 0xc5, 0x49, 0x48, 0xdb, 0x23, 0xa1, 0xe6,
	0x7f, 0x22, 0x21, 0xb8, 0x25, 0xa1, 0xd6, 0x96, 0x84, 0xda, 0xfb, 0x24, 0x64, 0xc1, 0xee, 0x69,
	0xc6, 0x71, 0xc4, 0x08, 0x8d, 0xde, 0x89, 0xe5, 0xaa, 0xb9, 0x5e, 0x05, 0xe5, 0x40, 0xfe, 0x0a,
	0xc0, 0x5b, 0x5b, 0x2b, 0xc2, 0xc3, 0x2c, 0xa6, 0x11, 0x93, 0x85, 0xca, 0x29, 0x0f, 0x8a, 0x21,
	0x2e, 0x07, 0xfb, 0xeb, 0x50, 0x99, 0xd3, 0x80, 0x19, 0x55, 0x59, 0xe4, 0xad, 0xdd, 0x22, 0x47,
	0x34, 0xf0, 0xa4, 0x8b, 0x7e, 0x13, 0xd6, 0x12, 0xcc, 0xa5, 0x66, 0xda, 0x9e, 0x38, 0xea, 0x1d,
	0xa8, 0xa5, 0xe1, 0x04, 0x27, 0x09, 0x4d, 0xca, 0xa9, 0xab, 0xa6, 0xe1, 0xa9, 0x30, 0x05, 0x24,
	0xc4, 0xb1, 0x60, 0x78, 0x5a, 0xb0, 0xea, 0xa9, 0x01, 0x62, 0xef, 0x31, 0x3c, 0x2d, 0xd2, 0x3c,
	0xfc, 0x04, 0xc0, 0xda, 0x23, 0x16, 0xe8, 0x1f, 0x41, 0xb8, 0xb1, 0xcd, 0xcc, 0xdd, 0x04, 0xb6,
	0x6a, 0xe9, 0xbe, 0xf6, 0x17, 0x0e, 0xeb, 0x62, 0xad, 0x57, 0x3e, 0xfe, 0xe1, 0xd7, 0x2f, 0xaa,
	0xa6, 0x75, 0xdb, 0xd9, 0x5d, 0xa7, 0xa5, 0xf7, 0x84, 0x67, 0xae, 0xfb, 0x64, 0xd5, 0x03, 0x4f,
	0x57, 0x3d, 0xf0, 0xcb, 0xaa, 0x07, 0x3e, 0xbf, 0xe8, 0x55, 0x9e, 0x5e, 0xf4, 0x2a, 0x3f, 0x5e,
	0xf4, 0x2a, 0xef, 0x0f, 0x36, 0xf4, 0xc4, 0x67, 0x28, 0x61, 0x84, 0x6d, 0x84, 0xca, 0x64, 0x30,
	0xa9, 0xaa, 0xb3, 0x86, 0x5c, 0xb6, 0x6f, 0xfe, 0x11, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x54, 0xa7,
	0x4c, 0x50, 0x08, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type MsgClient interface {

	EthereumTx(ctx context.Context, in *MsgEthereumTx, opts ...grpc.CallOption) (*MsgEthereumTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) EthereumTx(ctx context.Context, in *MsgEthereumTx, opts ...grpc.CallOption) (*MsgEthereumTxResponse, error) {
	out := new(MsgEthereumTxResponse)
	err := c.cc.Invoke(ctx, "/ethermint.evm.v1.Msg/EthereumTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type MsgServer interface {

	EthereumTx(context.Context, *MsgEthereumTx) (*MsgEthereumTxResponse, error)
}


type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) EthereumTx(ctx context.Context, req *MsgEthereumTx) (*MsgEthereumTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthereumTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_EthereumTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEthereumTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EthereumTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethermint.evm.v1.Msg/EthereumTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EthereumTx(ctx, req.(*MsgEthereumTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethermint.evm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EthereumTx",
			Handler:    _Msg_EthereumTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethermint/evm/v1/tx.proto",
}

func (m *MsgEthereumTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEthereumTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEthereumTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Size_ != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Size_))))
		i--
		dAtA[i] = 0x11
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LegacyTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.S) > 0 {
		i -= len(m.S)
		copy(dAtA[i:], m.S)
		i = encodeVarintTx(dAtA, i, uint64(len(m.S)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.R) > 0 {
		i -= len(m.R)
		copy(dAtA[i:], m.R)
		i = encodeVarintTx(dAtA, i, uint64(len(m.R)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.V) > 0 {
		i -= len(m.V)
		copy(dAtA[i:], m.V)
		i = encodeVarintTx(dAtA, i, uint64(len(m.V)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if m.Amount != nil {
		{
			size := m.Amount.Size()
			i -= size
			if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x22
	}
	if m.GasLimit != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.GasPrice != nil {
		{
			size := m.GasPrice.Size()
			i -= size
			if _, err := m.GasPrice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Nonce != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AccessListTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessListTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessListTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.S) > 0 {
		i -= len(m.S)
		copy(dAtA[i:], m.S)
		i = encodeVarintTx(dAtA, i, uint64(len(m.S)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.R) > 0 {
		i -= len(m.R)
		copy(dAtA[i:], m.R)
		i = encodeVarintTx(dAtA, i, uint64(len(m.R)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.V) > 0 {
		i -= len(m.V)
		copy(dAtA[i:], m.V)
		i = encodeVarintTx(dAtA, i, uint64(len(m.V)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Accesses) > 0 {
		for iNdEx := len(m.Accesses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accesses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Amount != nil {
		{
			size := m.Amount.Size()
			i -= size
			if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x2a
	}
	if m.GasLimit != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.GasPrice != nil {
		{
			size := m.GasPrice.Size()
			i -= size
			if _, err := m.GasPrice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Nonce != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainID != nil {
		{
			size := m.ChainID.Size()
			i -= size
			if _, err := m.ChainID.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DynamicFeeTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DynamicFeeTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DynamicFeeTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.S) > 0 {
		i -= len(m.S)
		copy(dAtA[i:], m.S)
		i = encodeVarintTx(dAtA, i, uint64(len(m.S)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.R) > 0 {
		i -= len(m.R)
		copy(dAtA[i:], m.R)
		i = encodeVarintTx(dAtA, i, uint64(len(m.R)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.V) > 0 {
		i -= len(m.V)
		copy(dAtA[i:], m.V)
		i = encodeVarintTx(dAtA, i, uint64(len(m.V)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Accesses) > 0 {
		for iNdEx := len(m.Accesses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accesses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x42
	}
	if m.Amount != nil {
		{
			size := m.Amount.Size()
			i -= size
			if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintTx(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x32
	}
	if m.GasLimit != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x28
	}
	if m.GasFeeCap != nil {
		{
			size := m.GasFeeCap.Size()
			i -= size
			if _, err := m.GasFeeCap.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.GasTipCap != nil {
		{
			size := m.GasTipCap.Size()
			i -= size
			if _, err := m.GasTipCap.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Nonce != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainID != nil {
		{
			size := m.ChainID.Size()
			i -= size
			if _, err := m.ChainID.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExtensionOptionsEthereumTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionOptionsEthereumTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtensionOptionsEthereumTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgEthereumTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEthereumTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEthereumTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasUsed != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x28
	}
	if len(m.VmError) > 0 {
		i -= len(m.VmError)
		copy(dAtA[i:], m.VmError)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VmError)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Ret) > 0 {
		i -= len(m.Ret)
		copy(dAtA[i:], m.Ret)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ret)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgEthereumTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Size_ != 0 {
		n += 9
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *LegacyTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovTx(uint64(m.Nonce))
	}
	if m.GasPrice != nil {
		l = m.GasPrice.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovTx(uint64(m.GasLimit))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.V)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.R)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.S)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *AccessListTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainID != nil {
		l = m.ChainID.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovTx(uint64(m.Nonce))
	}
	if m.GasPrice != nil {
		l = m.GasPrice.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovTx(uint64(m.GasLimit))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Accesses) > 0 {
		for _, e := range m.Accesses {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.V)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.R)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.S)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *DynamicFeeTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainID != nil {
		l = m.ChainID.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovTx(uint64(m.Nonce))
	}
	if m.GasTipCap != nil {
		l = m.GasTipCap.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GasFeeCap != nil {
		l = m.GasFeeCap.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GasLimit != 0 {
		n += 1 + sovTx(uint64(m.GasLimit))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Accesses) > 0 {
		for _, e := range m.Accesses {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.V)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.R)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.S)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *ExtensionOptionsEthereumTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgEthereumTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Ret)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.VmError)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GasUsed != 0 {
		n += 1 + sovTx(uint64(m.GasUsed))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgEthereumTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEthereumTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEthereumTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Size_ = float64(math.Float64frombits(v))
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *LegacyTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: LegacyTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.GasPrice = &v
			if err := m.GasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.Amount = &v
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.V = append(m.V[:0], dAtA[iNdEx:postIndex]...)
			if m.V == nil {
				m.V = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.R = append(m.R[:0], dAtA[iNdEx:postIndex]...)
			if m.R == nil {
				m.R = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.S = append(m.S[:0], dAtA[iNdEx:postIndex]...)
			if m.S == nil {
				m.S = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *AccessListTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: AccessListTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessListTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.ChainID = &v
			if err := m.ChainID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.GasPrice = &v
			if err := m.GasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.Amount = &v
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accesses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accesses = append(m.Accesses, AccessTuple{})
			if err := m.Accesses[len(m.Accesses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.V = append(m.V[:0], dAtA[iNdEx:postIndex]...)
			if m.V == nil {
				m.V = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.R = append(m.R[:0], dAtA[iNdEx:postIndex]...)
			if m.R == nil {
				m.R = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.S = append(m.S[:0], dAtA[iNdEx:postIndex]...)
			if m.S == nil {
				m.S = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *DynamicFeeTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: DynamicFeeTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DynamicFeeTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.ChainID = &v
			if err := m.ChainID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasTipCap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.GasTipCap = &v
			if err := m.GasTipCap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasFeeCap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.GasFeeCap = &v
			if err := m.GasFeeCap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.Amount = &v
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accesses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accesses = append(m.Accesses, AccessTuple{})
			if err := m.Accesses[len(m.Accesses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.V = append(m.V[:0], dAtA[iNdEx:postIndex]...)
			if m.V == nil {
				m.V = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.R = append(m.R[:0], dAtA[iNdEx:postIndex]...)
			if m.R == nil {
				m.R = []byte{}
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.S = append(m.S[:0], dAtA[iNdEx:postIndex]...)
			if m.S == nil {
				m.S = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *ExtensionOptionsEthereumTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: ExtensionOptionsEthereumTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionOptionsEthereumTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgEthereumTxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEthereumTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEthereumTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ret = append(m.Ret[:0], dAtA[iNdEx:postIndex]...)
			if m.Ret == nil {
				m.Ret = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VmError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
