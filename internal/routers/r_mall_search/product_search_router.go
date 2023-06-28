//	@author	hongjun500
//	@date	2023/6/19 17:57
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package r_mall_search

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services/s_mall_search"
)

type ProductSearchRouter struct {
	s_mall_search.ProductSearchService
}

func NewSearchRouter(service s_mall_search.ProductSearchService) *ProductSearchRouter {
	return &ProductSearchRouter{ProductSearchService: service}
}

// GroupProductRouter 搜索路由
func (router *ProductSearchRouter) GroupProductRouter(searchGroup *gin.RouterGroup) {
	{
		// 搜索商品
		// searchGroup.GET("/list", router.ProductSearchService.ProductSearchList)
		// 导入所有数据库中商品到ES
		searchGroup.POST("/importAll", router.ProductSearchService.ImportAll)
	}
}
