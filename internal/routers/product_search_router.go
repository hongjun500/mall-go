// @author hongjun500
// @date 2023/6/19 17:57
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/services"
)

type SearchRouter struct {
	services.ProductSearchService
}

func NewSearchRouter(service services.ProductSearchService) *SearchRouter {
	return &SearchRouter{ProductSearchService: service}
}

// GroupProductRouter 搜索路由
func (router *SearchRouter) GroupProductRouter(searchGroup *gin.RouterGroup) {
	{
		// 搜索商品
		// searchGroup.GET("/list", router.ProductSearchService.ProductSearchList)
	}
}
