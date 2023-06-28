//	@author	hongjun500
//	@date	2023/6/15 14:06
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 使用 casbin 封装的动态权限认证，并将其封装成 gin 的中间件

package mid

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/pkg/security"
)

// GinCasbinMiddleware 自定义动态权限中间件
func GinCasbinMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 跨域请求会先发送一个 OPTIONS 请求，这里我们给 OPTIONS 请求直接返回正常状态
		if context.Request.Method == http.MethodOptions {
			context.Next()
			return
		}
		url := context.Request.URL
		path := url.Path
		// fmt.Println("url = ", url)
		// fmt.Println("path = ", path)
		sub, _ := context.Get("username")
		// 检查资源权限
		policy := security.Enforcer.GetPolicy()
		fmt.Println("policy = ", policy)
		if !security.Enforcer.Enforce(sub, path, "*") {
			gin_common.CreateForbidden(context)
			context.Abort()
			return
		}
		context.Next()
	}
}
