package middlewares

import (
	"net/http"
	"strings"

	"github.com/rest_api/config"
	constantsrest "github.com/rest_api/constants/rest"
	responses "github.com/rest_api/handlers/rest/responses"

	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware return gin middleware for authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ParseToken(ctx)
		if err != nil {
			unauthorizedResponse := responses.Response{
				Errors: []string{err.Error()},
				Message: responses.GetErrorConstant(
					constantsrest.ResponseCodeUnauthorized).Message,
				Status: constantsrest.StatusFailed,
			}

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorizedResponse)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		ctx.Set(constantsrest.JwtPayload, claims)
		ctx.Next()
	}
}

// ParseToken :nodoc:
func ParseToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		return nil, ginJwt.ErrEmptyAuthHeader
	}

	parts := strings.SplitN(tokenString, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, ginJwt.ErrInvalidAuthHeader
	}

	tokenString = parts[1]
	jwtToken, err := jwt.Parse(tokenString, tokenParser)
	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}

func tokenParser(token *jwt.Token) (interface{}, error) {
	if jwt.GetSigningMethod(constantsrest.AuthSigningMethod) != token.Method {
		return nil, ginJwt.ErrInvalidSigningAlgorithm
	}
	var cfg = *config.GetConfig()
	return []byte(cfg.App.Key), nil
}
