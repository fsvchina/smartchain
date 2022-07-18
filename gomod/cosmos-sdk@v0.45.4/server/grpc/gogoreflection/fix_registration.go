package gogoreflection

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"reflect"

	_ "github.com/gogo/protobuf/gogoproto"
	gogoproto "github.com/gogo/protobuf/proto"


	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/regen-network/cosmos-proto"
)

var importsToFix = map[string]string{
	"gogo.proto":   "gogoproto/gogo.proto",
	"cosmos.proto": "cosmos_proto/cosmos.proto",
}






func fixRegistration(registeredAs, importedAs string) error {
	raw := gogoproto.FileDescriptor(registeredAs)
	if len(raw) == 0 {
		return fmt.Errorf("file descriptor not found for %s", registeredAs)
	}

	fd, err := decodeFileDesc(raw)
	if err != nil {
		return err
	}


	*fd.Name = importedAs
	fixedRaw, err := compress(fd)
	if err != nil {
		return fmt.Errorf("unable to compress: %w", err)
	}
	gogoproto.RegisterFile(importedAs, fixedRaw)
	return nil
}

func init() {




	for registeredAs, importedAs := range importsToFix {
		err := fixRegistration(registeredAs, importedAs)
		if err != nil {
			panic(err)
		}
	}
}



func compress(fd *dpb.FileDescriptorProto) ([]byte, error) {
	fdBytes, err := proto.Marshal(fd)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	cw := gzip.NewWriter(buf)
	_, err = cw.Write(fdBytes)
	if err != nil {
		return nil, err
	}
	err = cw.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getFileDescriptor(filePath string) []byte {


	fd := gogoproto.FileDescriptor(filePath)
	if len(fd) != 0 {
		return fd
	}

	return proto.FileDescriptor(filePath)
}

func getMessageType(name string) reflect.Type {
	typ := gogoproto.MessageType(name)
	if typ != nil {
		return typ
	}

	return proto.MessageType(name)
}

func getExtension(extID int32, m proto.Message) *gogoproto.ExtensionDesc {

	for id, desc := range gogoproto.RegisteredExtensions(m) {
		if id == extID {
			return desc
		}
	}


	for id, desc := range proto.RegisteredExtensions(m) {
		if id == extID {
			return &gogoproto.ExtensionDesc{
				ExtendedType:  desc.ExtendedType,
				ExtensionType: desc.ExtensionType,
				Field:         desc.Field,
				Name:          desc.Name,
				Tag:           desc.Tag,
				Filename:      desc.Filename,
			}
		}
	}

	return nil
}

func getExtensionsNumbers(m proto.Message) []int32 {
	gogoProtoExts := gogoproto.RegisteredExtensions(m)
	out := make([]int32, 0, len(gogoProtoExts))
	for id := range gogoProtoExts {
		out = append(out, id)
	}
	if len(out) != 0 {
		return out
	}

	protoExts := proto.RegisteredExtensions(m)
	out = make([]int32, 0, len(protoExts))
	for id := range protoExts {
		out = append(out, id)
	}
	return out
}
