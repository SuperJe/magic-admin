package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
)

type LoginVerify struct {
	api.Api
}

func (v LoginVerify) LoginVerify(context *gin.Context) {
	//svc := service.LoginVerify{}
	//req := &dto.LoginVerifyReq{}
}
