package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common"
	"net/http"
)

type Management struct {
	api.Api
}

func (p Management) UpdateCodeProblem(ctx *gin.Context) {
	if user.GetRoleName(ctx) != common.RoleAdmin {
		p.Logger.Error(fmt.Errorf("user role err"))
		p.Error(http.StatusBadRequest, fmt.Errorf("permission err"), "permission err")
		return
	}
	svc := service.Management{}
	req := &dto.UpdateCodeProblemReq{}
	if err := p.MakeContext(ctx).MakeOrm().Bind(req).MakeService(&svc.Service).Errors; err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	bs, _ := json.Marshal(req)
	fmt.Printf("req:%s\n\n\n", string(bs))
	if req.ID == 0 {
		err := fmt.Errorf("id empty")
		p.Logger.Error(err)
		p.Error(http.StatusBadRequest, err, err.Error())
		return
	}

	if err := svc.UpdateCodeProblem(ctx, req.ToCodeProblem()); err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	p.OK(&dto.UpdateCodeProblemRsp{
		BaseRsp: dto.BaseRsp{},
	}, "success")
}

func (p Management) GetCodeProblem(ctx *gin.Context) {
	if user.GetRoleName(ctx) != common.RoleAdmin {
		p.Logger.Error(fmt.Errorf("user role err"))
		p.Error(http.StatusBadRequest, fmt.Errorf("permission err"), "permission err")
		return
	}
	svc := service.Management{}
	req := &dto.GetCodeProblemReq{}
	if err := p.MakeContext(ctx).MakeOrm().Bind(req).MakeService(&svc.Service).Errors; err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	rsp, err := svc.GetCodeProblem(ctx, req.Offset, req.Limit, req.IsReverse)
	if err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	p.OK(rsp, "success")
}
