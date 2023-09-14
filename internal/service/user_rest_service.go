package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	aujwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wastill/petclinic/pkg/commons"
	"time"

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
	user, err := s.UserUseCase.RegisterUserWithPassword(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
		Roles:    []string{"user"},
	})

	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return &pb.CreateUserReply{}, nil
}

func (s *PetClinicRestService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := s.UserUseCase.AuthnWithUserPwd(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"roles":    "111",
		"exp":      time.Now().Unix() + int64(s.jwtValidSeconds),
	})
	reply := &pb.LoginReply{}
	if s, err := token.SignedString(commons.TestKey); err != nil {
		return nil, errors.New(500, "JWT_ERROR", err.Error())
	} else {
		reply.Authorization = s
	}
	return reply, nil
	//return nil, nil
}

func (s *PetClinicRestService) CurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.GetCurrentUserReply, error) {
	if c, ok := aujwt.FromContext(ctx); ok {
		fmt.Println(c)
	}
	return nil, status.Errorf(codes.Unimplemented, "method CurrentUser not implemented")
}
