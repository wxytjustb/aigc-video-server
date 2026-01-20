package app

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type AppUsersService struct{}

// CreateAppUsers 创建appUsers表记录
// Author [yourname](https://github.com/yourname)
func (appUsersService *AppUsersService) CreateAppUsers(ctx context.Context, appUsers *app.AppUsers) (err error) {
	err = global.GVA_DB.Create(appUsers).Error
	return err
}

// DeleteAppUsers 删除appUsers表记录
// Author [yourname](https://github.com/yourname)
func (appUsersService *AppUsersService) DeleteAppUsers(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&app.AppUsers{}, "id = ?", ID).Error
	return err
}

// DeleteAppUsersByIds 批量删除appUsers表记录
// Author [yourname](https://github.com/yourname)
func (appUsersService *AppUsersService) DeleteAppUsersByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppUsers{}, "id in ?", IDs).Error
	return err
}

// UpdateAppUsers 更新appUsers表记录
// Author [yourname](https://github.com/yourname)
func (appUsersService *AppUsersService) UpdateAppUsers(ctx context.Context, appUsers app.AppUsers) (err error) {
	err = global.GVA_DB.Model(&app.AppUsers{}).Where("id = ?", appUsers.ID).Updates(&appUsers).Error
	return err
}

// GetAppUsers 根据ID获取appUsers表记录
// Author [yourname](https://github.com/yourname)
func (appUsersService *AppUsersService) GetAppUsers(ctx context.Context, ID string) (appUsers app.AppUsers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&appUsers).Error
	return
}

// GetAppUsersInfoList 分页获取appUsers表记
func (appUsersService *AppUsersService) GetAppUsersInfoList(ctx context.Context, info appReq.AppUsersSearch) (list []app.AppUsers, total int64, err error) {
	limit := info.PageSize
	if limit == 0 {
		limit = 10
	}
	offset := (info.Page - 1) * limit

	var appUserss []app.AppUsers
	db := global.GVA_DB.WithContext(ctx).Model(&app.AppUsers{})

	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	// 获取总记录数
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页并获取数据
	err = db.Limit(limit).Offset(offset).Find(&appUserss).Error
	return appUserss, total, err
}
func (appUsersService *AppUsersService) GetAppUsersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
	// 请自行实现
}
