package apis

import (
	"encoding/json"
	"fmt"
	"github.com/SuperJe/coco/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/cmd/migrate/migration/models"
	"go-admin/common"
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
	req.UserID = int64(user.GetUserId(ctx))
	sysUser := &models.SysUser{}
	if len(req.Username) > 0 && req.Username != "undefined" {
		if err := p.Orm.Table("sys_user").Where("username = ?", req.Username).First(sysUser).Error; err != nil {
			p.Logger.Error(fmt.Errorf("get first err:%s", err.Error()))
			p.Error(http.StatusBadRequest, fmt.Errorf("get first err:%s", err.Error()), "get first err")
			return
		}
	}
	if len(sysUser.Username) > 0 {
		req.UserID = int64(sysUser.UserId)
	}
	list := make([]int32, 0)
	if err := json.Unmarshal([]byte(req.IDList), &list); err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusBadRequest, err, err.Error())
		return
	}
	rsp, err := svc.GetPracticeCode(ctx, list, int(req.UserID))
	// rsp, err := svc.GetPracticeCode(ctx, list, user.GetUserId(ctx))
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

	req.UserID = int64(user.GetUserId(ctx))
	sysUser := &models.SysUser{}
	if len(req.Username) > 0 && req.Username != "undefined" {
		if err := p.Orm.Table("sys_user").Where("username = ?", req.Username).First(sysUser).Error; err != nil {
			p.Logger.Error(fmt.Errorf("get first err:%s", err.Error()))
			p.Error(http.StatusBadRequest, fmt.Errorf("get first err:%s", err.Error()), "get first err")
			return
		}
	}
	if len(sysUser.Username) > 0 {
		req.UserID = int64(sysUser.UserId)
	}
	rsp, err := svc.SubmitPracticeCode(ctx, req.UserID, req.ID, req.Code, req.Lang)
	if err != nil {
		rsp = &dto.SubmitPracticeCodeRsp{
			BaseRsp: dto.BaseRsp{Code: -1, Msg: err.Error()},
		}
		p.OK(rsp, "success")
		return
	}
	p.OK(rsp, "success")
}

func (p Practice) GetQuestions(ctx *gin.Context) {
	svc := service.Practice{}
	req := &dto.GetQuestionsReq{}
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
	rsp, err := svc.GetQuestions(ctx, list)
	if err != nil {
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	p.OK(rsp, "success")
}

func (p Practice) QuestionSubmit(ctx *gin.Context) {
	svc := service.Practice{}
	req := &dto.QuestionSubmitReq{}
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
	rsp, err := svc.QuestionSubmit(ctx, req)
	if err != nil {
		rsp = &dto.QuestionSubmitRsp{
			BaseRsp: dto.BaseRsp{Code: -1, Msg: err.Error()},
		}
		p.OK(rsp, "success")
		return
	}
	p.OK(rsp, "success")
}

func (p Practice) AddCodeProblem(ctx *gin.Context) {
	if user.GetRoleName(ctx) != common.RoleAdmin && user.GetRoleName(ctx) != common.RoleTeacher {
		p.Logger.Error(fmt.Errorf("user role err"))
		p.Error(http.StatusBadRequest, fmt.Errorf("permission err"), "permission err")
		return
	}
	svc := service.Practice{}
	req := &dto.AddCodeProblemReq{}
	if err := p.MakeContext(ctx).MakeOrm().Bind(req, binding.JSON).MakeService(&svc.Service).Errors; err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusBadRequest, err, err.Error())
		return
	}
	if req.PID == 0 || util.ExistEmptyString(req.ExampleOutput, req.ExampleInput, req.Detail, req.Title) {
		p.Error(http.StatusBadRequest, fmt.Errorf("bad request"), "check params plz")
		return
	}

	if err := svc.AddCodeProblem(ctx, req); err != nil {
		p.OK(&dto.AddCodeProblemRsp{BaseRsp: dto.BaseRsp{Code: -1, Msg: err.Error()}}, "fail")
		return
	}
	p.OK(&dto.AddCodeProblemRsp{}, "success")
}

func (p Practice) GetTest(ctx *gin.Context) {
	svc := service.Practice{}
	req := &dto.GetTestReq{}
	if err := p.MakeContext(ctx).MakeOrm().Bind(req).MakeService(&svc.Service).Errors; err != nil {
		p.Logger.Error(err)
		p.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	fmt.Printf("\n\n\n\ntestid=%d", req.Id)
	if len(user.GetRoleName(ctx)) <= 0 {
		p.Logger.Error(fmt.Errorf("user role err"))
		p.Error(http.StatusBadRequest, fmt.Errorf("permission err"), "permission err")
		return
	}
	rsp, err := svc.GetTest(ctx, req)
	if err != nil {
		rsp = &dto.GetTestRsp{
			BaseRsp: dto.BaseRsp{Code: -1, Msg: err.Error()},
		}
		p.OK(rsp, "success")
		return
	}
	p.OK(rsp, "success")
}
