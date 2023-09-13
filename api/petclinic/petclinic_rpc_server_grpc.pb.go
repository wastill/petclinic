// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: petclinic/petclinic_rpc_server.proto

package petclinic

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// PetClinicRPCServiceClient is the client API for PetClinicRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetClinicRPCServiceClient interface {
}

type petClinicRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetClinicRPCServiceClient(cc grpc.ClientConnInterface) PetClinicRPCServiceClient {
	return &petClinicRPCServiceClient{cc}
}

// PetClinicRPCServiceServer is the server API for PetClinicRPCService service.
// All implementations must embed UnimplementedPetClinicRPCServiceServer
// for forward compatibility
type PetClinicRPCServiceServer interface {
	mustEmbedUnimplementedPetClinicRPCServiceServer()
}

// UnimplementedPetClinicRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPetClinicRPCServiceServer struct {
}

func (UnimplementedPetClinicRPCServiceServer) mustEmbedUnimplementedPetClinicRPCServiceServer() {}

// UnsafePetClinicRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetClinicRPCServiceServer will
// result in compilation errors.
type UnsafePetClinicRPCServiceServer interface {
	mustEmbedUnimplementedPetClinicRPCServiceServer()
}

func RegisterPetClinicRPCServiceServer(s grpc.ServiceRegistrar, srv PetClinicRPCServiceServer) {
	s.RegisterService(&PetClinicRPCService_ServiceDesc, srv)
}

// PetClinicRPCService_ServiceDesc is the grpc.ServiceDesc for PetClinicRPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetClinicRPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.petclinic.PetClinicRPCService",
	HandlerType: (*PetClinicRPCServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "petclinic/petclinic_rpc_server.proto",
}