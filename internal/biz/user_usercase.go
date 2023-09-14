package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wastill/petclinic/internal/util"
)

type UserRepo interface {
	GetUser(ctx context.Context, username string) (*User, error)
	CountUserByUsername(ctx context.Context, username string) (int, error)
	CreateUser(ctx context.Context, User *User) (*User, error)
}
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}
func (uc UserUseCase) getLogger(ctx context.Context) *log.Helper {
	return uc.log.WithContext(ctx)
}

func (uc UserUseCase) RegisterUserWithPassword(ctx context.Context, param *User) (*User, error) {
	userCount, err := uc.repo.CountUserByUsername(ctx, param.Username)
	if err != nil {
		uc.getLogger(ctx).Error("register with username and password", err.Error())
		return nil, err
	}
	if userCount > 0 {
		return nil, errors.New("username is repeated")
	}

	var encryptedPwd, salt string
	encryptedPwd, salt, err = util.NewPassword(param.Password)
	if err != nil {
		uc.getLogger(ctx).Error("register with username and password", err.Error())
		return nil, err
	}
	param.Password = encryptedPwd
	param.Salt = salt

	fmt.Println("register user ", param)

	newUser, err := uc.repo.CreateUser(ctx, param)
	if err != nil {
		uc.getLogger(ctx).Error("register with username and password", err.Error())
		return nil, err
	}
	return newUser, nil
}

func (uc UserUseCase) GetUser(ctx context.Context, username string) (*User, error) {
	user, err := uc.repo.GetUser(ctx, username)
	if err != nil {
		uc.getLogger(ctx).Error("get user", err.Error())
		return nil, err
	}
	return user, nil
}

func (uc UserUseCase) AuthnWithUserPwd(ctx context.Context, username, password string) (*User, error) {
	user, err := uc.repo.GetUser(ctx, username)
	fmt.Println("auth query user ", user)

	passed, err := util.CheckPassword(password, user.Password, user.Salt)
	if err != nil {
		uc.getLogger(ctx).Error("check user password", err.Error())
		return nil, err
	}
	if !passed {
		return nil, errors.New("密码错误!")
	}
	return user, nil
}
