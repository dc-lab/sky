// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: api/proto/cloud_manager/entity.proto

// build *.pb.go from project root
// protoc -I. --go_out=plugins=grpc,paths=source_relative:. api/proto/cloud_manager/*.proto

package cloud_manager

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

type ETransactionStatus int32

const (
	ETransactionStatus_UNKNOWN   ETransactionStatus = 0
	ETransactionStatus_ACTIVE    ETransactionStatus = 1
	ETransactionStatus_COMPLETED ETransactionStatus = 2
	ETransactionStatus_EXPIRED   ETransactionStatus = 3
	ETransactionStatus_CANCELLED ETransactionStatus = 4
)

// Enum value maps for ETransactionStatus.
var (
	ETransactionStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "ACTIVE",
		2: "COMPLETED",
		3: "EXPIRED",
		4: "CANCELLED",
	}
	ETransactionStatus_value = map[string]int32{
		"UNKNOWN":   0,
		"ACTIVE":    1,
		"COMPLETED": 2,
		"EXPIRED":   3,
		"CANCELLED": 4,
	}
)

func (x ETransactionStatus) Enum() *ETransactionStatus {
	p := new(ETransactionStatus)
	*p = x
	return p
}

func (x ETransactionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ETransactionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_cloud_manager_entity_proto_enumTypes[0].Descriptor()
}

func (ETransactionStatus) Type() protoreflect.EnumType {
	return &file_api_proto_cloud_manager_entity_proto_enumTypes[0]
}

func (x ETransactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ETransactionStatus.Descriptor instead.
func (ETransactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_cloud_manager_entity_proto_rawDescGZIP(), []int{0}
}

type TTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *TTransaction) Reset() {
	*x = TTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cloud_manager_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TTransaction) ProtoMessage() {}

func (x *TTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cloud_manager_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TTransaction.ProtoReflect.Descriptor instead.
func (*TTransaction) Descriptor() ([]byte, []int) {
	return file_api_proto_cloud_manager_entity_proto_rawDescGZIP(), []int{0}
}

func (x *TTransaction) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type TDockerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registry   string `protobuf:"bytes,1,opt,name=registry,proto3" json:"registry,omitempty"`
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Image      string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Tag        string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *TDockerImage) Reset() {
	*x = TDockerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cloud_manager_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDockerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDockerImage) ProtoMessage() {}

func (x *TDockerImage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cloud_manager_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDockerImage.ProtoReflect.Descriptor instead.
func (*TDockerImage) Descriptor() ([]byte, []int) {
	return file_api_proto_cloud_manager_entity_proto_rawDescGZIP(), []int{1}
}

func (x *TDockerImage) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *TDockerImage) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *TDockerImage) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *TDockerImage) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

var File_api_proto_cloud_manager_entity_proto protoreflect.FileDescriptor

var file_api_proto_cloud_manager_entity_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x35,
	0x0a, 0x0c, 0x54, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x0c, 0x54, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x2a, 0x58, 0x0a, 0x12, 0x45, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x04, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x63, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x73, 0x6b, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_cloud_manager_entity_proto_rawDescOnce sync.Once
	file_api_proto_cloud_manager_entity_proto_rawDescData = file_api_proto_cloud_manager_entity_proto_rawDesc
)

func file_api_proto_cloud_manager_entity_proto_rawDescGZIP() []byte {
	file_api_proto_cloud_manager_entity_proto_rawDescOnce.Do(func() {
		file_api_proto_cloud_manager_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_cloud_manager_entity_proto_rawDescData)
	})
	return file_api_proto_cloud_manager_entity_proto_rawDescData
}

var file_api_proto_cloud_manager_entity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_cloud_manager_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_cloud_manager_entity_proto_goTypes = []interface{}{
	(ETransactionStatus)(0), // 0: entity.ETransactionStatus
	(*TTransaction)(nil),    // 1: entity.TTransaction
	(*TDockerImage)(nil),    // 2: entity.TDockerImage
}
var file_api_proto_cloud_manager_entity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_cloud_manager_entity_proto_init() }
func file_api_proto_cloud_manager_entity_proto_init() {
	if File_api_proto_cloud_manager_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_cloud_manager_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TTransaction); i {
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
		file_api_proto_cloud_manager_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDockerImage); i {
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
			RawDescriptor: file_api_proto_cloud_manager_entity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_cloud_manager_entity_proto_goTypes,
		DependencyIndexes: file_api_proto_cloud_manager_entity_proto_depIdxs,
		EnumInfos:         file_api_proto_cloud_manager_entity_proto_enumTypes,
		MessageInfos:      file_api_proto_cloud_manager_entity_proto_msgTypes,
	}.Build()
	File_api_proto_cloud_manager_entity_proto = out.File
	file_api_proto_cloud_manager_entity_proto_rawDesc = nil
	file_api_proto_cloud_manager_entity_proto_goTypes = nil
	file_api_proto_cloud_manager_entity_proto_depIdxs = nil
}