//	@author	hongjun500
//	@date	2023/6/13 10:43
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台资源分类管理路由

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsResourceCategoryRouter struct {
	s_mall_admin.UmsResourceCategoryService
}

func NewUmsResourceCategoryRouter(service s_mall_admin.UmsResourceCategoryService) *UmsResourceCategoryRouter {
	return &UmsResourceCategoryRouter{UmsResourceCategoryService: service}
}

// GroupUmsResourceCategoryRouter 后台资源分类路由
func (router *UmsResourceCategoryRouter) GroupUmsResourceCategoryRouter(umsResourceCategoryGroup *gin.RouterGroup) {

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
