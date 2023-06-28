//	@author	hongjun500
//	@date	2023/6/14 10:01
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsMemberLevelRouter struct {
	s_mall_admin.UmsMemberLevelService
}

func NewUmsMemberLevelRouter(service s_mall_admin.UmsMemberLevelService) *UmsMemberLevelRouter {
	return &UmsMemberLevelRouter{UmsMemberLevelService: service}
}

// GroupUmsMemberLevelRouter 会员等级管理路由
func (router *UmsMemberLevelRouter) GroupUmsMemberLevelRouter(umsMemberLevelGroup *gin.RouterGroup) {
	{
		// 查看所有会员等级
		umsMemberLevelGroup.GET("/list", router.UmsMemberLevelService.UmsMemberLevelList)
	}
}
