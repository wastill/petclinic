// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: petclinic/pet_type.proto

package petclinic

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	common "github.com/wastill/petclinic/api/common"
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

type CreatePetTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // the name of string must be between 1 and 80 character
}

func (x *CreatePetTypeRequest) Reset() {
	*x = CreatePetTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetTypeRequest) ProtoMessage() {}

func (x *CreatePetTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetTypeRequest.ProtoReflect.Descriptor instead.
func (*CreatePetTypeRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePetTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreatePetTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetType *PetType `protobuf:"bytes,1,opt,name=pet_type,json=petType,proto3" json:"pet_type,omitempty"`
}

func (x *CreatePetTypeReply) Reset() {
	*x = CreatePetTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePetTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePetTypeReply) ProtoMessage() {}

func (x *CreatePetTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePetTypeReply.ProtoReflect.Descriptor instead.
func (*CreatePetTypeReply) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePetTypeReply) GetPetType() *PetType {
	if x != nil {
		return x.PetType
	}
	return nil
}

type UpdatePetTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // the name of string must be between 1 a
}

func (x *UpdatePetTypeRequest) Reset() {
	*x = UpdatePetTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetTypeRequest) ProtoMessage() {}

func (x *UpdatePetTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdatePetTypeRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePetTypeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePetTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdatePetTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetType *PetType `protobuf:"bytes,1,opt,name=pet_type,json=petType,proto3" json:"pet_type,omitempty"`
}

func (x *UpdatePetTypeReply) Reset() {
	*x = UpdatePetTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePetTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePetTypeReply) ProtoMessage() {}

func (x *UpdatePetTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePetTypeReply.ProtoReflect.Descriptor instead.
func (*UpdatePetTypeReply) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePetTypeReply) GetPetType() *PetType {
	if x != nil {
		return x.PetType
	}
	return nil
}

type DeletePetTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePetTypeRequest) Reset() {
	*x = DeletePetTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetTypeRequest) ProtoMessage() {}

func (x *DeletePetTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetTypeRequest.ProtoReflect.Descriptor instead.
func (*DeletePetTypeRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePetTypeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeletePetTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowEffect int32 `protobuf:"varint,1,opt,name=row_effect,json=rowEffect,proto3" json:"row_effect,omitempty"`
}

func (x *DeletePetTypeReply) Reset() {
	*x = DeletePetTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetTypeReply) ProtoMessage() {}

func (x *DeletePetTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetTypeReply.ProtoReflect.Descriptor instead.
func (*DeletePetTypeReply) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePetTypeReply) GetRowEffect() int32 {
	if x != nil {
		return x.RowEffect
	}
	return 0
}

type GetPetTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPetTypeRequest) Reset() {
	*x = GetPetTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetTypeRequest) ProtoMessage() {}

func (x *GetPetTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetTypeRequest.ProtoReflect.Descriptor instead.
func (*GetPetTypeRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{6}
}

func (x *GetPetTypeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPetTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetType *PetType `protobuf:"bytes,1,opt,name=pet_type,json=petType,proto3" json:"pet_type,omitempty"`
}

func (x *GetPetTypeReply) Reset() {
	*x = GetPetTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetTypeReply) ProtoMessage() {}

func (x *GetPetTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetTypeReply.ProtoReflect.Descriptor instead.
func (*GetPetTypeReply) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{7}
}

func (x *GetPetTypeReply) GetPetType() *PetType {
	if x != nil {
		return x.PetType
	}
	return nil
}

type ListPetTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int32 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListPetTypeRequest) Reset() {
	*x = ListPetTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetTypeRequest) ProtoMessage() {}

func (x *ListPetTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetTypeRequest.ProtoReflect.Descriptor instead.
func (*ListPetTypeRequest) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{8}
}

func (x *ListPetTypeRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListPetTypeRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListPetTypeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *common.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Data []*PetType   `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListPetTypeReply) Reset() {
	*x = ListPetTypeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPetTypeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPetTypeReply) ProtoMessage() {}

func (x *ListPetTypeReply) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPetTypeReply.ProtoReflect.Descriptor instead.
func (*ListPetTypeReply) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{9}
}

func (x *ListPetTypeReply) GetPage() *common.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListPetTypeReply) GetData() []*PetType {
	if x != nil {
		return x.Data
	}
	return nil
}

type PetType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PetType) Reset() {
	*x = PetType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_petclinic_pet_type_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PetType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetType) ProtoMessage() {}

func (x *PetType) ProtoReflect() protoreflect.Message {
	mi := &file_petclinic_pet_type_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetType.ProtoReflect.Descriptor instead.
func (*PetType) Descriptor() ([]byte, []int) {
	return file_petclinic_pet_type_proto_rawDescGZIP(), []int{10}
}

func (x *PetType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PetType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_petclinic_pet_type_proto protoreflect.FileDescriptor

var file_petclinic_pet_type_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2f, 0x70, 0x65, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x50, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c,
	0x69, 0x6e, 0x69, 0x63, 0x2e, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x70, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x50, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x08, 0x70,
	0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x50, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x70, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2f,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x33, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x77, 0x5f, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x6f, 0x77, 0x45, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65,
	0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x70, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x64, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x50,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2d, 0x0a, 0x07,
	0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x73, 0x74, 0x69, 0x6c,
	0x6c, 0x2f, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x3b,
	0x70, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_petclinic_pet_type_proto_rawDescOnce sync.Once
	file_petclinic_pet_type_proto_rawDescData = file_petclinic_pet_type_proto_rawDesc
)

func file_petclinic_pet_type_proto_rawDescGZIP() []byte {
	file_petclinic_pet_type_proto_rawDescOnce.Do(func() {
		file_petclinic_pet_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_petclinic_pet_type_proto_rawDescData)
	})
	return file_petclinic_pet_type_proto_rawDescData
}

var file_petclinic_pet_type_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_petclinic_pet_type_proto_goTypes = []interface{}{
	(*CreatePetTypeRequest)(nil), // 0: api.petclinic.CreatePetTypeRequest
	(*CreatePetTypeReply)(nil),   // 1: api.petclinic.CreatePetTypeReply
	(*UpdatePetTypeRequest)(nil), // 2: api.petclinic.UpdatePetTypeRequest
	(*UpdatePetTypeReply)(nil),   // 3: api.petclinic.UpdatePetTypeReply
	(*DeletePetTypeRequest)(nil), // 4: api.petclinic.DeletePetTypeRequest
	(*DeletePetTypeReply)(nil),   // 5: api.petclinic.DeletePetTypeReply
	(*GetPetTypeRequest)(nil),    // 6: api.petclinic.GetPetTypeRequest
	(*GetPetTypeReply)(nil),      // 7: api.petclinic.GetPetTypeReply
	(*ListPetTypeRequest)(nil),   // 8: api.petclinic.ListPetTypeRequest
	(*ListPetTypeReply)(nil),     // 9: api.petclinic.ListPetTypeReply
	(*PetType)(nil),              // 10: api.petclinic.PetType
	(*common.Page)(nil),          // 11: api.common.Page
}
var file_petclinic_pet_type_proto_depIdxs = []int32{
	10, // 0: api.petclinic.CreatePetTypeReply.pet_type:type_name -> api.petclinic.PetType
	10, // 1: api.petclinic.UpdatePetTypeReply.pet_type:type_name -> api.petclinic.PetType
	10, // 2: api.petclinic.GetPetTypeReply.pet_type:type_name -> api.petclinic.PetType
	11, // 3: api.petclinic.ListPetTypeReply.page:type_name -> api.common.Page
	10, // 4: api.petclinic.ListPetTypeReply.data:type_name -> api.petclinic.PetType
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_petclinic_pet_type_proto_init() }
func file_petclinic_pet_type_proto_init() {
	if File_petclinic_pet_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_petclinic_pet_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePetTypeRequest); i {
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
		file_petclinic_pet_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePetTypeReply); i {
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
		file_petclinic_pet_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetTypeRequest); i {
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
		file_petclinic_pet_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePetTypeReply); i {
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
		file_petclinic_pet_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetTypeRequest); i {
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
		file_petclinic_pet_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetTypeReply); i {
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
		file_petclinic_pet_type_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetTypeRequest); i {
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
		file_petclinic_pet_type_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetTypeReply); i {
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
		file_petclinic_pet_type_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPetTypeRequest); i {
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
		file_petclinic_pet_type_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPetTypeReply); i {
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
		file_petclinic_pet_type_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PetType); i {
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
			RawDescriptor: file_petclinic_pet_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_petclinic_pet_type_proto_goTypes,
		DependencyIndexes: file_petclinic_pet_type_proto_depIdxs,
		MessageInfos:      file_petclinic_pet_type_proto_msgTypes,
	}.Build()
	File_petclinic_pet_type_proto = out.File
	file_petclinic_pet_type_proto_rawDesc = nil
	file_petclinic_pet_type_proto_goTypes = nil
	file_petclinic_pet_type_proto_depIdxs = nil
}