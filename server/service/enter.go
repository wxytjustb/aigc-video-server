package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/app"
	"github.com/flipped-aurora/gin-vue-admin/server/service/app_core"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	AppServiceGroup      app.ServiceGroup
	App_coreServiceGroup app_core.ServiceGroup
}
