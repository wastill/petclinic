// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: petclinic/role.proto

package petclinic

import (
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

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{0}
}

type CreateRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRoleReply) Reset() {
	*x = CreateRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleReply) ProtoMessage() {}

func (x *CreateRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleReply.ProtoReflect.Descriptor instead.
func (*CreateRoleReply) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{1}
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{2}
}

type UpdateRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRoleReply) Reset() {
	*x = UpdateRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleReply) ProtoMessage() {}

func (x *UpdateRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleReply.ProtoReflect.Descriptor instead.
func (*UpdateRoleReply) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{3}
}

type DeleteRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{4}
}

type DeleteRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRoleReply) Reset() {
	*x = DeleteRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleReply) ProtoMessage() {}

func (x *DeleteRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleReply.ProtoReflect.Descriptor instead.
func (*DeleteRoleReply) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{5}
}

type GetRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRoleRequest) Reset() {
	*x = GetRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleRequest) ProtoMessage() {}

func (x *GetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleRequest.ProtoReflect.Descriptor instead.
func (*GetRoleRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{6}
}

type GetRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRoleReply) Reset() {
	*x = GetRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleReply) ProtoMessage() {}

func (x *GetRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleReply.ProtoReflect.Descriptor instead.
func (*GetRoleReply) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{7}
}

type ListRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRoleRequest) Reset() {
	*x = ListRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoleRequest) ProtoMessage() {}

func (x *ListRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoleRequest.ProtoReflect.Descriptor instead.
func (*ListRoleRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{8}
}

type ListRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRoleReply) Reset() {
	*x = ListRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_role_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoleReply) ProtoMessage() {}

func (x *ListRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_role_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoleReply.ProtoReflect.Descriptor instead.
func (*ListRoleReply) Descriptor() ([]byte, []int) {
	return file_petclinic_role_proto_rawDescGZIP(), []int{9}
}

var File_petclinic_role_proto protoreflect.FileDescriptor

var file_petclinic_role_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63,
	0x6c, 0x69, 0x6e, 0x69, 0x63, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0x87, 0x03, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65,
	0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74,
	0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x48, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x4d, 0x0a,
	0x0d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x50, 0x01,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x73,
	0x74, 0x69, 0x6c, 0x6c, 0x2f, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e,
	0x69, 0x63, 0x3b, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_petclinic_role_proto_rawDescOnce sync.Once
	file_petclinic_role_proto_rawDescData = file_petclinic_role_proto_rawDesc
)

func file_petclinic_role_proto_rawDescGZIP() []byte {
	file_petclinic_role_proto_rawDescOnce.Do(func() {
		file_petclinic_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_petclinic_role_proto_rawDescData)
	})
	return file_petclinic_role_proto_rawDescData
}

var file_petclinic_role_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_petclinic_role_proto_goTypes = []interface{}{
	(*CreateRoleRequest)(nil), // 0: api.petclinic.CreateRoleRequest
	(*CreateRoleReply)(nil),   // 1: api.petclinic.CreateRoleReply
	(*UpdateRoleRequest)(nil), // 2: api.petclinic.UpdateRoleRequest
	(*UpdateRoleReply)(nil),   // 3: api.petclinic.UpdateRoleReply
	(*DeleteRoleRequest)(nil), // 4: api.petclinic.DeleteRoleRequest
	(*DeleteRoleReply)(nil),   // 5: api.petclinic.DeleteRoleReply
	(*GetRoleRequest)(nil),    // 6: api.petclinic.GetRoleRequest
	(*GetRoleReply)(nil),      // 7: api.petclinic.GetRoleReply
	(*ListRoleRequest)(nil),   // 8: api.petclinic.ListRoleRequest
	(*ListRoleReply)(nil),     // 9: api.petclinic.ListRoleReply
}
var file_petclinic_role_proto_depIdxs = []int32{
	0, // 0: api.petclinic.Role.CreateRole:input_type -> api.petclinic.CreateRoleRequest
	2, // 1: api.petclinic.Role.UpdateRole:input_type -> api.petclinic.UpdateRoleRequest
	4, // 2: api.petclinic.Role.DeleteRole:input_type -> api.petclinic.DeleteRoleRequest
	6, // 3: api.petclinic.Role.GetRole:input_type -> api.petclinic.GetRoleRequest
	8, // 4: api.petclinic.Role.ListRole:input_type -> api.petclinic.ListRoleRequest
	1, // 5: api.petclinic.Role.CreateRole:output_type -> api.petclinic.CreateRoleReply
	3, // 6: api.petclinic.Role.UpdateRole:output_type -> api.petclinic.UpdateRoleReply
	5, // 7: api.petclinic.Role.DeleteRole:output_type -> api.petclinic.DeleteRoleReply
	7, // 8: api.petclinic.Role.GetRole:output_type -> api.petclinic.GetRoleReply
	9, // 9: api.petclinic.Role.ListRole:output_type -> api.petclinic.ListRoleReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_petclinic_role_proto_init() }
func file_petclinic_role_proto_init() {
	if File_petclinic_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_petclinic_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_petclinic_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleReply); i {
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
		file_petclinic_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleRequest); i {
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
		file_petclinic_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleReply); i {
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
		file_petclinic_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRequest); i {
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
		file_petclinic_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleReply); i {
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
		file_petclinic_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleRequest); i {
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
		file_petclinic_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleReply); i {
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
		file_petclinic_role_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoleRequest); i {
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
		file_petclinic_role_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoleReply); i {
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
			RawDescriptor: file_petclinic_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petclinic_role_proto_goTypes,
		DependencyIndexes: file_petclinic_role_proto_depIdxs,
		MessageInfos:      file_petclinic_role_proto_msgTypes,
	}.Build()
	File_petclinic_role_proto = out.File
	file_petclinic_role_proto_rawDesc = nil
	file_petclinic_role_proto_goTypes = nil
	file_petclinic_role_proto_depIdxs = nil
}
