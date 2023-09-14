package authn_jwt

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

// type authKey struct{}

// type IdentityInfo struct {
// 	// RequestId string
// 	UserId    int64
// 	OrgId     int64
// 	RoleNames string
// }

const (

	// bearerWord the bearer key word for authorization
	bearerWord string = "Bearer"

	// bearerFormat authorization token format
	bearerFormat string = "Bearer %s"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	authorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

// Option is jwt option.
type Option func(*options)

// Parser is a jwt parser
type options struct {
	signingMethod jwt.SigningMethod
	claims        func() jwt.Claims
	tokenHeader   map[string]interface{}
	publicKey     *rsa.PublicKey
}

// WithSigningMethod with signing method option.
func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// WithClaims with customer claim
// If you use it in Server, f needs to return a new jwt.Claims object each time to avoid concurrent write problems
// If you use it in Client, f only needs to return a single object to provide performance
func WithClaims(f func() jwt.Claims) Option {
	return func(o *options) {
		o.claims = f
	}
}

// WithTokenHeader withe customer tokenHeader for client side
func WithTokenHeader(header map[string]interface{}) Option {
	return func(o *options) {
		o.tokenHeader = header
	}
}

// use rsa
// Server is a server auth middleware. Check the token and extract the info from token.
func Server(
	// keyFunc jwt.Keyfunc,
	opts ...Option) middleware.Middleware {
	o := &options{
		// signingMethod: jwt.SigningMethodHS256,
		signingMethod: jwt.SigningMethodRS256,
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				// if keyFunc == nil {
				// 	return nil, ErrMissingKeyFunc
				// }
				auths := strings.SplitN(header.RequestHeader().Get(authorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
					return nil, ErrMissingJwtToken
				}
				jwtToken := auths[1]
				var (
					tokenInfo *jwt.Token
					err       error
				)
				// if o.claims != nil {
				// 	tokenInfo, err = jwt.ParseWithClaims(jwtToken, o.claims(), keyFunc)
				// } else {
				//  tokenInfo, err = jwt.Parse(jwtToken, keyFunc)
				tokenInfo, err = jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
						return o.publicKey, nil
					}
					return nil, fmt.Errorf("unsupport jwt method: %v", token.Header["alg"])
				})
				// }
				if err != nil {
					if ve, ok := err.(*jwt.ValidationError); ok {
						if ve.Errors&jwt.ValidationErrorMalformed != 0 {
							return nil, ErrTokenInvalid
						} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
							return nil, ErrTokenExpired
						} else {
							return nil, ErrTokenParseFail
						}
					}
					return nil, errors.Unauthorized(reason, err.Error())
				} else if !tokenInfo.Valid {
					return nil, ErrTokenInvalid
				} else if tokenInfo.Method != o.signingMethod {
					return nil, ErrUnSupportSigningMethod
				}
				if claims, ok := tokenInfo.Claims.(jwt.MapClaims); ok {
					//idInfo := &appcommons.IdentityInfo{
					//	UserId: int64(claims["uid"].(float64)),
					//}
					//roleClaims := claims["roles"]
					//if roleClaims != nil {
					//	idInfo.RoleNames = roleClaims.(string)
					//}
					//ctx = appcommons.WithIdentityInfo(ctx, idInfo)
					fmt.Println(claims)
					return handler(ctx, req)
				}
				return nil, ErrTokenInvalid
			}
			return nil, ErrWrongContext
		}
	}
}

// Client is a client jwt middleware.
func Client(keyProvider jwt.Keyfunc, opts ...Option) middleware.Middleware {
	claims := jwt.RegisteredClaims{}
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
		claims:        func() jwt.Claims { return claims },
	}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if keyProvider == nil {
				return nil, ErrNeedTokenProvider
			}
			token := jwt.NewWithClaims(o.signingMethod, o.claims())
			if o.tokenHeader != nil {
				for k, v := range o.tokenHeader {
					token.Header[k] = v
				}
			}
			key, err := keyProvider(token)
			if err != nil {
				return nil, ErrGetKey
			}
			tokenStr, err := token.SignedString(key)
			if err != nil {
				return nil, ErrSignToken
			}
			if clientContext, ok := transport.FromClientContext(ctx); ok {
				clientContext.RequestHeader().Set(authorizationKey, fmt.Sprintf(bearerFormat, tokenStr))
				return handler(ctx, req)
			}
			return nil, ErrWrongContext
		}
	}
}

// // NewContext put auth info into context
// func NewContext(ctx context.Context, info *IdentityInfo) context.Context {
// 	return context.WithValue(ctx, authKey{}, info)
// }

// // FromContext extract auth info from context
// func FromContext(ctx context.Context) (token *IdentityInfo, ok bool) {
// 	token, ok = ctx.Value(authKey{}).(*IdentityInfo)
// 	return
// }
