// @author hongjun500
// @date 2023/6/14 10:01
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsMemberLevelRouter struct {
	services.UmsMemberLevelService
}

func NewUmsMemberLevelRouter(service services.UmsMemberLevelService) *UmsMemberLevelRouter {
	return &UmsMemberLevelRouter{UmsMemberLevelService: service}
}

// GroupUmsMemberLevelRouter 会员等级管理路由
func (router *UmsMemberLevelRouter) GroupUmsMemberLevelRouter(routerEngine *gin.Engine) {
	umsMemberLevelGroup := routerEngine.Group("/memberLevel").Use(mid.GinJWTMiddleware())
	{
		// 查看所有会员等级
		umsMemberLevelGroup.GET("/list", router.UmsMemberLevelService.UmsMemberLevelList)
	}
}
