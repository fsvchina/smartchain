


package types

import (
	fmt "fmt"
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


type Snapshot struct {
	Height   uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Format   uint32   `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"`
	Chunks   uint32   `protobuf:"varint,3,opt,name=chunks,proto3" json:"chunks,omitempty"`
	Hash     []byte   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	Metadata Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{0}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Snapshot) GetFormat() uint32 {
	if m != nil {
		return m.Format
	}
	return 0
}

func (m *Snapshot) GetChunks() uint32 {
	if m != nil {
		return m.Chunks
	}
	return 0
}

func (m *Snapshot) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Snapshot) GetMetadata() Metadata {
	if m != nil {
		return m.Metadata
	}
	return Metadata{}
}


type Metadata struct {
	ChunkHashes [][]byte `protobuf:"bytes,1,rep,name=chunk_hashes,json=chunkHashes,proto3" json:"chunk_hashes,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{1}
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

func (m *Metadata) GetChunkHashes() [][]byte {
	if m != nil {
		return m.ChunkHashes
	}
	return nil
}


type SnapshotItem struct {

	//





	Item isSnapshotItem_Item `protobuf_oneof:"item"`
}

func (m *SnapshotItem) Reset()         { *m = SnapshotItem{} }
func (m *SnapshotItem) String() string { return proto.CompactTextString(m) }
func (*SnapshotItem) ProtoMessage()    {}
func (*SnapshotItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{2}
}
func (m *SnapshotItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotItem.Merge(m, src)
}
func (m *SnapshotItem) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotItem.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotItem proto.InternalMessageInfo

type isSnapshotItem_Item interface {
	isSnapshotItem_Item()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SnapshotItem_Store struct {
	Store *SnapshotStoreItem `protobuf:"bytes,1,opt,name=store,proto3,oneof" json:"store,omitempty"`
}
type SnapshotItem_IAVL struct {
	IAVL *SnapshotIAVLItem `protobuf:"bytes,2,opt,name=iavl,proto3,oneof" json:"iavl,omitempty"`
}
type SnapshotItem_Extension struct {
	Extension *SnapshotExtensionMeta `protobuf:"bytes,3,opt,name=extension,proto3,oneof" json:"extension,omitempty"`
}
type SnapshotItem_ExtensionPayload struct {
	ExtensionPayload *SnapshotExtensionPayload `protobuf:"bytes,4,opt,name=extension_payload,json=extensionPayload,proto3,oneof" json:"extension_payload,omitempty"`
}

func (*SnapshotItem_Store) isSnapshotItem_Item()            {}
func (*SnapshotItem_IAVL) isSnapshotItem_Item()             {}
func (*SnapshotItem_Extension) isSnapshotItem_Item()        {}
func (*SnapshotItem_ExtensionPayload) isSnapshotItem_Item() {}

func (m *SnapshotItem) GetItem() isSnapshotItem_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *SnapshotItem) GetStore() *SnapshotStoreItem {
	if x, ok := m.GetItem().(*SnapshotItem_Store); ok {
		return x.Store
	}
	return nil
}

func (m *SnapshotItem) GetIAVL() *SnapshotIAVLItem {
	if x, ok := m.GetItem().(*SnapshotItem_IAVL); ok {
		return x.IAVL
	}
	return nil
}

func (m *SnapshotItem) GetExtension() *SnapshotExtensionMeta {
	if x, ok := m.GetItem().(*SnapshotItem_Extension); ok {
		return x.Extension
	}
	return nil
}

func (m *SnapshotItem) GetExtensionPayload() *SnapshotExtensionPayload {
	if x, ok := m.GetItem().(*SnapshotItem_ExtensionPayload); ok {
		return x.ExtensionPayload
	}
	return nil
}


func (*SnapshotItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SnapshotItem_Store)(nil),
		(*SnapshotItem_IAVL)(nil),
		(*SnapshotItem_Extension)(nil),
		(*SnapshotItem_ExtensionPayload)(nil),
	}
}


type SnapshotStoreItem struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *SnapshotStoreItem) Reset()         { *m = SnapshotStoreItem{} }
func (m *SnapshotStoreItem) String() string { return proto.CompactTextString(m) }
func (*SnapshotStoreItem) ProtoMessage()    {}
func (*SnapshotStoreItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{3}
}
func (m *SnapshotStoreItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotStoreItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotStoreItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotStoreItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotStoreItem.Merge(m, src)
}
func (m *SnapshotStoreItem) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotStoreItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotStoreItem.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotStoreItem proto.InternalMessageInfo

func (m *SnapshotStoreItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}


type SnapshotIAVLItem struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`

	Version int64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`

	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *SnapshotIAVLItem) Reset()         { *m = SnapshotIAVLItem{} }
func (m *SnapshotIAVLItem) String() string { return proto.CompactTextString(m) }
func (*SnapshotIAVLItem) ProtoMessage()    {}
func (*SnapshotIAVLItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{4}
}
func (m *SnapshotIAVLItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotIAVLItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotIAVLItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotIAVLItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotIAVLItem.Merge(m, src)
}
func (m *SnapshotIAVLItem) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotIAVLItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotIAVLItem.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotIAVLItem proto.InternalMessageInfo

func (m *SnapshotIAVLItem) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SnapshotIAVLItem) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SnapshotIAVLItem) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SnapshotIAVLItem) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}


type SnapshotExtensionMeta struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Format uint32 `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"`
}

func (m *SnapshotExtensionMeta) Reset()         { *m = SnapshotExtensionMeta{} }
func (m *SnapshotExtensionMeta) String() string { return proto.CompactTextString(m) }
func (*SnapshotExtensionMeta) ProtoMessage()    {}
func (*SnapshotExtensionMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{5}
}
func (m *SnapshotExtensionMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotExtensionMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotExtensionMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotExtensionMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotExtensionMeta.Merge(m, src)
}
func (m *SnapshotExtensionMeta) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotExtensionMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotExtensionMeta.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotExtensionMeta proto.InternalMessageInfo

func (m *SnapshotExtensionMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SnapshotExtensionMeta) GetFormat() uint32 {
	if m != nil {
		return m.Format
	}
	return 0
}


type SnapshotExtensionPayload struct {
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *SnapshotExtensionPayload) Reset()         { *m = SnapshotExtensionPayload{} }
func (m *SnapshotExtensionPayload) String() string { return proto.CompactTextString(m) }
func (*SnapshotExtensionPayload) ProtoMessage()    {}
func (*SnapshotExtensionPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7a3c9b0a19e1ee, []int{6}
}
func (m *SnapshotExtensionPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotExtensionPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotExtensionPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotExtensionPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotExtensionPayload.Merge(m, src)
}
func (m *SnapshotExtensionPayload) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotExtensionPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotExtensionPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotExtensionPayload proto.InternalMessageInfo

func (m *SnapshotExtensionPayload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Snapshot)(nil), "cosmos.base.snapshots.v1beta1.Snapshot")
	proto.RegisterType((*Metadata)(nil), "cosmos.base.snapshots.v1beta1.Metadata")
	proto.RegisterType((*SnapshotItem)(nil), "cosmos.base.snapshots.v1beta1.SnapshotItem")
	proto.RegisterType((*SnapshotStoreItem)(nil), "cosmos.base.snapshots.v1beta1.SnapshotStoreItem")
	proto.RegisterType((*SnapshotIAVLItem)(nil), "cosmos.base.snapshots.v1beta1.SnapshotIAVLItem")
	proto.RegisterType((*SnapshotExtensionMeta)(nil), "cosmos.base.snapshots.v1beta1.SnapshotExtensionMeta")
	proto.RegisterType((*SnapshotExtensionPayload)(nil), "cosmos.base.snapshots.v1beta1.SnapshotExtensionPayload")
}

func init() {
	proto.RegisterFile("cosmos/base/snapshots/v1beta1/snapshot.proto", fileDescriptor_dd7a3c9b0a19e1ee)
}

var fileDescriptor_dd7a3c9b0a19e1ee = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x8d, 0xd7, 0xb4, 0x74, 0x37, 0x41, 0xea, 0xac, 0x81, 0x22, 0x24, 0xb2, 0x92, 0x97, 0xf5,
	0x61, 0x4b, 0x58, 0x99, 0xc4, 0x33, 0x45, 0xa0, 0x54, 0x02, 0x81, 0x3c, 0xc4, 0x03, 0x2f, 0x93,
	0xdb, 0x7a, 0x4d, 0xd4, 0x26, 0xae, 0x6a, 0xb7, 0xa2, 0x7f, 0xc1, 0xaf, 0xf0, 0x17, 0x7b, 0xdc,
	0x23, 0x4f, 0x13, 0x6a, 0x3f, 0x80, 0x5f, 0x40, 0xb6, 0x93, 0x30, 0x8d, 0x0d, 0xb6, 0xa7, 0xde,
	0x73, 0x7a, 0xee, 0xf1, 0xf5, 0xc9, 0x35, 0x1c, 0x0c, 0xb9, 0xc8, 0xb8, 0x88, 0x06, 0x54, 0xb0,
	0x48, 0xe4, 0x74, 0x26, 0x12, 0x2e, 0x45, 0xb4, 0x3c, 0x1a, 0x30, 0x49, 0x8f, 0x2a, 0x26, 0x9c,
	0xcd, 0xb9, 0xe4, 0xf8, 0xa9, 0x51, 0x87, 0x4a, 0x1d, 0x56, 0xea, 0xb0, 0x50, 0x3f, 0xd9, 0x1d,
	0xf3, 0x31, 0xd7, 0xca, 0x48, 0x55, 0xa6, 0x29, 0xf8, 0x8e, 0xa0, 0x79, 0x52, 0x68, 0xf1, 0x63,
	0x68, 0x24, 0x2c, 0x1d, 0x27, 0xd2, 0x43, 0x6d, 0xd4, 0xb1, 0x49, 0x81, 0x14, 0x7f, 0xc6, 0xe7,
	0x19, 0x95, 0xde, 0x56, 0x1b, 0x75, 0x1e, 0x92, 0x02, 0x29, 0x7e, 0x98, 0x2c, 0xf2, 0x89, 0xf0,
	0x6a, 0x86, 0x37, 0x08, 0x63, 0xb0, 0x13, 0x2a, 0x12, 0xcf, 0x6e, 0xa3, 0x8e, 0x4b, 0x74, 0x8d,
	0xfb, 0xd0, 0xcc, 0x98, 0xa4, 0x23, 0x2a, 0xa9, 0x57, 0x6f, 0xa3, 0x8e, 0xd3, 0xdd, 0x0f, 0xff,
	0x39, 0x70, 0xf8, 0xbe, 0x90, 0xf7, 0xec, 0xf3, 0xcb, 0x3d, 0x8b, 0x54, 0xed, 0xc1, 0x21, 0x34,
	0xcb, 0xff, 0xf0, 0x33, 0x70, 0xf5, 0xa1, 0xa7, 0xea, 0x10, 0x26, 0x3c, 0xd4, 0xae, 0x75, 0x5c,
	0xe2, 0x68, 0x2e, 0xd6, 0x54, 0xf0, 0x6b, 0x0b, 0xdc, 0xf2, 0x8a, 0x7d, 0xc9, 0x32, 0x1c, 0x43,
	0x5d, 0x48, 0x3e, 0x67, 0xfa, 0x96, 0x4e, 0xf7, 0xf9, 0x7f, 0xe6, 0x28, 0x7b, 0x4f, 0x54, 0x8f,
	0x32, 0x88, 0x2d, 0x62, 0x0c, 0xf0, 0x07, 0xb0, 0x53, 0xba, 0x9c, 0xea, 0x58, 0x9c, 0x6e, 0x74,
	0x47, 0xa3, 0xfe, 0xab, 0xcf, 0xef, 0x94, 0x4f, 0xaf, 0xb9, 0xbe, 0xdc, 0xb3, 0x15, 0x8a, 0x2d,
	0xa2, 0x8d, 0xf0, 0x27, 0xd8, 0x66, 0x5f, 0x25, 0xcb, 0x45, 0xca, 0x73, 0x1d, 0xaa, 0xd3, 0x3d,
	0xbe, 0xa3, 0xeb, 0x9b, 0xb2, 0x4f, 0x65, 0x13, 0x5b, 0xe4, 0x8f, 0x11, 0x3e, 0x83, 0x9d, 0x0a,
	0x9c, 0xce, 0xe8, 0x6a, 0xca, 0xe9, 0x48, 0x7f, 0x1c, 0xa7, 0xfb, 0xf2, 0xbe, 0xee, 0x1f, 0x4d,
	0x7b, 0x6c, 0x91, 0x16, 0xbb, 0xc6, 0xf5, 0x1a, 0x60, 0xa7, 0x92, 0x65, 0xc1, 0x3e, 0xec, 0xfc,
	0x15, 0x9a, 0x5a, 0x8a, 0x9c, 0x66, 0x26, 0xf4, 0x6d, 0xa2, 0xeb, 0x60, 0x0a, 0xad, 0xeb, 0xa1,
	0xe0, 0x16, 0xd4, 0x26, 0x6c, 0xa5, 0x65, 0x2e, 0x51, 0x25, 0xde, 0x85, 0xfa, 0x92, 0x4e, 0x17,
	0x4c, 0xc7, 0xec, 0x12, 0x03, 0xb0, 0x07, 0x0f, 0x96, 0x6c, 0x5e, 0x05, 0x55, 0x23, 0x25, 0xbc,
	0xb2, 0xc6, 0xea, 0x8e, 0xf5, 0x72, 0x8d, 0x83, 0xd7, 0xf0, 0xe8, 0xc6, 0xb0, 0x6e, 0x1a, 0xed,
	0xb6, 0x9d, 0x0f, 0x8e, 0xc1, 0xbb, 0x2d, 0x13, 0x35, 0x52, 0x99, 0xae, 0x19, 0xbf, 0x84, 0xbd,
	0xb7, 0xe7, 0x6b, 0x1f, 0x5d, 0xac, 0x7d, 0xf4, 0x73, 0xed, 0xa3, 0x6f, 0x1b, 0xdf, 0xba, 0xd8,
	0xf8, 0xd6, 0x8f, 0x8d, 0x6f, 0x7d, 0x39, 0x18, 0xa7, 0x32, 0x59, 0x0c, 0xc2, 0x21, 0xcf, 0xa2,
	0xe2, 0xb9, 0x9b, 0x9f, 0x43, 0x31, 0x9a, 0x5c, 0x79, 0xf4, 0x72, 0x35, 0x63, 0x62, 0xd0, 0xd0,
	0xaf, 0xf6, 0xc5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x90, 0xf8, 0x1f, 0x1a, 0x04, 0x00,
	0x00,
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSnapshot(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x22
	}
	if m.Chunks != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Chunks))
		i--
		dAtA[i] = 0x18
	}
	if m.Format != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Format))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
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
	if len(m.ChunkHashes) > 0 {
		for iNdEx := len(m.ChunkHashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChunkHashes[iNdEx])
			copy(dAtA[i:], m.ChunkHashes[iNdEx])
			i = encodeVarintSnapshot(dAtA, i, uint64(len(m.ChunkHashes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Item != nil {
		{
			size := m.Item.Size()
			i -= size
			if _, err := m.Item.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotItem_Store) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotItem_Store) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Store != nil {
		{
			size, err := m.Store.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSnapshot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *SnapshotItem_IAVL) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotItem_IAVL) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IAVL != nil {
		{
			size, err := m.IAVL.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSnapshot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *SnapshotItem_Extension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotItem_Extension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Extension != nil {
		{
			size, err := m.Extension.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSnapshot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *SnapshotItem_ExtensionPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotItem_ExtensionPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ExtensionPayload != nil {
		{
			size, err := m.ExtensionPayload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSnapshot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *SnapshotStoreItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotStoreItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotStoreItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotIAVLItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotIAVLItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotIAVLItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if m.Version != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotExtensionMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotExtensionMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotExtensionMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Format != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Format))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotExtensionPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotExtensionPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotExtensionPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	offset -= sovSnapshot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovSnapshot(uint64(m.Height))
	}
	if m.Format != 0 {
		n += 1 + sovSnapshot(uint64(m.Format))
	}
	if m.Chunks != 0 {
		n += 1 + sovSnapshot(uint64(m.Chunks))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChunkHashes) > 0 {
		for _, b := range m.ChunkHashes {
			l = len(b)
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	return n
}

func (m *SnapshotItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Item != nil {
		n += m.Item.Size()
	}
	return n
}

func (m *SnapshotItem_Store) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Store != nil {
		l = m.Store.Size()
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}
func (m *SnapshotItem_IAVL) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IAVL != nil {
		l = m.IAVL.Size()
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}
func (m *SnapshotItem_Extension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Extension != nil {
		l = m.Extension.Size()
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}
func (m *SnapshotItem_ExtensionPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExtensionPayload != nil {
		l = m.ExtensionPayload.Size()
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}
func (m *SnapshotStoreItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func (m *SnapshotIAVLItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovSnapshot(uint64(m.Version))
	}
	if m.Height != 0 {
		n += 1 + sovSnapshot(uint64(m.Height))
	}
	return n
}

func (m *SnapshotExtensionMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	if m.Format != 0 {
		n += 1 + sovSnapshot(uint64(m.Format))
	}
	return n
}

func (m *SnapshotExtensionPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func sovSnapshot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			m.Format = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Format |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			m.Chunks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Chunks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
				return ErrIntOverflowSnapshot
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
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkHashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChunkHashes = append(m.ChunkHashes, make([]byte, postIndex-iNdEx))
			copy(m.ChunkHashes[len(m.ChunkHashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *SnapshotItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SnapshotStoreItem{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &SnapshotItem_Store{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IAVL", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SnapshotIAVLItem{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &SnapshotItem_IAVL{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extension", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SnapshotExtensionMeta{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &SnapshotItem_Extension{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtensionPayload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SnapshotExtensionPayload{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &SnapshotItem_ExtensionPayload{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *SnapshotStoreItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotStoreItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotStoreItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *SnapshotIAVLItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotIAVLItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotIAVLItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *SnapshotExtensionMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotExtensionMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotExtensionMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			m.Format = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Format |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *SnapshotExtensionPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotExtensionPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotExtensionPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
				return 0, ErrInvalidLengthSnapshot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSnapshot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSnapshot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSnapshot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSnapshot = fmt.Errorf("proto: unexpected end of group")
)
