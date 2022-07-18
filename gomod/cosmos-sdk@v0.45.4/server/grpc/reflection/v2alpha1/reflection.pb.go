


package v2alpha1

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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


type AppDescriptor struct {


	Authn *AuthnDescriptor `protobuf:"bytes,1,opt,name=authn,proto3" json:"authn,omitempty"`

	Chain *ChainDescriptor `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`

	Codec *CodecDescriptor `protobuf:"bytes,3,opt,name=codec,proto3" json:"codec,omitempty"`

	Configuration *ConfigurationDescriptor `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`

	QueryServices *QueryServicesDescriptor `protobuf:"bytes,5,opt,name=query_services,json=queryServices,proto3" json:"query_services,omitempty"`

	Tx *TxDescriptor `protobuf:"bytes,6,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *AppDescriptor) Reset()         { *m = AppDescriptor{} }
func (m *AppDescriptor) String() string { return proto.CompactTextString(m) }
func (*AppDescriptor) ProtoMessage()    {}
func (*AppDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{0}
}
func (m *AppDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppDescriptor.Merge(m, src)
}
func (m *AppDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *AppDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_AppDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_AppDescriptor proto.InternalMessageInfo

func (m *AppDescriptor) GetAuthn() *AuthnDescriptor {
	if m != nil {
		return m.Authn
	}
	return nil
}

func (m *AppDescriptor) GetChain() *ChainDescriptor {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *AppDescriptor) GetCodec() *CodecDescriptor {
	if m != nil {
		return m.Codec
	}
	return nil
}

func (m *AppDescriptor) GetConfiguration() *ConfigurationDescriptor {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *AppDescriptor) GetQueryServices() *QueryServicesDescriptor {
	if m != nil {
		return m.QueryServices
	}
	return nil
}

func (m *AppDescriptor) GetTx() *TxDescriptor {
	if m != nil {
		return m.Tx
	}
	return nil
}


type TxDescriptor struct {



	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`

	Msgs []*MsgDescriptor `protobuf:"bytes,2,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (m *TxDescriptor) Reset()         { *m = TxDescriptor{} }
func (m *TxDescriptor) String() string { return proto.CompactTextString(m) }
func (*TxDescriptor) ProtoMessage()    {}
func (*TxDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{1}
}
func (m *TxDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxDescriptor.Merge(m, src)
}
func (m *TxDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *TxDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TxDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TxDescriptor proto.InternalMessageInfo

func (m *TxDescriptor) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *TxDescriptor) GetMsgs() []*MsgDescriptor {
	if m != nil {
		return m.Msgs
	}
	return nil
}



type AuthnDescriptor struct {

	SignModes []*SigningModeDescriptor `protobuf:"bytes,1,rep,name=sign_modes,json=signModes,proto3" json:"sign_modes,omitempty"`
}

func (m *AuthnDescriptor) Reset()         { *m = AuthnDescriptor{} }
func (m *AuthnDescriptor) String() string { return proto.CompactTextString(m) }
func (*AuthnDescriptor) ProtoMessage()    {}
func (*AuthnDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{2}
}
func (m *AuthnDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthnDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthnDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthnDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthnDescriptor.Merge(m, src)
}
func (m *AuthnDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *AuthnDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthnDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_AuthnDescriptor proto.InternalMessageInfo

func (m *AuthnDescriptor) GetSignModes() []*SigningModeDescriptor {
	if m != nil {
		return m.SignModes
	}
	return nil
}





type SigningModeDescriptor struct {

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`

	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`


	AuthnInfoProviderMethodFullname string `protobuf:"bytes,3,opt,name=authn_info_provider_method_fullname,json=authnInfoProviderMethodFullname,proto3" json:"authn_info_provider_method_fullname,omitempty"`
}

func (m *SigningModeDescriptor) Reset()         { *m = SigningModeDescriptor{} }
func (m *SigningModeDescriptor) String() string { return proto.CompactTextString(m) }
func (*SigningModeDescriptor) ProtoMessage()    {}
func (*SigningModeDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{3}
}
func (m *SigningModeDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigningModeDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigningModeDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SigningModeDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningModeDescriptor.Merge(m, src)
}
func (m *SigningModeDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *SigningModeDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningModeDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_SigningModeDescriptor proto.InternalMessageInfo

func (m *SigningModeDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SigningModeDescriptor) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *SigningModeDescriptor) GetAuthnInfoProviderMethodFullname() string {
	if m != nil {
		return m.AuthnInfoProviderMethodFullname
	}
	return ""
}


type ChainDescriptor struct {

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ChainDescriptor) Reset()         { *m = ChainDescriptor{} }
func (m *ChainDescriptor) String() string { return proto.CompactTextString(m) }
func (*ChainDescriptor) ProtoMessage()    {}
func (*ChainDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{4}
}
func (m *ChainDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainDescriptor.Merge(m, src)
}
func (m *ChainDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *ChainDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ChainDescriptor proto.InternalMessageInfo

func (m *ChainDescriptor) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}


type CodecDescriptor struct {

	Interfaces []*InterfaceDescriptor `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
}

func (m *CodecDescriptor) Reset()         { *m = CodecDescriptor{} }
func (m *CodecDescriptor) String() string { return proto.CompactTextString(m) }
func (*CodecDescriptor) ProtoMessage()    {}
func (*CodecDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{5}
}
func (m *CodecDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodecDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodecDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodecDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodecDescriptor.Merge(m, src)
}
func (m *CodecDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *CodecDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_CodecDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_CodecDescriptor proto.InternalMessageInfo

func (m *CodecDescriptor) GetInterfaces() []*InterfaceDescriptor {
	if m != nil {
		return m.Interfaces
	}
	return nil
}


type InterfaceDescriptor struct {

	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`


	InterfaceAcceptingMessages []*InterfaceAcceptingMessageDescriptor `protobuf:"bytes,2,rep,name=interface_accepting_messages,json=interfaceAcceptingMessages,proto3" json:"interface_accepting_messages,omitempty"`

	InterfaceImplementers []*InterfaceImplementerDescriptor `protobuf:"bytes,3,rep,name=interface_implementers,json=interfaceImplementers,proto3" json:"interface_implementers,omitempty"`
}

func (m *InterfaceDescriptor) Reset()         { *m = InterfaceDescriptor{} }
func (m *InterfaceDescriptor) String() string { return proto.CompactTextString(m) }
func (*InterfaceDescriptor) ProtoMessage()    {}
func (*InterfaceDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{6}
}
func (m *InterfaceDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterfaceDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterfaceDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterfaceDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceDescriptor.Merge(m, src)
}
func (m *InterfaceDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *InterfaceDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceDescriptor proto.InternalMessageInfo

func (m *InterfaceDescriptor) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *InterfaceDescriptor) GetInterfaceAcceptingMessages() []*InterfaceAcceptingMessageDescriptor {
	if m != nil {
		return m.InterfaceAcceptingMessages
	}
	return nil
}

func (m *InterfaceDescriptor) GetInterfaceImplementers() []*InterfaceImplementerDescriptor {
	if m != nil {
		return m.InterfaceImplementers
	}
	return nil
}


type InterfaceImplementerDescriptor struct {

	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`




	TypeUrl string `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
}

func (m *InterfaceImplementerDescriptor) Reset()         { *m = InterfaceImplementerDescriptor{} }
func (m *InterfaceImplementerDescriptor) String() string { return proto.CompactTextString(m) }
func (*InterfaceImplementerDescriptor) ProtoMessage()    {}
func (*InterfaceImplementerDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{7}
}
func (m *InterfaceImplementerDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterfaceImplementerDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterfaceImplementerDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterfaceImplementerDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceImplementerDescriptor.Merge(m, src)
}
func (m *InterfaceImplementerDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *InterfaceImplementerDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceImplementerDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceImplementerDescriptor proto.InternalMessageInfo

func (m *InterfaceImplementerDescriptor) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *InterfaceImplementerDescriptor) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}



type InterfaceAcceptingMessageDescriptor struct {

	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`



	FieldDescriptorNames []string `protobuf:"bytes,2,rep,name=field_descriptor_names,json=fieldDescriptorNames,proto3" json:"field_descriptor_names,omitempty"`
}

func (m *InterfaceAcceptingMessageDescriptor) Reset()         { *m = InterfaceAcceptingMessageDescriptor{} }
func (m *InterfaceAcceptingMessageDescriptor) String() string { return proto.CompactTextString(m) }
func (*InterfaceAcceptingMessageDescriptor) ProtoMessage()    {}
func (*InterfaceAcceptingMessageDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{8}
}
func (m *InterfaceAcceptingMessageDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterfaceAcceptingMessageDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterfaceAcceptingMessageDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterfaceAcceptingMessageDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceAcceptingMessageDescriptor.Merge(m, src)
}
func (m *InterfaceAcceptingMessageDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *InterfaceAcceptingMessageDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceAcceptingMessageDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceAcceptingMessageDescriptor proto.InternalMessageInfo

func (m *InterfaceAcceptingMessageDescriptor) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *InterfaceAcceptingMessageDescriptor) GetFieldDescriptorNames() []string {
	if m != nil {
		return m.FieldDescriptorNames
	}
	return nil
}


type ConfigurationDescriptor struct {

	Bech32AccountAddressPrefix string `protobuf:"bytes,1,opt,name=bech32_account_address_prefix,json=bech32AccountAddressPrefix,proto3" json:"bech32_account_address_prefix,omitempty"`
}

func (m *ConfigurationDescriptor) Reset()         { *m = ConfigurationDescriptor{} }
func (m *ConfigurationDescriptor) String() string { return proto.CompactTextString(m) }
func (*ConfigurationDescriptor) ProtoMessage()    {}
func (*ConfigurationDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{9}
}
func (m *ConfigurationDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigurationDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigurationDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigurationDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationDescriptor.Merge(m, src)
}
func (m *ConfigurationDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *ConfigurationDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationDescriptor proto.InternalMessageInfo

func (m *ConfigurationDescriptor) GetBech32AccountAddressPrefix() string {
	if m != nil {
		return m.Bech32AccountAddressPrefix
	}
	return ""
}


type MsgDescriptor struct {

	MsgTypeUrl string `protobuf:"bytes,1,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
}

func (m *MsgDescriptor) Reset()         { *m = MsgDescriptor{} }
func (m *MsgDescriptor) String() string { return proto.CompactTextString(m) }
func (*MsgDescriptor) ProtoMessage()    {}
func (*MsgDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{10}
}
func (m *MsgDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDescriptor.Merge(m, src)
}
func (m *MsgDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *MsgDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDescriptor proto.InternalMessageInfo

func (m *MsgDescriptor) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}


type GetAuthnDescriptorRequest struct {
}

func (m *GetAuthnDescriptorRequest) Reset()         { *m = GetAuthnDescriptorRequest{} }
func (m *GetAuthnDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthnDescriptorRequest) ProtoMessage()    {}
func (*GetAuthnDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{11}
}
func (m *GetAuthnDescriptorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthnDescriptorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthnDescriptorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthnDescriptorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthnDescriptorRequest.Merge(m, src)
}
func (m *GetAuthnDescriptorRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthnDescriptorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthnDescriptorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthnDescriptorRequest proto.InternalMessageInfo


type GetAuthnDescriptorResponse struct {

	Authn *AuthnDescriptor `protobuf:"bytes,1,opt,name=authn,proto3" json:"authn,omitempty"`
}

func (m *GetAuthnDescriptorResponse) Reset()         { *m = GetAuthnDescriptorResponse{} }
func (m *GetAuthnDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthnDescriptorResponse) ProtoMessage()    {}
func (*GetAuthnDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{12}
}
func (m *GetAuthnDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthnDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthnDescriptorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthnDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthnDescriptorResponse.Merge(m, src)
}
func (m *GetAuthnDescriptorResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthnDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthnDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthnDescriptorResponse proto.InternalMessageInfo

func (m *GetAuthnDescriptorResponse) GetAuthn() *AuthnDescriptor {
	if m != nil {
		return m.Authn
	}
	return nil
}


type GetChainDescriptorRequest struct {
}

func (m *GetChainDescriptorRequest) Reset()         { *m = GetChainDescriptorRequest{} }
func (m *GetChainDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*GetChainDescriptorRequest) ProtoMessage()    {}
func (*GetChainDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{13}
}
func (m *GetChainDescriptorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetChainDescriptorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetChainDescriptorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetChainDescriptorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChainDescriptorRequest.Merge(m, src)
}
func (m *GetChainDescriptorRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetChainDescriptorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChainDescriptorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChainDescriptorRequest proto.InternalMessageInfo


type GetChainDescriptorResponse struct {

	Chain *ChainDescriptor `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *GetChainDescriptorResponse) Reset()         { *m = GetChainDescriptorResponse{} }
func (m *GetChainDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*GetChainDescriptorResponse) ProtoMessage()    {}
func (*GetChainDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{14}
}
func (m *GetChainDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetChainDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetChainDescriptorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetChainDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChainDescriptorResponse.Merge(m, src)
}
func (m *GetChainDescriptorResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetChainDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChainDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetChainDescriptorResponse proto.InternalMessageInfo

func (m *GetChainDescriptorResponse) GetChain() *ChainDescriptor {
	if m != nil {
		return m.Chain
	}
	return nil
}


type GetCodecDescriptorRequest struct {
}

func (m *GetCodecDescriptorRequest) Reset()         { *m = GetCodecDescriptorRequest{} }
func (m *GetCodecDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*GetCodecDescriptorRequest) ProtoMessage()    {}
func (*GetCodecDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{15}
}
func (m *GetCodecDescriptorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCodecDescriptorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCodecDescriptorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCodecDescriptorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCodecDescriptorRequest.Merge(m, src)
}
func (m *GetCodecDescriptorRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetCodecDescriptorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCodecDescriptorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCodecDescriptorRequest proto.InternalMessageInfo


type GetCodecDescriptorResponse struct {

	Codec *CodecDescriptor `protobuf:"bytes,1,opt,name=codec,proto3" json:"codec,omitempty"`
}

func (m *GetCodecDescriptorResponse) Reset()         { *m = GetCodecDescriptorResponse{} }
func (m *GetCodecDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*GetCodecDescriptorResponse) ProtoMessage()    {}
func (*GetCodecDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{16}
}
func (m *GetCodecDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetCodecDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetCodecDescriptorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetCodecDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCodecDescriptorResponse.Merge(m, src)
}
func (m *GetCodecDescriptorResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetCodecDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCodecDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCodecDescriptorResponse proto.InternalMessageInfo

func (m *GetCodecDescriptorResponse) GetCodec() *CodecDescriptor {
	if m != nil {
		return m.Codec
	}
	return nil
}


type GetConfigurationDescriptorRequest struct {
}

func (m *GetConfigurationDescriptorRequest) Reset()         { *m = GetConfigurationDescriptorRequest{} }
func (m *GetConfigurationDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigurationDescriptorRequest) ProtoMessage()    {}
func (*GetConfigurationDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{17}
}
func (m *GetConfigurationDescriptorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetConfigurationDescriptorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetConfigurationDescriptorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetConfigurationDescriptorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigurationDescriptorRequest.Merge(m, src)
}
func (m *GetConfigurationDescriptorRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetConfigurationDescriptorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigurationDescriptorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigurationDescriptorRequest proto.InternalMessageInfo


type GetConfigurationDescriptorResponse struct {

	Config *ConfigurationDescriptor `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *GetConfigurationDescriptorResponse) Reset()         { *m = GetConfigurationDescriptorResponse{} }
func (m *GetConfigurationDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*GetConfigurationDescriptorResponse) ProtoMessage()    {}
func (*GetConfigurationDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{18}
}
func (m *GetConfigurationDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetConfigurationDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetConfigurationDescriptorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetConfigurationDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigurationDescriptorResponse.Merge(m, src)
}
func (m *GetConfigurationDescriptorResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetConfigurationDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigurationDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigurationDescriptorResponse proto.InternalMessageInfo

func (m *GetConfigurationDescriptorResponse) GetConfig() *ConfigurationDescriptor {
	if m != nil {
		return m.Config
	}
	return nil
}


type GetQueryServicesDescriptorRequest struct {
}

func (m *GetQueryServicesDescriptorRequest) Reset()         { *m = GetQueryServicesDescriptorRequest{} }
func (m *GetQueryServicesDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*GetQueryServicesDescriptorRequest) ProtoMessage()    {}
func (*GetQueryServicesDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{19}
}
func (m *GetQueryServicesDescriptorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetQueryServicesDescriptorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetQueryServicesDescriptorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetQueryServicesDescriptorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQueryServicesDescriptorRequest.Merge(m, src)
}
func (m *GetQueryServicesDescriptorRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetQueryServicesDescriptorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQueryServicesDescriptorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQueryServicesDescriptorRequest proto.InternalMessageInfo


type GetQueryServicesDescriptorResponse struct {

	Queries *QueryServicesDescriptor `protobuf:"bytes,1,opt,name=queries,proto3" json:"queries,omitempty"`
}

func (m *GetQueryServicesDescriptorResponse) Reset()         { *m = GetQueryServicesDescriptorResponse{} }
func (m *GetQueryServicesDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*GetQueryServicesDescriptorResponse) ProtoMessage()    {}
func (*GetQueryServicesDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{20}
}
func (m *GetQueryServicesDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetQueryServicesDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetQueryServicesDescriptorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetQueryServicesDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQueryServicesDescriptorResponse.Merge(m, src)
}
func (m *GetQueryServicesDescriptorResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetQueryServicesDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQueryServicesDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetQueryServicesDescriptorResponse proto.InternalMessageInfo

func (m *GetQueryServicesDescriptorResponse) GetQueries() *QueryServicesDescriptor {
	if m != nil {
		return m.Queries
	}
	return nil
}


type GetTxDescriptorRequest struct {
}

func (m *GetTxDescriptorRequest) Reset()         { *m = GetTxDescriptorRequest{} }
func (m *GetTxDescriptorRequest) String() string { return proto.CompactTextString(m) }
func (*GetTxDescriptorRequest) ProtoMessage()    {}
func (*GetTxDescriptorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{21}
}
func (m *GetTxDescriptorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTxDescriptorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTxDescriptorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTxDescriptorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxDescriptorRequest.Merge(m, src)
}
func (m *GetTxDescriptorRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetTxDescriptorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxDescriptorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxDescriptorRequest proto.InternalMessageInfo


type GetTxDescriptorResponse struct {


	Tx *TxDescriptor `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *GetTxDescriptorResponse) Reset()         { *m = GetTxDescriptorResponse{} }
func (m *GetTxDescriptorResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxDescriptorResponse) ProtoMessage()    {}
func (*GetTxDescriptorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{22}
}
func (m *GetTxDescriptorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetTxDescriptorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetTxDescriptorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetTxDescriptorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxDescriptorResponse.Merge(m, src)
}
func (m *GetTxDescriptorResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetTxDescriptorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxDescriptorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxDescriptorResponse proto.InternalMessageInfo

func (m *GetTxDescriptorResponse) GetTx() *TxDescriptor {
	if m != nil {
		return m.Tx
	}
	return nil
}


type QueryServicesDescriptor struct {

	QueryServices []*QueryServiceDescriptor `protobuf:"bytes,1,rep,name=query_services,json=queryServices,proto3" json:"query_services,omitempty"`
}

func (m *QueryServicesDescriptor) Reset()         { *m = QueryServicesDescriptor{} }
func (m *QueryServicesDescriptor) String() string { return proto.CompactTextString(m) }
func (*QueryServicesDescriptor) ProtoMessage()    {}
func (*QueryServicesDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{23}
}
func (m *QueryServicesDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServicesDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServicesDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServicesDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServicesDescriptor.Merge(m, src)
}
func (m *QueryServicesDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *QueryServicesDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServicesDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServicesDescriptor proto.InternalMessageInfo

func (m *QueryServicesDescriptor) GetQueryServices() []*QueryServiceDescriptor {
	if m != nil {
		return m.QueryServices
	}
	return nil
}


type QueryServiceDescriptor struct {

	Fullname string `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`

	IsModule bool `protobuf:"varint,2,opt,name=is_module,json=isModule,proto3" json:"is_module,omitempty"`

	Methods []*QueryMethodDescriptor `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (m *QueryServiceDescriptor) Reset()         { *m = QueryServiceDescriptor{} }
func (m *QueryServiceDescriptor) String() string { return proto.CompactTextString(m) }
func (*QueryServiceDescriptor) ProtoMessage()    {}
func (*QueryServiceDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{24}
}
func (m *QueryServiceDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceDescriptor.Merge(m, src)
}
func (m *QueryServiceDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceDescriptor proto.InternalMessageInfo

func (m *QueryServiceDescriptor) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *QueryServiceDescriptor) GetIsModule() bool {
	if m != nil {
		return m.IsModule
	}
	return false
}

func (m *QueryServiceDescriptor) GetMethods() []*QueryMethodDescriptor {
	if m != nil {
		return m.Methods
	}
	return nil
}




type QueryMethodDescriptor struct {

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`


	FullQueryPath string `protobuf:"bytes,2,opt,name=full_query_path,json=fullQueryPath,proto3" json:"full_query_path,omitempty"`
}

func (m *QueryMethodDescriptor) Reset()         { *m = QueryMethodDescriptor{} }
func (m *QueryMethodDescriptor) String() string { return proto.CompactTextString(m) }
func (*QueryMethodDescriptor) ProtoMessage()    {}
func (*QueryMethodDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_15c91f0b8d6bf3d0, []int{25}
}
func (m *QueryMethodDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMethodDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMethodDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMethodDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMethodDescriptor.Merge(m, src)
}
func (m *QueryMethodDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *QueryMethodDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMethodDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMethodDescriptor proto.InternalMessageInfo

func (m *QueryMethodDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueryMethodDescriptor) GetFullQueryPath() string {
	if m != nil {
		return m.FullQueryPath
	}
	return ""
}

func init() {
	proto.RegisterType((*AppDescriptor)(nil), "cosmos.base.reflection.v2alpha1.AppDescriptor")
	proto.RegisterType((*TxDescriptor)(nil), "cosmos.base.reflection.v2alpha1.TxDescriptor")
	proto.RegisterType((*AuthnDescriptor)(nil), "cosmos.base.reflection.v2alpha1.AuthnDescriptor")
	proto.RegisterType((*SigningModeDescriptor)(nil), "cosmos.base.reflection.v2alpha1.SigningModeDescriptor")
	proto.RegisterType((*ChainDescriptor)(nil), "cosmos.base.reflection.v2alpha1.ChainDescriptor")
	proto.RegisterType((*CodecDescriptor)(nil), "cosmos.base.reflection.v2alpha1.CodecDescriptor")
	proto.RegisterType((*InterfaceDescriptor)(nil), "cosmos.base.reflection.v2alpha1.InterfaceDescriptor")
	proto.RegisterType((*InterfaceImplementerDescriptor)(nil), "cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor")
	proto.RegisterType((*InterfaceAcceptingMessageDescriptor)(nil), "cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor")
	proto.RegisterType((*ConfigurationDescriptor)(nil), "cosmos.base.reflection.v2alpha1.ConfigurationDescriptor")
	proto.RegisterType((*MsgDescriptor)(nil), "cosmos.base.reflection.v2alpha1.MsgDescriptor")
	proto.RegisterType((*GetAuthnDescriptorRequest)(nil), "cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest")
	proto.RegisterType((*GetAuthnDescriptorResponse)(nil), "cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse")
	proto.RegisterType((*GetChainDescriptorRequest)(nil), "cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest")
	proto.RegisterType((*GetChainDescriptorResponse)(nil), "cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse")
	proto.RegisterType((*GetCodecDescriptorRequest)(nil), "cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest")
	proto.RegisterType((*GetCodecDescriptorResponse)(nil), "cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse")
	proto.RegisterType((*GetConfigurationDescriptorRequest)(nil), "cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest")
	proto.RegisterType((*GetConfigurationDescriptorResponse)(nil), "cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse")
	proto.RegisterType((*GetQueryServicesDescriptorRequest)(nil), "cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest")
	proto.RegisterType((*GetQueryServicesDescriptorResponse)(nil), "cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse")
	proto.RegisterType((*GetTxDescriptorRequest)(nil), "cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest")
	proto.RegisterType((*GetTxDescriptorResponse)(nil), "cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse")
	proto.RegisterType((*QueryServicesDescriptor)(nil), "cosmos.base.reflection.v2alpha1.QueryServicesDescriptor")
	proto.RegisterType((*QueryServiceDescriptor)(nil), "cosmos.base.reflection.v2alpha1.QueryServiceDescriptor")
	proto.RegisterType((*QueryMethodDescriptor)(nil), "cosmos.base.reflection.v2alpha1.QueryMethodDescriptor")
}

func init() {
	proto.RegisterFile("cosmos/base/reflection/v2alpha1/reflection.proto", fileDescriptor_15c91f0b8d6bf3d0)
}

var fileDescriptor_15c91f0b8d6bf3d0 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x58, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xce, 0x3a, 0xbf, 0x5f, 0x9b, 0x46, 0x0c, 0x24, 0x71, 0xdd, 0xe2, 0xa6, 0x1b, 0x09, 0xf5,
	0x52, 0xbb, 0x49, 0xa3, 0xb4, 0x82, 0x94, 0xca, 0x69, 0x68, 0x15, 0x89, 0xa0, 0xe0, 0xa4, 0x80,
	0x10, 0xea, 0x6a, 0xbd, 0x3b, 0x5e, 0x8f, 0xf0, 0xee, 0x6c, 0x76, 0xc6, 0xc1, 0xb9, 0x72, 0xe0,
	0x0c, 0xe2, 0x4f, 0xe0, 0xc0, 0x9d, 0xbf, 0x02, 0xc1, 0xa5, 0x12, 0x17, 0x8e, 0x28, 0x41, 0xe2,
	0x00, 0x7f, 0x04, 0x9a, 0x1f, 0x76, 0xc6, 0xce, 0xda, 0xde, 0x24, 0x3d, 0x25, 0xb3, 0xef, 0xbd,
	0x6f, 0xbe, 0xef, 0xed, 0xe8, 0x7d, 0xb3, 0x86, 0x07, 0x1e, 0x65, 0x21, 0x65, 0xe5, 0x9a, 0xcb,
	0x70, 0x39, 0xc1, 0xf5, 0x26, 0xf6, 0x38, 0xa1, 0x51, 0xf9, 0x68, 0xcd, 0x6d, 0xc6, 0x0d, 0x77,
	0xd5, 0x78, 0x56, 0x8a, 0x13, 0xca, 0x29, 0xba, 0xa3, 0x2a, 0x4a, 0xa2, 0xa2, 0x64, 0x44, 0x3b,
	0x15, 0x85, 0xdb, 0x01, 0xa5, 0x41, 0x13, 0x97, 0xdd, 0x98, 0x94, 0xdd, 0x28, 0xa2, 0xdc, 0x15,
	0x71, 0xa6, 0xca, 0xed, 0x7f, 0xc6, 0x61, 0xae, 0x12, 0xc7, 0xdb, 0x98, 0x79, 0x09, 0x89, 0x39,
	0x4d, 0xd0, 0x73, 0x98, 0x74, 0x5b, 0xbc, 0x11, 0xe5, 0xad, 0x65, 0xeb, 0xde, 0xb5, 0xb5, 0x07,
	0xa5, 0x11, 0x1b, 0x94, 0x2a, 0x22, 0xfb, 0x0c, 0xa0, 0xaa, 0xca, 0x05, 0x8e, 0xd7, 0x70, 0x49,
	0x94, 0xcf, 0x65, 0xc4, 0x79, 0x26, 0xb2, 0x4d, 0x1c, 0x59, 0x2e, 0x71, 0xa8, 0x8f, 0xbd, 0xfc,
	0x78, 0x56, 0x1c, 0x91, 0xdd, 0x83, 0x23, 0x1e, 0xa0, 0x57, 0x30, 0xe7, 0xd1, 0xa8, 0x4e, 0x82,
	0x56, 0x22, 0x3b, 0x90, 0x9f, 0x90, 0x78, 0x8f, 0x33, 0xe0, 0x19, 0x55, 0x06, 0x6e, 0x2f, 0x1c,
	0x72, 0xe0, 0xc6, 0x61, 0x0b, 0x27, 0xc7, 0x0e, 0xc3, 0xc9, 0x11, 0xf1, 0x30, 0xcb, 0x4f, 0x66,
	0xdc, 0xe0, 0x53, 0x51, 0xb6, 0xaf, 0xab, 0xcc, 0x0d, 0x0e, 0xcd, 0x00, 0x7a, 0x02, 0x39, 0xde,
	0xce, 0x4f, 0x49, 0xd0, 0xfb, 0x23, 0x41, 0x0f, 0xda, 0x06, 0x52, 0x8e, 0xb7, 0xed, 0x08, 0xae,
	0x9b, 0xcf, 0x50, 0x01, 0x66, 0xea, 0xad, 0x66, 0x33, 0x72, 0x43, 0x2c, 0x5f, 0xf5, 0x6c, 0xb5,
	0xbb, 0x46, 0x5b, 0x30, 0x11, 0xb2, 0x80, 0xe5, 0x73, 0xcb, 0xe3, 0xf7, 0xae, 0xad, 0x95, 0x46,
	0x6e, 0xb6, 0xcb, 0x02, 0x63, 0x37, 0x59, 0x6b, 0x37, 0x60, 0xbe, 0xef, 0x64, 0xa0, 0x97, 0x00,
	0x8c, 0x04, 0x91, 0x13, 0x52, 0x1f, 0xb3, 0xbc, 0x25, 0xc1, 0x37, 0x46, 0x82, 0xef, 0x93, 0x20,
	0x22, 0x51, 0xb0, 0x4b, 0x7d, 0x6c, 0x6c, 0x32, 0x2b, 0x90, 0xc4, 0x33, 0x66, 0xff, 0x60, 0xc1,
	0x42, 0x6a, 0x12, 0x42, 0x30, 0x61, 0xe8, 0x93, 0xff, 0xa3, 0x45, 0x98, 0x8a, 0x5a, 0x61, 0x0d,
	0x27, 0xf2, 0x60, 0x4e, 0x56, 0xf5, 0x0a, 0x7d, 0x0c, 0x2b, 0xf2, 0xe0, 0x3a, 0x24, 0xaa, 0x53,
	0x27, 0x4e, 0xe8, 0x11, 0xf1, 0x71, 0xe2, 0x84, 0x98, 0x37, 0xa8, 0xef, 0x74, 0x5b, 0x35, 0x2e,
	0xa1, 0xee, 0xc8, 0xd4, 0x9d, 0xa8, 0x4e, 0xf7, 0x74, 0xe2, 0xae, 0xcc, 0x7b, 0xae, 0xd3, 0xec,
	0xbb, 0x30, 0xdf, 0x77, 0x9e, 0xd1, 0x0d, 0xc8, 0x11, 0x5f, 0x53, 0xc9, 0x11, 0xdf, 0x0e, 0x60,
	0xbe, 0xef, 0xa8, 0xa2, 0x03, 0x00, 0x12, 0x71, 0x9c, 0xd4, 0x5d, 0xaf, 0xdb, 0xa0, 0xf5, 0x91,
	0x0d, 0xda, 0xe9, 0x94, 0x18, 0xed, 0x31, 0x70, 0xec, 0x5f, 0x72, 0xf0, 0x76, 0x4a, 0xce, 0xd0,
	0x13, 0xf0, 0x9d, 0x05, 0xb7, 0xbb, 0x10, 0x8e, 0xeb, 0x79, 0x38, 0xe6, 0x24, 0x0a, 0x9c, 0x10,
	0x33, 0xe6, 0x06, 0xb8, 0x73, 0x34, 0xb6, 0xb3, 0x93, 0xab, 0x74, 0x30, 0x76, 0x15, 0x84, 0x41,
	0xb6, 0x40, 0x06, 0x25, 0x31, 0x74, 0x04, 0x8b, 0x67, 0x3c, 0x48, 0x18, 0x37, 0x71, 0x88, 0xc5,
	0x9a, 0xe5, 0xc7, 0x25, 0x83, 0xa7, 0xd9, 0x19, 0xec, 0x9c, 0x55, 0x1b, 0x9b, 0x2f, 0x90, 0x94,
	0x38, 0xb3, 0x3f, 0x87, 0xe2, 0xf0, 0xc2, 0xa1, 0xed, 0xbb, 0x09, 0x33, 0xfc, 0x38, 0xc6, 0x4e,
	0x2b, 0x69, 0xca, 0x63, 0x36, 0x5b, 0x9d, 0x16, 0xeb, 0x97, 0x49, 0xd3, 0xfe, 0x06, 0x56, 0x32,
	0xf4, 0x64, 0x28, 0xfa, 0x3a, 0x2c, 0xd6, 0x09, 0x6e, 0xfa, 0x8e, 0xdf, 0xcd, 0x77, 0x44, 0x40,
	0xbd, 0x95, 0xd9, 0xea, 0x3b, 0x32, 0x7a, 0x06, 0xf6, 0x89, 0x88, 0xd9, 0x5f, 0xc1, 0xd2, 0x80,
	0x51, 0x86, 0x2a, 0xf0, 0x6e, 0x0d, 0x7b, 0x8d, 0x87, 0x6b, 0xe2, 0x4d, 0xd3, 0x56, 0xc4, 0x1d,
	0xd7, 0xf7, 0x13, 0xcc, 0x98, 0x13, 0x27, 0xb8, 0x4e, 0xda, 0x9a, 0x41, 0x41, 0x25, 0x55, 0x54,
	0x4e, 0x45, 0xa5, 0xec, 0xc9, 0x0c, 0x7b, 0x15, 0xe6, 0x7a, 0xa6, 0x00, 0x5a, 0x86, 0xeb, 0x21,
	0x0b, 0x9c, 0x6e, 0x1b, 0x14, 0x04, 0x84, 0x2c, 0x38, 0xd0, 0x9d, 0xb8, 0x05, 0x37, 0x5f, 0x60,
	0xde, 0x6f, 0x1f, 0xf8, 0xb0, 0x85, 0x19, 0xb7, 0x7d, 0x28, 0xa4, 0x05, 0x59, 0x4c, 0x23, 0x86,
	0xdf, 0x94, 0x49, 0x69, 0x0a, 0xfd, 0xce, 0xd3, 0x43, 0xe1, 0x5c, 0xf0, 0x8c, 0x82, 0xf2, 0x37,
	0xeb, 0x4a, 0xfe, 0xd6, 0xa1, 0xd0, 0x67, 0x5a, 0xbd, 0x14, 0xfa, 0x83, 0x06, 0x05, 0x69, 0x8d,
	0xd6, 0x95, 0xac, 0xd1, 0x5e, 0x81, 0xbb, 0x72, 0x97, 0x74, 0x9f, 0xd3, 0x54, 0x8e, 0xc0, 0x1e,
	0x96, 0xa4, 0x29, 0xed, 0xc1, 0x94, 0xb2, 0x45, 0xcd, 0xe9, 0xf2, 0xf6, 0xaa, 0x71, 0x34, 0xb9,
	0x41, 0x1e, 0xa9, 0xc9, 0xb5, 0x25, 0xb9, 0x81, 0x49, 0x9a, 0x5c, 0x15, 0xa6, 0x85, 0xa5, 0x12,
	0x39, 0x5b, 0xaf, 0xe6, 0xcd, 0x1d, 0x20, 0x3b, 0x0f, 0x8b, 0x2f, 0x30, 0xef, 0x71, 0x5b, 0xcd,
	0xe9, 0x0b, 0x58, 0x3a, 0x17, 0xd1, 0x44, 0x94, 0x95, 0x5b, 0x97, 0xb5, 0xf2, 0x63, 0x58, 0x1a,
	0xc0, 0x0b, 0xbd, 0x3a, 0x77, 0x0b, 0x51, 0x2e, 0xf2, 0xe8, 0x42, 0x4a, 0x07, 0x5e, 0x42, 0xec,
	0x9f, 0x2c, 0x58, 0x4c, 0xcf, 0x1c, 0x3a, 0xb1, 0x6e, 0xc1, 0x2c, 0x61, 0xc2, 0xf7, 0x5b, 0x4d,
	0x2c, 0x07, 0xe2, 0x4c, 0x75, 0x86, 0xb0, 0x5d, 0xb9, 0x46, 0x7b, 0x30, 0xad, 0x5c, 0xb6, 0x33,
	0xd3, 0x37, 0xb2, 0x91, 0x55, 0x96, 0x6b, 0xbe, 0x14, 0x0d, 0x63, 0xef, 0xc3, 0x42, 0x6a, 0x46,
	0xea, 0x85, 0xe0, 0x3d, 0x98, 0x17, 0x3c, 0x1d, 0xd5, 0xb7, 0xd8, 0xe5, 0x0d, 0x3d, 0xb2, 0xe7,
	0xc4, 0x63, 0x89, 0xb3, 0xe7, 0xf2, 0xc6, 0xda, 0xcf, 0x00, 0x6f, 0x55, 0xbb, 0x5c, 0xb4, 0x7e,
	0xf4, 0xbb, 0x05, 0xe8, 0xfc, 0xa0, 0x42, 0xef, 0x8f, 0x94, 0x30, 0x70, 0xf4, 0x15, 0x3e, 0xb8,
	0x54, 0xad, 0x3a, 0x5a, 0xf6, 0xe6, 0xb7, 0x7f, 0xfc, 0xfd, 0x63, 0x6e, 0x03, 0xad, 0x97, 0x07,
	0x7d, 0x4a, 0xac, 0xd6, 0x30, 0x77, 0x57, 0xcb, 0x6e, 0x1c, 0x1b, 0xfe, 0x51, 0x56, 0x97, 0x76,
	0xad, 0xa6, 0xff, 0xea, 0x92, 0x49, 0x4d, 0xfa, 0x14, 0xcd, 0xa6, 0x66, 0xc0, 0x90, 0xbd, 0xb4,
	0x1a, 0xf5, 0xe9, 0xd0, 0x51, 0xd3, 0x77, 0xcb, 0xca, 0xa6, 0x26, 0x75, 0x20, 0x67, 0x54, 0x93,
	0x3e, 0xaf, 0x2f, 0xaf, 0x46, 0x7e, 0xc0, 0xfc, 0x6b, 0x69, 0x33, 0x48, 0xf7, 0xf0, 0xad, 0x6c,
	0xcc, 0x86, 0xcd, 0xf8, 0xc2, 0xb3, 0x2b, 0x61, 0x68, 0x95, 0xdb, 0x52, 0xe5, 0x87, 0x68, 0xf3,
	0xc2, 0x2a, 0xcd, 0xcf, 0xa9, 0xff, 0x94, 0xda, 0x41, 0x73, 0x2e, 0x93, 0xda, 0xe1, 0xa6, 0x91,
	0x4d, 0xed, 0x08, 0x4f, 0xb1, 0x3f, 0x92, 0x6a, 0x9f, 0xa2, 0x27, 0x17, 0x54, 0xdb, 0x3b, 0xa5,
	0xd1, 0x6f, 0x16, 0xcc, 0xf7, 0xb9, 0x05, 0x7a, 0x94, 0x85, 0x5f, 0x8a, 0xf3, 0x14, 0x1e, 0x5f,
	0xbc, 0xf0, 0x8a, 0xef, 0x8e, 0xb7, 0x8d, 0xd5, 0xd6, 0x67, 0xbf, 0x9e, 0x14, 0xad, 0xd7, 0x27,
	0x45, 0xeb, 0xaf, 0x93, 0xa2, 0xf5, 0xfd, 0x69, 0x71, 0xec, 0xf5, 0x69, 0x71, 0xec, 0xcf, 0xd3,
	0xe2, 0xd8, 0x97, 0x9b, 0x01, 0xe1, 0x8d, 0x56, 0xad, 0xe4, 0xd1, 0xb0, 0xb3, 0x83, 0xfa, 0x73,
	0x9f, 0xf9, 0x5f, 0x97, 0x45, 0x37, 0x70, 0x52, 0x0e, 0x92, 0xd8, 0x4b, 0xfb, 0xf1, 0xa3, 0x36,
	0x25, 0x7f, 0xb3, 0x78, 0xf8, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xdb, 0xec, 0xac, 0x26,
	0x11, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type ReflectionServiceClient interface {



	GetAuthnDescriptor(ctx context.Context, in *GetAuthnDescriptorRequest, opts ...grpc.CallOption) (*GetAuthnDescriptorResponse, error)

	GetChainDescriptor(ctx context.Context, in *GetChainDescriptorRequest, opts ...grpc.CallOption) (*GetChainDescriptorResponse, error)

	GetCodecDescriptor(ctx context.Context, in *GetCodecDescriptorRequest, opts ...grpc.CallOption) (*GetCodecDescriptorResponse, error)

	GetConfigurationDescriptor(ctx context.Context, in *GetConfigurationDescriptorRequest, opts ...grpc.CallOption) (*GetConfigurationDescriptorResponse, error)

	GetQueryServicesDescriptor(ctx context.Context, in *GetQueryServicesDescriptorRequest, opts ...grpc.CallOption) (*GetQueryServicesDescriptorResponse, error)

	GetTxDescriptor(ctx context.Context, in *GetTxDescriptorRequest, opts ...grpc.CallOption) (*GetTxDescriptorResponse, error)
}

type reflectionServiceClient struct {
	cc grpc1.ClientConn
}

func NewReflectionServiceClient(cc grpc1.ClientConn) ReflectionServiceClient {
	return &reflectionServiceClient{cc}
}

func (c *reflectionServiceClient) GetAuthnDescriptor(ctx context.Context, in *GetAuthnDescriptorRequest, opts ...grpc.CallOption) (*GetAuthnDescriptorResponse, error) {
	out := new(GetAuthnDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetAuthnDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetChainDescriptor(ctx context.Context, in *GetChainDescriptorRequest, opts ...grpc.CallOption) (*GetChainDescriptorResponse, error) {
	out := new(GetChainDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetChainDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetCodecDescriptor(ctx context.Context, in *GetCodecDescriptorRequest, opts ...grpc.CallOption) (*GetCodecDescriptorResponse, error) {
	out := new(GetCodecDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetCodecDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetConfigurationDescriptor(ctx context.Context, in *GetConfigurationDescriptorRequest, opts ...grpc.CallOption) (*GetConfigurationDescriptorResponse, error) {
	out := new(GetConfigurationDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetConfigurationDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetQueryServicesDescriptor(ctx context.Context, in *GetQueryServicesDescriptorRequest, opts ...grpc.CallOption) (*GetQueryServicesDescriptorResponse, error) {
	out := new(GetQueryServicesDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetQueryServicesDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) GetTxDescriptor(ctx context.Context, in *GetTxDescriptorRequest, opts ...grpc.CallOption) (*GetTxDescriptorResponse, error) {
	out := new(GetTxDescriptorResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v2alpha1.ReflectionService/GetTxDescriptor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type ReflectionServiceServer interface {



	GetAuthnDescriptor(context.Context, *GetAuthnDescriptorRequest) (*GetAuthnDescriptorResponse, error)

	GetChainDescriptor(context.Context, *GetChainDescriptorRequest) (*GetChainDescriptorResponse, error)

	GetCodecDescriptor(context.Context, *GetCodecDescriptorRequest) (*GetCodecDescriptorResponse, error)

	GetConfigurationDescriptor(context.Context, *GetConfigurationDescriptorRequest) (*GetConfigurationDescriptorResponse, error)

	GetQueryServicesDescriptor(context.Context, *GetQueryServicesDescriptorRequest) (*GetQueryServicesDescriptorResponse, error)

	GetTxDescriptor(context.Context, *GetTxDescriptorRequest) (*GetTxDescriptorResponse, error)
}


type UnimplementedReflectionServiceServer struct {
}

func (*UnimplementedReflectionServiceServer) GetAuthnDescriptor(ctx context.Context, req *GetAuthnDescriptorRequest) (*GetAuthnDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthnDescriptor not implemented")
}
func (*UnimplementedReflectionServiceServer) GetChainDescriptor(ctx context.Context, req *GetChainDescriptorRequest) (*GetChainDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainDescriptor not implemented")
}
func (*UnimplementedReflectionServiceServer) GetCodecDescriptor(ctx context.Context, req *GetCodecDescriptorRequest) (*GetCodecDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodecDescriptor not implemented")
}
func (*UnimplementedReflectionServiceServer) GetConfigurationDescriptor(ctx context.Context, req *GetConfigurationDescriptorRequest) (*GetConfigurationDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigurationDescriptor not implemented")
}
func (*UnimplementedReflectionServiceServer) GetQueryServicesDescriptor(ctx context.Context, req *GetQueryServicesDescriptorRequest) (*GetQueryServicesDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueryServicesDescriptor not implemented")
}
func (*UnimplementedReflectionServiceServer) GetTxDescriptor(ctx context.Context, req *GetTxDescriptorRequest) (*GetTxDescriptorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxDescriptor not implemented")
}

func RegisterReflectionServiceServer(s grpc1.Server, srv ReflectionServiceServer) {
	s.RegisterService(&_ReflectionService_serviceDesc, srv)
}

func _ReflectionService_GetAuthnDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthnDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetAuthnDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetAuthnDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetAuthnDescriptor(ctx, req.(*GetAuthnDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetChainDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChainDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetChainDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetChainDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetChainDescriptor(ctx, req.(*GetChainDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetCodecDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodecDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetCodecDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetCodecDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetCodecDescriptor(ctx, req.(*GetCodecDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetConfigurationDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetConfigurationDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetConfigurationDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetConfigurationDescriptor(ctx, req.(*GetConfigurationDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetQueryServicesDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryServicesDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetQueryServicesDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetQueryServicesDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetQueryServicesDescriptor(ctx, req.(*GetQueryServicesDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_GetTxDescriptor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxDescriptorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).GetTxDescriptor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v2alpha1.ReflectionService/GetTxDescriptor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).GetTxDescriptor(ctx, req.(*GetTxDescriptorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReflectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.reflection.v2alpha1.ReflectionService",
	HandlerType: (*ReflectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthnDescriptor",
			Handler:    _ReflectionService_GetAuthnDescriptor_Handler,
		},
		{
			MethodName: "GetChainDescriptor",
			Handler:    _ReflectionService_GetChainDescriptor_Handler,
		},
		{
			MethodName: "GetCodecDescriptor",
			Handler:    _ReflectionService_GetCodecDescriptor_Handler,
		},
		{
			MethodName: "GetConfigurationDescriptor",
			Handler:    _ReflectionService_GetConfigurationDescriptor_Handler,
		},
		{
			MethodName: "GetQueryServicesDescriptor",
			Handler:    _ReflectionService_GetQueryServicesDescriptor_Handler,
		},
		{
			MethodName: "GetTxDescriptor",
			Handler:    _ReflectionService_GetTxDescriptor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/reflection/v2alpha1/reflection.proto",
}

func (m *AppDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.QueryServices != nil {
		{
			size, err := m.QueryServices.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Configuration != nil {
		{
			size, err := m.Configuration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Codec != nil {
		{
			size, err := m.Codec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Chain != nil {
		{
			size, err := m.Chain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Authn != nil {
		{
			size, err := m.Authn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReflection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Fullname) > 0 {
		i -= len(m.Fullname)
		copy(dAtA[i:], m.Fullname)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Fullname)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthnDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthnDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthnDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SignModes) > 0 {
		for iNdEx := len(m.SignModes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SignModes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReflection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SigningModeDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigningModeDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SigningModeDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthnInfoProviderMethodFullname) > 0 {
		i -= len(m.AuthnInfoProviderMethodFullname)
		copy(dAtA[i:], m.AuthnInfoProviderMethodFullname)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.AuthnInfoProviderMethodFullname)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Number != 0 {
		i = encodeVarintReflection(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChainDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CodecDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodecDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodecDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Interfaces) > 0 {
		for iNdEx := len(m.Interfaces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Interfaces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReflection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *InterfaceDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterfaceDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterfaceDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InterfaceImplementers) > 0 {
		for iNdEx := len(m.InterfaceImplementers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InterfaceImplementers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReflection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.InterfaceAcceptingMessages) > 0 {
		for iNdEx := len(m.InterfaceAcceptingMessages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InterfaceAcceptingMessages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReflection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Fullname) > 0 {
		i -= len(m.Fullname)
		copy(dAtA[i:], m.Fullname)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Fullname)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InterfaceImplementerDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterfaceImplementerDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterfaceImplementerDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TypeUrl) > 0 {
		i -= len(m.TypeUrl)
		copy(dAtA[i:], m.TypeUrl)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.TypeUrl)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Fullname) > 0 {
		i -= len(m.Fullname)
		copy(dAtA[i:], m.Fullname)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Fullname)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InterfaceAcceptingMessageDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterfaceAcceptingMessageDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterfaceAcceptingMessageDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FieldDescriptorNames) > 0 {
		for iNdEx := len(m.FieldDescriptorNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FieldDescriptorNames[iNdEx])
			copy(dAtA[i:], m.FieldDescriptorNames[iNdEx])
			i = encodeVarintReflection(dAtA, i, uint64(len(m.FieldDescriptorNames[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Fullname) > 0 {
		i -= len(m.Fullname)
		copy(dAtA[i:], m.Fullname)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Fullname)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfigurationDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigurationDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigurationDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bech32AccountAddressPrefix) > 0 {
		i -= len(m.Bech32AccountAddressPrefix)
		copy(dAtA[i:], m.Bech32AccountAddressPrefix)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Bech32AccountAddressPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAuthnDescriptorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthnDescriptorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthnDescriptorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetAuthnDescriptorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthnDescriptorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthnDescriptorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Authn != nil {
		{
			size, err := m.Authn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetChainDescriptorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetChainDescriptorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetChainDescriptorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetChainDescriptorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetChainDescriptorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetChainDescriptorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Chain != nil {
		{
			size, err := m.Chain.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetCodecDescriptorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCodecDescriptorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetCodecDescriptorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetCodecDescriptorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetCodecDescriptorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetCodecDescriptorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Codec != nil {
		{
			size, err := m.Codec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetConfigurationDescriptorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetConfigurationDescriptorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetConfigurationDescriptorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetConfigurationDescriptorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetConfigurationDescriptorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetConfigurationDescriptorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetQueryServicesDescriptorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetQueryServicesDescriptorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetQueryServicesDescriptorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetQueryServicesDescriptorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetQueryServicesDescriptorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetQueryServicesDescriptorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Queries != nil {
		{
			size, err := m.Queries.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetTxDescriptorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTxDescriptorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTxDescriptorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetTxDescriptorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetTxDescriptorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetTxDescriptorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintReflection(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryServicesDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServicesDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServicesDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueryServices) > 0 {
		for iNdEx := len(m.QueryServices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.QueryServices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReflection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryServiceDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for iNdEx := len(m.Methods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Methods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReflection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.IsModule {
		i--
		if m.IsModule {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Fullname) > 0 {
		i -= len(m.Fullname)
		copy(dAtA[i:], m.Fullname)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Fullname)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryMethodDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMethodDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMethodDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FullQueryPath) > 0 {
		i -= len(m.FullQueryPath)
		copy(dAtA[i:], m.FullQueryPath)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.FullQueryPath)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintReflection(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReflection(dAtA []byte, offset int, v uint64) int {
	offset -= sovReflection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AppDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Authn != nil {
		l = m.Authn.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	if m.Chain != nil {
		l = m.Chain.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	if m.Codec != nil {
		l = m.Codec.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	if m.Configuration != nil {
		l = m.Configuration.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	if m.QueryServices != nil {
		l = m.QueryServices.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *TxDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fullname)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func (m *AuthnDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SignModes) > 0 {
		for _, e := range m.SignModes {
			l = e.Size()
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func (m *SigningModeDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	if m.Number != 0 {
		n += 1 + sovReflection(uint64(m.Number))
	}
	l = len(m.AuthnInfoProviderMethodFullname)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *ChainDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *CodecDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Interfaces) > 0 {
		for _, e := range m.Interfaces {
			l = e.Size()
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func (m *InterfaceDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fullname)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	if len(m.InterfaceAcceptingMessages) > 0 {
		for _, e := range m.InterfaceAcceptingMessages {
			l = e.Size()
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	if len(m.InterfaceImplementers) > 0 {
		for _, e := range m.InterfaceImplementers {
			l = e.Size()
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func (m *InterfaceImplementerDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fullname)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *InterfaceAcceptingMessageDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fullname)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	if len(m.FieldDescriptorNames) > 0 {
		for _, s := range m.FieldDescriptorNames {
			l = len(s)
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func (m *ConfigurationDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bech32AccountAddressPrefix)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *MsgDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *GetAuthnDescriptorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetAuthnDescriptorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Authn != nil {
		l = m.Authn.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *GetChainDescriptorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetChainDescriptorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chain != nil {
		l = m.Chain.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *GetCodecDescriptorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetCodecDescriptorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Codec != nil {
		l = m.Codec.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *GetConfigurationDescriptorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetConfigurationDescriptorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *GetQueryServicesDescriptorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetQueryServicesDescriptorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Queries != nil {
		l = m.Queries.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *GetTxDescriptorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetTxDescriptorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func (m *QueryServicesDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.QueryServices) > 0 {
		for _, e := range m.QueryServices {
			l = e.Size()
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func (m *QueryServiceDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Fullname)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	if m.IsModule {
		n += 2
	}
	if len(m.Methods) > 0 {
		for _, e := range m.Methods {
			l = e.Size()
			n += 1 + l + sovReflection(uint64(l))
		}
	}
	return n
}

func (m *QueryMethodDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	l = len(m.FullQueryPath)
	if l > 0 {
		n += 1 + l + sovReflection(uint64(l))
	}
	return n
}

func sovReflection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReflection(x uint64) (n int) {
	return sovReflection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: AppDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authn == nil {
				m.Authn = &AuthnDescriptor{}
			}
			if err := m.Authn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Chain == nil {
				m.Chain = &ChainDescriptor{}
			}
			if err := m.Chain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Codec == nil {
				m.Codec = &CodecDescriptor{}
			}
			if err := m.Codec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Configuration == nil {
				m.Configuration = &ConfigurationDescriptor{}
			}
			if err := m.Configuration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryServices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.QueryServices == nil {
				m.QueryServices = &QueryServicesDescriptor{}
			}
			if err := m.QueryServices.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &TxDescriptor{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *TxDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: TxDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fullname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &MsgDescriptor{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *AuthnDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: AuthnDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthnDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignModes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignModes = append(m.SignModes, &SigningModeDescriptor{})
			if err := m.SignModes[len(m.SignModes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *SigningModeDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: SigningModeDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigningModeDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthnInfoProviderMethodFullname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthnInfoProviderMethodFullname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *ChainDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: ChainDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *CodecDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: CodecDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodecDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interfaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Interfaces = append(m.Interfaces, &InterfaceDescriptor{})
			if err := m.Interfaces[len(m.Interfaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *InterfaceDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: InterfaceDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterfaceDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fullname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterfaceAcceptingMessages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterfaceAcceptingMessages = append(m.InterfaceAcceptingMessages, &InterfaceAcceptingMessageDescriptor{})
			if err := m.InterfaceAcceptingMessages[len(m.InterfaceAcceptingMessages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterfaceImplementers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterfaceImplementers = append(m.InterfaceImplementers, &InterfaceImplementerDescriptor{})
			if err := m.InterfaceImplementers[len(m.InterfaceImplementers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *InterfaceImplementerDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: InterfaceImplementerDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterfaceImplementerDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fullname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *InterfaceAcceptingMessageDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: InterfaceAcceptingMessageDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterfaceAcceptingMessageDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fullname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldDescriptorNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldDescriptorNames = append(m.FieldDescriptorNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *ConfigurationDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: ConfigurationDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigurationDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32AccountAddressPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32AccountAddressPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *MsgDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: MsgDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetAuthnDescriptorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetAuthnDescriptorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthnDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetAuthnDescriptorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetAuthnDescriptorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthnDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authn == nil {
				m.Authn = &AuthnDescriptor{}
			}
			if err := m.Authn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetChainDescriptorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetChainDescriptorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetChainDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetChainDescriptorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetChainDescriptorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetChainDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Chain == nil {
				m.Chain = &ChainDescriptor{}
			}
			if err := m.Chain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetCodecDescriptorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetCodecDescriptorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCodecDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetCodecDescriptorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetCodecDescriptorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetCodecDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Codec == nil {
				m.Codec = &CodecDescriptor{}
			}
			if err := m.Codec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetConfigurationDescriptorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetConfigurationDescriptorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetConfigurationDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetConfigurationDescriptorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetConfigurationDescriptorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetConfigurationDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &ConfigurationDescriptor{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetQueryServicesDescriptorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetQueryServicesDescriptorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetQueryServicesDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetQueryServicesDescriptorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetQueryServicesDescriptorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetQueryServicesDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Queries == nil {
				m.Queries = &QueryServicesDescriptor{}
			}
			if err := m.Queries.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetTxDescriptorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetTxDescriptorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTxDescriptorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *GetTxDescriptorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: GetTxDescriptorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetTxDescriptorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &TxDescriptor{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *QueryServicesDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: QueryServicesDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServicesDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryServices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryServices = append(m.QueryServices, &QueryServiceDescriptor{})
			if err := m.QueryServices[len(m.QueryServices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *QueryServiceDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: QueryServiceDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fullname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fullname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsModule", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
			m.IsModule = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Methods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Methods = append(m.Methods, &QueryMethodDescriptor{})
			if err := m.Methods[len(m.Methods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func (m *QueryMethodDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReflection
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
			return fmt.Errorf("proto: QueryMethodDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMethodDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullQueryPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReflection
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
				return ErrInvalidLengthReflection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReflection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FullQueryPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReflection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReflection
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
func skipReflection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReflection
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
					return 0, ErrIntOverflowReflection
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
					return 0, ErrIntOverflowReflection
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
				return 0, ErrInvalidLengthReflection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReflection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReflection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReflection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReflection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReflection = fmt.Errorf("proto: unexpected end of group")
)
