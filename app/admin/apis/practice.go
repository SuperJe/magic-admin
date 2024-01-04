package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"net/http"
)

type Practice struct {
	api.Api
}

func (p Practice) GetPracticeCode(ctx *gin.Context) {
	svc := service.Practice{}
	req := &dto.GetPracticeCodeReq{}
	if err := p.MakeContext(ctx).MakeOrm().Bind(req).MakeService(&svc.Service).Errors; err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	if len(user.GetRoleName(ctx)) <= 0 {
		p.Logger.Error(fmt.Errorf("user role err"))
		p.Error(http.StatusBadRequest, fmt.Errorf("permission err"), "permission err")
		return
	}

	list := make([]int32, 0)
	if err := json.Unmarshal([]byte(req.IDList), &list); err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusBadRequest, err, err.Error())
		return
	}
	rsp, err := svc.GetPracticeCode(ctx, list, user.GetUserId(ctx))
	if err != nil {
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	p.OK(rsp, "success")
}

func (p Practice) SubmitPracticeCode(ctx *gin.Context) {
	svc := service.Practice{}
	req := &dto.SubmitPracticeCodeReq{}
	if err := p.MakeContext(ctx).MakeOrm().Bind(req, binding.JSON).MakeService(&svc.Service).Errors; err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	if len(user.GetRoleName(ctx)) <= 0 {
		p.Logger.Error(fmt.Errorf("user role err"))
		p.Error(http.StatusBadRequest, fmt.Errorf("permission err"), "permission err")
		return
	}

	rsp, err := svc.SubmitPracticeCode(ctx, int64(user.GetUserId(ctx)), req.ID, req.Code)
	if err != nil {
		rsp = &dto.SubmitPracticeCodeRsp{
			BaseRsp: dto.BaseRsp{Code: -1, Msg: err.Error()},
		}
		p.OK(rsp, "success")
		return
	}
	p.OK(rsp, "success")
}
