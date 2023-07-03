//	@author	hongjun500
//	@date	2023/6/13 11:10
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/request/ums_admin_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsResourceRouter struct {
	s_mall_admin.UmsResourceService
}

func NewUmsResourceRouter(service s_mall_admin.UmsResourceService) *UmsResourceRouter {
	return &UmsResourceRouter{UmsResourceService: service}
}

// GroupUmsResourceRouter 后台资源管理路由
func (router *UmsResourceRouter) GroupUmsResourceRouter(umsResourceGroup *gin.RouterGroup) {

	{
		// 添加后台资源
		umsResourceGroup.POST("/create", router.create)
		// 修改后台资源
		umsResourceGroup.POST("/update/:id", router.update)
		// 根据ID获取资源详情
		umsResourceGroup.GET("/:id", router.detail)
		// 根据ID删除后台资源
		umsResourceGroup.POST("/delete/:id", router.delete)
		// 分页模糊查询后台资源
		umsResourceGroup.GET("/list", router.listPage)
		// 查询所有后台资源
		umsResourceGroup.GET("/listAll", router.listAll)

	}
}

// create 添加后台资源
//
//	@Description	添加后台资源
//	@Summary		添加后台资源
//	@Tags			后台资源管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body	ums_admin_dto.UmsResourceCreateDTO	true	"添加后台资源"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resource/create [post]
func (router *UmsResourceRouter) create(context *gin.Context) {
	var dto ums_admin_dto.UmsResourceCreateDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsResourceCreate(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// update 修改后台资源
//
//	@Description	修改后台资源
//	@Summary		修改后台资源
//	@Tags			后台资源管理
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int								true	"资源id"
//	@Param			request	body	ums_admin_dto.UmsResourceCreateDTO	true	"修改后台资源"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resource/update/{id} [post]
func (router *UmsResourceRouter) update(context *gin.Context) {
	var dto ums_admin_dto.UmsResourceCreateDTO
	var pathVariableDTO base_dto.PathVariableDTO
	err := context.ShouldBind(&dto)
	err = context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsResourceUpdate(pathVariableDTO.Id, dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// detail 根据ID获取资源详情
//
//	@Description	根据ID获取资源详情
//	@Summary		根据ID获取资源详情
//	@Tags			后台资源管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"资源id"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resource/{id} [get]
func (router *UmsResourceRouter) detail(context *gin.Context) {
	var pathVariableDTO base_dto.PathVariableDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	umsResource, err := router.UmsResourceItem(pathVariableDTO.Id)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, umsResource)
}

// delete 根据ID删除后台资源
//
//	@Description	根据ID删除后台资源
//	@Summary		根据ID删除后台资源
//	@Tags			后台资源管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"资源id"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resource/delete/{id} [post]
func (router *UmsResourceRouter) delete(context *gin.Context) {
	var pathVariableDTO base_dto.PathVariableDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsResourceDelete(pathVariableDTO.Id)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// listPage 分页模糊查询后台资源
//
//	@Description	分页模糊查询后台资源
//	@Summary		分页模糊查询后台资源
//	@Tags			后台资源管理
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			categoryId	query	int64	false	"资源分类ID"
//	@Param			nameKeyword	query	string	false	"资源名称"
//	@Param			urlKeyword	query	string	false	"资源URL"
//	@Param			pageNum		query	int64	true	"页码"
//	@Param			pageSize	query	int64	true	"每页数量"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resource/list [get]
func (router *UmsResourceRouter) listPage(context *gin.Context) {
	var dto ums_admin_dto.UmsResourcePageListDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	pageInfo, err := router.UmsResourcePageList(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, pageInfo)
}

// listAll 查询所有后台资源
//
//	@Description	查询所有后台资源
//	@Summary		查询所有后台资源
//	@Tags			后台资源管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/resource/listAll [get]
func (router *UmsResourceRouter) listAll(context *gin.Context) {
	m := new(models.UmsResource)
	list, err := m.SelectAll(router.UmsResourceService.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(context, gin_common.DatabaseError)
		return
	}
	gin_common.CreateSuccess(context, list)
}
