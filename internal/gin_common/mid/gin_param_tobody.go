//	@author	hongjun500
//	@date	2023/6/9 15:27
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 解析请求头中的参数，将参数解析加到请求体中

package mid

import "github.com/gin-gonic/gin"

func GinParamToBody() gin.HandlerFunc {
	return func(context *gin.Context) {
		// todo 从请求头中获取参数，将参数解析加到请求体中
		// context.Request.Body
		// 请求前
		context.Next()
		// 请求后

	}
}
