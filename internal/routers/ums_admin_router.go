package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsAdminRouter struct {
	services.UmsAdminService
}

func CreateUmsAdminRouter(service services.UmsAdminService) *UmsAdminRouter {
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

	}

	authGroup := routerEngine.Group("/admin").Use(mid.GinJWTMiddleware())
	{
		// 刷新 token
		authGroup.POST("/refreshToken", router.UmsAdminService.UmsAdminRefreshToken)
		// 根据用户 ID 获取用户信息
		authGroup.GET("/info", router.UmsAdminService.UmsAdminInfo)
		// authGroup.GET("/info/:user_id", router.UmsAdminService.UmsAdminInfo)
		// 用户列表分页
		authGroup.POST("/list", router.UmsAdminService.UmsAdminListPage)
		// 指定用户 ID 获取用户信息
		authGroup.GET("/:user_id", router.UmsAdminService.UmsAdminItem)
		// 修改指定用户信息
		authGroup.POST("/update/:user_id", router.UmsAdminService.UmsAdminUpdate)
	}

}
