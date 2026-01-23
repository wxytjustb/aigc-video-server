package app

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	appModel "github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"gorm.io/gorm"
)

type LoginService struct{}

// CasdoorLogin Casdoor OAuth 2.0 登录 - 只处理用户信息，不生成 token
func (s *LoginService) CasdoorLogin(ctx context.Context, code string) (*appModel.AppUsers, error) {
	// 验证 code 参数
	if code == "" {
		return nil, errors.New("authorization code 不能为空")
	}

	cfg := global.GVA_CONFIG.Casdoor

	if cfg.Endpoint == "" || cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, errors.New("casdoor 配置不完整")
	}

	// 初始化 SDK 并换取 token
	casdoorsdk.InitConfig(cfg.Endpoint, cfg.ClientID, cfg.ClientSecret, cfg.Certificate, cfg.OrganizationName, cfg.ApplicationName)
	token, err := casdoorsdk.GetOAuthToken(code, "")
	if err != nil {
		return nil, fmt.Errorf("换取 token 失败: %w", err)
	}

	// 解析并验证 token
	claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("解析 token 失败: %w", err)
	}

	// 验证必要字段
	if claims.User.Id == "" {
		return nil, errors.New("token 缺少必要的用户信息")
	}

	db := global.GVA_DB.WithContext(ctx)

	user := appModel.AppUsers{
		CasdoorOwner: &claims.User.Owner,
		CasdoorName:  &claims.User.Name,
		CasdoorId:    &claims.User.Id,
		Nickname:     &claims.User.DisplayName,
		Avatar:       &claims.Avatar,
		Email:        &claims.User.Email,
		Phone:        &claims.User.Phone,
		Source:       &claims.SigninMethod,
	}

	// 先尝试查找用户
	err = db.Where("casdoor_id = ?", claims.User.Id).First(&user).Error
	if err != nil {
		// 如果用户不存在，创建新用户
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 设置默认登录次数
			loginCount := 1
			now := time.Now()
			user.LastLoginAt = &now
			user.LoginCount = &loginCount
			if err = db.Create(&user).Error; err != nil {
				// 如果创建失败是因为唯一索引冲突（可能是并发创建或数据库已有重复数据），重新查找
				if strings.Contains(err.Error(), "Duplicate entry") {
					// 尝试重新查找用户
					if findErr := db.Where("casdoor_id = ?", claims.User.Id).First(&user).Error; findErr != nil {
						return nil, fmt.Errorf("创建用户失败且查询用户也失败: %w, 查询错误: %w", err, findErr)
					}
					// 找到了用户，继续后续流程
				} else {
					return nil, fmt.Errorf("创建用户失败: %w", err)
				}
			}
		} else {
			return nil, fmt.Errorf("查询用户失败: %w", err)
		}
	} else {
		// 如果用户已存在但缺少信息，则更新非空字段
		updateData := appModel.AppUsers{}
		needUpdate := false

		if user.CasdoorId == nil && claims.User.Id != "" {
			updateData.CasdoorId = &claims.User.Id
			needUpdate = true
		}
		if user.CasdoorOwner == nil && claims.User.Owner != "" {
			updateData.CasdoorOwner = &claims.User.Owner
			needUpdate = true
		}
		if user.CasdoorName == nil && claims.User.Name != "" {
			updateData.CasdoorName = &claims.User.Name
			needUpdate = true
		}
		if user.Nickname == nil && claims.User.DisplayName != "" {
			updateData.Nickname = &claims.User.DisplayName
			needUpdate = true
		}
		if user.Avatar == nil && claims.User.Avatar != "" {
			updateData.Avatar = &claims.User.Avatar
			needUpdate = true
		}
		if user.Email == nil && claims.User.Email != "" {
			updateData.Email = &claims.User.Email
			needUpdate = true
		}
		if user.Phone == nil && claims.User.Phone != "" {
			updateData.Phone = &claims.User.Phone
			needUpdate = true
		}
		if user.Source == nil && claims.SigninMethod != "" {
			updateData.Source = &claims.SigninMethod
			needUpdate = true
		}

		if needUpdate {
			if err := db.Model(&user).Where("casdoor_id = ?", claims.User.Id).Updates(&updateData).Error; err != nil {
				return nil, fmt.Errorf("更新用户信息失败: %w", err)
			}
		}

		// 更新登录时间和登录次数
		now := time.Now()
		loginCount := 1
		if user.LoginCount != nil {
			loginCount = *user.LoginCount + 1
		}
		loginUpdate := appModel.AppUsers{
			LastLoginAt: &now,
			LoginCount:  &loginCount,
		}
		if err := db.Model(&user).Where("casdoor_id = ?", claims.User.Id).Updates(&loginUpdate).Error; err != nil {
			return nil, fmt.Errorf("更新登录统计失败: %w", err)
		}

		// 重新查询用户信息以获取最新数据
		if err := db.Where("casdoor_id = ?", claims.User.Id).First(&user).Error; err != nil {
			return nil, fmt.Errorf("查询用户信息失败: %w", err)
		}
	}

	return &user, nil
}
