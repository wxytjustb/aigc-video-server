package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/app"
	"github.com/flipped-aurora/gin-vue-admin/server/router/app_core"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System   system.RouterGroup
	App      app.RouterGroup
	App_core app_core.RouterGroup
}
