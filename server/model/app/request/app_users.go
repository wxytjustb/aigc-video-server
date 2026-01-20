package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppUsersSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
}
