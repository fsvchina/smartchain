package codec

import (
	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec/types"
)

type (




	//


	Codec interface {
		BinaryCodec
		JSONCodec
	}

	BinaryCodec interface {

		Marshal(o ProtoMarshaler) ([]byte, error)

		MustMarshal(o ProtoMarshaler) []byte


		MarshalLengthPrefixed(o ProtoMarshaler) ([]byte, error)


		MustMarshalLengthPrefixed(o ProtoMarshaler) []byte



		Unmarshal(bz []byte, ptr ProtoMarshaler) error

		MustUnmarshal(bz []byte, ptr ProtoMarshaler)



		UnmarshalLengthPrefixed(bz []byte, ptr ProtoMarshaler) error


		MustUnmarshalLengthPrefixed(bz []byte, ptr ProtoMarshaler)



		MarshalInterface(i proto.Message) ([]byte, error)



		UnmarshalInterface(bz []byte, ptr interface{}) error

		types.AnyUnpacker
	}

	JSONCodec interface {

		MarshalJSON(o proto.Message) ([]byte, error)

		MustMarshalJSON(o proto.Message) []byte


		MarshalInterfaceJSON(i proto.Message) ([]byte, error)



		UnmarshalInterfaceJSON(bz []byte, ptr interface{}) error



		UnmarshalJSON(bz []byte, ptr proto.Message) error

		MustUnmarshalJSON(bz []byte, ptr proto.Message)
	}



	ProtoMarshaler interface {
		proto.Message

		Marshal() ([]byte, error)
		MarshalTo(data []byte) (n int, err error)
		MarshalToSizedBuffer(dAtA []byte) (int, error)
		Size() int
		Unmarshal(data []byte) error
	}



	AminoMarshaler interface {
		MarshalAmino() ([]byte, error)
		UnmarshalAmino([]byte) error
		MarshalAminoJSON() ([]byte, error)
		UnmarshalAminoJSON([]byte) error
	}
)
