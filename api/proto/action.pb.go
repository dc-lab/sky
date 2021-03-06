// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: api/proto/action.proto

// build *.pb.go from project root
// protoc -I. --go_out=plugins=grpc,paths=source_relative:. api/proto/common/*.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Result_ResultCode int32

const (
	Result_NONE     Result_ResultCode = 0
	Result_WAIT     Result_ResultCode = 1
	Result_RUN      Result_ResultCode = 2
	Result_FAILED   Result_ResultCode = 3
	Result_CANCELED Result_ResultCode = 4
	Result_DELETED  Result_ResultCode = 5
	Result_SUCCESS  Result_ResultCode = 6
)

// Enum value maps for Result_ResultCode.
var (
	Result_ResultCode_name = map[int32]string{
		0: "NONE",
		1: "WAIT",
		2: "RUN",
		3: "FAILED",
		4: "CANCELED",
		5: "DELETED",
		6: "SUCCESS",
	}
	Result_ResultCode_value = map[string]int32{
		"NONE":     0,
		"WAIT":     1,
		"RUN":      2,
		"FAILED":   3,
		"CANCELED": 4,
		"DELETED":  5,
		"SUCCESS":  6,
	}
)

func (x Result_ResultCode) Enum() *Result_ResultCode {
	p := new(Result_ResultCode)
	*p = x
	return p
}

func (x Result_ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result_ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_action_proto_enumTypes[0].Descriptor()
}

func (Result_ResultCode) Type() protoreflect.EnumType {
	return &file_api_proto_action_proto_enumTypes[0]
}

func (x Result_ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result_ResultCode.Descriptor instead.
func (Result_ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_action_proto_rawDescGZIP(), []int{0, 0}
}

type Result_ErrorCode int32

const (
	Result_ERROR_UNKNOWN    Result_ErrorCode = 0
	Result_INTERNAL         Result_ErrorCode = 1
	Result_INVALID_ARGUMENT Result_ErrorCode = 2
	Result_UNAUTHENTICATED  Result_ErrorCode = 3
	Result_UNAUTHORIZED     Result_ErrorCode = 4
	Result_NOT_IMPLEMENTED  Result_ErrorCode = 5
)

// Enum value maps for Result_ErrorCode.
var (
	Result_ErrorCode_name = map[int32]string{
		0: "ERROR_UNKNOWN",
		1: "INTERNAL",
		2: "INVALID_ARGUMENT",
		3: "UNAUTHENTICATED",
		4: "UNAUTHORIZED",
		5: "NOT_IMPLEMENTED",
	}
	Result_ErrorCode_value = map[string]int32{
		"ERROR_UNKNOWN":    0,
		"INTERNAL":         1,
		"INVALID_ARGUMENT": 2,
		"UNAUTHENTICATED":  3,
		"UNAUTHORIZED":     4,
		"NOT_IMPLEMENTED":  5,
	}
)

func (x Result_ErrorCode) Enum() *Result_ErrorCode {
	p := new(Result_ErrorCode)
	*p = x
	return p
}

func (x Result_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_action_proto_enumTypes[1].Descriptor()
}

func (Result_ErrorCode) Type() protoreflect.EnumType {
	return &file_api_proto_action_proto_enumTypes[1]
}

func (x Result_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result_ErrorCode.Descriptor instead.
func (Result_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_action_proto_rawDescGZIP(), []int{0, 1}
}

// FIXME: WTF?
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode Result_ResultCode `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3,enum=pb.Result_ResultCode" json:"result_code,omitempty"` // required
	ErrorCode  Result_ErrorCode  `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3,enum=pb.Result_ErrorCode" json:"error_code,omitempty"`     // optional
	ErrorText  string            `protobuf:"bytes,3,opt,name=error_text,json=errorText,proto3" json:"error_text,omitempty"`                               // optional
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_proto_action_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetResultCode() Result_ResultCode {
	if x != nil {
		return x.ResultCode
	}
	return Result_NONE
}

func (x *Result) GetErrorCode() Result_ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return Result_ERROR_UNKNOWN
}

func (x *Result) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

var File_api_proto_action_proto protoreflect.FileDescriptor

var file_api_proto_action_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xf3, 0x02, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x33, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54,
	0x65, 0x78, 0x74, 0x22, 0x5d, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57,
	0x41, 0x49, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x55, 0x4e, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x06, 0x22, 0x7e, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48,
	0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x55,
	0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x04, 0x12, 0x13, 0x0a,
	0x0f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44,
	0x10, 0x05, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x63, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x73, 0x6b, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_action_proto_rawDescOnce sync.Once
	file_api_proto_action_proto_rawDescData = file_api_proto_action_proto_rawDesc
)

func file_api_proto_action_proto_rawDescGZIP() []byte {
	file_api_proto_action_proto_rawDescOnce.Do(func() {
		file_api_proto_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_action_proto_rawDescData)
	})
	return file_api_proto_action_proto_rawDescData
}

var file_api_proto_action_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_proto_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_action_proto_goTypes = []interface{}{
	(Result_ResultCode)(0), // 0: pb.Result.ResultCode
	(Result_ErrorCode)(0),  // 1: pb.Result.ErrorCode
	(*Result)(nil),         // 2: pb.Result
}
var file_api_proto_action_proto_depIdxs = []int32{
	0, // 0: pb.Result.result_code:type_name -> pb.Result.ResultCode
	1, // 1: pb.Result.error_code:type_name -> pb.Result.ErrorCode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_action_proto_init() }
func file_api_proto_action_proto_init() {
	if File_api_proto_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_action_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_action_proto_goTypes,
		DependencyIndexes: file_api_proto_action_proto_depIdxs,
		EnumInfos:         file_api_proto_action_proto_enumTypes,
		MessageInfos:      file_api_proto_action_proto_msgTypes,
	}.Build()
	File_api_proto_action_proto = out.File
	file_api_proto_action_proto_rawDesc = nil
	file_api_proto_action_proto_goTypes = nil
	file_api_proto_action_proto_depIdxs = nil
}
