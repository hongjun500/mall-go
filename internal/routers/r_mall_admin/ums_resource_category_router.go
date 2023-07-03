//	@author	hongjun500
//	@date	2023/6/13 10:43
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台资源分类管理路由

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/request/ums_admin_dto"
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
		umsResourceCategoryGroup.POST("/create", router.create)
		// 修改后台资源分类
		umsResourceCategoryGroup.POST("/update/:id", router.update)
		// 根据ID删除后台资源分类
		umsResourceCategoryGroup.POST("/delete/:id", router.delete)
		// 查询所有后台资源分类
		umsResourceCategoryGroup.GET("/listAll", router.listAll)
	}
}

// create 添加后台资源分类
//
//	@Description	添加后台资源分类
//	@Summary		添加后台资源分类
//	@Tags			后台资源分类管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body	ums_admin_dto.UmsResourceCategoryCreateDTO	true	"添加后台资源分类"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resourceCategory/create [post]
func (router *UmsResourceCategoryRouter) create(context *gin.Context) {
	var dto ums_admin_dto.UmsResourceCategoryCreateDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsResourceCategoryCreate(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// update 修改后台资源分类
//
//	@Description	修改后台资源分类
//	@Summary		修改后台资源分类
//	@Tags			后台资源分类管理
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int										true	"id"
//	@Param			request	body	ums_admin_dto.UmsResourceCategoryCreateDTO	true	"修改后台资源分类"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resourceCategory/update/{id} [post]
func (router *UmsResourceCategoryRouter) update(context *gin.Context) {
	var dto ums_admin_dto.UmsResourceCategoryCreateDTO
	var pathVariableDTO base_dto.PathVariableDTO
	err := context.ShouldBind(&dto)
	err = context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsResourceCategoryUpdate(pathVariableDTO, dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// delete 删除后台资源分类
//
//	@Description	删除后台资源分类
//	@Summary		删除后台资源分类
//	@Tags			后台资源分类管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"id"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resourceCategory/delete/{id} [post]
func (router *UmsResourceCategoryRouter) delete(context *gin.Context) {
	var pathVariableDTO base_dto.PathVariableDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsResourceCategoryDelete(pathVariableDTO.Id)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// listAll 查询所有后台资源分类
//
//	@Description	查询所有后台资源分类
//	@Summary		查询所有后台资源分类
//	@Tags			后台资源分类管理
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resourceCategory/listAll [get]
func (router *UmsResourceCategoryRouter) listAll(context *gin.Context) {
	list, err := router.UmsResourceCategoryService.UmsResourceCategoryList()
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}
