// @author hongjun500
// @date 2023/7/21 14:04
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type PmsProductAttributeRouter struct {
	s_mall_admin.PmsProductAttributeService
}

func NewPmsProductAttributeRouter(service s_mall_admin.PmsProductAttributeService) *PmsProductAttributeRouter {
	return &PmsProductAttributeRouter{
		PmsProductAttributeService: service,
	}
}

func (router *PmsProductAttributeRouter) GroupPmsProductAttributeRouter(productAttributeGroup *gin.RouterGroup) {
	{

		// 添加商品属性信息
		productAttributeGroup.POST("/create", router.createProductAttribute)
		// 根据分类查询属性列表或参数列表
		productAttributeGroup.GET("/list/:cid", router.list)
		// 修改商品属性信息
		productAttributeGroup.POST("/update/:id", router.updateProductAttribute)
		// 查询单个商品属性
		productAttributeGroup.GET("/:id", router.getProductAttribute)
		// 批量删除商品属性
		productAttributeGroup.POST("/delete", router.deleteProductAttribute)
		// 根据商品分类的id获取商品属性及属性分类
		productAttributeGroup.GET("/attrInfo/:productCategoryId", router.listProductAttributeInfo)
	}
}

// createProductAttribute 添加商品属性信息
//
//	@Summary		添加商品属性信息
//	@Description	添加商品属性信息
//	@Tags			商品属性管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		admin_dto.PmsProductAttributeDTO  true	"修改指定商品属性"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/productAttribute/create	[post]
func (router *PmsProductAttributeRouter) createProductAttribute(context *gin.Context) {
	var dto admin_dto.PmsProductAttributeDTO
	if err := context.ShouldBind(&dto); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	count, err := router.PmsProductAttributeService.Create(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, count)
}

// list 根据分类查询属性列表或参数列表
//
//	@Summary		根据分类查询属性列表或参数列表
//	@Description	根据分类查询属性列表或参数列表
//	@Tags			商品属性管理
//	@Accept			json
//	@Produce		json
//	@Param			type	query		string		true	"0表示属性，1表示参数"
//	@Param			cid		query		int			true	"分类id"
//	@Param			pageNum		query		int			true	"页码"
//	@Param			pageSize	query		int			true	"每页数量"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/productAttribute/list/{cid}	[get]
func (router *PmsProductAttributeRouter) list(context *gin.Context) {
	var dto admin_dto.PmsProductAttributeListDTO
	if err := context.ShouldBind(&dto); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	list, err := router.PmsProductAttributeService.ListPage(dto.CategoryId, dto.Type, dto.PageNum, dto.PageSize)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// updateProductAttribute 修改商品属性信息
//
//	@Summary		修改商品属性信息
//	@Description	修改商品属性信息
//	@Tags			商品属性管理
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"商品属性id"
//	@Param			request	body		admin_dto.PmsProductAttributeDTO true	"修改指定商品属性"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/productAttribute/update/{id}	[post]
func (router *PmsProductAttributeRouter) updateProductAttribute(context *gin.Context) {
	var dto admin_dto.PmsProductAttributeDTO
	var pathDTO base_dto.PathVariableDTO
	err := context.ShouldBindUri(&pathDTO)
	err = context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	count, err := router.PmsProductAttributeService.Update(pathDTO.Id, dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, count)
}

// getProductAttribute 查询单个商品属性
//
//	@Summary		查询单个商品属性
//	@Description	查询单个商品属性
//	@Tags			商品属性管理
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"商品属性id"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/productAttribute/{id}	[get]
func (router *PmsProductAttributeRouter) getProductAttribute(context *gin.Context) {
	var pathDTO base_dto.PathVariableDTO

	if err := context.ShouldBindUri(&pathDTO); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	productAttribute, err := router.PmsProductAttributeService.GetItem(pathDTO.Id)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, productAttribute)
}

// deleteProductAttribute 批量删除商品属性
//
//	@Summary		批量删除商品属性
//	@Description	批量删除商品属性
//	@Tags			商品属性管理
//	@Accept			json
//	@Produce		json
//	@Param			ids		formData		[]int		true	"商品属性id集合"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/productAttribute/delete	[post]
func (router *PmsProductAttributeRouter) deleteProductAttribute(context *gin.Context) {
	var ids base_dto.IdsDTO
	if err := context.ShouldBind(&ids); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	count, err := router.PmsProductAttributeService.Delete(ids.Ids)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, count)
}

// listProductAttributeInfo 根据分类查询属性列表或参数列表
//
//	@Summary		根据分类查询属性列表或参数列表
//	@Description	根据分类查询属性列表或参数列表
//	@Tags			商品属性管理
//	@Accept			json
//	@Produce		json
//	@Param			productCategoryId		path		int			true	"分类id"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	admin_dto.PmsProductAttributeInfoDTO
//	@Router			/productAttribute/attrInfo/{productCategoryId}	[get]
func (router *PmsProductAttributeRouter) listProductAttributeInfo(context *gin.Context) {
	var pathDTO admin_dto.PmsProductAttributeInfoDTO
	if err := context.ShouldBindUri(&pathDTO); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	productAttributeInfo, err := router.PmsProductAttributeService.ListFromProductAttrInfo(pathDTO.CategoryId)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, productAttributeInfo)
}
