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

func UpdateProfile(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := ctx.GetMainDbConnection()
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewUpdateProfileBiz(store)
		tokenPayload, err := apicommon.GetTokenPayload(c)
		if err != nil {
			panic(err)
		}
		err = biz.UpdateUser(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSuccessResponse(http.StatusOK, "success", tokenPayload.UserId, nil, nil))

	}
}
