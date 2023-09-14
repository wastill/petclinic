package data

import (
	"github.com/wastill/petclinic/internal/biz"
	"github.com/wastill/petclinic/internal/data/ent"
)

func UserPosToBos(po []*ent.User) []*biz.User {
	var bos []*biz.User
	for _, user := range po {
		bos = append(bos, UserPoToBo(user))
	}
	return bos
}
func UserPoToBo(po *ent.User) *biz.User {
	user := &biz.User{
		Username: po.Username,
		Password: po.Password,
		Enable:   po.Enable,
		Salt:     po.Salt,
	}

	if len(po.Edges.Roles) > 0 {
		for _, role := range po.Edges.Roles {
			user.Roles = append(user.Roles, role.Role)
		}
	}
	return user
}
