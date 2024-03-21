package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerPracticeRouter)
}

func registerPracticeRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Practice{}
	r := v1.Group("/practice").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/cpp", actions.PermissionAction(), api.GetPracticeCode)
		r.POST("/cpp", actions.PermissionAction(), api.SubmitPracticeCode)
		r.GET("/get_questions", actions.PermissionAction(), api.GetQuestions)
		r.POST("question_submit", actions.PermissionAction(), api.QuestionSubmit)
	}
}
