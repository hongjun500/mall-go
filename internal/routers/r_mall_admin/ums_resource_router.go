//	@author	hongjun500
//	@date	2023/6/13 11:10
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsResourceRouter struct {
	s_mall_admin.UmsResourceService
}

func NewUmsResourceRouter(service s_mall_admin.UmsResourceService) *UmsResourceRouter {
	return &UmsResourceRouter{UmsResourceService: service}
}

// GroupUmsResourceRouter 后台资源管理路由
func (router *UmsResourceRouter) GroupUmsResourceRouter(umsResourceGroup *gin.RouterGroup) {

	{
		// 添加后台资源
		umsResourceGroup.POST("/create", router.UmsResourceService.UmsResourceCreate)
		// 修改后台资源
		umsResourceGroup.POST("/update/:id", router.UmsResourceService.UmsResourceUpdate)
		// 根据ID获取资源详情
		umsResourceGroup.GET("/:id", router.UmsResourceService.UmsResourceItem)
		// 根据ID删除后台资源
		umsResourceGroup.POST("/delete/:id", router.UmsResourceService.UmsResourceDelete)
		// 分页模糊查询后台资源
		umsResourceGroup.GET("/list/resource", router.UmsResourceService.UmsResourcePageList)
		// 查询所有后台资源
		umsResourceGroup.GET("/listAll", router.UmsResourceService.UmsResourceList)

	}
}
