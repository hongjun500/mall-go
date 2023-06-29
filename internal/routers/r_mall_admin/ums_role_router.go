//	@author	hongjun500
//	@date	2023/6/13 15:47
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsRoleRouter struct {
	s_mall_admin.UmsRoleService
}

func NewUmsRoleRouter(service s_mall_admin.UmsRoleService) *UmsRoleRouter {
	return &UmsRoleRouter{UmsRoleService: service}
}

// GroupUmsRoleRouter 后台角色管理路由
func (router *UmsRoleRouter) GroupUmsRoleRouter(umsRoleGroup *gin.RouterGroup) {

	{
		// 添加角色
		umsRoleGroup.POST("/create", router.UmsRoleService.UmsRoleCreate)
		// 修改角色
		umsRoleGroup.POST("/update/:id", router.UmsRoleService.UmsRoleUpdate)
		// 批量删除角色
		umsRoleGroup.POST("/delete", router.UmsRoleService.UmsRoleDelete)
		// 获取所有角色
		umsRoleGroup.GET("/listAll", router.UmsRoleService.UmsRoleListAll)
		// 根据角色名称分页获取角色列表
		umsRoleGroup.GET("/list", router.UmsRoleService.UmsRoleList)
		// 修改角色状态
		umsRoleGroup.POST("/updateStatus/:id", router.UmsRoleService.UmsRoleUpdateStatus)
		// 获取角色相关菜单
		umsRoleGroup.GET("/listMenu/:roleId", router.UmsRoleService.UmsRoleListMenu)
		// 获取角色相关资源
		umsRoleGroup.GET("/listResource/:roleId", router.UmsRoleService.UmsRoleListResource)
		// 给角色分配菜单
		umsRoleGroup.POST("/allocMenu", router.UmsRoleService.UmsRoleAllocMenu)
		// 给角色分配资源
		umsRoleGroup.POST("/allocResource", router.UmsRoleService.UmsRoleAllocResource)
	}
}