


package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3


type Params struct {
	SendEnabled        []*SendEnabled `protobuf:"bytes,1,rep,name=send_enabled,json=sendEnabled,proto3" json:"send_enabled,omitempty" yaml:"send_enabled,omitempty"`
	DefaultSendEnabled bool           `protobuf:"varint,2,opt,name=default_send_enabled,json=defaultSendEnabled,proto3" json:"default_send_enabled,omitempty" yaml:"default_send_enabled,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd052eee12edf988, []int{0}
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

func (m *Params) GetSendEnabled() []*SendEnabled {
	if m != nil {
		return m.SendEnabled
	}
	return nil
}

func (m *Params) GetDefaultSendEnabled() bool {
	if m != nil {
		return m.DefaultSendEnabled
	}
	return false
}



type SendEnabled struct {
	Denom   string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Enabled bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *SendEnabled) Reset()      { *m = SendEnabled{} }
func (*SendEnabled) ProtoMessage() {}
func (*SendEnabled) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd052eee12edf988, []int{1}
}
func (m *SendEnabled) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendEnabled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendEnabled.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendEnabled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEnabled.Merge(m, src)
}
func (m *SendEnabled) XXX_Size() int {
	return m.Size()
}
func (m *SendEnabled) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEnabled.DiscardUnknown(m)
}

var xxx_messageInfo_SendEnabled proto.InternalMessageInfo

func (m *SendEnabled) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *SendEnabled) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}


type Input struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coins   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd052eee12edf988, []int{2}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Input.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return m.Size()
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo


type Output struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coins   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd052eee12edf988, []int{3}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Output.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return m.Size()
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo




//

type Supply struct {
	Total github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=total,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total"`
}

func (m *Supply) Reset()         { *m = Supply{} }
func (m *Supply) String() string { return proto.CompactTextString(m) }
func (*Supply) ProtoMessage()    {}
func (*Supply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd052eee12edf988, []int{4}
}
func (m *Supply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Supply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Supply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supply.Merge(m, src)
}
func (m *Supply) XXX_Size() int {
	return m.Size()
}
func (m *Supply) XXX_DiscardUnknown() {
	xxx_messageInfo_Supply.DiscardUnknown(m)
}

var xxx_messageInfo_Supply proto.InternalMessageInfo



type DenomUnit struct {

	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`





	Exponent uint32 `protobuf:"varint,2,opt,name=exponent,proto3" json:"exponent,omitempty"`

	Aliases []string `protobuf:"bytes,3,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (m *DenomUnit) Reset()         { *m = DenomUnit{} }
func (m *DenomUnit) String() string { return proto.CompactTextString(m) }
func (*DenomUnit) ProtoMessage()    {}
func (*DenomUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd052eee12edf988, []int{5}
}
func (m *DenomUnit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomUnit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomUnit.Merge(m, src)
}
func (m *DenomUnit) XXX_Size() int {
	return m.Size()
}
func (m *DenomUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomUnit.DiscardUnknown(m)
}

var xxx_messageInfo_DenomUnit proto.InternalMessageInfo

func (m *DenomUnit) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DenomUnit) GetExponent() uint32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func (m *DenomUnit) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}



type Metadata struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`

	DenomUnits []*DenomUnit `protobuf:"bytes,2,rep,name=denom_units,json=denomUnits,proto3" json:"denom_units,omitempty"`

	Base string `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`


	Display string `protobuf:"bytes,4,opt,name=display,proto3" json:"display,omitempty"`

	//

	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`


	//

	Symbol string `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd052eee12edf988, []int{6}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Metadata) GetDenomUnits() []*DenomUnit {
	if m != nil {
		return m.DenomUnits
	}
	return nil
}

func (m *Metadata) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Metadata) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "cosmos.bank.v1beta1.Params")
	proto.RegisterType((*SendEnabled)(nil), "cosmos.bank.v1beta1.SendEnabled")
	proto.RegisterType((*Input)(nil), "cosmos.bank.v1beta1.Input")
	proto.RegisterType((*Output)(nil), "cosmos.bank.v1beta1.Output")
	proto.RegisterType((*Supply)(nil), "cosmos.bank.v1beta1.Supply")
	proto.RegisterType((*DenomUnit)(nil), "cosmos.bank.v1beta1.DenomUnit")
	proto.RegisterType((*Metadata)(nil), "cosmos.bank.v1beta1.Metadata")
}

func init() { proto.RegisterFile("cosmos/bank/v1beta1/bank.proto", fileDescriptor_dd052eee12edf988) }

var fileDescriptor_dd052eee12edf988 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0x35, 0x8d, 0x49, 0x2f, 0xb0, 0x1c, 0x15, 0x72, 0x23, 0x61, 0x1b, 0x4b, 0x48, 0x29,
	0xa2, 0x4e, 0x0a, 0x0c, 0x28, 0x0b, 0x52, 0xca, 0x0f, 0x75, 0x40, 0x20, 0x57, 0x08, 0x09, 0x86,
	0xe8, 0x9c, 0xbb, 0x06, 0xab, 0xf6, 0x9d, 0x95, 0x3b, 0x57, 0xf5, 0x7f, 0xc0, 0x04, 0x8c, 0x8c,
	0x9d, 0x59, 0xe1, 0x7f, 0xa0, 0x63, 0x05, 0x0b, 0x53, 0x40, 0xc9, 0xc2, 0xdc, 0xbf, 0x00, 0xf9,
	0xce, 0xf9, 0x51, 0x29, 0x20, 0x06, 0x06, 0xa6, 0xbc, 0xef, 0xbd, 0xef, 0x7d, 0xef, 0xe9, 0xbb,
	0xe7, 0x40, 0xbb, 0xcf, 0x45, 0xc2, 0x45, 0x2b, 0xc4, 0xec, 0xa0, 0x75, 0xb8, 0x1d, 0x52, 0x89,
	0xb7, 0x15, 0xf0, 0xd3, 0x21, 0x97, 0x1c, 0x5d, 0xd6, 0x75, 0x5f, 0xa5, 0xca, 0x7a, 0x63, 0x7d,
	0xc0, 0x07, 0x5c, 0xd5, 0x5b, 0x45, 0xa4, 0xa9, 0x8d, 0x0d, 0x4d, 0xed, 0xe9, 0x42, 0xd9, 0xa7,
	0x4b, 0xf3, 0x29, 0x82, 0xce, 0xa6, 0xf4, 0x79, 0xc4, 0x74, 0xdd, 0xfb, 0x0a, 0xa0, 0xf9, 0x14,
	0x0f, 0x71, 0x22, 0xd0, 0x3e, 0xbc, 0x28, 0x28, 0x23, 0x3d, 0xca, 0x70, 0x18, 0x53, 0x62, 0x01,
	0xb7, 0xd2, 0xac, 0xdf, 0x72, 0xfd, 0x25, 0x7b, 0xf8, 0x7b, 0x94, 0x91, 0x07, 0x9a, 0xd7, 0xbd,
	0x76, 0x36, 0x72, 0xae, 0xe6, 0x38, 0x89, 0x3b, 0xde, 0x62, 0xff, 0x4d, 0x9e, 0x44, 0x92, 0x26,
	0xa9, 0xcc, 0xbd, 0xa0, 0x2e, 0xe6, 0x7c, 0xf4, 0x12, 0xae, 0x13, 0xba, 0x8f, 0xb3, 0x58, 0xf6,
	0xce, 0xcd, 0x5b, 0x71, 0x41, 0xb3, 0xd6, 0xdd, 0x3c, 0x1b, 0x39, 0xd7, 0xb5, 0xda, 0x32, 0xd6,
	0xa2, 0x2a, 0x2a, 0x09, 0x0b, 0xcb, 0x74, 0x56, 0xdf, 0x1f, 0x3b, 0x86, 0xf7, 0x08, 0xd6, 0x17,
	0x92, 0x68, 0x1d, 0x56, 0x09, 0x65, 0x3c, 0xb1, 0x80, 0x0b, 0x9a, 0x6b, 0x81, 0x06, 0xc8, 0x82,
	0x17, 0xce, 0x8d, 0x0e, 0xa6, 0xb0, 0x53, 0x2b, 0x44, 0x7e, 0x1e, 0x3b, 0xc0, 0x7b, 0x03, 0x60,
	0x75, 0x97, 0xa5, 0x99, 0x2c, 0xd8, 0x98, 0x90, 0x21, 0x15, 0xa2, 0x54, 0x99, 0x42, 0x84, 0x61,
	0xb5, 0x30, 0x54, 0x58, 0x2b, 0xca, 0xb0, 0x8d, 0xb9, 0x61, 0x82, 0xce, 0x0c, 0xdb, 0xe1, 0x11,
	0xeb, 0xb6, 0x4f, 0x46, 0x8e, 0xf1, 0xe1, 0xbb, 0xd3, 0x1c, 0x44, 0xf2, 0x55, 0x16, 0xfa, 0x7d,
	0x9e, 0x94, 0xaf, 0x55, 0xfe, 0x6c, 0x09, 0x72, 0xd0, 0x92, 0x79, 0x4a, 0x85, 0x6a, 0x10, 0x81,
	0x56, 0xee, 0xd4, 0x5e, 0xeb, 0x85, 0x0c, 0xef, 0x2d, 0x80, 0xe6, 0x93, 0x4c, 0xfe, 0x47, 0x1b,
	0x7d, 0x04, 0xd0, 0xdc, 0xcb, 0xd2, 0x34, 0xce, 0x8b, 0xb9, 0x92, 0x4b, 0x1c, 0x97, 0xa7, 0xf3,
	0x6f, 0xe7, 0x2a, 0xe5, 0xce, 0xc3, 0x72, 0x2e, 0xf8, 0xf2, 0x69, 0xeb, 0xee, 0x8d, 0x3f, 0x76,
	0x1f, 0xe9, 0x4f, 0x2b, 0xa6, 0x03, 0xdc, 0xcf, 0x5b, 0x87, 0xed, 0x3b, 0x6d, 0x5f, 0xef, 0xb9,
	0x6b, 0x01, 0xef, 0x39, 0x5c, 0xbb, 0x5f, 0x5c, 0xc1, 0x33, 0x16, 0xc9, 0xdf, 0xdc, 0x47, 0x03,
	0xd6, 0xe8, 0x51, 0xca, 0x19, 0x65, 0x52, 0x1d, 0xc8, 0xa5, 0x60, 0x86, 0x95, 0xf7, 0x71, 0x84,
	0x05, 0x15, 0x56, 0xc5, 0xad, 0x28, 0xef, 0x35, 0xf4, 0x3e, 0x03, 0x58, 0x7b, 0x4c, 0x25, 0x26,
	0x58, 0x62, 0xe4, 0xc2, 0x3a, 0xa1, 0xa2, 0x3f, 0x8c, 0x52, 0x19, 0x71, 0x56, 0xca, 0x2f, 0xa6,
	0xd0, 0xbd, 0x82, 0xc1, 0x78, 0xd2, 0xcb, 0x58, 0x24, 0xa7, 0x0f, 0x66, 0x2f, 0xfd, 0xe6, 0x66,
	0xfb, 0x06, 0x90, 0x4c, 0x43, 0x81, 0x10, 0x5c, 0x2d, 0xec, 0xb5, 0x2a, 0x4a, 0x5b, 0xc5, 0xc5,
	0x76, 0x24, 0x12, 0x69, 0x8c, 0x73, 0x6b, 0x55, 0x5f, 0x46, 0x09, 0x0b, 0x36, 0xc3, 0x09, 0xb5,
	0xaa, 0x9a, 0x5d, 0xc4, 0xe8, 0x0a, 0x34, 0x45, 0x9e, 0x84, 0x3c, 0xb6, 0x4c, 0x95, 0x2d, 0x51,
	0x77, 0xe7, 0x64, 0x6c, 0x83, 0xd3, 0xb1, 0x0d, 0x7e, 0x8c, 0x6d, 0xf0, 0x6e, 0x62, 0x1b, 0xa7,
	0x13, 0xdb, 0xf8, 0x36, 0xb1, 0x8d, 0x17, 0x9b, 0x7f, 0xe3, 0xbb, 0x7a, 0xbc, 0xd0, 0x54, 0x7f,
	0x33, 0xb7, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x03, 0xbf, 0xe9, 0xee, 0x04, 0x00, 0x00,
}

func (this *SendEnabled) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SendEnabled)
	if !ok {
		that2, ok := that.(SendEnabled)
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
	if this.Denom != that1.Denom {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	return true
}
func (this *Supply) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Supply)
	if !ok {
		that2, ok := that.(Supply)
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
	if len(this.Total) != len(that1.Total) {
		return false
	}
	for i := range this.Total {
		if !this.Total[i].Equal(&that1.Total[i]) {
			return false
		}
	}
	return true
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
	if m.DefaultSendEnabled {
		i--
		if m.DefaultSendEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.SendEnabled) > 0 {
		for iNdEx := len(m.SendEnabled) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SendEnabled[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBank(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SendEnabled) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendEnabled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendEnabled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Input) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBank(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Output) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Output) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Output) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBank(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Supply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Total) > 0 {
		for iNdEx := len(m.Total) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Total[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBank(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DenomUnit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomUnit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomUnit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Aliases) > 0 {
		for iNdEx := len(m.Aliases) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Aliases[iNdEx])
			copy(dAtA[i:], m.Aliases[iNdEx])
			i = encodeVarintBank(dAtA, i, uint64(len(m.Aliases[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Exponent != 0 {
		i = encodeVarintBank(dAtA, i, uint64(m.Exponent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Display) > 0 {
		i -= len(m.Display)
		copy(dAtA[i:], m.Display)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Display)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Base) > 0 {
		i -= len(m.Base)
		copy(dAtA[i:], m.Base)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Base)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DenomUnits) > 0 {
		for iNdEx := len(m.DenomUnits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomUnits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBank(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintBank(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBank(dAtA []byte, offset int, v uint64) int {
	offset -= sovBank(v)
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
	if len(m.SendEnabled) > 0 {
		for _, e := range m.SendEnabled {
			l = e.Size()
			n += 1 + l + sovBank(uint64(l))
		}
	}
	if m.DefaultSendEnabled {
		n += 2
	}
	return n
}

func (m *SendEnabled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *Input) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovBank(uint64(l))
		}
	}
	return n
}

func (m *Output) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovBank(uint64(l))
		}
	}
	return n
}

func (m *Supply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Total) > 0 {
		for _, e := range m.Total {
			l = e.Size()
			n += 1 + l + sovBank(uint64(l))
		}
	}
	return n
}

func (m *DenomUnit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	if m.Exponent != 0 {
		n += 1 + sovBank(uint64(m.Exponent))
	}
	if len(m.Aliases) > 0 {
		for _, s := range m.Aliases {
			l = len(s)
			n += 1 + l + sovBank(uint64(l))
		}
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	if len(m.DenomUnits) > 0 {
		for _, e := range m.DenomUnits {
			l = e.Size()
			n += 1 + l + sovBank(uint64(l))
		}
	}
	l = len(m.Base)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	l = len(m.Display)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovBank(uint64(l))
	}
	return n
}

func sovBank(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBank(x uint64) (n int) {
	return sovBank(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendEnabled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SendEnabled = append(m.SendEnabled, &SendEnabled{})
			if err := m.SendEnabled[len(m.SendEnabled)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultSendEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
			m.DefaultSendEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func (m *SendEnabled) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: SendEnabled: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendEnabled: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func (m *Output) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: Output: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Output: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func (m *Supply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: Supply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Total = append(m.Total, types.Coin{})
			if err := m.Total[len(m.Total)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func (m *DenomUnit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: DenomUnit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomUnit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			m.Exponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exponent |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aliases", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aliases = append(m.Aliases, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBank
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomUnits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomUnits = append(m.DenomUnits, &DenomUnit{})
			if err := m.DenomUnits[len(m.DenomUnits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Base = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Display", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Display = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBank
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
				return ErrInvalidLengthBank
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBank
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBank(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBank
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
func skipBank(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBank
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
					return 0, ErrIntOverflowBank
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
					return 0, ErrIntOverflowBank
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
				return 0, ErrInvalidLengthBank
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBank
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBank
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBank        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBank          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBank = fmt.Errorf("proto: unexpected end of group")
)
