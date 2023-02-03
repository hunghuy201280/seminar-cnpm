package ginuser

import (
	"auth-service/common"
	"auth-service/component"
	"auth-service/modules/user/userbiz"
	"auth-service/modules/user/usermodel"
	"auth-service/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := ctx.GetMainDbConnection()
		store := userstorage.NewSQLStore(db)
		hasher := common.NewMd5Hash()
		biz := userbiz.NewLoginBiz(store, hasher, ctx.GetJWTProvider(), ctx.GetJWTRefreshProvider())
		token, refreshToken, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(http.StatusOK, "success", common.JS{
			"token":         token,
			"refresh_token": refreshToken,
		}, nil, nil))

	}
}
