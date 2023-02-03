package middleware

import (
	"github.com/gin-gonic/gin"
	"s3-service/common"
	"s3-service/component"
)

func Recover(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			c.Header("Content-Type", "application/json")

			if appErr, ok := err.(*common.AppError); ok {
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
			appErr := common.ErrInternal(err.(error))
			c.AbortWithStatusJSON(appErr.StatusCode, appErr)
			panic(err)
			return

		}()
		c.Next()

	}
}
