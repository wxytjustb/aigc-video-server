package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUsersRouter struct{}

// InitAppUsersRouter 初始化 appUsers表 路由信息
func (s *AppUsersRouter) InitAppUsersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	appUsersRouter := Router.Group("appUsers").Use(middleware.OperationRecord())
	appUsersRouterWithoutRecord := Router.Group("appUsers")
	appUsersRouterWithoutAuth := PublicRouter.Group("appUsers")
	{
		appUsersRouter.POST("createAppUsers", appUsersApi.CreateAppUsers)             // 新建appUsers表
		appUsersRouter.DELETE("deleteAppUsers", appUsersApi.DeleteAppUsers)           // 删除appUsers表
		appUsersRouter.DELETE("deleteAppUsersByIds", appUsersApi.DeleteAppUsersByIds) // 批量删除appUsers表
		appUsersRouter.PUT("updateAppUsers", appUsersApi.UpdateAppUsers)              // 更新appUsers表
	}
	{
		appUsersRouterWithoutRecord.GET("findAppUsers", appUsersApi.FindAppUsers)       // 根据ID获取appUsers表
		appUsersRouterWithoutRecord.GET("getAppUsersList", appUsersApi.GetAppUsersList) // 获取appUsers表列表
	}
	{
		appUsersRouterWithoutAuth.GET("getAppUsersPublic", appUsersApi.GetAppUsersPublic) // appUsers表开放接口
	}
}
