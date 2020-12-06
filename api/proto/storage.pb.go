// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: api/proto/storage.proto

// build *.pb.go from project root
// protoc -I. --go_out=plugins=grpc,paths=source_relative:. api/proto/data_manager/*.proto

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

type ValidateUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId      string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	FileId      string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	UserId      string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UploadToken string `protobuf:"bytes,4,opt,name=upload_token,json=uploadToken,proto3" json:"upload_token,omitempty"`
}

func (x *ValidateUploadRequest) Reset() {
	*x = ValidateUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUploadRequest) ProtoMessage() {}

func (x *ValidateUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUploadRequest.ProtoReflect.Descriptor instead.
func (*ValidateUploadRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateUploadRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *ValidateUploadRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ValidateUploadRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ValidateUploadRequest) GetUploadToken() string {
	if x != nil {
		return x.UploadToken
	}
	return ""
}

type ValidateUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow bool `protobuf:"varint,1,opt,name=allow,proto3" json:"allow,omitempty"`
}

func (x *ValidateUploadResponse) Reset() {
	*x = ValidateUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUploadResponse) ProtoMessage() {}

func (x *ValidateUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUploadResponse.ProtoReflect.Descriptor instead.
func (*ValidateUploadResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateUploadResponse) GetAllow() bool {
	if x != nil {
		return x.Allow
	}
	return false
}

type SubmitFileHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	FileId string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Hash   string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *SubmitFileHashRequest) Reset() {
	*x = SubmitFileHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitFileHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitFileHashRequest) ProtoMessage() {}

func (x *SubmitFileHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitFileHashRequest.ProtoReflect.Descriptor instead.
func (*SubmitFileHashRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitFileHashRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *SubmitFileHashRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *SubmitFileHashRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SubmitFileHashRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type SubmitFileHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow bool `protobuf:"varint,1,opt,name=allow,proto3" json:"allow,omitempty"`
}

func (x *SubmitFileHashResponse) Reset() {
	*x = SubmitFileHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitFileHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitFileHashResponse) ProtoMessage() {}

func (x *SubmitFileHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitFileHashResponse.ProtoReflect.Descriptor instead.
func (*SubmitFileHashResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitFileHashResponse) GetAllow() bool {
	if x != nil {
		return x.Allow
	}
	return false
}

type GetFileHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FileId string `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (x *GetFileHashRequest) Reset() {
	*x = GetFileHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileHashRequest) ProtoMessage() {}

func (x *GetFileHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileHashRequest.ProtoReflect.Descriptor instead.
func (*GetFileHashRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{4}
}

func (x *GetFileHashRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *GetFileHashRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetFileHashRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type GetFileHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow bool   `protobuf:"varint,1,opt,name=allow,proto3" json:"allow,omitempty"`
	Hash  string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetFileHashResponse) Reset() {
	*x = GetFileHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileHashResponse) ProtoMessage() {}

func (x *GetFileHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileHashResponse.ProtoReflect.Descriptor instead.
func (*GetFileHashResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{5}
}

func (x *GetFileHashResponse) GetAllow() bool {
	if x != nil {
		return x.Allow
	}
	return false
}

func (x *GetFileHashResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type NodeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId     string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	FreeSpace  int64    `protobuf:"varint,2,opt,name=free_space,json=freeSpace,proto3" json:"free_space,omitempty"`
	BlobHashes []string `protobuf:"bytes,3,rep,name=blob_hashes,json=blobHashes,proto3" json:"blob_hashes,omitempty"`
}

func (x *NodeStatus) Reset() {
	*x = NodeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatus) ProtoMessage() {}

func (x *NodeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatus.ProtoReflect.Descriptor instead.
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{6}
}

func (x *NodeStatus) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeStatus) GetFreeSpace() int64 {
	if x != nil {
		return x.FreeSpace
	}
	return 0
}

func (x *NodeStatus) GetBlobHashes() []string {
	if x != nil {
		return x.BlobHashes
	}
	return nil
}

type NodeTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlobHashes []string `protobuf:"bytes,1,rep,name=blob_hashes,json=blobHashes,proto3" json:"blob_hashes,omitempty"`
}

func (x *NodeTarget) Reset() {
	*x = NodeTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeTarget) ProtoMessage() {}

func (x *NodeTarget) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeTarget.ProtoReflect.Descriptor instead.
func (*NodeTarget) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{7}
}

func (x *NodeTarget) GetBlobHashes() []string {
	if x != nil {
		return x.BlobHashes
	}
	return nil
}

type ResolveBlobReplicasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ResolveBlobReplicasRequest) Reset() {
	*x = ResolveBlobReplicasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveBlobReplicasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveBlobReplicasRequest) ProtoMessage() {}

func (x *ResolveBlobReplicasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveBlobReplicasRequest.ProtoReflect.Descriptor instead.
func (*ResolveBlobReplicasRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{8}
}

func (x *ResolveBlobReplicasRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type ResolveBlobReplicasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []string `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *ResolveBlobReplicasResponse) Reset() {
	*x = ResolveBlobReplicasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_storage_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveBlobReplicasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveBlobReplicasResponse) ProtoMessage() {}

func (x *ResolveBlobReplicasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_storage_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveBlobReplicasResponse.ProtoReflect.Descriptor instead.
func (*ResolveBlobReplicasResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_storage_proto_rawDescGZIP(), []int{9}
}

func (x *ResolveBlobReplicasResponse) GetLocations() []string {
	if x != nil {
		return x.Locations
	}
	return nil
}

var File_api_proto_storage_proto protoreflect.FileDescriptor

var file_api_proto_storage_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x85, 0x01,
	0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2e, 0x0a, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0x76, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x2e, 0x0a,
	0x16, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0x5f, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3f,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22,
	0x65, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x72, 0x65, 0x65,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x62,
	0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x48,
	0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x1a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x3b, 0x0a, 0x1b, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe4, 0x02, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x49, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x4c, 0x6f, 0x6f, 0x70, 0x12,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22,
	0x00, 0x12, 0x58, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x63, 0x2d, 0x6c, 0x61, 0x62,
	0x2f, 0x73, 0x6b, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_storage_proto_rawDescOnce sync.Once
	file_api_proto_storage_proto_rawDescData = file_api_proto_storage_proto_rawDesc
)

func file_api_proto_storage_proto_rawDescGZIP() []byte {
	file_api_proto_storage_proto_rawDescOnce.Do(func() {
		file_api_proto_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_storage_proto_rawDescData)
	})
	return file_api_proto_storage_proto_rawDescData
}

var file_api_proto_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_storage_proto_goTypes = []interface{}{
	(*ValidateUploadRequest)(nil),       // 0: pb.ValidateUploadRequest
	(*ValidateUploadResponse)(nil),      // 1: pb.ValidateUploadResponse
	(*SubmitFileHashRequest)(nil),       // 2: pb.SubmitFileHashRequest
	(*SubmitFileHashResponse)(nil),      // 3: pb.SubmitFileHashResponse
	(*GetFileHashRequest)(nil),          // 4: pb.GetFileHashRequest
	(*GetFileHashResponse)(nil),         // 5: pb.GetFileHashResponse
	(*NodeStatus)(nil),                  // 6: pb.NodeStatus
	(*NodeTarget)(nil),                  // 7: pb.NodeTarget
	(*ResolveBlobReplicasRequest)(nil),  // 8: pb.ResolveBlobReplicasRequest
	(*ResolveBlobReplicasResponse)(nil), // 9: pb.ResolveBlobReplicasResponse
}
var file_api_proto_storage_proto_depIdxs = []int32{
	0, // 0: pb.Master.ValidateUpload:input_type -> pb.ValidateUploadRequest
	2, // 1: pb.Master.SubmitFileHash:input_type -> pb.SubmitFileHashRequest
	4, // 2: pb.Master.GetFileHash:input_type -> pb.GetFileHashRequest
	6, // 3: pb.Master.Loop:input_type -> pb.NodeStatus
	8, // 4: pb.Master.ResolveBlobReplicas:input_type -> pb.ResolveBlobReplicasRequest
	1, // 5: pb.Master.ValidateUpload:output_type -> pb.ValidateUploadResponse
	3, // 6: pb.Master.SubmitFileHash:output_type -> pb.SubmitFileHashResponse
	5, // 7: pb.Master.GetFileHash:output_type -> pb.GetFileHashResponse
	7, // 8: pb.Master.Loop:output_type -> pb.NodeTarget
	9, // 9: pb.Master.ResolveBlobReplicas:output_type -> pb.ResolveBlobReplicasResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_storage_proto_init() }
func file_api_proto_storage_proto_init() {
	if File_api_proto_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUploadRequest); i {
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
		file_api_proto_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateUploadResponse); i {
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
		file_api_proto_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitFileHashRequest); i {
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
		file_api_proto_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitFileHashResponse); i {
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
		file_api_proto_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileHashRequest); i {
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
		file_api_proto_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileHashResponse); i {
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
		file_api_proto_storage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatus); i {
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
		file_api_proto_storage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeTarget); i {
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
		file_api_proto_storage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveBlobReplicasRequest); i {
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
		file_api_proto_storage_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveBlobReplicasResponse); i {
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
			RawDescriptor: file_api_proto_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_storage_proto_goTypes,
		DependencyIndexes: file_api_proto_storage_proto_depIdxs,
		MessageInfos:      file_api_proto_storage_proto_msgTypes,
	}.Build()
	File_api_proto_storage_proto = out.File
	file_api_proto_storage_proto_rawDesc = nil
	file_api_proto_storage_proto_goTypes = nil
	file_api_proto_storage_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	ValidateUpload(ctx context.Context, in *ValidateUploadRequest, opts ...grpc.CallOption) (*ValidateUploadResponse, error)
	SubmitFileHash(ctx context.Context, in *SubmitFileHashRequest, opts ...grpc.CallOption) (*SubmitFileHashResponse, error)
	GetFileHash(ctx context.Context, in *GetFileHashRequest, opts ...grpc.CallOption) (*GetFileHashResponse, error)
	Loop(ctx context.Context, in *NodeStatus, opts ...grpc.CallOption) (*NodeTarget, error)
	ResolveBlobReplicas(ctx context.Context, in *ResolveBlobReplicasRequest, opts ...grpc.CallOption) (*ResolveBlobReplicasResponse, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) ValidateUpload(ctx context.Context, in *ValidateUploadRequest, opts ...grpc.CallOption) (*ValidateUploadResponse, error) {
	out := new(ValidateUploadResponse)
	err := c.cc.Invoke(ctx, "/pb.Master/ValidateUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) SubmitFileHash(ctx context.Context, in *SubmitFileHashRequest, opts ...grpc.CallOption) (*SubmitFileHashResponse, error) {
	out := new(SubmitFileHashResponse)
	err := c.cc.Invoke(ctx, "/pb.Master/SubmitFileHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetFileHash(ctx context.Context, in *GetFileHashRequest, opts ...grpc.CallOption) (*GetFileHashResponse, error) {
	out := new(GetFileHashResponse)
	err := c.cc.Invoke(ctx, "/pb.Master/GetFileHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Loop(ctx context.Context, in *NodeStatus, opts ...grpc.CallOption) (*NodeTarget, error) {
	out := new(NodeTarget)
	err := c.cc.Invoke(ctx, "/pb.Master/Loop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) ResolveBlobReplicas(ctx context.Context, in *ResolveBlobReplicasRequest, opts ...grpc.CallOption) (*ResolveBlobReplicasResponse, error) {
	out := new(ResolveBlobReplicasResponse)
	err := c.cc.Invoke(ctx, "/pb.Master/ResolveBlobReplicas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	ValidateUpload(context.Context, *ValidateUploadRequest) (*ValidateUploadResponse, error)
	SubmitFileHash(context.Context, *SubmitFileHashRequest) (*SubmitFileHashResponse, error)
	GetFileHash(context.Context, *GetFileHashRequest) (*GetFileHashResponse, error)
	Loop(context.Context, *NodeStatus) (*NodeTarget, error)
	ResolveBlobReplicas(context.Context, *ResolveBlobReplicasRequest) (*ResolveBlobReplicasResponse, error)
}

// UnimplementedMasterServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (*UnimplementedMasterServer) ValidateUpload(context.Context, *ValidateUploadRequest) (*ValidateUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUpload not implemented")
}
func (*UnimplementedMasterServer) SubmitFileHash(context.Context, *SubmitFileHashRequest) (*SubmitFileHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitFileHash not implemented")
}
func (*UnimplementedMasterServer) GetFileHash(context.Context, *GetFileHashRequest) (*GetFileHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileHash not implemented")
}
func (*UnimplementedMasterServer) Loop(context.Context, *NodeStatus) (*NodeTarget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Loop not implemented")
}
func (*UnimplementedMasterServer) ResolveBlobReplicas(context.Context, *ResolveBlobReplicasRequest) (*ResolveBlobReplicasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveBlobReplicas not implemented")
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_ValidateUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ValidateUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Master/ValidateUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ValidateUpload(ctx, req.(*ValidateUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_SubmitFileHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitFileHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).SubmitFileHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Master/SubmitFileHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).SubmitFileHash(ctx, req.(*SubmitFileHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetFileHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetFileHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Master/GetFileHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetFileHash(ctx, req.(*GetFileHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Loop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Loop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Master/Loop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Loop(ctx, req.(*NodeStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_ResolveBlobReplicas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveBlobReplicasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).ResolveBlobReplicas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Master/ResolveBlobReplicas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).ResolveBlobReplicas(ctx, req.(*ResolveBlobReplicasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateUpload",
			Handler:    _Master_ValidateUpload_Handler,
		},
		{
			MethodName: "SubmitFileHash",
			Handler:    _Master_SubmitFileHash_Handler,
		},
		{
			MethodName: "GetFileHash",
			Handler:    _Master_GetFileHash_Handler,
		},
		{
			MethodName: "Loop",
			Handler:    _Master_Loop_Handler,
		},
		{
			MethodName: "ResolveBlobReplicas",
			Handler:    _Master_ResolveBlobReplicas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/storage.proto",
}