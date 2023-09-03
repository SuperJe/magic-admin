package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"net/http"
)

type Courses struct {
	api.Api
}

func (c Courses) GetCourseDetail(ctx *gin.Context) {
	svc := service.Courses{}
	req := &dto.GetCoursesDetailReq{}
	if err := c.MakeContext(ctx).MakeOrm().Bind(req).MakeService(&svc.Service).Errors; err != nil {
		c.Logger.Error(err)
		c.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	if len(user.GetRoleName(ctx)) <= 0 {
		c.Logger.Error(fmt.Errorf("user role err"))
		c.Error(http.StatusBadRequest, fmt.Errorf("permission err"), "permission err")
		return
	}

	rsp, err := svc.GetCourseDetail(ctx, req.CourseType)
	if err != nil {
		c.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	c.OK(rsp, "success")
}

func (c Courses) GetLearnedLessons(ctx *gin.Context) {
	svc := service.Courses{}
	req := &dto.GetLearnedReq{}
	if err := c.MakeContext(ctx).MakeOrm().Bind(req).MakeService(&svc.Service).Errors; err != nil {
		c.Logger.Error(err)
		c.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	if user.GetUserId(ctx) == 0 {
		c.Logger.Error(fmt.Errorf("user id err"))
		c.Error(http.StatusBadRequest, fmt.Errorf("user id err"), "user id err")
		return
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("req:", req)
	fmt.Println()
	fmt.Println()

	req.UserID = int64(user.GetUserId(ctx))
	rsp, err := svc.GetLearnedLessons(ctx, req)
	if err != nil {
		fmt.Printf("\n\nerr:%s\n\n\n", err.Error())
		c.Logger.Error(err)
		c.Error(http.StatusInternalServerError, err, "server busy")
		return
	}
	c.OK(rsp, "success")
}
