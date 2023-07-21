// @author hongjun500
// @date 2023/7/14 13:45
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package r_mall_admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type PmsProductCategoryRouter struct {
	s_mall_admin.PmsProductCategoryService
}

func NewPmsProductCategoryRouter(service s_mall_admin.PmsProductCategoryService) *PmsProductCategoryRouter {
	return &PmsProductCategoryRouter{
		PmsProductCategoryService: service,
	}
}

// GroupPmsProductCategoryRouter 商品分类路由组
func (router *PmsProductCategoryRouter) GroupPmsProductCategoryRouter(productCategoryGroup *gin.RouterGroup) {

	{
		// 添加商品分类
		productCategoryGroup.POST("/create", router.CreateProductCategory)
		// 修改商品分类
		productCategoryGroup.POST("/update/:id", router.UpdateProductCategory)
		// 分页查询商品分类
		productCategoryGroup.GET("/list/:parentId", router.ListProductCategory)
		// 根据id获取商品分类
		productCategoryGroup.GET("/:id", router.GetProductCategory)
		// 删除商品分类
		productCategoryGroup.POST("/delete/:id", router.DeleteProductCategory)
		// 批量修改导航状态
		productCategoryGroup.POST("/update/navStatus", router.UpdateNavStatus)
		// 批量修改显示状态
		productCategoryGroup.POST("/update/showStatus", router.UpdateShowStatus)
		// 查询所有一级分类及子分类
		productCategoryGroup.GET("/list/withChildren", router.ListWithChildren)
	}
}

// CreateProductCategory 添加商品分类
//
//	@Summary		添加商品分类
//	@Description	添加商品分类
//	@Tags			商品分类
//	@Accept			json
//	@Produce		json
//	@Param			request	body		admin_dto.PmsProductCategoryDTO true "添加商品分类"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/productCategory/create [post]
func (router *PmsProductCategoryRouter) CreateProductCategory(context *gin.Context) {
	var dto admin_dto.PmsProductCategoryDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	count, err := router.PmsProductCategoryService.CreateProductCategory(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, count)
}

// UpdateProductCategory 修改商品分类
//
//	@Summary		修改商品分类
//	@Description	修改商品分类
//	@Tags			商品分类
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"商品分类id"
//	@Param			request	body		admin_dto.PmsProductCategoryDTO true "修改商品分类"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/productCategory/update/{id} [post]
func (router *PmsProductCategoryRouter) UpdateProductCategory(context *gin.Context) {
	var dto admin_dto.PmsProductCategoryDTO
	var path base_dto.PathVariableDTO
	err := context.ShouldBind(&dto)
	err = context.ShouldBindUri(&path)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	count, err := router.PmsProductCategoryService.UpdateProductCategory(path.Id, dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, count)
}

// ListProductCategory 分页查询商品分类
//
//	@Summary		分页查询商品分类
//	@Description	分页查询商品分类
//	@Tags			商品分类
//	@Accept			json
//	@Produce		json
//	@Param			parentId		path	int		true	"父分类的编号"
//	@Param			pageNum			query	int		true	"页码"
//	@Param			pageSize		query	int		true	"每页数量"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/productCategory/list/{parentId} [get]
func (router *PmsProductCategoryRouter) ListProductCategory(context *gin.Context) {
	var path base_dto.PathVariableDTO
	err := context.ShouldBindUri(&path)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	var page base_dto.PageDTO
	err = context.ShouldBind(&page)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	list, err := router.PmsProductCategoryService.ListProductCategory(path.Id, page)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// GetProductCategory 根据id获取商品分类
//
//	@Summary		根据id获取商品分类
//	@Description	根据id获取商品分类
//	@Tags			商品分类
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true "商品分类id"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/productCategory/{id} [get]
func (router *PmsProductCategoryRouter) GetProductCategory(context *gin.Context) {
	var path base_dto.PathVariableDTO
	err := context.ShouldBindUri(&path)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	productCategory := new(models.PmsProductCategory)
	item, err := productCategory.SelectById(router.PmsProductCategoryService.DbFactory.GormMySQL, path.Id)
	if err != nil {
		gin_common.CreateFail(context, gin_common.DatabaseError)
		return
	}
	gin_common.CreateSuccess(context, item)
}

// DeleteProductCategory 根据id删除商品分类
//
//	@Summary		根据id删除商品分类
//	@Description	根据id删除商品分类
//	@Tags			商品分类
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"商品分类id"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/productCategory/delete/{id} [post]
func (router *PmsProductCategoryRouter) DeleteProductCategory(context *gin.Context) {
	var path base_dto.PathVariableDTO
	err := context.ShouldBindUri(&path)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	count, err := new(models.PmsProductCategory).DeleteById(router.PmsProductCategoryService.DbFactory.GormMySQL, path.Id)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, count)
}

// UpdateNavStatus 批量修改导航状态
//
//	@Summary		批量修改导航状态
//	@Description	批量修改导航状态
//	@Tags			商品分类
//	@Accept			json
//	@Produce		json
//	@Param			ids	formData []int	true "商品分类id集合"
//	@Param			navStatus formData	int		true "导航状态"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/productCategory/update/navStatus [post]
func (router *PmsProductCategoryRouter) UpdateNavStatus(context *gin.Context) {
	var dto base_dto.IdsDTO
	err := context.ShouldBind(&dto)
	navStatus := context.PostForm("navStatus")
	if err != nil || navStatus == "" {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	navStatusi, _ := strconv.Atoi(navStatus)
	if err != nil {
		return
	}
	count, err := router.PmsProductCategoryService.UpdateNavStatus(dto.Ids, navStatusi)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, count)
}

// UpdateShowStatus 批量修改导航状态
//
//	@Summary		批量修改导航状态
//	@Description	批量修改导航状态
//	@Tags			商品分类
//	@Accept			json
//	@Produce		json
//	@Param			ids	formData	[]int		true	"商品分类id集合"
//	@Param			showStatus formData	int		true	"显示状态"
//	@Security		GinJWTMiddleware
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/productCategory/update/showStatus [post]
func (router *PmsProductCategoryRouter) UpdateShowStatus(context *gin.Context) {
	var dto base_dto.IdsDTO
	err := context.ShouldBind(&dto)
	showStatus := context.PostForm("showStatus")
	if err != nil || showStatus == "" {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	showStatusi, _ := strconv.Atoi(showStatus)
	if err != nil {
		return
	}
	count, err := router.PmsProductCategoryService.UpdateShowStatus(dto.Ids, showStatusi)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, count)
}

// ListWithChildren 查询所有一级分类及子分类
//
//	@Summary		查询所有一级分类及子分类
//	@Description	查询所有一级分类及子分类
//	@Tags			商品分类
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Security		GinJWTMiddleware
//	@Router			/productCategory/list/withChildren [get]
func (router *PmsProductCategoryRouter) ListWithChildren(context *gin.Context) {
	list, err := router.PmsProductCategoryService.ListWithChildren()
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}
