//	@author	hongjun500
//	@date	2023/6/11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台菜单管理路由

package mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsMenuRouter struct {
	services.UmsMenuService
}

func NewUmsMenuRouter(service services.UmsMenuService) *UmsMenuRouter {
	return &UmsMenuRouter{UmsMenuService: service}
}

// GroupUmsMenuRouter 后台菜单路由
func (router *UmsMenuRouter) GroupUmsMenuRouter(umsMenuGroup *gin.RouterGroup) {
	{
		// 新增菜单
		umsMenuGroup.POST("/create", router.UmsMenuService.UmsMenuCreate)
		// 修改菜单
		umsMenuGroup.POST("/update/:id", router.UmsMenuService.UmsMenuUpdate)
		// 删除菜单
		umsMenuGroup.POST("/delete/:id", router.UmsMenuService.UmsMenuDelete)
		// 根据ID获取菜单详情
		umsMenuGroup.GET("/:id", router.UmsMenuService.UmsMenuItem)
		// 分页获取菜单列表
		umsMenuGroup.GET("/list/:parentId", router.UmsMenuService.UmsMenuPageList)
		// 修改菜单显示状态
		umsMenuGroup.POST("/updateHidden/:id", router.UmsMenuService.UmsMenuUpdateHidden)
		// 树形结构返回所有菜单列表
		umsMenuGroup.GET("/treeList", router.UmsMenuService.UmsMenuTreeList)
	}
}
