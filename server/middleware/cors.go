package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 直接放行所有跨域请求并放行所有 OPTIONS 方法
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		
		// 如果没有Origin头，允许所有来源（用于非浏览器请求）
		if origin == "" {
			c.Header("Access-Control-Allow-Origin", "*")
		} else {
			c.Header("Access-Control-Allow-Origin", origin)
		}
		
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id,Origin,X-Requested-With,Accept")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT, PATCH, HEAD")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400") // 预检请求缓存24小时

		// 放行所有OPTIONS方法（预检请求）
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 处理请求
		c.Next()
	}
}

// CorsByRules 按照配置处理跨域请求
func CorsByRules() gin.HandlerFunc {
	// 放行全部
	if global.GVA_CONFIG.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		whitelist := checkCors(origin)

		// 通过检查, 添加请求头
		if whitelist != nil {
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			if whitelist.AllowHeaders != "" {
				c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			} else {
				// 默认请求头
				c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id,Origin,X-Requested-With,Accept")
			}
			if whitelist.AllowMethods != "" {
				c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			} else {
				// 默认方法
				c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT, PATCH, HEAD")
			}
			if whitelist.ExposeHeaders != "" {
				c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			} else {
				// 默认暴露头
				c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
			}
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
			c.Header("Access-Control-Max-Age", "86400") // 预检请求缓存24小时
		}

		// 处理预检请求（OPTIONS）
		if c.Request.Method == http.MethodOptions {
			// 如果通过了白名单检查，或者不是严格模式，都允许预检请求
			if whitelist != nil || global.GVA_CONFIG.Cors.Mode != "strict-whitelist" {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
			// 严格模式且未通过检查，拒绝预检请求
			if whitelist == nil && global.GVA_CONFIG.Cors.Mode == "strict-whitelist" {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}

		// 严格白名单模式且未通过检查，直接拒绝处理请求（排除健康检查）
		if whitelist == nil && global.GVA_CONFIG.Cors.Mode == "strict-whitelist" && 
			!(c.Request.Method == "GET" && c.Request.URL.Path == "/health") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// 处理请求
		c.Next()
	}
}

func checkCors(currentOrigin string) *config.CORSWhitelist {
	if currentOrigin == "" {
		return nil
	}
	for _, whitelist := range global.GVA_CONFIG.Cors.Whitelist {
		// 遍历配置中的跨域头，寻找匹配项
		// 支持精确匹配和通配符匹配（如 http://*.example.com）
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
		// 如果配置为 "*"，允许所有来源
		if whitelist.AllowOrigin == "*" {
			return &whitelist
		}
	}
	return nil
}
