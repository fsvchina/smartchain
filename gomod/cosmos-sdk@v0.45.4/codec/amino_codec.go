package codec

import (
	"github.com/gogo/protobuf/proto"
)



type AminoCodec struct {
	*LegacyAmino
}

var _ Codec = &AminoCodec{}


func NewAminoCodec(codec *LegacyAmino) *AminoCodec {
	return &AminoCodec{LegacyAmino: codec}
}


func (ac *AminoCodec) Marshal(o ProtoMarshaler) ([]byte, error) {
	return ac.LegacyAmino.Marshal(o)
}


func (ac *AminoCodec) MustMarshal(o ProtoMarshaler) []byte {
	return ac.LegacyAmino.MustMarshal(o)
}


func (ac *AminoCodec) MarshalLengthPrefixed(o ProtoMarshaler) ([]byte, error) {
	return ac.LegacyAmino.MarshalLengthPrefixed(o)
}


func (ac *AminoCodec) MustMarshalLengthPrefixed(o ProtoMarshaler) []byte {
	return ac.LegacyAmino.MustMarshalLengthPrefixed(o)
}


func (ac *AminoCodec) Unmarshal(bz []byte, ptr ProtoMarshaler) error {
	return ac.LegacyAmino.Unmarshal(bz, ptr)
}


func (ac *AminoCodec) MustUnmarshal(bz []byte, ptr ProtoMarshaler) {
	ac.LegacyAmino.MustUnmarshal(bz, ptr)
}


func (ac *AminoCodec) UnmarshalLengthPrefixed(bz []byte, ptr ProtoMarshaler) error {
	return ac.LegacyAmino.UnmarshalLengthPrefixed(bz, ptr)
}


func (ac *AminoCodec) MustUnmarshalLengthPrefixed(bz []byte, ptr ProtoMarshaler) {
	ac.LegacyAmino.MustUnmarshalLengthPrefixed(bz, ptr)
}



func (ac *AminoCodec) MarshalJSON(o proto.Message) ([]byte, error) {
	return ac.LegacyAmino.MarshalJSON(o)
}



func (ac *AminoCodec) MustMarshalJSON(o proto.Message) []byte {
	return ac.LegacyAmino.MustMarshalJSON(o)
}



func (ac *AminoCodec) UnmarshalJSON(bz []byte, ptr proto.Message) error {
	return ac.LegacyAmino.UnmarshalJSON(bz, ptr)
}



func (ac *AminoCodec) MustUnmarshalJSON(bz []byte, ptr proto.Message) {
	ac.LegacyAmino.MustUnmarshalJSON(bz, ptr)
}




func (ac *AminoCodec) MarshalInterface(i proto.Message) ([]byte, error) {
	if err := assertNotNil(i); err != nil {
		return nil, err
	}
	return ac.LegacyAmino.Marshal(i)
}




//



func (ac *AminoCodec) UnmarshalInterface(bz []byte, ptr interface{}) error {
	return ac.LegacyAmino.Unmarshal(bz, ptr)
}




func (ac *AminoCodec) MarshalInterfaceJSON(i proto.Message) ([]byte, error) {
	if err := assertNotNil(i); err != nil {
		return nil, err
	}
	return ac.LegacyAmino.MarshalJSON(i)
}




//



func (ac *AminoCodec) UnmarshalInterfaceJSON(bz []byte, ptr interface{}) error {
	return ac.LegacyAmino.UnmarshalJSON(bz, ptr)
}
