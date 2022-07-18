

grpc/reflection"

	s := grpc.NewServer()
	pb.RegisterYourOwnServer(s, &server{})


	reflection.Register(s)

	s.Serve(lis)

*/
package gogoreflection

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"sync"


	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/status"
)

type serverReflectionServer struct {
	rpb.UnimplementedServerReflectionServer
	s *grpc.Server

	initSymbols  sync.Once
	serviceNames []string
	symbols      map[string]*dpb.FileDescriptorProto
}


func Register(s *grpc.Server) {
	rpb.RegisterServerReflectionServer(s, &serverReflectionServer{
		s: s,
	})
}





type protoMessage interface {
	Descriptor() ([]byte, []int)
}

func (s *serverReflectionServer) getSymbols() (svcNames []string, symbolIndex map[string]*dpb.FileDescriptorProto) {
	s.initSymbols.Do(func() {
		serviceInfo := s.s.GetServiceInfo()

		s.symbols = map[string]*dpb.FileDescriptorProto{}
		s.serviceNames = make([]string, 0, len(serviceInfo))
		processed := map[string]struct{}{}
		for svc, info := range serviceInfo {
			s.serviceNames = append(s.serviceNames, svc)
			fdenc, ok := parseMetadata(info.Metadata)
			if !ok {
				continue
			}
			fd, err := decodeFileDesc(fdenc)
			if err != nil {
				continue
			}
			s.processFile(fd, processed)
		}
		sort.Strings(s.serviceNames)
	})

	return s.serviceNames, s.symbols
}

func (s *serverReflectionServer) processFile(fd *dpb.FileDescriptorProto, processed map[string]struct{}) {
	filename := fd.GetName()
	if _, ok := processed[filename]; ok {
		return
	}
	processed[filename] = struct{}{}

	prefix := fd.GetPackage()

	for _, msg := range fd.MessageType {
		s.processMessage(fd, prefix, msg)
	}
	for _, en := range fd.EnumType {
		s.processEnum(fd, prefix, en)
	}
	for _, ext := range fd.Extension {
		s.processField(fd, prefix, ext)
	}
	for _, svc := range fd.Service {
		svcName := fqn(prefix, svc.GetName())
		s.symbols[svcName] = fd
		for _, meth := range svc.Method {
			name := fqn(svcName, meth.GetName())
			s.symbols[name] = fd
		}
	}

	for _, dep := range fd.Dependency {
		fdenc := getFileDescriptor(dep)
		fdDep, err := decodeFileDesc(fdenc)
		if err != nil {
			continue
		}
		s.processFile(fdDep, processed)
	}
}

func (s *serverReflectionServer) processMessage(fd *dpb.FileDescriptorProto, prefix string, msg *dpb.DescriptorProto) {
	msgName := fqn(prefix, msg.GetName())
	s.symbols[msgName] = fd

	for _, nested := range msg.NestedType {
		s.processMessage(fd, msgName, nested)
	}
	for _, en := range msg.EnumType {
		s.processEnum(fd, msgName, en)
	}
	for _, ext := range msg.Extension {
		s.processField(fd, msgName, ext)
	}
	for _, fld := range msg.Field {
		s.processField(fd, msgName, fld)
	}
	for _, oneof := range msg.OneofDecl {
		oneofName := fqn(msgName, oneof.GetName())
		s.symbols[oneofName] = fd
	}
}

func (s *serverReflectionServer) processEnum(fd *dpb.FileDescriptorProto, prefix string, en *dpb.EnumDescriptorProto) {
	enName := fqn(prefix, en.GetName())
	s.symbols[enName] = fd

	for _, val := range en.Value {
		valName := fqn(enName, val.GetName())
		s.symbols[valName] = fd
	}
}

func (s *serverReflectionServer) processField(fd *dpb.FileDescriptorProto, prefix string, fld *dpb.FieldDescriptorProto) {
	fldName := fqn(prefix, fld.GetName())
	s.symbols[fldName] = fd
}

func fqn(prefix, name string) string {
	if prefix == "" {
		return name
	}
	return prefix + "." + name
}



func (s *serverReflectionServer) fileDescForType(st reflect.Type) (*dpb.FileDescriptorProto, error) {
	m, ok := reflect.Zero(reflect.PtrTo(st)).Interface().(protoMessage)
	if !ok {
		return nil, fmt.Errorf("failed to create message from type: %v", st)
	}
	enc, _ := m.Descriptor()

	return decodeFileDesc(enc)
}



func decodeFileDesc(enc []byte) (*dpb.FileDescriptorProto, error) {
	raw, err := decompress(enc)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress enc: %v", err)
	}

	fd := new(dpb.FileDescriptorProto)
	if err := proto.Unmarshal(raw, fd); err != nil {
		return nil, fmt.Errorf("bad descriptor: %v", err)
	}
	return fd, nil
}


func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	return out, nil
}

func typeForName(name string) (reflect.Type, error) {
	pt := getMessageType(name)
	if pt == nil {
		return nil, fmt.Errorf("unknown type: %q", name)
	}
	st := pt.Elem()

	return st, nil
}

func fileDescContainingExtension(st reflect.Type, ext int32) (*dpb.FileDescriptorProto, error) {
	m, ok := reflect.Zero(reflect.PtrTo(st)).Interface().(proto.Message)
	if !ok {
		return nil, fmt.Errorf("failed to create message from type: %v", st)
	}

	extDesc := getExtension(ext, m)

	if extDesc == nil {
		return nil, fmt.Errorf("failed to find registered extension for extension number %v", ext)
	}

	return decodeFileDesc(getFileDescriptor(extDesc.Filename))
}

func (s *serverReflectionServer) allExtensionNumbersForType(st reflect.Type) ([]int32, error) {
	m, ok := reflect.Zero(reflect.PtrTo(st)).Interface().(proto.Message)
	if !ok {
		return nil, fmt.Errorf("failed to create message from type: %v", st)
	}

	out := getExtensionsNumbers(m)
	return out, nil
}




func fileDescWithDependencies(fd *dpb.FileDescriptorProto, sentFileDescriptors map[string]bool) ([][]byte, error) {
	r := [][]byte{}
	queue := []*dpb.FileDescriptorProto{fd}
	for len(queue) > 0 {
		currentfd := queue[0]
		queue = queue[1:]
		if sent := sentFileDescriptors[currentfd.GetName()]; len(r) == 0 || !sent {
			sentFileDescriptors[currentfd.GetName()] = true
			currentfdEncoded, err := proto.Marshal(currentfd)
			if err != nil {
				return nil, err
			}
			r = append(r, currentfdEncoded)
		}
		for _, dep := range currentfd.Dependency {
			fdenc := getFileDescriptor(dep)
			fdDep, err := decodeFileDesc(fdenc)
			if err != nil {
				continue
			}
			queue = append(queue, fdDep)
		}
	}
	return r, nil
}




func (s *serverReflectionServer) fileDescEncodingByFilename(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
	enc := getFileDescriptor(name)
	if enc == nil {
		return nil, fmt.Errorf("unknown file: %v", name)
	}
	fd, err := decodeFileDesc(enc)
	if err != nil {
		return nil, err
	}
	return fileDescWithDependencies(fd, sentFileDescriptors)
}





func parseMetadata(meta interface{}) ([]byte, bool) {

	if fileNameForMeta, ok := meta.(string); ok {
		return getFileDescriptor(fileNameForMeta), true
	}


	if enc, ok := meta.([]byte); ok {
		return enc, true
	}

	return nil, false
}





func (s *serverReflectionServer) fileDescEncodingContainingSymbol(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
	_, symbols := s.getSymbols()
	fd := symbols[name]
	if fd == nil {


		if st, err := typeForName(name); err == nil {
			fd, err = s.fileDescForType(st)
			if err != nil {
				return nil, err
			}
		}
	}

	if fd == nil {
		return nil, fmt.Errorf("unknown symbol: %v", name)
	}

	return fileDescWithDependencies(fd, sentFileDescriptors)
}




func (s *serverReflectionServer) fileDescEncodingContainingExtension(typeName string, extNum int32, sentFileDescriptors map[string]bool) ([][]byte, error) {
	st, err := typeForName(typeName)
	if err != nil {
		return nil, err
	}
	fd, err := fileDescContainingExtension(st, extNum)
	if err != nil {
		return nil, err
	}
	return fileDescWithDependencies(fd, sentFileDescriptors)
}


func (s *serverReflectionServer) allExtensionNumbersForTypeName(name string) ([]int32, error) {
	st, err := typeForName(name)
	if err != nil {
		return nil, err
	}
	extNums, err := s.allExtensionNumbersForType(st)
	if err != nil {
		return nil, err
	}
	return extNums, nil
}


func (s *serverReflectionServer) ServerReflectionInfo(stream rpb.ServerReflection_ServerReflectionInfoServer) error {
	sentFileDescriptors := make(map[string]bool)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		out := &rpb.ServerReflectionResponse{
			ValidHost:       in.Host,
			OriginalRequest: in,
		}
		switch req := in.MessageRequest.(type) {
		case *rpb.ServerReflectionRequest_FileByFilename:
			b, err := s.fileDescEncodingByFilename(req.FileByFilename, sentFileDescriptors)
			if err != nil {
				out.MessageResponse = &rpb.ServerReflectionResponse_ErrorResponse{
					ErrorResponse: &rpb.ErrorResponse{
						ErrorCode:    int32(codes.NotFound),
						ErrorMessage: err.Error(),
					},
				}
			} else {
				out.MessageResponse = &rpb.ServerReflectionResponse_FileDescriptorResponse{
					FileDescriptorResponse: &rpb.FileDescriptorResponse{FileDescriptorProto: b},
				}
			}
		case *rpb.ServerReflectionRequest_FileContainingSymbol:
			b, err := s.fileDescEncodingContainingSymbol(req.FileContainingSymbol, sentFileDescriptors)
			if err != nil {
				out.MessageResponse = &rpb.ServerReflectionResponse_ErrorResponse{
					ErrorResponse: &rpb.ErrorResponse{
						ErrorCode:    int32(codes.NotFound),
						ErrorMessage: err.Error(),
					},
				}
			} else {
				out.MessageResponse = &rpb.ServerReflectionResponse_FileDescriptorResponse{
					FileDescriptorResponse: &rpb.FileDescriptorResponse{FileDescriptorProto: b},
				}
			}
		case *rpb.ServerReflectionRequest_FileContainingExtension:
			typeName := req.FileContainingExtension.ContainingType
			extNum := req.FileContainingExtension.ExtensionNumber
			b, err := s.fileDescEncodingContainingExtension(typeName, extNum, sentFileDescriptors)
			if err != nil {
				out.MessageResponse = &rpb.ServerReflectionResponse_ErrorResponse{
					ErrorResponse: &rpb.ErrorResponse{
						ErrorCode:    int32(codes.NotFound),
						ErrorMessage: err.Error(),
					},
				}
			} else {
				out.MessageResponse = &rpb.ServerReflectionResponse_FileDescriptorResponse{
					FileDescriptorResponse: &rpb.FileDescriptorResponse{FileDescriptorProto: b},
				}
			}
		case *rpb.ServerReflectionRequest_AllExtensionNumbersOfType:
			extNums, err := s.allExtensionNumbersForTypeName(req.AllExtensionNumbersOfType)
			if err != nil {
				out.MessageResponse = &rpb.ServerReflectionResponse_ErrorResponse{
					ErrorResponse: &rpb.ErrorResponse{
						ErrorCode:    int32(codes.NotFound),
						ErrorMessage: err.Error(),
					},
				}
				log.Printf("OH NO: %s", err)
			} else {
				out.MessageResponse = &rpb.ServerReflectionResponse_AllExtensionNumbersResponse{
					AllExtensionNumbersResponse: &rpb.ExtensionNumberResponse{
						BaseTypeName:    req.AllExtensionNumbersOfType,
						ExtensionNumber: extNums,
					},
				}
			}
		case *rpb.ServerReflectionRequest_ListServices:
			svcNames, _ := s.getSymbols()
			serviceResponses := make([]*rpb.ServiceResponse, len(svcNames))
			for i, n := range svcNames {
				serviceResponses[i] = &rpb.ServiceResponse{
					Name: n,
				}
			}
			out.MessageResponse = &rpb.ServerReflectionResponse_ListServicesResponse{
				ListServicesResponse: &rpb.ListServiceResponse{
					Service: serviceResponses,
				},
			}
		default:
			return status.Errorf(codes.InvalidArgument, "invalid MessageRequest: %v", in.MessageRequest)
		}
		if err := stream.Send(out); err != nil {
			return err
		}
	}
}
