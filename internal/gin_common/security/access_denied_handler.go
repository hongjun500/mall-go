//	@author	hongjun500
//	@date	2023/6/15 10:06
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 自定义处理无权限返回结果

package security

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
)

func AccessDeniedHandler(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Cache-Control", "no-cache")
	context.Header("Content-Type", "application/json; charset=utf-8")
	gin_common.CreateForbidden(context)
	context.Abort()
}
