package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	v1 "github.com/wastill/petclinic/api/petclinic"
	"github.com/wastill/petclinic/internal/conf"
	"github.com/wastill/petclinic/internal/service"
	"github.com/wastill/petclinic/middleware/cors"
	"github.com/wastill/petclinic/pkg/commons"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.petclinic.PetClinicRestService/Login"] = struct{}{}
	whiteList["/api.petclinic.PetClinicRestService/Register"] = struct{}{}
	whiteList["/api.petclinic.PetClinicRestService/GetUser"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, petclinic *service.PetClinicRestService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			logging.Server(logger),
			recovery.Recovery(),
			mmd.Server(mmd.WithPropagatedPrefix("x-token-")),
			// 使用
			selector.Server(
				jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
					return []byte(commons.TestKey), nil
				}),
			).
				Match(NewWhiteListMatcher()).
				Build(),
		),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	// ================ 跨域 ============
	opts = append(opts, http.Filter(cors.CORS(
		cors.AllowedHeaders([]string{"*"}),
		cors.AllowedOrigins([]string{"*"}),
		cors.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}),
	)))
	// ================================
	srv := http.NewServer(opts...)
	h := openapiv2.NewHandler()
	//将/q/路由放在最前匹配
	srv.HandlePrefix("/q/", h)
	v1.RegisterPetClinicRestServiceHTTPServer(srv, petclinic)
	return srv
}
