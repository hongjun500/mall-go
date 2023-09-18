// @author hongjun500
// @date 2023/7/20 14:02
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type PmsProductAttributeCategoryRouter struct {
	s_mall_admin.PmsProductAttributeCategoryService
}

func NewPmsProductAttributeCategoryRouter(service s_mall_admin.PmsProductAttributeCategoryService) *PmsProductAttributeCategoryRouter {
	return &PmsProductAttributeCategoryRouter{
		PmsProductAttributeCategoryService: service,
	}
}

func (router *PmsProductAttributeCategoryRouter) GroupPmsProductAttributeCategoryRouter(productAttributeCategoryGroup *gin.RouterGroup) {
	{
		// 添加商品属性分类
		productAttributeCategoryGroup.POST("/create", router.create)
		// 修改商品属性分类
		productAttributeCategoryGroup.POST("/update/:id", router.update)
		// 删除商品属性分类
		productAttributeCategoryGroup.POST("/delete/:id", router.delete)
		// 获取单个商品属性分类信息
		productAttributeCategoryGroup.GET("/:id", router.getItem)
		// 分页获取所有商品属性分类
		productAttributeCategoryGroup.GET("/list", router.list)
		// 获取所有商品属性分类及其下属性
		productAttributeCategoryGroup.GET("/list/withAttr", router.listWithAttr)
	}
}

// create 添加商品属性分类
//
//	@Summary		添加商品属性分类
//	@Description	添加商品属性分类
//	@Tags			商品属性分类
//	@Accept			json
//	@Produce		json
//	@Param			name	formData	string	true	"属性分类名称"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /productAttribute/category/create [post]
func (router *PmsProductAttributeCategoryRouter) create(context *gin.Context) {
	var name string
	if name = context.PostForm("name"); name == "" {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	create, err := router.PmsProductAttributeCategoryService.Create(name)
	if err != nil || create == 0 {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, create)
}

// update 修改商品属性分类
//
//	@Summary		修改商品属性分类
//	@Description	修改商品属性分类
//	@Tags			商品属性分类
//	@Accept			json
//	@Produce		json
//	@Param			name	formData	string	true	"属性分类名称"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /productAttribute/category/update/{id} [post]
func (router *PmsProductAttributeCategoryRouter) update(context *gin.Context) {
	var pathId base_dto.PathVariableDTO
	name := context.PostForm("name")
	if err := context.ShouldBindUri(&pathId); err != nil || name == "" {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	update, err := router.PmsProductAttributeCategoryService.Update(pathId.Id, name)
	if err != nil || update == 0 {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, update)
}

// delete 删除商品属性分类
//
//	@Summary		删除商品属性分类
//	@Description	删除商品属性分类
//	@Tags			商品属性分类
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"属性分类id"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /productAttribute/category/delete/{id} [post]
func (router *PmsProductAttributeCategoryRouter) delete(context *gin.Context) {
	var pathId base_dto.PathVariableDTO
	if err := context.ShouldBindUri(&pathId); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.PmsProductAttributeCategoryService.Delete(pathId.Id)
	if err != nil || rows == 0 {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// getItem 获取单个商品属性分类信息
//
//	@Summary		获取单个商品属性分类信息
//	@Description	获取单个商品属性分类信息
//	@Tags			商品属性分类
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"属性分类id"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /productAttribute/category/{id} [get]
func (router *PmsProductAttributeCategoryRouter) getItem(context *gin.Context) {
	var pathId base_dto.PathVariableDTO
	if err := context.ShouldBindUri(&pathId); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	item, err := router.PmsProductAttributeCategoryService.GetByID(pathId.Id)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, item)
}

// list 分页获取所有商品属性分类
//
//	@Summary		分页获取所有商品属性分类
//	@Description	分页获取所有商品属性分类
//	@Tags			商品属性分类
//	@Accept			json
//	@Produce		json
//	@Param			pageSize	query	int		true	"pageSize"
//	@Param			pageNum		query	int		true	"pageNum"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /productAttribute/category/list [get]
func (router *PmsProductAttributeCategoryRouter) list(context *gin.Context) {
	var pageDto base_dto.PageDTO
	if err := context.ShouldBind(&pageDto); err != nil {
		pageDto = base_dto.PageDTO{PageNum: 1, PageSize: 5}
	}
	page, err := router.PmsProductAttributeCategoryService.ListPage(pageDto.PageNum, pageDto.PageSize)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, page)
}

// listWithAttr 获取所有商品属性分类及其下属性
//
//	@Summary		获取所有商品属性分类及其下属性
//	@Description	获取所有商品属性分类及其下属性
//	@Tags			商品属性分类
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router /productAttribute/category/list/withAttr [get]
func (router *PmsProductAttributeCategoryRouter) listWithAttr(context *gin.Context) {
	list, err := router.PmsProductAttributeCategoryService.ListWithAttr()
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}
