package unknownproto

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"sync"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"google.golang.org/protobuf/encoding/protowire"

	"github.com/cosmos/cosmos-sdk/codec/types"
)

const bit11NonCritical = 1 << 10

type descriptorIface interface {
	Descriptor() ([]byte, []int)
}




func RejectUnknownFieldsStrict(bz []byte, msg proto.Message, resolver jsonpb.AnyResolver) error {
	_, err := RejectUnknownFields(bz, msg, false, resolver)
	return err
}







func RejectUnknownFields(bz []byte, msg proto.Message, allowUnknownNonCriticals bool, resolver jsonpb.AnyResolver) (hasUnknownNonCriticals bool, err error) {
	if len(bz) == 0 {
		return hasUnknownNonCriticals, nil
	}

	desc, ok := msg.(descriptorIface)
	if !ok {
		return hasUnknownNonCriticals, fmt.Errorf("%T does not have a Descriptor() method", msg)
	}

	fieldDescProtoFromTagNum, _, err := getDescriptorInfo(desc, msg)
	if err != nil {
		return hasUnknownNonCriticals, err
	}

	for len(bz) > 0 {
		tagNum, wireType, m := protowire.ConsumeTag(bz)
		if m < 0 {
			return hasUnknownNonCriticals, errors.New("invalid length")
		}

		fieldDescProto, ok := fieldDescProtoFromTagNum[int32(tagNum)]
		switch {
		case ok:

			if !canEncodeType(wireType, fieldDescProto.GetType()) {
				return hasUnknownNonCriticals, &errMismatchedWireType{
					Type:         reflect.ValueOf(msg).Type().String(),
					TagNum:       tagNum,
					GotWireType:  wireType,
					WantWireType: protowire.Type(fieldDescProto.WireType()),
				}
			}

		default:
			isCriticalField := tagNum&bit11NonCritical == 0

			if !isCriticalField {
				hasUnknownNonCriticals = true
			}

			if isCriticalField || !allowUnknownNonCriticals {

				return hasUnknownNonCriticals, &errUnknownField{
					Type:     reflect.ValueOf(msg).Type().String(),
					TagNum:   tagNum,
					WireType: wireType,
				}
			}
		}


		bz = bz[m:]
		n := protowire.ConsumeFieldValue(tagNum, wireType, bz)
		if n < 0 {
			err = fmt.Errorf("could not consume field value for tagNum: %d, wireType: %q; %w",
				tagNum, wireTypeToString(wireType), protowire.ParseError(n))
			return hasUnknownNonCriticals, err
		}
		fieldBytes := bz[:n]
		bz = bz[n:]


		if fieldDescProto == nil || fieldDescProto.IsScalar() {
			continue
		}

		protoMessageName := fieldDescProto.GetTypeName()
		if protoMessageName == "" {
			switch typ := fieldDescProto.GetType(); typ {
			case descriptor.FieldDescriptorProto_TYPE_STRING, descriptor.FieldDescriptorProto_TYPE_BYTES:



			default:
				return hasUnknownNonCriticals, fmt.Errorf("failed to get typename for message of type %v, can only be TYPE_STRING or TYPE_BYTES", typ)
			}
			continue
		}




		_, o := protowire.ConsumeVarint(fieldBytes)
		fieldBytes = fieldBytes[o:]

		var msg proto.Message
		var err error

		if protoMessageName == ".google.protobuf.Any" {

			hasUnknownNonCriticalsChild, err := RejectUnknownFields(fieldBytes, (*types.Any)(nil), allowUnknownNonCriticals, resolver)
			hasUnknownNonCriticals = hasUnknownNonCriticals || hasUnknownNonCriticalsChild
			if err != nil {
				return hasUnknownNonCriticals, err
			}

			any := new(types.Any)
			if err := proto.Unmarshal(fieldBytes, any); err != nil {
				return hasUnknownNonCriticals, err
			}
			protoMessageName = any.TypeUrl
			fieldBytes = any.Value
			msg, err = resolver.Resolve(protoMessageName)
			if err != nil {
				return hasUnknownNonCriticals, err
			}
		} else {
			msg, err = protoMessageForTypeName(protoMessageName[1:])
			if err != nil {
				return hasUnknownNonCriticals, err
			}
		}

		hasUnknownNonCriticalsChild, err := RejectUnknownFields(fieldBytes, msg, allowUnknownNonCriticals, resolver)
		hasUnknownNonCriticals = hasUnknownNonCriticals || hasUnknownNonCriticalsChild
		if err != nil {
			return hasUnknownNonCriticals, err
		}
	}

	return hasUnknownNonCriticals, nil
}

var protoMessageForTypeNameMu sync.RWMutex
var protoMessageForTypeNameCache = make(map[string]proto.Message)



func protoMessageForTypeName(protoMessageName string) (proto.Message, error) {
	protoMessageForTypeNameMu.RLock()
	msg, ok := protoMessageForTypeNameCache[protoMessageName]
	protoMessageForTypeNameMu.RUnlock()
	if ok {
		return msg, nil
	}

	concreteGoType := proto.MessageType(protoMessageName)
	if concreteGoType == nil {
		return nil, fmt.Errorf("failed to retrieve the message of type %q", protoMessageName)
	}

	value := reflect.New(concreteGoType).Elem()
	msg, ok = value.Interface().(proto.Message)
	if !ok {
		return nil, fmt.Errorf("%q does not implement proto.Message", protoMessageName)
	}


	protoMessageForTypeNameMu.Lock()
	protoMessageForTypeNameCache[protoMessageName] = msg
	protoMessageForTypeNameMu.Unlock()

	return msg, nil
}




var checks = [...]map[descriptor.FieldDescriptorProto_Type]bool{

	0: {
		descriptor.FieldDescriptorProto_TYPE_INT32:  true,
		descriptor.FieldDescriptorProto_TYPE_INT64:  true,
		descriptor.FieldDescriptorProto_TYPE_UINT32: true,
		descriptor.FieldDescriptorProto_TYPE_UINT64: true,
		descriptor.FieldDescriptorProto_TYPE_SINT32: true,
		descriptor.FieldDescriptorProto_TYPE_SINT64: true,
		descriptor.FieldDescriptorProto_TYPE_BOOL:   true,
		descriptor.FieldDescriptorProto_TYPE_ENUM:   true,
	},


	1: {
		descriptor.FieldDescriptorProto_TYPE_FIXED64:  true,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: true,
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:   true,
	},


	2: {
		descriptor.FieldDescriptorProto_TYPE_STRING:  true,
		descriptor.FieldDescriptorProto_TYPE_BYTES:   true,
		descriptor.FieldDescriptorProto_TYPE_MESSAGE: true,



		descriptor.FieldDescriptorProto_TYPE_INT32:    true,
		descriptor.FieldDescriptorProto_TYPE_INT64:    true,
		descriptor.FieldDescriptorProto_TYPE_UINT32:   true,
		descriptor.FieldDescriptorProto_TYPE_UINT64:   true,
		descriptor.FieldDescriptorProto_TYPE_SINT32:   true,
		descriptor.FieldDescriptorProto_TYPE_SINT64:   true,
		descriptor.FieldDescriptorProto_TYPE_BOOL:     true,
		descriptor.FieldDescriptorProto_TYPE_ENUM:     true,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:  true,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64: true,
		descriptor.FieldDescriptorProto_TYPE_DOUBLE:   true,
	},


	3: {
		descriptor.FieldDescriptorProto_TYPE_GROUP: true,
	},


	4: {
		descriptor.FieldDescriptorProto_TYPE_GROUP: true,
	},


	5: {
		descriptor.FieldDescriptorProto_TYPE_FIXED32:  true,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32: true,
		descriptor.FieldDescriptorProto_TYPE_FLOAT:    true,
	},
}



func canEncodeType(wireType protowire.Type, descType descriptor.FieldDescriptorProto_Type) bool {
	if iwt := int(wireType); iwt < 0 || iwt >= len(checks) {
		return false
	}
	return checks[wireType][descType]
}



type errMismatchedWireType struct {
	Type         string
	GotWireType  protowire.Type
	WantWireType protowire.Type
	TagNum       protowire.Number
}


func (mwt *errMismatchedWireType) String() string {
	return fmt.Sprintf("Mismatched %q: {TagNum: %d, GotWireType: %q != WantWireType: %q}",
		mwt.Type, mwt.TagNum, wireTypeToString(mwt.GotWireType), wireTypeToString(mwt.WantWireType))
}


func (mwt *errMismatchedWireType) Error() string {
	return mwt.String()
}

var _ error = (*errMismatchedWireType)(nil)

func wireTypeToString(wt protowire.Type) string {
	switch wt {
	case 0:
		return "varint"
	case 1:
		return "fixed64"
	case 2:
		return "bytes"
	case 3:
		return "start_group"
	case 4:
		return "end_group"
	case 5:
		return "fixed32"
	default:
		return fmt.Sprintf("unknown type: %d", wt)
	}
}



type errUnknownField struct {
	Type     string
	TagNum   protowire.Number
	WireType protowire.Type
}


func (twt *errUnknownField) String() string {
	return fmt.Sprintf("errUnknownField %q: {TagNum: %d, WireType:%q}",
		twt.Type, twt.TagNum, wireTypeToString(twt.WireType))
}


func (twt *errUnknownField) Error() string {
	return twt.String()
}

var _ error = (*errUnknownField)(nil)

var (
	protoFileToDesc   = make(map[string]*descriptor.FileDescriptorProto)
	protoFileToDescMu sync.RWMutex
)

func unnestDesc(mdescs []*descriptor.DescriptorProto, indices []int) *descriptor.DescriptorProto {
	mdesc := mdescs[indices[0]]
	for _, index := range indices[1:] {
		mdesc = mdesc.NestedType[index]
	}
	return mdesc
}



func extractFileDescMessageDesc(desc descriptorIface) (*descriptor.FileDescriptorProto, *descriptor.DescriptorProto, error) {
	gzippedPb, indices := desc.Descriptor()

	protoFileToDescMu.RLock()
	cached, ok := protoFileToDesc[string(gzippedPb)]
	protoFileToDescMu.RUnlock()

	if ok {
		return cached, unnestDesc(cached.MessageType, indices), nil
	}


	gzr, err := gzip.NewReader(bytes.NewReader(gzippedPb))
	if err != nil {
		return nil, nil, err
	}
	protoBlob, err := ioutil.ReadAll(gzr)
	if err != nil {
		return nil, nil, err
	}

	fdesc := new(descriptor.FileDescriptorProto)
	if err := proto.Unmarshal(protoBlob, fdesc); err != nil {
		return nil, nil, err
	}


	protoFileToDescMu.Lock()
	protoFileToDesc[string(gzippedPb)] = fdesc
	protoFileToDescMu.Unlock()


	return fdesc, unnestDesc(fdesc.MessageType, indices), nil
}

type descriptorMatch struct {
	cache map[int32]*descriptor.FieldDescriptorProto
	desc  *descriptor.DescriptorProto
}

var descprotoCacheMu sync.RWMutex
var descprotoCache = make(map[reflect.Type]*descriptorMatch)


func getDescriptorInfo(desc descriptorIface, msg proto.Message) (map[int32]*descriptor.FieldDescriptorProto, *descriptor.DescriptorProto, error) {
	key := reflect.ValueOf(msg).Type()

	descprotoCacheMu.RLock()
	got, ok := descprotoCache[key]
	descprotoCacheMu.RUnlock()

	if ok {
		return got.cache, got.desc, nil
	}


	_, md, err := extractFileDescMessageDesc(desc)
	if err != nil {
		return nil, nil, err
	}

	tagNumToTypeIndex := make(map[int32]*descriptor.FieldDescriptorProto)
	for _, field := range md.Field {
		tagNumToTypeIndex[field.GetNumber()] = field
	}

	descprotoCacheMu.Lock()
	descprotoCache[key] = &descriptorMatch{
		cache: tagNumToTypeIndex,
		desc:  md,
	}
	descprotoCacheMu.Unlock()

	return tagNumToTypeIndex, md, nil
}



type DefaultAnyResolver struct{}

var _ jsonpb.AnyResolver = DefaultAnyResolver{}


func (d DefaultAnyResolver) Resolve(typeURL string) (proto.Message, error) {

	mname := typeURL
	if slash := strings.LastIndex(mname, "/"); slash >= 0 {
		mname = mname[slash+1:]
	}
	mt := proto.MessageType(mname)
	if mt == nil {
		return nil, fmt.Errorf("unknown message type %q", mname)
	}
	return reflect.New(mt.Elem()).Interface().(proto.Message), nil
}
