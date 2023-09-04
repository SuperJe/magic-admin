package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCoursesRouter)
}

func registerCoursesRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Courses{}
	r := v1.Group("/courses").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/detail", actions.PermissionAction(), api.GetCourseDetail)
		r.GET("/learned", actions.PermissionAction(), api.GetLearnedLessons)
		r.POST("/sign", actions.PermissionAction(), api.SignLesson)
	}
}
