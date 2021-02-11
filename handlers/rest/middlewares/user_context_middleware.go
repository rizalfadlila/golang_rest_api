package middlewares

import (
	"context"

	constantsrest "github.com/rest_api/constants/rest"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserContextMiddleware extracts user context from JWT and injects it to request context
func UserContextMiddleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		authHeader := ginCtx.GetHeader("Authorization")
		if len(authHeader) < 7 {
			ginCtx.Next()
			return
		}

		var claims jwt.MapClaims
		_, _ = jwt.ParseWithClaims(authHeader[7:], &claims, nil)

		requestCtx := ginCtx.Request.Context()

		requestCtx = context.WithValue(requestCtx, constantsrest.KeyRequestHeaderAuthorizarion, authHeader)

		if value, ok := claims["user"].(float64); ok {
			requestCtx = context.WithValue(requestCtx, constantsrest.KeyUserID, uint(value))
		}

		ginCtx.Request = ginCtx.Request.WithContext(requestCtx)
		ginCtx.Next()
	}
}

// IsAgent function return true if user is an agent.
// NOTE: Should apply UserContextMiddleware before use this middleware.
func IsAgent(ginCtx *gin.Context) bool {
	userType, ok := ginCtx.Request.Context().Value(constantsrest.KeyUserType).(string)
	if ok && userType == "BA" {
		return true
	}
	return false
}
