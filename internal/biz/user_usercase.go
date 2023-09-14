package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetUser(ctx context.Context, username string) (User, error)
	CreateUser(ctx context.Context, User User) (User, error)
}
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc UserUseCase) CreateUser(ctx context.Context, Param *User) (*User, error) {

	return nil, nil
}
