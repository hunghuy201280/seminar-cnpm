package apicommon

import (
	"github.com/gin-gonic/gin"
	"s3-service/common"
	"s3-service/component/tokenprovider"
)

func GetTokenPayload(c *gin.Context) (*tokenprovider.TokenPayload, error) {
	_tokenPayload, exist := c.Get(TOKEN_KEY)
	tokenPayload := _tokenPayload.(*tokenprovider.TokenPayload)
	if !exist {
		return nil, common.NewUnauthorized()
	}
	return tokenPayload, nil
}
