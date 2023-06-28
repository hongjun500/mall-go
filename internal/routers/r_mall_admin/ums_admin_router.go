//	@author	hongjun500
//	@date	2023/6/11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台用户路由

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsAdminRouter struct {
	s_mall_admin.UmsAdminService
}

func NewUmsAdminRouter(service s_mall_admin.UmsAdminService) *UmsAdminRouter {
	return &UmsAdminRouter{UmsAdminService: service}
}

// GroupUmsAdminRouter 用户管理路由
func (router *UmsAdminRouter) GroupUmsAdminRouter(umsAdminGroup *gin.RouterGroup) {
	// umsAdminGroup := routerEngine.Group("/admin")

	// 刷新 token
	umsAdminGroup.POST("/refreshToken", router.UmsAdminService.UmsAdminRefreshToken)
	// 根据用户 ID 获取用户信息
	umsAdminGroup.GET("/info", router.UmsAdminService.UmsAdminInfo)
	// umsAdminGroup.GET("/info/:user_id", router.UmsAdminService.UmsAdminInfo)
	// 用户列表分页
	umsAdminGroup.GET("/list", router.UmsAdminService.UmsAdminListPage)
	// 获取指定用户信息
	umsAdminGroup.GET("/:id", router.UmsAdminService.UmsAdminItem)
	// 修改指定用户信息
	umsAdminGroup.POST("/update/:id", router.UmsAdminService.UmsAdminUpdate)
	// 删除指定用户
	umsAdminGroup.POST("/delete/:id", router.UmsAdminService.UmsAdminDelete)
	// 修改指定用户状态
	umsAdminGroup.POST("/updateStatus/:id", router.UmsAdminService.UmsAdminUpdateStatus)
	// 给用户分配角色
	umsAdminGroup.POST("/role/update", router.UmsAdminService.UmsAdminRoleUpdate)
	// 获取指定用户的角色
	umsAdminGroup.GET("/role/:adminId", router.UmsAdminService.UmsAdminRoleItem)
	// 修改指定用户密码
	umsAdminGroup.POST("/updatePassword", router.UmsAdminService.UmsAdminUpdatePassword)

}
