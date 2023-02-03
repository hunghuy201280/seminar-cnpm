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

func Register(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := ctx.GetMainDbConnection()
		store := userstorage.NewSQLStore(db)
		hasher := common.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, hasher)
		userId, err := biz.Register(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSuccessResponse(http.StatusCreated, "success", userId, nil, nil))

	}
}
