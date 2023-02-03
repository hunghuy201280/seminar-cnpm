package apicommon

import (
	"auth-service/common"
	"auth-service/component/tokenprovider"
	"github.com/gin-gonic/gin"
)

func GetTokenPayload(c *gin.Context) (*tokenprovider.TokenPayload, error) {
	_tokenPayload, exist := c.Get(TOKEN_KEY)
	tokenPayload := _tokenPayload.(*tokenprovider.TokenPayload)
	if !exist {
		return nil, common.NewUnauthorized()
	}
	return tokenPayload, nil
}
