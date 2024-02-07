package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/cmd/migrate/migration/models"
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
	req.UserID = int64(user.GetUserId(ctx))
	sysUser := &models.SysUser{}
	if len(req.Username) > 0 {
		if err := c.Orm.Table("sys_user").Where("username = ?", req.Username).First(sysUser).Error; err != nil {
			c.Logger.Error(fmt.Errorf("get first err:%s", err.Error()))
			c.Error(http.StatusBadRequest, fmt.Errorf("get first err:%s", err.Error()), "get first err")
			return
		}
	}
	fmt.Printf("\n\n\n\nsysUser%+v\n\n\n\n", sysUser)
	if len(sysUser.Username) > 0 {
		req.UserID = int64(sysUser.UserId)
	}
	//if req.Username == "小明" {
	//	req.UserID = 2
	//}
	rsp, err := svc.GetLearnedLessons(ctx, req)
	if err != nil {
		c.Logger.Error(err)
		c.Error(http.StatusInternalServerError, err, "server busy")
		return
	}
	c.OK(rsp, "success")
}

func (c Courses) SignLesson(ctx *gin.Context) {
	s := service.Courses{}
	req := &dto.SignLessonReq{}
	err := c.MakeContext(ctx).MakeOrm().Bind(req, binding.JSON).MakeService(&s.Service).Errors
	if err != nil {
		c.Logger.Error(err)
		c.Error(http.StatusBadRequest, err, err.Error())
		return
	}
	if user.GetUserId(ctx) == 0 {
		c.Logger.Error(fmt.Errorf("user id err"))
		c.Error(http.StatusBadRequest, fmt.Errorf("user id err"), "user id err")
		return
	}

	req.UserID = int64(user.GetUserId(ctx))
	rsp, err := s.SignLesson(ctx, req)
	if err != nil {
		c.Logger.Error(err)
		c.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	c.OK(rsp, "success")
}

func (c Courses) AddLessonRecord(ctx *gin.Context) {
	svc := service.Courses{}
	req := &dto.AddLessonRecordReq{}
	if err := c.MakeContext(ctx).MakeOrm().Bind(req).MakeService(&svc.Service).Errors; err != nil {
		c.Logger.Error(err)
		c.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	// TODO: role name判断
	if user.GetRoleName(ctx) != "teacher" && user.GetRoleName(ctx) != "admin" {
		c.Logger.Error(fmt.Errorf("user role err"))
		c.Error(http.StatusBadRequest, fmt.Errorf("permission err"), "permission err")
		return
	}
	// TODO: 请求字段校验
	if req.CourseType != 1 && req.CourseType != 2 {
		c.Logger.Error(fmt.Errorf("courseType err"))
		c.Error(http.StatusBadRequest, fmt.Errorf("courseType err"), "courseType err")
		return
	}
	if len(req.KnowledgeTags) <= 0 || len(req.Remark) <= 0 || len(req.Teacher) <= 0 {
		c.Logger.Error(fmt.Errorf("missing field"))
		c.Error(http.StatusBadRequest, fmt.Errorf("missing field"), "missing field")
		return
	}
	sysUser := &models.SysUser{}
	if err := c.Orm.Table("sys_user").Where("username = ?", req.Name).First(sysUser).Error; err != nil {
		c.Logger.Error(fmt.Errorf("get first err:%s", err.Error()))
		c.Error(http.StatusBadRequest, fmt.Errorf("get first err:%s", err.Error()), "get first err")
		return
	}
	fmt.Printf("\n\n\n\nusername=%s, sysUser:%+v, req：%+v\n\n\n\n", req.Name, sysUser, req)
	// 没有报错说明找到了
	req.UserID = sysUser.UserId
	rsp, err := svc.AddLessonRecord(ctx, req)
	if err != nil {
		c.Logger.Error(err)
		c.Error(http.StatusInternalServerError, err, err.Error())
		return
	}
	c.OK(rsp, "success")
}

//func (c Courses) GetStudentName(ctx *gin.Context) {
//	svc := service.Courses{}
//}
