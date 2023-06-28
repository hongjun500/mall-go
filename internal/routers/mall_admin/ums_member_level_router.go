//	@author	hongjun500
//	@date	2023/6/14 10:01
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsMemberLevelRouter struct {
	services.UmsMemberLevelService
}

func NewUmsMemberLevelRouter(service services.UmsMemberLevelService) *UmsMemberLevelRouter {
	return &UmsMemberLevelRouter{UmsMemberLevelService: service}
}

// GroupUmsMemberLevelRouter 会员等级管理路由
func (router *UmsMemberLevelRouter) GroupUmsMemberLevelRouter(umsMemberLevelGroup *gin.RouterGroup) {
	{
		// 查看所有会员等级
		umsMemberLevelGroup.GET("/list", router.UmsMemberLevelService.UmsMemberLevelList)
	}
}
