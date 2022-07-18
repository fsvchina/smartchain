


package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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



type MsgSetWithdrawAddress struct {
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	WithdrawAddress  string `protobuf:"bytes,2,opt,name=withdraw_address,json=withdrawAddress,proto3" json:"withdraw_address,omitempty" yaml:"withdraw_address"`
}

func (m *MsgSetWithdrawAddress) Reset()         { *m = MsgSetWithdrawAddress{} }
func (m *MsgSetWithdrawAddress) String() string { return proto.CompactTextString(m) }
func (*MsgSetWithdrawAddress) ProtoMessage()    {}
func (*MsgSetWithdrawAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed4f433d965e58ca, []int{0}
}
func (m *MsgSetWithdrawAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetWithdrawAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetWithdrawAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetWithdrawAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetWithdrawAddress.Merge(m, src)
}
func (m *MsgSetWithdrawAddress) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetWithdrawAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetWithdrawAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetWithdrawAddress proto.InternalMessageInfo


type MsgSetWithdrawAddressResponse struct {
}

func (m *MsgSetWithdrawAddressResponse) Reset()         { *m = MsgSetWithdrawAddressResponse{} }
func (m *MsgSetWithdrawAddressResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetWithdrawAddressResponse) ProtoMessage()    {}
func (*MsgSetWithdrawAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed4f433d965e58ca, []int{1}
}
func (m *MsgSetWithdrawAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetWithdrawAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetWithdrawAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetWithdrawAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetWithdrawAddressResponse.Merge(m, src)
}
func (m *MsgSetWithdrawAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetWithdrawAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetWithdrawAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetWithdrawAddressResponse proto.InternalMessageInfo



type MsgWithdrawDelegatorReward struct {
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
}

func (m *MsgWithdrawDelegatorReward) Reset()         { *m = MsgWithdrawDelegatorReward{} }
func (m *MsgWithdrawDelegatorReward) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawDelegatorReward) ProtoMessage()    {}
func (*MsgWithdrawDelegatorReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed4f433d965e58ca, []int{2}
}
func (m *MsgWithdrawDelegatorReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawDelegatorReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawDelegatorReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawDelegatorReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawDelegatorReward.Merge(m, src)
}
func (m *MsgWithdrawDelegatorReward) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawDelegatorReward) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawDelegatorReward.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawDelegatorReward proto.InternalMessageInfo


type MsgWithdrawDelegatorRewardResponse struct {
}

func (m *MsgWithdrawDelegatorRewardResponse) Reset()         { *m = MsgWithdrawDelegatorRewardResponse{} }
func (m *MsgWithdrawDelegatorRewardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawDelegatorRewardResponse) ProtoMessage()    {}
func (*MsgWithdrawDelegatorRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed4f433d965e58ca, []int{3}
}
func (m *MsgWithdrawDelegatorRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawDelegatorRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawDelegatorRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawDelegatorRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawDelegatorRewardResponse.Merge(m, src)
}
func (m *MsgWithdrawDelegatorRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawDelegatorRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawDelegatorRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawDelegatorRewardResponse proto.InternalMessageInfo



type MsgWithdrawValidatorCommission struct {
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
}

func (m *MsgWithdrawValidatorCommission) Reset()         { *m = MsgWithdrawValidatorCommission{} }
func (m *MsgWithdrawValidatorCommission) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawValidatorCommission) ProtoMessage()    {}
func (*MsgWithdrawValidatorCommission) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed4f433d965e58ca, []int{4}
}
func (m *MsgWithdrawValidatorCommission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawValidatorCommission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawValidatorCommission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawValidatorCommission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawValidatorCommission.Merge(m, src)
}
func (m *MsgWithdrawValidatorCommission) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawValidatorCommission) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawValidatorCommission.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawValidatorCommission proto.InternalMessageInfo


type MsgWithdrawValidatorCommissionResponse struct {
}

func (m *MsgWithdrawValidatorCommissionResponse) Reset() {
	*m = MsgWithdrawValidatorCommissionResponse{}
}
func (m *MsgWithdrawValidatorCommissionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawValidatorCommissionResponse) ProtoMessage()    {}
func (*MsgWithdrawValidatorCommissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed4f433d965e58ca, []int{5}
}
func (m *MsgWithdrawValidatorCommissionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawValidatorCommissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawValidatorCommissionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawValidatorCommissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawValidatorCommissionResponse.Merge(m, src)
}
func (m *MsgWithdrawValidatorCommissionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawValidatorCommissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawValidatorCommissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawValidatorCommissionResponse proto.InternalMessageInfo



type MsgFundCommunityPool struct {
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	Depositor string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
}

func (m *MsgFundCommunityPool) Reset()         { *m = MsgFundCommunityPool{} }
func (m *MsgFundCommunityPool) String() string { return proto.CompactTextString(m) }
func (*MsgFundCommunityPool) ProtoMessage()    {}
func (*MsgFundCommunityPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed4f433d965e58ca, []int{6}
}
func (m *MsgFundCommunityPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundCommunityPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundCommunityPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundCommunityPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundCommunityPool.Merge(m, src)
}
func (m *MsgFundCommunityPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundCommunityPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundCommunityPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundCommunityPool proto.InternalMessageInfo


type MsgFundCommunityPoolResponse struct {
}

func (m *MsgFundCommunityPoolResponse) Reset()         { *m = MsgFundCommunityPoolResponse{} }
func (m *MsgFundCommunityPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFundCommunityPoolResponse) ProtoMessage()    {}
func (*MsgFundCommunityPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed4f433d965e58ca, []int{7}
}
func (m *MsgFundCommunityPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundCommunityPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundCommunityPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundCommunityPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundCommunityPoolResponse.Merge(m, src)
}
func (m *MsgFundCommunityPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundCommunityPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundCommunityPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundCommunityPoolResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetWithdrawAddress)(nil), "cosmos.distribution.v1beta1.MsgSetWithdrawAddress")
	proto.RegisterType((*MsgSetWithdrawAddressResponse)(nil), "cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse")
	proto.RegisterType((*MsgWithdrawDelegatorReward)(nil), "cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward")
	proto.RegisterType((*MsgWithdrawDelegatorRewardResponse)(nil), "cosmos.distribution.v1beta1.MsgWithdrawDelegatorRewardResponse")
	proto.RegisterType((*MsgWithdrawValidatorCommission)(nil), "cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission")
	proto.RegisterType((*MsgWithdrawValidatorCommissionResponse)(nil), "cosmos.distribution.v1beta1.MsgWithdrawValidatorCommissionResponse")
	proto.RegisterType((*MsgFundCommunityPool)(nil), "cosmos.distribution.v1beta1.MsgFundCommunityPool")
	proto.RegisterType((*MsgFundCommunityPoolResponse)(nil), "cosmos.distribution.v1beta1.MsgFundCommunityPoolResponse")
}

func init() {
	proto.RegisterFile("cosmos/distribution/v1beta1/tx.proto", fileDescriptor_ed4f433d965e58ca)
}

var fileDescriptor_ed4f433d965e58ca = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x2d, 0xaa, 0xe8, 0x31, 0x90, 0x58, 0x45, 0x0d, 0x4e, 0x38, 0x57, 0x56, 0x85,
	0xb2, 0x60, 0x93, 0x30, 0x20, 0xc2, 0x80, 0x48, 0x50, 0xa5, 0x0e, 0x11, 0xc8, 0x48, 0x20, 0xb1,
	0x20, 0x3b, 0x77, 0x72, 0x4f, 0xc4, 0xbe, 0xc8, 0x77, 0x6e, 0x9a, 0x11, 0x89, 0x81, 0x11, 0x89,
	0x0f, 0x40, 0x25, 0x16, 0xc4, 0xcc, 0xc8, 0x07, 0xc8, 0xd8, 0x91, 0x29, 0xa0, 0x64, 0x61, 0xee,
	0x27, 0x40, 0xf1, 0x3f, 0x92, 0xda, 0x49, 0x29, 0xe9, 0x94, 0xe8, 0xbd, 0xe7, 0xf9, 0xf9, 0x79,
	0x95, 0xe7, 0x62, 0xb8, 0xdb, 0x61, 0xdc, 0x65, 0xdc, 0xc0, 0x94, 0x0b, 0x9f, 0xda, 0x81, 0xa0,
	0xcc, 0x33, 0x0e, 0x6b, 0x36, 0x11, 0x56, 0xcd, 0x10, 0x47, 0x7a, 0xcf, 0x67, 0x82, 0xc9, 0xe5,
	0x48, 0xa5, 0xcf, 0xaa, 0xf4, 0x58, 0xa5, 0x6c, 0x39, 0xcc, 0x61, 0xa1, 0xce, 0x98, 0x7e, 0x8b,
	0x2c, 0x0a, 0x8a, 0xc1, 0xb6, 0xc5, 0x49, 0x0a, 0xec, 0x30, 0xea, 0x45, 0xe7, 0xda, 0x37, 0x00,
	0x6f, 0xb4, 0xb9, 0xf3, 0x9c, 0x88, 0x97, 0x54, 0x1c, 0x60, 0xdf, 0xea, 0x3f, 0xc6, 0xd8, 0x27,
	0x9c, 0xcb, 0xfb, 0xb0, 0x88, 0x49, 0x97, 0x38, 0x96, 0x60, 0xfe, 0x6b, 0x2b, 0x1a, 0x96, 0xc0,
	0x0e, 0xa8, 0x6e, 0x36, 0x2b, 0xa7, 0x23, 0xb5, 0x34, 0xb0, 0xdc, 0x6e, 0x43, 0xcb, 0x48, 0x34,
	0xb3, 0x90, 0xce, 0x12, 0xd4, 0x1e, 0x2c, 0xf4, 0x63, 0x7a, 0x4a, 0x5a, 0x0b, 0x49, 0xe5, 0xd3,
	0x91, 0xba, 0x1d, 0x91, 0xce, 0x2a, 0x34, 0xf3, 0x7a, 0x7f, 0x3e, 0x52, 0xe3, 0xea, 0xfb, 0x63,
	0x55, 0xfa, 0x7d, 0xac, 0x4a, 0x9a, 0x0a, 0x6f, 0xe5, 0xa6, 0x36, 0x09, 0xef, 0x31, 0x8f, 0x13,
	0xed, 0x3b, 0x80, 0x4a, 0x9b, 0x3b, 0xc9, 0xf1, 0x93, 0x24, 0x92, 0x49, 0xfa, 0x96, 0x8f, 0x2f,
	0x73, 0xb9, 0x7d, 0x58, 0x3c, 0xb4, 0xba, 0x14, 0xcf, 0xa1, 0xd6, 0xce, 0xa2, 0x32, 0x12, 0xcd,
	0x2c, 0xa4, 0xb3, 0xec, 0x7e, 0xbb, 0x50, 0x5b, 0x9c, 0x3e, 0x5d, 0x32, 0x80, 0x68, 0x46, 0xf5,
	0x22, 0xc1, 0xb5, 0x98, 0xeb, 0x52, 0xce, 0x29, 0xf3, 0xf2, 0xc3, 0x81, 0x15, 0xc3, 0x55, 0xe1,
	0xed, 0xe5, 0x8f, 0x4d, 0x03, 0x7e, 0x06, 0x70, 0xab, 0xcd, 0x9d, 0xbd, 0xc0, 0xc3, 0xd3, 0xd3,
	0xc0, 0xa3, 0x62, 0xf0, 0x8c, 0xb1, 0xae, 0xdc, 0x81, 0x1b, 0x96, 0xcb, 0x02, 0x4f, 0x94, 0xc0,
	0xce, 0x7a, 0xf5, 0x5a, 0xfd, 0xa6, 0x1e, 0x57, 0x7b, 0xda, 0xd3, 0xa4, 0xd2, 0x7a, 0x8b, 0x51,
	0xaf, 0x79, 0x77, 0x38, 0x52, 0xa5, 0xaf, 0x3f, 0xd5, 0xaa, 0x43, 0xc5, 0x41, 0x60, 0xeb, 0x1d,
	0xe6, 0x1a, 0x71, 0xa9, 0xa3, 0x8f, 0x3b, 0x1c, 0xbf, 0x31, 0xc4, 0xa0, 0x47, 0x78, 0x68, 0xe0,
	0x66, 0x8c, 0x96, 0x2b, 0x70, 0x13, 0x93, 0x1e, 0xe3, 0x54, 0x30, 0x3f, 0xfa, 0x45, 0xcc, 0xbf,
	0x83, 0x99, 0x7d, 0x10, 0xac, 0xe4, 0x85, 0x4c, 0xb6, 0xa8, 0x0f, 0xaf, 0xc0, 0xf5, 0x36, 0x77,
	0xe4, 0x77, 0x00, 0xca, 0x39, 0x17, 0xa5, 0xae, 0x2f, 0xb9, 0x96, 0x7a, 0x6e, 0x4d, 0x95, 0xc6,
	0xc5, 0x3d, 0x49, 0x1c, 0xf9, 0x23, 0x80, 0xdb, 0x8b, 0x7a, 0x7d, 0xff, 0x3c, 0xee, 0x02, 0xa3,
	0xf2, 0xe8, 0x3f, 0x8d, 0x69, 0xaa, 0x4f, 0x00, 0x96, 0x97, 0x35, 0xf1, 0xe1, 0xbf, 0x3e, 0x20,
	0xc7, 0xac, 0xb4, 0x56, 0x30, 0xa7, 0x09, 0xdf, 0x02, 0x58, 0xcc, 0x36, 0xb1, 0x76, 0x1e, 0x3a,
	0x63, 0x51, 0x1e, 0x5c, 0xd8, 0x92, 0x64, 0x68, 0x3e, 0xfd, 0x32, 0x46, 0x60, 0x38, 0x46, 0xe0,
	0x64, 0x8c, 0xc0, 0xaf, 0x31, 0x02, 0x1f, 0x26, 0x48, 0x3a, 0x99, 0x20, 0xe9, 0xc7, 0x04, 0x49,
	0xaf, 0x6a, 0x4b, 0x2b, 0x7e, 0x34, 0xff, 0x76, 0x08, 0x1b, 0x6f, 0x6f, 0x84, 0x7f, 0xe3, 0xf7,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x18, 0x71, 0xb0, 0x41, 0x06, 0x00, 0x00,
}

func (this *MsgSetWithdrawAddressResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgSetWithdrawAddressResponse)
	if !ok {
		that2, ok := that.(MsgSetWithdrawAddressResponse)
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
	return true
}
func (this *MsgWithdrawDelegatorRewardResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgWithdrawDelegatorRewardResponse)
	if !ok {
		that2, ok := that.(MsgWithdrawDelegatorRewardResponse)
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
	return true
}
func (this *MsgWithdrawValidatorCommissionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgWithdrawValidatorCommissionResponse)
	if !ok {
		that2, ok := that.(MsgWithdrawValidatorCommissionResponse)
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
	return true
}
func (this *MsgFundCommunityPoolResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgFundCommunityPoolResponse)
	if !ok {
		that2, ok := that.(MsgFundCommunityPoolResponse)
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
	return true
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type MsgClient interface {


	SetWithdrawAddress(ctx context.Context, in *MsgSetWithdrawAddress, opts ...grpc.CallOption) (*MsgSetWithdrawAddressResponse, error)


	WithdrawDelegatorReward(ctx context.Context, in *MsgWithdrawDelegatorReward, opts ...grpc.CallOption) (*MsgWithdrawDelegatorRewardResponse, error)


	WithdrawValidatorCommission(ctx context.Context, in *MsgWithdrawValidatorCommission, opts ...grpc.CallOption) (*MsgWithdrawValidatorCommissionResponse, error)


	FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetWithdrawAddress(ctx context.Context, in *MsgSetWithdrawAddress, opts ...grpc.CallOption) (*MsgSetWithdrawAddressResponse, error) {
	out := new(MsgSetWithdrawAddressResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Msg/SetWithdrawAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawDelegatorReward(ctx context.Context, in *MsgWithdrawDelegatorReward, opts ...grpc.CallOption) (*MsgWithdrawDelegatorRewardResponse, error) {
	out := new(MsgWithdrawDelegatorRewardResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Msg/WithdrawDelegatorReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawValidatorCommission(ctx context.Context, in *MsgWithdrawValidatorCommission, opts ...grpc.CallOption) (*MsgWithdrawValidatorCommissionResponse, error) {
	out := new(MsgWithdrawValidatorCommissionResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Msg/WithdrawValidatorCommission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error) {
	out := new(MsgFundCommunityPoolResponse)
	err := c.cc.Invoke(ctx, "/cosmos.distribution.v1beta1.Msg/FundCommunityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type MsgServer interface {


	SetWithdrawAddress(context.Context, *MsgSetWithdrawAddress) (*MsgSetWithdrawAddressResponse, error)


	WithdrawDelegatorReward(context.Context, *MsgWithdrawDelegatorReward) (*MsgWithdrawDelegatorRewardResponse, error)


	WithdrawValidatorCommission(context.Context, *MsgWithdrawValidatorCommission) (*MsgWithdrawValidatorCommissionResponse, error)


	FundCommunityPool(context.Context, *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error)
}


type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetWithdrawAddress(ctx context.Context, req *MsgSetWithdrawAddress) (*MsgSetWithdrawAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWithdrawAddress not implemented")
}
func (*UnimplementedMsgServer) WithdrawDelegatorReward(ctx context.Context, req *MsgWithdrawDelegatorReward) (*MsgWithdrawDelegatorRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawDelegatorReward not implemented")
}
func (*UnimplementedMsgServer) WithdrawValidatorCommission(ctx context.Context, req *MsgWithdrawValidatorCommission) (*MsgWithdrawValidatorCommissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawValidatorCommission not implemented")
}
func (*UnimplementedMsgServer) FundCommunityPool(ctx context.Context, req *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundCommunityPool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetWithdrawAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetWithdrawAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetWithdrawAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Msg/SetWithdrawAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetWithdrawAddress(ctx, req.(*MsgSetWithdrawAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawDelegatorReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawDelegatorReward)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawDelegatorReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Msg/WithdrawDelegatorReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawDelegatorReward(ctx, req.(*MsgWithdrawDelegatorReward))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawValidatorCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawValidatorCommission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawValidatorCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Msg/WithdrawValidatorCommission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawValidatorCommission(ctx, req.(*MsgWithdrawValidatorCommission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FundCommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundCommunityPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundCommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Msg/FundCommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundCommunityPool(ctx, req.(*MsgFundCommunityPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.distribution.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetWithdrawAddress",
			Handler:    _Msg_SetWithdrawAddress_Handler,
		},
		{
			MethodName: "WithdrawDelegatorReward",
			Handler:    _Msg_WithdrawDelegatorReward_Handler,
		},
		{
			MethodName: "WithdrawValidatorCommission",
			Handler:    _Msg_WithdrawValidatorCommission_Handler,
		},
		{
			MethodName: "FundCommunityPool",
			Handler:    _Msg_FundCommunityPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/distribution/v1beta1/tx.proto",
}

func (m *MsgSetWithdrawAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetWithdrawAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetWithdrawAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawAddress) > 0 {
		i -= len(m.WithdrawAddress)
		copy(dAtA[i:], m.WithdrawAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.WithdrawAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetWithdrawAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetWithdrawAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetWithdrawAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawDelegatorReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawDelegatorReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawDelegatorReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawDelegatorRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawDelegatorRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawDelegatorRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawValidatorCommission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawValidatorCommission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawValidatorCommission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawValidatorCommissionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawValidatorCommissionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawValidatorCommissionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgFundCommunityPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundCommunityPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundCommunityPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgFundCommunityPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundCommunityPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundCommunityPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgSetWithdrawAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.WithdrawAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetWithdrawAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawDelegatorReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawDelegatorRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawValidatorCommission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawValidatorCommissionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgFundCommunityPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgFundCommunityPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetWithdrawAddress) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetWithdrawAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetWithdrawAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
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
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddress", wireType)
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
			m.WithdrawAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSetWithdrawAddressResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSetWithdrawAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetWithdrawAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgWithdrawDelegatorReward) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawDelegatorReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawDelegatorReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
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
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MsgWithdrawDelegatorRewardResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawDelegatorRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawDelegatorRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgWithdrawValidatorCommission) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawValidatorCommission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawValidatorCommission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MsgWithdrawValidatorCommissionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawValidatorCommissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawValidatorCommissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgFundCommunityPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFundCommunityPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundCommunityPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
			m.Depositor = string(dAtA[iNdEx:postIndex])
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
func (m *MsgFundCommunityPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFundCommunityPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundCommunityPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
