package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerManagementRouter)
}

func registerManagementRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Management{}
	r := v1.Group("/management").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/all_code_problem", actions.PermissionAction(), api.GetCodeProblem)
		r.POST("/update_code_problem", actions.PermissionAction(), api.UpdateCodeProblem)
	}
}
