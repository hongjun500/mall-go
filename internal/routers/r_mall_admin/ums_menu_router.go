//	@author	hongjun500
//	@date	2023/6/11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台菜单管理路由

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/request/ums_admin_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsMenuRouter struct {
	s_mall_admin.UmsMenuService
}

func NewUmsMenuRouter(service s_mall_admin.UmsMenuService) *UmsMenuRouter {
	return &UmsMenuRouter{UmsMenuService: service}
}

// GroupUmsMenuRouter 后台菜单路由
func (router *UmsMenuRouter) GroupUmsMenuRouter(umsMenuGroup *gin.RouterGroup) {
	{
		// 新增菜单
		umsMenuGroup.POST("/create", router.create)
		// 修改菜单
		umsMenuGroup.POST("/update/:id", router.update)
		// 删除菜单
		umsMenuGroup.POST("/delete/:id", router.delete)
		// 根据ID获取菜单详情
		umsMenuGroup.GET("/:id", router.detail)
		// 分页获取菜单列表
		umsMenuGroup.GET("/list/:parentId", router.listPage)
		// 修改菜单显示状态
		umsMenuGroup.POST("/updateHidden/:id", router.updateHidden)
		// 树形结构返回所有菜单列表
		umsMenuGroup.GET("/treeList", router.treeList)
	}
}

// create 添加后台菜单
//
//	@Summary		添加后台菜单
//	@Description	添加后台菜单
//	@Tags			后台菜单管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body	ums_admin_dto.UmsMenuCreateDTO	true	"添加后台菜单"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/menu/create [post]
func (router *UmsMenuRouter) create(context *gin.Context) {
	var umsMenuCreateDTO ums_admin_dto.UmsMenuCreateDTO
	err := context.ShouldBind(&umsMenuCreateDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsMenuService.UmsMenuCreate(umsMenuCreateDTO)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// update 修改后台菜单
//
//	@Summary		修改后台菜单
//	@Description	修改后台菜单
//	@Tags			后台菜单管理
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int64						true	"菜单ID"
//	@Param			request	body	ums_admin_dto.UmsMenuCreateDTO	true	"修改后台菜单"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/menu/update/{id} [post]
func (router *UmsMenuRouter) update(context *gin.Context) {
	var umsMenuCreateDTO ums_admin_dto.UmsMenuCreateDTO
	err := context.ShouldBind(&umsMenuCreateDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsMenuService.UmsMenuUpdate(umsMenuCreateDTO)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// delete 删除后台菜单
//
//	@Description	删除后台菜单
//	@Summary		删除后台菜单
//	@Tags			后台菜单管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int64	true	"菜单ID"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/menu/delete/{id} [post]
func (router *UmsMenuRouter) delete(context *gin.Context) {
	var dto ums_admin_dto.UmsMenuCreateDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsMenuService.UmsMenuDelete(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// detail 根据ID获取菜单详情
//
//	@Summary		根据ID获取菜单详情
//	@Description	根据ID获取菜单详情
//	@Tags			后台菜单管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int64	true	"菜单ID"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/menu/{id} [get]
func (router *UmsMenuRouter) detail(context *gin.Context) {
	var dto ums_admin_dto.UmsMenuCreateDTO
	err := context.ShouldBindUri(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	umsMenu, err := router.UmsMenuService.UmsMenuItem(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, umsMenu)
}

// listPage 分页查询后台菜单
//
//	@Summary		分页查询后台菜单
//	@Description	分页查询后台菜单
//	@Tags			后台菜单管理
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			parentId	path	int64	true	"父级菜单ID"
//	@Param			pageNum		query	int64	true	"页码"
//	@Param			pageSize	query	int64	true	"每页数量"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/menu/list/{parentId} [get]
func (router *UmsMenuRouter) listPage(context *gin.Context) {
	//	@Param	pageNum		formData	int64	true	"页码"
	//	@Param	pageSize	formData	int64	true	"每页数量"
	// 这会导致swagger文档中的参数不正确
	var pageDTO base_dto.PageDTO
	var parentIdDTO ums_admin_dto.UmsMenuListDTO
	err := context.ShouldBind(&pageDTO)
	err = context.ShouldBindUri(&parentIdDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	page, err := router.UmsMenuService.UmsMenuPageList(pageDTO, parentIdDTO)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, page)
}

// updateHidden 修改菜单显示状态
//
//	@Summary		修改菜单显示状态
//	@Description	修改菜单显示状态
//	@Tags			后台菜单管理
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int64	true	"菜单ID"
//	@Param			hidden	formData	int64	true	"是否隐藏"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/menu/updateHidden/{id} [post]
func (router *UmsMenuRouter) updateHidden(context *gin.Context) {
	var dto ums_admin_dto.UmsMenuHiddenDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsMenuService.UmsMenuUpdateHidden(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// treeList 树形结构返回所有菜单列表
//
//	@Description	树形结构返回所有菜单列表
//	@Summary		树形结构返回所有菜单列表
//	@Tags			后台菜单管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/menu/treeList [get]
func (router *UmsMenuRouter) treeList(context *gin.Context) {
	list, err := router.UmsMenuService.UmsMenuTreeList()
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}
