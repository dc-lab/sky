// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: api/proto/common/task_results.proto

// build *.pb.go from project root
// protoc -I. --go_out=plugins=grpc,paths=source_relative:. api/proto/common/*.proto

package common

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TaskFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path             string               `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Size             uint64               `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Hash             string               `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	CreationTime     *timestamp.Timestamp `protobuf:"bytes,4,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	ModificationTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=modification_time,json=modificationTime,proto3" json:"modification_time,omitempty"`
}

func (x *TaskFile) Reset() {
	*x = TaskFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_task_results_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskFile) ProtoMessage() {}

func (x *TaskFile) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_task_results_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskFile.ProtoReflect.Descriptor instead.
func (*TaskFile) Descriptor() ([]byte, []int) {
	return file_api_proto_common_task_results_proto_rawDescGZIP(), []int{0}
}

func (x *TaskFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TaskFile) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TaskFile) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *TaskFile) GetCreationTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *TaskFile) GetModificationTime() *timestamp.Timestamp {
	if x != nil {
		return x.ModificationTime
	}
	return nil
}

var File_api_proto_common_task_results_proto protoreflect.FileDescriptor

var file_api_proto_common_task_results_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0,
	0x01, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x10, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x63, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x73, 0x6b, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_common_task_results_proto_rawDescOnce sync.Once
	file_api_proto_common_task_results_proto_rawDescData = file_api_proto_common_task_results_proto_rawDesc
)

func file_api_proto_common_task_results_proto_rawDescGZIP() []byte {
	file_api_proto_common_task_results_proto_rawDescOnce.Do(func() {
		file_api_proto_common_task_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_common_task_results_proto_rawDescData)
	})
	return file_api_proto_common_task_results_proto_rawDescData
}

var file_api_proto_common_task_results_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_common_task_results_proto_goTypes = []interface{}{
	(*TaskFile)(nil),            // 0: common.TaskFile
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_api_proto_common_task_results_proto_depIdxs = []int32{
	1, // 0: common.TaskFile.creation_time:type_name -> google.protobuf.Timestamp
	1, // 1: common.TaskFile.modification_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_common_task_results_proto_init() }
func file_api_proto_common_task_results_proto_init() {
	if File_api_proto_common_task_results_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_common_task_results_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskFile); i {
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
			RawDescriptor: file_api_proto_common_task_results_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_common_task_results_proto_goTypes,
		DependencyIndexes: file_api_proto_common_task_results_proto_depIdxs,
		MessageInfos:      file_api_proto_common_task_results_proto_msgTypes,
	}.Build()
	File_api_proto_common_task_results_proto = out.File
	file_api_proto_common_task_results_proto_rawDesc = nil
	file_api_proto_common_task_results_proto_goTypes = nil
	file_api_proto_common_task_results_proto_depIdxs = nil
}
