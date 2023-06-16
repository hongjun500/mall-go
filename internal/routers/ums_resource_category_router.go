// @author hongjun500
// @date 2023/6/13 10:43
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description: 后台资源分类管理路由

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/services"
)

type UmsResourceCategoryRouter struct {
	services.UmsResourceCategoryService
}

func NewUmsResourceCategoryRouter(service services.UmsResourceCategoryService) *UmsResourceCategoryRouter {
	return &UmsResourceCategoryRouter{UmsResourceCategoryService: service}
}

// GroupUmsResourceCategoryRouter 后台资源分类路由
func (router *UmsResourceCategoryRouter) GroupUmsResourceCategoryRouter(routerEngine *gin.Engine) {
	umsResourceCategoryGroup := routerEngine.Group("/resourceCategory").Use(mid.GinJWTMiddleware()).Use(mid.GinCasbinMiddleware())
	{
		// 添加后台资源分类
		umsResourceCategoryGroup.POST("/create", router.UmsResourceCategoryService.UmsResourceCategoryCreate)
		// 修改后台资源分类
		umsResourceCategoryGroup.POST("/update/:id", router.UmsResourceCategoryService.UmsResourceCategoryUpdate)
		// 根据ID删除后台资源分类
		umsResourceCategoryGroup.POST("/delete/:id", router.UmsResourceCategoryService.UmsResourceCategoryDelete)
		// 查询所有后台资源分类
		umsResourceCategoryGroup.GET("/listAll", router.UmsResourceCategoryService.UmsResourceCategoryList)
	}
}
