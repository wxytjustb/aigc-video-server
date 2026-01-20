package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/app"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/app_core"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	AppApiGroup      app.ApiGroup
	App_coreApiGroup app_core.ApiGroup
}
