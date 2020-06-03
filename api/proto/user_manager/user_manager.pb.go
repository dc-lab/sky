// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: api/proto/user_manager/user_manager.proto

// build *.pb.go from project root
// protoc -I. --go_out=plugins=grpc,paths=source_relative:. api/proto/user_manager/*.proto

package user_manager

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // required
}

func (x *TUser) Reset() {
	*x = TUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_manager_user_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TUser) ProtoMessage() {}

func (x *TUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_manager_user_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TUser.ProtoReflect.Descriptor instead.
func (*TUser) Descriptor() ([]byte, []int) {
	return file_api_proto_user_manager_user_manager_proto_rawDescGZIP(), []int{0}
}

func (x *TUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`       // required
	Name  string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   // required
	Users []string `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"` // required
}

func (x *TGroup) Reset() {
	*x = TGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_manager_user_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TGroup) ProtoMessage() {}

func (x *TGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_manager_user_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TGroup.ProtoReflect.Descriptor instead.
func (*TGroup) Descriptor() ([]byte, []int) {
	return file_api_proto_user_manager_user_manager_proto_rawDescGZIP(), []int{1}
}

func (x *TGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TGroup) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_api_proto_user_manager_user_manager_proto protoreflect.FileDescriptor

var file_api_proto_user_manager_user_manager_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x17, 0x0a, 0x05, 0x54, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x42, 0x0a, 0x06, 0x54, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x4d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x63, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x73, 0x6b, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_user_manager_user_manager_proto_rawDescOnce sync.Once
	file_api_proto_user_manager_user_manager_proto_rawDescData = file_api_proto_user_manager_user_manager_proto_rawDesc
)

func file_api_proto_user_manager_user_manager_proto_rawDescGZIP() []byte {
	file_api_proto_user_manager_user_manager_proto_rawDescOnce.Do(func() {
		file_api_proto_user_manager_user_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_user_manager_user_manager_proto_rawDescData)
	})
	return file_api_proto_user_manager_user_manager_proto_rawDescData
}

var file_api_proto_user_manager_user_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_user_manager_user_manager_proto_goTypes = []interface{}{
	(*TUser)(nil),  // 0: user_manager.TUser
	(*TGroup)(nil), // 1: user_manager.TGroup
}
var file_api_proto_user_manager_user_manager_proto_depIdxs = []int32{
	0, // 0: user_manager.UserManager.GetUserGroups:input_type -> user_manager.TUser
	1, // 1: user_manager.UserManager.GetUserGroups:output_type -> user_manager.TGroup
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_user_manager_user_manager_proto_init() }
func file_api_proto_user_manager_user_manager_proto_init() {
	if File_api_proto_user_manager_user_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_user_manager_user_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TUser); i {
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
		file_api_proto_user_manager_user_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TGroup); i {
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
			RawDescriptor: file_api_proto_user_manager_user_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_user_manager_user_manager_proto_goTypes,
		DependencyIndexes: file_api_proto_user_manager_user_manager_proto_depIdxs,
		MessageInfos:      file_api_proto_user_manager_user_manager_proto_msgTypes,
	}.Build()
	File_api_proto_user_manager_user_manager_proto = out.File
	file_api_proto_user_manager_user_manager_proto_rawDesc = nil
	file_api_proto_user_manager_user_manager_proto_goTypes = nil
	file_api_proto_user_manager_user_manager_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserManagerClient is the client API for UserManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserManagerClient interface {
	GetUserGroups(ctx context.Context, in *TUser, opts ...grpc.CallOption) (UserManager_GetUserGroupsClient, error)
}

type userManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagerClient(cc grpc.ClientConnInterface) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) GetUserGroups(ctx context.Context, in *TUser, opts ...grpc.CallOption) (UserManager_GetUserGroupsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserManager_serviceDesc.Streams[0], "/user_manager.UserManager/GetUserGroups", opts...)
	if err != nil {
		return nil, err
	}
	x := &userManagerGetUserGroupsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserManager_GetUserGroupsClient interface {
	Recv() (*TGroup, error)
	grpc.ClientStream
}

type userManagerGetUserGroupsClient struct {
	grpc.ClientStream
}

func (x *userManagerGetUserGroupsClient) Recv() (*TGroup, error) {
	m := new(TGroup)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserManagerServer is the server API for UserManager service.
type UserManagerServer interface {
	GetUserGroups(*TUser, UserManager_GetUserGroupsServer) error
}

// UnimplementedUserManagerServer can be embedded to have forward compatible implementations.
type UnimplementedUserManagerServer struct {
}

func (*UnimplementedUserManagerServer) GetUserGroups(*TUser, UserManager_GetUserGroupsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserGroups not implemented")
}

func RegisterUserManagerServer(s *grpc.Server, srv UserManagerServer) {
	s.RegisterService(&_UserManager_serviceDesc, srv)
}

func _UserManager_GetUserGroups_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TUser)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserManagerServer).GetUserGroups(m, &userManagerGetUserGroupsServer{stream})
}

type UserManager_GetUserGroupsServer interface {
	Send(*TGroup) error
	grpc.ServerStream
}

type userManagerGetUserGroupsServer struct {
	grpc.ServerStream
}

func (x *userManagerGetUserGroupsServer) Send(m *TGroup) error {
	return x.ServerStream.SendMsg(m)
}

var _UserManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user_manager.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserGroups",
			Handler:       _UserManager_GetUserGroups_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/proto/user_manager/user_manager.proto",
}
