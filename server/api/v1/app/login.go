package app

import (
	"fmt"
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

	// 记录请求信息
	global.GVA_LOG.Info("Casdoor登录请求开始",
		zap.String("method", c.Request.Method),
		zap.String("url", c.Request.URL.String()),
		zap.String("remote_addr", c.ClientIP()),
		zap.String("user_agent", c.Request.UserAgent()),
	)

	// 尝试从多个位置获取 code 参数：query、form、body
	code := c.Query("code")
	if code == "" {
		code = c.PostForm("code")
	}
	if code == "" {
		// 尝试从 JSON body 中获取
		var jsonBody struct {
			Code  string `json:"code"`
			State string `json:"state"`
		}
		if err := c.ShouldBindJSON(&jsonBody); err == nil {
			code = jsonBody.Code
		}
	}

	state := c.Query("state")
	if state == "" {
		state = c.PostForm("state")
	}

	// 记录接收到的参数（不记录敏感信息）
	global.GVA_LOG.Info("Casdoor登录请求参数",
		zap.String("code_length", fmt.Sprintf("%d", len(code))),
		zap.Bool("code_empty", code == ""),
		zap.String("state", state),
		zap.Any("query_params", c.Request.URL.Query()),
	)

	// 验证 code 参数
	if code == "" {
		global.GVA_LOG.Error("Casdoor登录失败: authorization code 为空",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.Any("query_params", c.Request.URL.Query()),
			zap.String("content_type", c.GetHeader("Content-Type")),
		)
		response.FailWithMessage("登录失败: authorization code 不能为空，请确保从 Casdoor 授权回调中正确获取 code 参数", c)
		return
	}

	// 调用服务层获取用户信息
	user, err := loginService.CasdoorLogin(ctx, code)
	if err != nil {
		global.GVA_LOG.Error("Casdoor登录失败",
			zap.String("code_length", fmt.Sprintf("%d", len(code))),
			zap.Error(err),
		)
		response.FailWithMessage("登录失败: "+err.Error(), c)
		return
	}

	global.GVA_LOG.Info("Casdoor登录成功",
		zap.String("casdoor_id", func() string {
			if user.CasdoorId != nil {
				return *user.CasdoorId
			}
			return ""
		}()),
		zap.Uint("user_id", user.ID),
	)

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
