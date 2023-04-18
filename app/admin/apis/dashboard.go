package apis

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
)

type Dashboard struct {
	api.Api
}

func (d Dashboard) All(c *gin.Context) {
	fmt.Println("Dashboard info info info info info info")
	svc := service.Dashboard{}
	if err := d.MakeContext(c).MakeService(&svc.Service).Errors; err != nil {
		d.Logger.Error(err)
		d.Error(500, err, err.Error())
		return
	}
	name := user.GetUserName(c)
	rsp, err := svc.All(c, name)
	if err != nil {
		d.Error(500, err, err.Error())
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	bs, _ := json.Marshal(rsp)
	fmt.Println("ok:", string(bs))
	d.OK(rsp, "success")
	// req := dto.ArticleGetReq{}
	// s := service.Article{}
	// err := e.MakeContext(c).
	// 	MakeOrm().
	// 	Bind(&req).
	// 	MakeService(&s.Service).
	// 	Errors
	// if err != nil {
	// 	e.Logger.Error(err)
	// 	e.Error(500, err, err.Error())
	// 	return
	// }
	// var object models.Article

	// p := actions.GetPermissionFromContext(c)
	// err = s.Get(&req, p, &object)
	// if err != nil {
	// 	e.Error(500, err, fmt.Sprintf("获取文章失败，\r\n失败信息 %s", err.Error()))
	// 	return
	// }

	// e.OK( object, "查询成功")
}
