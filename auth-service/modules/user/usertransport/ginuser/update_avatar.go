package ginuser

import (
	"auth-service/common"
	"auth-service/common/apicommon"
	"auth-service/component"
	"auth-service/modules/user/userbiz"
	"auth-service/modules/user/usermodel"
	"auth-service/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateAvatar(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UpdateAvatarRequest
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := ctx.GetMainDbConnection()
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewUpdateAvatarBiz(store)
		tokenPayload, err := apicommon.GetTokenPayload(c)
		if err != nil {
			panic(err)
		}
		err = biz.UpdateAvatar(c.Request.Context(), &data, tokenPayload.UserId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSuccessResponse(http.StatusOK, "success", tokenPayload.UserId, nil, nil))

	}
}
