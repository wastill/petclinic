package service

import (
	pb "github.com/wastill/petclinic/api/petclinic"
	"github.com/wastill/petclinic/internal/biz"
)

type PetClinicRestService struct {
	pb.UnimplementedPetClinicRestServiceServer
	UserUseCase *biz.UserUseCase
}

func NewPetClinicRestService(UserUseCase *biz.UserUseCase) *PetClinicRestService {
	return &PetClinicRestService{
		UserUseCase: UserUseCase,
	}
}
