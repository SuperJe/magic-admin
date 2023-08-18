package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
)

type Dashboard struct {
	api.Api
}

func (d Dashboard) All(c *gin.Context) {
	svc := service.Dashboard{}
	if err := d.MakeContext(c).MakeOrm().MakeService(&svc.Service).Errors; err != nil {
		d.Logger.Error(err)
		d.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	name := user.GetUserName(c)
	id := user.GetUserId(c)
	rsp, err := svc.All(c, name, id)
	if err != nil {
		d.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	d.OK(rsp, "success")
}
