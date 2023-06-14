// @author hongjun500
// @date 2023/6/11
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台用户路由

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsAdminRouter struct {
	services.UmsAdminService
}

func NewUmsAdminRouter(service services.UmsAdminService) *UmsAdminRouter {
	return &UmsAdminRouter{UmsAdminService: service}
}

// GroupUmsAdminRouter 用户管理路由
func (router *UmsAdminRouter) GroupUmsAdminRouter(routerEngine *gin.Engine) {
	umsAdminGroup := routerEngine.Group("/admin")
	{
		// 用户注册
		umsAdminGroup.POST("/register", router.UmsAdminService.UmsAdminRegister)
		/*adminGroup.POST("/register", router.UmsAdminRegister) 这样写也可以*/

		// 用户登录
		umsAdminGroup.POST("/login", router.UmsAdminService.UmsAdminLogin)

		// 用户登出
		umsAdminGroup.POST("/logout", router.UmsAdminService.UmsAdminLogout)
		umsAdminGroup.GET("/authTest", router.UmsAdminService.UmsAdminAuthTest)
	}

	authGroup := routerEngine.Group("/admin").Use(mid.GinJWTMiddleware())
	{
		// 刷新 token
		authGroup.POST("/refreshToken", router.UmsAdminService.UmsAdminRefreshToken)
		// 根据用户 ID 获取用户信息
		umsAdminGroup.GET("/info", router.UmsAdminService.UmsAdminInfo)
		// authGroup.GET("/info/:user_id", router.UmsAdminService.UmsAdminInfo)
		// 用户列表分页
		authGroup.GET("/list", router.UmsAdminService.UmsAdminListPage)
		// 获取指定用户信息
		authGroup.GET("/:id", router.UmsAdminService.UmsAdminItem)
		// 修改指定用户信息
		authGroup.POST("/update/:id", router.UmsAdminService.UmsAdminUpdate)
		// 删除指定用户
		authGroup.POST("/delete/:id", router.UmsAdminService.UmsAdminDelete)
		// 修改指定用户状态
		authGroup.POST("/updateStatus/:id", router.UmsAdminService.UmsAdminUpdateStatus)
		// 给用户分配角色
		authGroup.POST("/role/update", router.UmsAdminService.UmsAdminRoleUpdate)
		// 获取指定用户的角色
		authGroup.GET("/role/:adminId", router.UmsAdminService.UmsAdminRoleItem)
		// 修改指定用户密码
		authGroup.POST("/updatePassword", router.UmsAdminService.UmsAdminUpdatePassword)
	}

}
