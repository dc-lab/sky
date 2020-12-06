// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: api/proto/user_manager.proto

// build *.pb.go from project root
// protoc -I. --go_out=plugins=grpc,paths=source_relative:. api/proto/user_manager/*.proto

package proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // required
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_proto_user_manager_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`       // required
	Name  string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   // required
	Users []string `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"` // required
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_api_proto_user_manager_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_api_proto_user_manager_proto protoreflect.FileDescriptor

var file_api_proto_user_manager_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x37, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0x00, 0x30, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x63, 0x2d, 0x6c, 0x61, 0x62, 0x2f, 0x73, 0x6b, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_user_manager_proto_rawDescOnce sync.Once
	file_api_proto_user_manager_proto_rawDescData = file_api_proto_user_manager_proto_rawDesc
)

func file_api_proto_user_manager_proto_rawDescGZIP() []byte {
	file_api_proto_user_manager_proto_rawDescOnce.Do(func() {
		file_api_proto_user_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_user_manager_proto_rawDescData)
	})
	return file_api_proto_user_manager_proto_rawDescData
}

var file_api_proto_user_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_user_manager_proto_goTypes = []interface{}{
	(*User)(nil),  // 0: pb.User
	(*Group)(nil), // 1: pb.Group
}
var file_api_proto_user_manager_proto_depIdxs = []int32{
	0, // 0: pb.UserManager.GetUserGroups:input_type -> pb.User
	1, // 1: pb.UserManager.GetUserGroups:output_type -> pb.Group
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_user_manager_proto_init() }
func file_api_proto_user_manager_proto_init() {
	if File_api_proto_user_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_user_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_api_proto_user_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
			RawDescriptor: file_api_proto_user_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_user_manager_proto_goTypes,
		DependencyIndexes: file_api_proto_user_manager_proto_depIdxs,
		MessageInfos:      file_api_proto_user_manager_proto_msgTypes,
	}.Build()
	File_api_proto_user_manager_proto = out.File
	file_api_proto_user_manager_proto_rawDesc = nil
	file_api_proto_user_manager_proto_goTypes = nil
	file_api_proto_user_manager_proto_depIdxs = nil
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
	GetUserGroups(ctx context.Context, in *User, opts ...grpc.CallOption) (UserManager_GetUserGroupsClient, error)
}

type userManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagerClient(cc grpc.ClientConnInterface) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) GetUserGroups(ctx context.Context, in *User, opts ...grpc.CallOption) (UserManager_GetUserGroupsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserManager_serviceDesc.Streams[0], "/pb.UserManager/GetUserGroups", opts...)
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
	Recv() (*Group, error)
	grpc.ClientStream
}

type userManagerGetUserGroupsClient struct {
	grpc.ClientStream
}

func (x *userManagerGetUserGroupsClient) Recv() (*Group, error) {
	m := new(Group)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserManagerServer is the server API for UserManager service.
type UserManagerServer interface {
	GetUserGroups(*User, UserManager_GetUserGroupsServer) error
}

// UnimplementedUserManagerServer can be embedded to have forward compatible implementations.
type UnimplementedUserManagerServer struct {
}

func (*UnimplementedUserManagerServer) GetUserGroups(*User, UserManager_GetUserGroupsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserGroups not implemented")
}

func RegisterUserManagerServer(s *grpc.Server, srv UserManagerServer) {
	s.RegisterService(&_UserManager_serviceDesc, srv)
}

func _UserManager_GetUserGroups_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserManagerServer).GetUserGroups(m, &userManagerGetUserGroupsServer{stream})
}

type UserManager_GetUserGroupsServer interface {
	Send(*Group) error
	grpc.ServerStream
}

type userManagerGetUserGroupsServer struct {
	grpc.ServerStream
}

func (x *userManagerGetUserGroupsServer) Send(m *Group) error {
	return x.ServerStream.SendMsg(m)
}

var _UserManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserGroups",
			Handler:       _UserManager_GetUserGroups_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/proto/user_manager.proto",
}