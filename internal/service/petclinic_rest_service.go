package service

import (
	pb "github.com/wastill/petclinic/api/petclinic"
	"github.com/wastill/petclinic/internal/biz"
	"github.com/wastill/petclinic/internal/conf"
)

type PetClinicRestService struct {
	pb.UnimplementedPetClinicRestServiceServer
	UserUseCase     *biz.UserUseCase
	jwtValidSeconds int32
}

func NewPetClinicRestService(c *conf.Bootstrap, UserUseCase *biz.UserUseCase) *PetClinicRestService {
	return &PetClinicRestService{
		UserUseCase:     UserUseCase,
		jwtValidSeconds: c.Jwt.ValidSeconds,
	}
}
