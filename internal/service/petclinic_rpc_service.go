package service

import (
	pb "github.com/wastill/petclinic/api/petclinic"
)

type PetClinicRpcService struct {
	pb.UnimplementedPetClinicRPCServiceServer
}

func NewPetClinicRPCService() *PetClinicRpcService {
	return &PetClinicRpcService{}
}
