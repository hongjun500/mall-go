package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsAdminRouter struct {
	*services.UmsAdminService
}

func NewUmsAdminRouter(service *services.UmsAdminService) *UmsAdminRouter {
	return &UmsAdminRouter{UmsAdminService: service}
}

// GroupUmsAdminRouter 用户管理路由
func (router *UmsAdminRouter) GroupUmsAdminRouter(routerEngine *gin.Engine) {
	adminGroup := routerEngine.Group("/admin")
	{
		// 用户注册
		adminGroup.POST("/register", router.UmsAdminRegister)
		/*adminGroup.POST("/register", router.UmsAdminRegister) 这样写也可以*/

		// 用户登录
		// umsAdminGroup.POST("/login", router.UmsAdminService.UmsAdminLogin)
	}

	routerEngine.GET("/ping", router.UmsAdminRegister)
}
