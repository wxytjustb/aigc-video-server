package app

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AppUsersRouter
	LoginApiRouter
}

var (
	appUsersApi = api.ApiGroupApp.AppApiGroup.AppUsersApi
	loginApi    = api.ApiGroupApp.AppApiGroup.LoginApi
)
