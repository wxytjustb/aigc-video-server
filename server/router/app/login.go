package app

import "github.com/gin-gonic/gin"

type LoginApiRouter struct{}

// InitAppAuthRouter 初始化 C端认证路由（Casdoor）
func (s *LoginApiRouter) InitAppAuthRouter(_ *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	loginRouter := PublicRouter.Group("login")
	{
		loginRouter.POST("casdoorLogin", loginApi.CasdoorLogin)
	}
}
