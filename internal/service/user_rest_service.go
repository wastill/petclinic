package service

import (
	"context"
	"fmt"
	pb "github.com/wastill/petclinic/api/petclinic"
	"github.com/wastill/petclinic/internal/biz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PetClinicRestService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func (s *PetClinicRestService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}

func (s *PetClinicRestService) Register(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user, err := s.UserUseCase.CreateUser(ctx, &biz.User{})
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (s *PetClinicRestService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (s *PetClinicRestService) CurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.GetCurrentUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentUser not implemented")
}
