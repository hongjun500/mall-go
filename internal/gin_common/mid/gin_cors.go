// @author hongjun500
// @date 2023/6/14 17:14
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package mid

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GinCORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, "+
			"Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}

		context.Next()
	}
}

func corss() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 使用 cors 中间件
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true       // 允许所有来源
		config.AllowCredentials = true      // 允许发送 cookie
		config.AllowHeaders = []string{"*"} // 放行全部原始头信息
		config.AllowMethods = []string{"*"} // 允许所有请求方法跨域调用
		context.Next()
	}
}
