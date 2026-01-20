package app

import (
	"time"

	appModel "github.com/flipped-aurora/gin-vue-admin/server/model/app"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	appRes "github.com/flipped-aurora/gin-vue-admin/server/model/app/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

// 打包

type LoginApi struct{}

// CasdoorLogin
// @Tags     Login
// @Summary  Casdoor OAuth 2.0 登录
// @Accept   application/json
// @Produce  application/json
// @Param    code query string false "授权码（Authorization Code），从 Casdoor 授权回调中获取"
// @Param    state query string false "状态参数，用于防止 CSRF 攻击（可选）"
// @Success  200 {object} response.Response{data=appRes.LoginResponse,msg=string} "登录成功"
// @Failure  400 {object} response.Response "参数错误"
// @Failure  500 {object} response.Response "服务器错误"
// @Router   /login/casdoorLogin [post]
func (a *LoginApi) CasdoorLogin(c *gin.Context) {
	ctx := c.Request.Context()
	code := c.Query("code")

	// 调用服务层获取用户信息
	user, err := loginService.CasdoorLogin(ctx, code)
	if err != nil {
		response.FailWithMessage("登录失败: "+err.Error(), c)
		return
	}

	// 调用 TokenNext 处理 token 生成和登录响应
	a.TokenNext(c, *user)
}

func (a *LoginApi) TokenNext(c *gin.Context, user appModel.AppUsers) {
	// 生成 JWT token
	token, claims, err := utils.LoginAppUserToken(user)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	userResp := appModel.AppUsers{
		GVA_MODEL:    user.GVA_MODEL,
		CasdoorOwner: user.CasdoorOwner,
		CasdoorName:  user.CasdoorName,
		CasdoorId:    user.CasdoorId,
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		Email:        user.Email,
		Phone:        user.Phone,
		Source:       user.Source,
		LastLoginAt:  user.LastLoginAt,
		LoginCount:   user.LoginCount,
	}

	// 使用 Casdoor 用户标识作为唯一标识
	userKey := "app:" + *user.CasdoorId
	maxAgeSeconds := int(claims.RegisteredClaims.ExpiresAt.Unix() - time.Now().Unix())

	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, maxAgeSeconds)
		response.OkWithDetailed(appRes.LoginResponse{
			User:      userResp,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	// 多点登录：检查 Redis 中是否已有 token
	if jwtStr, err := jwtService.GetRedisJWT(userKey); err == redis.Nil {
		// Redis 中没有，设置新 token
		if err := jwtService.SetRedisJWT(token, userKey); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, maxAgeSeconds)
		response.OkWithDetailed(appRes.LoginResponse{
			User:      userResp,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("获取登录状态失败!", zap.Error(err))
		response.FailWithMessage("获取登录状态失败", c)
	} else {
		// Redis 中已有 token，将旧 token 加入黑名单，设置新 token
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, userKey); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, maxAgeSeconds)
		response.OkWithDetailed(appRes.LoginResponse{
			User:      userResp,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}
