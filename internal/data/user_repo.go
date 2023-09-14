package data

import (
	"context"
	_ "entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/wastill/petclinic/internal/biz"
	"github.com/wastill/petclinic/internal/data/ent"
	"github.com/wastill/petclinic/internal/data/ent/user"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func (repo UserRepo) GetDB() *ent.Client {
	return repo.data.db
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {

	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo UserRepo) CountUserByUsername(ctx context.Context, username string) (int, error) {
	count, err := repo.GetDB().User.Query().Where(user.UsernameEQ(username)).Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (u UserRepo) GetUser(ctx context.Context, username string) (*biz.User, error) {

	first, err := u.GetDB().User.Query().Where(user.UsernameEQ(username)).WithRoles().First(ctx)
	if err != nil {
		return &biz.User{}, err
	}
	return UserPoToBo(first), nil
}

func (repo UserRepo) CreateUser(ctx context.Context, param *biz.User) (*biz.User, error) {
	err := repo.GetDB().User.Create().SetUsername(param.Username).
		SetPassword(param.Password).
		SetSalt(param.Salt).SetEnable(true).Exec(ctx)
	if err != nil {
		return nil, err
	}
	first, err := repo.GetDB().User.Query().Where(user.UsernameEQ(param.Username)).WithRoles().First(ctx)
	if err != nil {
		return nil, err
	}
	return UserPoToBo(first), nil
}
