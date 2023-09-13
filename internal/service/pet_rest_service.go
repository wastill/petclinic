package service

import (
	"context"
	pb "github.com/wastill/petclinic/api/petclinic"
)

func (s *PetClinicRestService) CreatePet(ctx context.Context, req *pb.CreatePetRequest) (*pb.CreatePetReply, error) {
	return &pb.CreatePetReply{}, nil
}
func (s *PetClinicRestService) UpdatePet(ctx context.Context, req *pb.UpdatePetRequest) (*pb.UpdatePetReply, error) {
	return &pb.UpdatePetReply{}, nil
}
func (s *PetClinicRestService) DeletePet(ctx context.Context, req *pb.DeletePetRequest) (*pb.DeletePetReply, error) {
	return &pb.DeletePetReply{}, nil
}
func (s *PetClinicRestService) GetPet(ctx context.Context, req *pb.GetPetRequest) (*pb.GetPetReply, error) {
	return &pb.GetPetReply{}, nil
}
func (s *PetClinicRestService) ListPet(ctx context.Context, req *pb.ListPetRequest) (*pb.ListPetReply, error) {
	return &pb.ListPetReply{}, nil
}
func (s *PetClinicRestService) CreateOwnerPet(ctx context.Context, req *pb.CreateOwnerPetRequest) (*pb.CreateOwnerPetReply, error) {
	return &pb.CreateOwnerPetReply{}, nil
}
func (s *PetClinicRestService) DeleteOwnerPet(ctx context.Context, req *pb.DeleteOwnerPetRequest) (*pb.DeleteOwnerPetReply, error) {
	return &pb.DeleteOwnerPetReply{}, nil
}
func (s *PetClinicRestService) GetOwnerPet(ctx context.Context, req *pb.GetOwnerPetRequest) (*pb.GetOwnerPetReply, error) {
	return &pb.GetOwnerPetReply{}, nil
}
