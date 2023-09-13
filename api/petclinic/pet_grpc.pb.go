// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: petclinic/pet.proto

package petclinic

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Pet_CreatePet_FullMethodName = "/api.petclinic.Pet/CreatePet"
	Pet_UpdatePet_FullMethodName = "/api.petclinic.Pet/UpdatePet"
	Pet_DeletePet_FullMethodName = "/api.petclinic.Pet/DeletePet"
	Pet_GetPet_FullMethodName    = "/api.petclinic.Pet/GetPet"
	Pet_ListPet_FullMethodName   = "/api.petclinic.Pet/ListPet"
)

// PetClient is the client API for Pet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetClient interface {
	CreatePet(ctx context.Context, in *CreatePetRequest, opts ...grpc.CallOption) (*CreatePetReply, error)
	UpdatePet(ctx context.Context, in *UpdatePetRequest, opts ...grpc.CallOption) (*UpdatePetReply, error)
	DeletePet(ctx context.Context, in *DeletePetRequest, opts ...grpc.CallOption) (*DeletePetReply, error)
	GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetReply, error)
	ListPet(ctx context.Context, in *ListPetRequest, opts ...grpc.CallOption) (*ListPetReply, error)
}

type petClient struct {
	cc grpc.ClientConnInterface
}

func NewPetClient(cc grpc.ClientConnInterface) PetClient {
	return &petClient{cc}
}

func (c *petClient) CreatePet(ctx context.Context, in *CreatePetRequest, opts ...grpc.CallOption) (*CreatePetReply, error) {
	out := new(CreatePetReply)
	err := c.cc.Invoke(ctx, Pet_CreatePet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petClient) UpdatePet(ctx context.Context, in *UpdatePetRequest, opts ...grpc.CallOption) (*UpdatePetReply, error) {
	out := new(UpdatePetReply)
	err := c.cc.Invoke(ctx, Pet_UpdatePet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petClient) DeletePet(ctx context.Context, in *DeletePetRequest, opts ...grpc.CallOption) (*DeletePetReply, error) {
	out := new(DeletePetReply)
	err := c.cc.Invoke(ctx, Pet_DeletePet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petClient) GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetReply, error) {
	out := new(GetPetReply)
	err := c.cc.Invoke(ctx, Pet_GetPet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petClient) ListPet(ctx context.Context, in *ListPetRequest, opts ...grpc.CallOption) (*ListPetReply, error) {
	out := new(ListPetReply)
	err := c.cc.Invoke(ctx, Pet_ListPet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetServer is the server API for Pet service.
// All implementations must embed UnimplementedPetServer
// for forward compatibility
type PetServer interface {
	CreatePet(context.Context, *CreatePetRequest) (*CreatePetReply, error)
	UpdatePet(context.Context, *UpdatePetRequest) (*UpdatePetReply, error)
	DeletePet(context.Context, *DeletePetRequest) (*DeletePetReply, error)
	GetPet(context.Context, *GetPetRequest) (*GetPetReply, error)
	ListPet(context.Context, *ListPetRequest) (*ListPetReply, error)
	mustEmbedUnimplementedPetServer()
}

// UnimplementedPetServer must be embedded to have forward compatible implementations.
type UnimplementedPetServer struct {
}

func (UnimplementedPetServer) CreatePet(context.Context, *CreatePetRequest) (*CreatePetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePet not implemented")
}
func (UnimplementedPetServer) UpdatePet(context.Context, *UpdatePetRequest) (*UpdatePetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePet not implemented")
}
func (UnimplementedPetServer) DeletePet(context.Context, *DeletePetRequest) (*DeletePetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePet not implemented")
}
func (UnimplementedPetServer) GetPet(context.Context, *GetPetRequest) (*GetPetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPet not implemented")
}
func (UnimplementedPetServer) ListPet(context.Context, *ListPetRequest) (*ListPetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPet not implemented")
}
func (UnimplementedPetServer) mustEmbedUnimplementedPetServer() {}

// UnsafePetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetServer will
// result in compilation errors.
type UnsafePetServer interface {
	mustEmbedUnimplementedPetServer()
}

func RegisterPetServer(s grpc.ServiceRegistrar, srv PetServer) {
	s.RegisterService(&Pet_ServiceDesc, srv)
}

func _Pet_CreatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServer).CreatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pet_CreatePet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServer).CreatePet(ctx, req.(*CreatePetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pet_UpdatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServer).UpdatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pet_UpdatePet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServer).UpdatePet(ctx, req.(*UpdatePetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pet_DeletePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServer).DeletePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pet_DeletePet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServer).DeletePet(ctx, req.(*DeletePetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pet_GetPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServer).GetPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pet_GetPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServer).GetPet(ctx, req.(*GetPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pet_ListPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServer).ListPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pet_ListPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServer).ListPet(ctx, req.(*ListPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pet_ServiceDesc is the grpc.ServiceDesc for Pet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.petclinic.Pet",
	HandlerType: (*PetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePet",
			Handler:    _Pet_CreatePet_Handler,
		},
		{
			MethodName: "UpdatePet",
			Handler:    _Pet_UpdatePet_Handler,
		},
		{
			MethodName: "DeletePet",
			Handler:    _Pet_DeletePet_Handler,
		},
		{
			MethodName: "GetPet",
			Handler:    _Pet_GetPet_Handler,
		},
		{
			MethodName: "ListPet",
			Handler:    _Pet_ListPet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petclinic/pet.proto",
}
