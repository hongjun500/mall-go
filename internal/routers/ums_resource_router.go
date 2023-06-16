// @author hongjun500
// @date 2023/6/13 11:10
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsResourceRouter struct {
	services.UmsResourceService
}

func NewUmsResourceRouter(service services.UmsResourceService) *UmsResourceRouter {
	return &UmsResourceRouter{UmsResourceService: service}
}

// GroupUmsResourceRouter 后台资源管理路由
func (router *UmsResourceRouter) GroupUmsResourceRouter(routerEngine *gin.Engine) {
	umsResourceGroup := routerEngine.Group("/resource").Use(mid.GinJWTMiddleware()).Use(mid.GinCasbinMiddleware())
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
