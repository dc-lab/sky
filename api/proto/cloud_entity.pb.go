// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: api/proto/cloud_entity.proto

// build *.pb.go from project root
// protoc -I. --go_out=plugins=grpc,paths=source_relative:. api/proto/cloud_manager/*.proto

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

type TransactionStatus int32

const (
	TransactionStatus_UNKNOWN   TransactionStatus = 0
	TransactionStatus_ACTIVE    TransactionStatus = 1
	TransactionStatus_COMPLETED TransactionStatus = 2
	TransactionStatus_EXPIRED   TransactionStatus = 3
	TransactionStatus_CANCELLED TransactionStatus = 4
)

// Enum value maps for TransactionStatus.
var (
	TransactionStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "ACTIVE",
		2: "COMPLETED",
		3: "EXPIRED",
		4: "CANCELLED",
	}
	TransactionStatus_value = map[string]int32{
		"UNKNOWN":   0,
		"ACTIVE":    1,
		"COMPLETED": 2,
		"EXPIRED":   3,
		"CANCELLED": 4,
	}
)

func (x TransactionStatus) Enum() *TransactionStatus {
	p := new(TransactionStatus)
	*p = x
	return p
}

func (x TransactionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_cloud_entity_proto_enumTypes[0].Descriptor()
}

func (TransactionStatus) Type() protoreflect.EnumType {
	return &file_api_proto_cloud_entity_proto_enumTypes[0]
}

func (x TransactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionStatus.Descriptor instead.
func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_cloud_entity_proto_rawDescGZIP(), []int{0}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cloud_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cloud_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_api_proto_cloud_entity_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type DockerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registry   string `protobuf:"bytes,1,opt,name=registry,proto3" json:"registry,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Image      string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Tag        string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *DockerImage) Reset() {
	*x = DockerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cloud_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerImage) ProtoMessage() {}

func (x *DockerImage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cloud_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerImage.ProtoReflect.Descriptor instead.
func (*DockerImage) Descriptor() ([]byte, []int) {
	return file_api_proto_cloud_entity_proto_rawDescGZIP(), []int{1}
}

func (x *DockerImage) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *DockerImage) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *DockerImage) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *DockerImage) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

var File_api_proto_cloud_entity_proto protoreflect.FileDescriptor

var file_api_proto_cloud_entity_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x34, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x0b, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x2a, 0x57, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x10, 0x04, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x63, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x73, 0x6b, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_cloud_entity_proto_rawDescOnce sync.Once
	file_api_proto_cloud_entity_proto_rawDescData = file_api_proto_cloud_entity_proto_rawDesc
)

func file_api_proto_cloud_entity_proto_rawDescGZIP() []byte {
	file_api_proto_cloud_entity_proto_rawDescOnce.Do(func() {
		file_api_proto_cloud_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_cloud_entity_proto_rawDescData)
	})
	return file_api_proto_cloud_entity_proto_rawDescData
}

var file_api_proto_cloud_entity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_cloud_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_cloud_entity_proto_goTypes = []interface{}{
	(TransactionStatus)(0), // 0: pb.TransactionStatus
	(*Transaction)(nil),    // 1: pb.Transaction
	(*DockerImage)(nil),    // 2: pb.DockerImage
}
var file_api_proto_cloud_entity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_cloud_entity_proto_init() }
func file_api_proto_cloud_entity_proto_init() {
	if File_api_proto_cloud_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_cloud_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_api_proto_cloud_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerImage); i {
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
			RawDescriptor: file_api_proto_cloud_entity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_cloud_entity_proto_goTypes,
		DependencyIndexes: file_api_proto_cloud_entity_proto_depIdxs,
		EnumInfos:         file_api_proto_cloud_entity_proto_enumTypes,
		MessageInfos:      file_api_proto_cloud_entity_proto_msgTypes,
	}.Build()
	File_api_proto_cloud_entity_proto = out.File
	file_api_proto_cloud_entity_proto_rawDesc = nil
	file_api_proto_cloud_entity_proto_goTypes = nil
	file_api_proto_cloud_entity_proto_depIdxs = nil
}
