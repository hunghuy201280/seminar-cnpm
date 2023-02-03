package ginuser

import (
	"auth-service/common"
	"auth-service/common/apicommon"
	"auth-service/component"
	"auth-service/modules/user/userbiz"
	"auth-service/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenPayload, err := apicommon.GetTokenPayload(c)

		if err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		userId := tokenPayload.UserId
		db := ctx.GetMainDbConnection()
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewGetUserInfoBiz(store)
		user, err := biz.GetUserInfo(c.Request.Context(), userId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(http.StatusOK, "success", user, nil, nil))

	}
}
