// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: petclinic/speciality.proto

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

type CreateSpecialityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSpecialityRequest) Reset() {
	*x = CreateSpecialityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSpecialityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSpecialityRequest) ProtoMessage() {}

func (x *CreateSpecialityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSpecialityRequest.ProtoReflect.Descriptor instead.
func (*CreateSpecialityRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{0}
}

type CreateSpecialityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSpecialityReply) Reset() {
	*x = CreateSpecialityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSpecialityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSpecialityReply) ProtoMessage() {}

func (x *CreateSpecialityReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSpecialityReply.ProtoReflect.Descriptor instead.
func (*CreateSpecialityReply) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{1}
}

type UpdateSpecialityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSpecialityRequest) Reset() {
	*x = UpdateSpecialityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSpecialityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSpecialityRequest) ProtoMessage() {}

func (x *UpdateSpecialityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSpecialityRequest.ProtoReflect.Descriptor instead.
func (*UpdateSpecialityRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{2}
}

type UpdateSpecialityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSpecialityReply) Reset() {
	*x = UpdateSpecialityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSpecialityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSpecialityReply) ProtoMessage() {}

func (x *UpdateSpecialityReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSpecialityReply.ProtoReflect.Descriptor instead.
func (*UpdateSpecialityReply) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{3}
}

type DeleteSpecialityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSpecialityRequest) Reset() {
	*x = DeleteSpecialityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSpecialityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSpecialityRequest) ProtoMessage() {}

func (x *DeleteSpecialityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSpecialityRequest.ProtoReflect.Descriptor instead.
func (*DeleteSpecialityRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{4}
}

type DeleteSpecialityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSpecialityReply) Reset() {
	*x = DeleteSpecialityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSpecialityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSpecialityReply) ProtoMessage() {}

func (x *DeleteSpecialityReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSpecialityReply.ProtoReflect.Descriptor instead.
func (*DeleteSpecialityReply) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{5}
}

type GetSpecialityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSpecialityRequest) Reset() {
	*x = GetSpecialityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpecialityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpecialityRequest) ProtoMessage() {}

func (x *GetSpecialityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpecialityRequest.ProtoReflect.Descriptor instead.
func (*GetSpecialityRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{6}
}

type GetSpecialityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSpecialityReply) Reset() {
	*x = GetSpecialityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSpecialityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSpecialityReply) ProtoMessage() {}

func (x *GetSpecialityReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSpecialityReply.ProtoReflect.Descriptor instead.
func (*GetSpecialityReply) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{7}
}

type ListSpecialityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSpecialityRequest) Reset() {
	*x = ListSpecialityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSpecialityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSpecialityRequest) ProtoMessage() {}

func (x *ListSpecialityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSpecialityRequest.ProtoReflect.Descriptor instead.
func (*ListSpecialityRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{8}
}

type ListSpecialityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSpecialityReply) Reset() {
	*x = ListSpecialityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_speciality_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSpecialityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSpecialityReply) ProtoMessage() {}

func (x *ListSpecialityReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_speciality_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSpecialityReply.ProtoReflect.Descriptor instead.
func (*ListSpecialityReply) Descriptor() ([]byte, []int) {
	return file_petclinic_speciality_proto_rawDescGZIP(), []int{9}
}

var File_petclinic_speciality_proto protoreflect.FileDescriptor

var file_petclinic_speciality_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x22, 0x19, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xe7, 0x03, 0x0a, 0x0a, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x60, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x60, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e,
	0x69, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x60, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c,
	0x69, 0x6e, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x57, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63,
	0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x4d, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63,
	0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x61, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2f, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x74, 0x63, 0x6c,
	0x69, 0x6e, 0x69, 0x63, 0x3b, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_petclinic_speciality_proto_rawDescOnce sync.Once
	file_petclinic_speciality_proto_rawDescData = file_petclinic_speciality_proto_rawDesc
)

func file_petclinic_speciality_proto_rawDescGZIP() []byte {
	file_petclinic_speciality_proto_rawDescOnce.Do(func() {
		file_petclinic_speciality_proto_rawDescData = protoimpl.X.CompressGZIP(file_petclinic_speciality_proto_rawDescData)
	})
	return file_petclinic_speciality_proto_rawDescData
}

var file_petclinic_speciality_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_petclinic_speciality_proto_goTypes = []interface{}{
	(*CreateSpecialityRequest)(nil), // 0: api.petclinic.CreateSpecialityRequest
	(*CreateSpecialityReply)(nil),   // 1: api.petclinic.CreateSpecialityReply
	(*UpdateSpecialityRequest)(nil), // 2: api.petclinic.UpdateSpecialityRequest
	(*UpdateSpecialityReply)(nil),   // 3: api.petclinic.UpdateSpecialityReply
	(*DeleteSpecialityRequest)(nil), // 4: api.petclinic.DeleteSpecialityRequest
	(*DeleteSpecialityReply)(nil),   // 5: api.petclinic.DeleteSpecialityReply
	(*GetSpecialityRequest)(nil),    // 6: api.petclinic.GetSpecialityRequest
	(*GetSpecialityReply)(nil),      // 7: api.petclinic.GetSpecialityReply
	(*ListSpecialityRequest)(nil),   // 8: api.petclinic.ListSpecialityRequest
	(*ListSpecialityReply)(nil),     // 9: api.petclinic.ListSpecialityReply
}
var file_petclinic_speciality_proto_depIdxs = []int32{
	0, // 0: api.petclinic.Speciality.CreateSpeciality:input_type -> api.petclinic.CreateSpecialityRequest
	2, // 1: api.petclinic.Speciality.UpdateSpeciality:input_type -> api.petclinic.UpdateSpecialityRequest
	4, // 2: api.petclinic.Speciality.DeleteSpeciality:input_type -> api.petclinic.DeleteSpecialityRequest
	6, // 3: api.petclinic.Speciality.GetSpeciality:input_type -> api.petclinic.GetSpecialityRequest
	8, // 4: api.petclinic.Speciality.ListSpeciality:input_type -> api.petclinic.ListSpecialityRequest
	1, // 5: api.petclinic.Speciality.CreateSpeciality:output_type -> api.petclinic.CreateSpecialityReply
	3, // 6: api.petclinic.Speciality.UpdateSpeciality:output_type -> api.petclinic.UpdateSpecialityReply
	5, // 7: api.petclinic.Speciality.DeleteSpeciality:output_type -> api.petclinic.DeleteSpecialityReply
	7, // 8: api.petclinic.Speciality.GetSpeciality:output_type -> api.petclinic.GetSpecialityReply
	9, // 9: api.petclinic.Speciality.ListSpeciality:output_type -> api.petclinic.ListSpecialityReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_petclinic_speciality_proto_init() }
func file_petclinic_speciality_proto_init() {
	if File_petclinic_speciality_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_petclinic_speciality_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSpecialityRequest); i {
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
		file_petclinic_speciality_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSpecialityReply); i {
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
		file_petclinic_speciality_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSpecialityRequest); i {
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
		file_petclinic_speciality_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSpecialityReply); i {
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
		file_petclinic_speciality_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSpecialityRequest); i {
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
		file_petclinic_speciality_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSpecialityReply); i {
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
		file_petclinic_speciality_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpecialityRequest); i {
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
		file_petclinic_speciality_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSpecialityReply); i {
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
		file_petclinic_speciality_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSpecialityRequest); i {
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
		file_petclinic_speciality_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSpecialityReply); i {
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
			RawDescriptor: file_petclinic_speciality_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_petclinic_speciality_proto_goTypes,
		DependencyIndexes: file_petclinic_speciality_proto_depIdxs,
		MessageInfos:      file_petclinic_speciality_proto_msgTypes,
	}.Build()
	File_petclinic_speciality_proto = out.File
	file_petclinic_speciality_proto_rawDesc = nil
	file_petclinic_speciality_proto_goTypes = nil
	file_petclinic_speciality_proto_depIdxs = nil
}