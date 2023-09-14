package commons

import (
	"context"
	"strconv"
)

type authKey struct{}

type IdentityInfo struct {
	UserId    int64
	OrgId     int64
	RoleNames string
}

func WithIdentityInfo(ctx context.Context, info *IdentityInfo) context.Context {
	return context.WithValue(ctx, authKey{}, info)
}

func IdentityInfoValue(ctx context.Context) (token *IdentityInfo, ok bool) {
	token, ok = ctx.Value(authKey{}).(*IdentityInfo)
	return
}

func GetOperatorStr(ctx context.Context) string {
	if idInfo, ok := IdentityInfoValue(ctx); ok {
		return "u" + strconv.FormatInt(idInfo.UserId, 10)
	}
	return "empty"
}
