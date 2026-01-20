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
// GetAppUsersInfoList 分页获取appUsers表记录
func (appUsersService *AppUsersService) GetAppUsersInfoList(ctx context.Context, info appReq.AppUsersSearch) (list []app.AppUsers, total int64, err error) {
func (appUsersService *AppUsersService)GetAppUsersInfoList(ctx context.Context, info appReq.AppUsersSearch) (list []app.AppUsers, total int64, err error) {
	limit := info.PageSize
	// 创建db
    // 创建db
	var appUserss []app.AppUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

    
	if err != nil {
		return
	}
    }

		db = db.Limit(limit).Offset(offset)
	}
    }

	return appUserss, total, err
	return  appUserss, total, err
func (appUsersService *AppUsersService) GetAppUsersPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
    // 请自行实现
}
