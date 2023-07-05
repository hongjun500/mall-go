// @author	hongjun500
// @date	2023/6/19 17:57
// @tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package r_mall_search

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_search"
)

type ProductSearchRouter struct {
	s_mall_search.ProductSearchService
}

// 弃用这种方式，使用 NewSearchRouter() 方法
/*var (
	productSearchService s_mall_search.ProductSearchService
)*/

func NewSearchRouter(service s_mall_search.ProductSearchService) *ProductSearchRouter {
	// productSearchService = service
	return &ProductSearchRouter{ProductSearchService: service}
}

// GroupProductRouter 搜索路由组
func (router *ProductSearchRouter) GroupProductRouter(searchGroup *gin.RouterGroup) {
	{

		// 导入所有数据库中商品到ES
		searchGroup.POST("/importAll", router.importAll)
		// 根据id删除商品
		searchGroup.GET("/delete/:id", router.delete)
		// 根据商品id推荐商品
		searchGroup.GET("/recommend/:id", router.recommend)
	}
}

// importAll 将数据库中的商品信息导入到 es
// @Summary		将数据库中的商品信息导入到 es
// @Description	将数据库中的商品信息导入到 es
// @Tags			搜索商品管理
// @Accept			application/json
// @Produce		application/json
// @Security 		GinJWTMiddleware
// @Success		200	{object}	gin_common.GinCommonResponse
// @Router			/product/importAll [post]
func (router *ProductSearchRouter) importAll(context *gin.Context) {
	err := router.ProductSearchService.ImportAll()
	if err != nil {
		gin_common.CreateFail(context, err.Error())
		return
	}
	gin_common.Create(context)
}

// delete 根据id删除商品
// @Summary		将数据库中的商品信息导入到 es
// @Description	将数据库中的商品信息导入到 es
// @Tags		搜索商品管理
// @Accept		application/json
// @Produce		application/json
// @Param		id	path	int	true	"id"
// @Security 	GinJWTMiddleware
// @Success		200	{object}	gin_common.GinCommonResponse
// @Router		/product/delete/{id} [get]
func (router *ProductSearchRouter) delete(context *gin.Context) {
	var pathVariableDTO base_dto.PathVariableDTO
	err := context.BindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, _ := router.ProductSearchService.Delete(pathVariableDTO.Id)
	gin_common.CreateSuccess(context, rows)
}

// recommend 根据商品id推荐商品
// @Summary		根据商品id推荐商品
// @Description	根据商品id推荐商品
// @Tags		搜索商品管理
// @Accept		application/json
// @Produce		application/json
// @Param		id	path	int	true	"id"
// @Security 	GinJWTMiddleware
// @Success		200	{object}	gin_common.GinCommonResponse
// @Router		/product/recommend/{id} [get]
func (router *ProductSearchRouter) recommend(context *gin.Context) {
	var pathVariableDTO base_dto.PathVariableDTO
	var pageDTO base_dto.PageDTO
	err := context.BindUri(&pathVariableDTO)
	err = context.BindQuery(&pageDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	page, _ := router.ProductSearchService.SearchById(pathVariableDTO.Id, pageDTO.PageNum, pageDTO.PageSize)
	gin_common.CreateSuccess(context, page)
}
