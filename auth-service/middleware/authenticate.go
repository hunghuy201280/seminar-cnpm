package middleware

import (
	"auth-service/common"
	"auth-service/component"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authenticate(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtProvider := ctx.GetJWTProvider()
		authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			panic(common.NewUnauthorized())
		}
		tokenString := authHeader[1]
		if tokenString == "" {
			panic(common.NewUnauthorized())
		}
		tokenPayload, err := jwtProvider.Validate(tokenString)
		if err != nil {
			panic(err)
		}
		c.Set("tokenPayload", tokenPayload)
		c.Next()
	}
}
