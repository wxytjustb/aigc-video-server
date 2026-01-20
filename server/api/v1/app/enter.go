package app

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AppUsersApi
	LoginApi
}

var (
	appUsersService = service.ServiceGroupApp.AppServiceGroup.AppUsersService
	loginService    = service.ServiceGroupApp.AppServiceGroup.LoginService
	jwtService      = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
