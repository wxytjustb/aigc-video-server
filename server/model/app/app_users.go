// 自动生成模板AppUsers
package app

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// appUsers表 结构体  AppUsers
type AppUsers struct {
	global.GVA_MODEL
	// 唯一身份识别码
	CasdoorId    *string `json:"casdoorId" form:"casdoorId" gorm:"comment:Casdoor用户ID;column:casdoor_id;size:36;"`                     // Casdoor提供的用户ID
	CasdoorOwner *string `json:"casdoorOwner" form:"casdoorOwner" gorm:"not null;comment:Casdoor组织名;column:casdoor_owner;size:100;"`   // 组织名
	CasdoorName  *string `json:"casdoorName" form:"casdoorName" gorm:"not null;comment:Casdoor用户名/唯一标识;column:casdoor_name;size:100;"` // 用户名/唯一标识
	// 基本信息
	Nickname *string `json:"nickname" form:"nickname" gorm:"comment:昵称;column:nickname;size:255;"` // 昵称
	Avatar   *string `json:"avatar" form:"avatar" gorm:"comment:头像;column:avatar;size:500;"`       // 头像
	Email    *string `json:"email" form:"email" gorm:"comment:邮箱;column:email;size:100;"`          // 邮箱
	Phone    *string `json:"phone" form:"phone" gorm:"comment:手机号;column:phone;size:20;"`          // 手机号
	// 业务逻辑
	Source *string `json:"source" form:"source" gorm:"comment:来源;column:source;size:50;"` // 来源: "google", "wechat", "password"

	// 登录统计
	LastLoginAt *time.Time `json:"lastLoginAt" form:"lastLoginAt" gorm:"comment:最近登录时间;column:last_login_at;"`     // 最近登录时间
	LoginCount  *int       `json:"loginCount" form:"loginCount" gorm:"default:0;comment:登录次数;column:login_count;"` // 登录次数
}

// TableName appUsers表 AppUsers自定义表名 app_users
func (AppUsers) TableName() string {
	return "app_users"
}
