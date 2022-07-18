package types

import (
	fmt "fmt"

	"github.com/gogo/protobuf/proto"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type Any struct {






	//




	//








	//



	//




	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`

	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`


	XXX_NoUnkeyedLiteral struct{} `json:"-"`


	XXX_unrecognized []byte `json:"-"`


	XXX_sizecache int32 `json:"-"`

	cachedValue interface{}

	compat *anyCompat
}





func NewAnyWithValue(v proto.Message) (*Any, error) {
	if v == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPackAny, "Expecting non nil value to create a new Any")
	}

	bz, err := proto.Marshal(v)
	if err != nil {
		return nil, err
	}

	return &Any{
		TypeUrl:     "/" + proto.MessageName(v),
		Value:       bz,
		cachedValue: v,
	}, nil
}






func UnsafePackAny(x interface{}) *Any {
	if msg, ok := x.(proto.Message); ok {
		any, err := NewAnyWithValue(msg)
		if err == nil {
			return any
		}
	}
	return &Any{cachedValue: x}
}




func (any *Any) pack(x proto.Message) error {
	any.TypeUrl = "/" + proto.MessageName(x)
	bz, err := proto.Marshal(x)
	if err != nil {
		return err
	}

	any.Value = bz
	any.cachedValue = x

	return nil
}


func (any *Any) GetCachedValue() interface{} {
	return any.cachedValue
}



func (any *Any) GoString() string {
	if any == nil {
		return "nil"
	}
	extra := ""
	if any.XXX_unrecognized != nil {
		extra = fmt.Sprintf(",\n  XXX_unrecognized: %#v,\n", any.XXX_unrecognized)
	}
	return fmt.Sprintf("&Any{TypeUrl: %#v,\n  Value: %#v%s\n}",
		any.TypeUrl, any.Value, extra)
}


func (any *Any) String() string {
	if any == nil {
		return "nil"
	}
	return fmt.Sprintf("&Any{TypeUrl:%v,Value:%v,XXX_unrecognized:%v}",
		any.TypeUrl, any.Value, any.XXX_unrecognized)
}
