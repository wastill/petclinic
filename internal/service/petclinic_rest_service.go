package service

import (
	pb "github.com/wastill/petclinic/api/petclinic"
)

type PetClinicRestService struct {
	pb.UnimplementedPetClinicRestServiceServer
}

func NewPetClinicServiceService() *PetClinicRestService {
	return &PetClinicRestService{}
}
