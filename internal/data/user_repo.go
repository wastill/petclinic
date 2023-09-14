package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wastill/petclinic/internal/biz"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {

	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u UserRepo) GetUser(ctx context.Context, username string) (biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) CreateUser(ctx context.Context, User biz.User) (biz.User, error) {
	//TODO implement me
	panic("implement me")
}
