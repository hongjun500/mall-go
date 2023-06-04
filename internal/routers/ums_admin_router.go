package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsAdminRouter struct {
	*services.UmsAdminService
}

func CreateUmsAdminRouter(service *services.UmsAdminService) *UmsAdminRouter {
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
		// 刷新 token
		umsAdminGroup.POST("/refreshToken", router.UmsAdminService.UmsAdminRefreshToken)
		// 根据用户 ID 获取用户信息
		umsAdminGroup.GET("/info", router.UmsAdminService.UmsAdminInfo)
		// 用户列表分页
		umsAdminGroup.GET("/list", router.UmsAdminService.UmsAdminListPage)
	}
	authGroup := routerEngine.Group("/auth").Use(mid.GinJWTMiddleware())
	{
		authGroup.GET("/ping", func(context *gin.Context) {
			gin_common.CreateSuccess("ok", context)
		})
	}

	unAuthGroup := routerEngine.Group("/unauth")
	{
		unAuthGroup.GET("/ping", func(context *gin.Context) {
			gin_common.CreateSuccess("ok", context)
		})
	}
}
