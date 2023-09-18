//	@author	hongjun500
//	@date	2023/6/13 15:47
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsRoleRouter struct {
	s_mall_admin.UmsRoleService
}

func NewUmsRoleRouter(service s_mall_admin.UmsRoleService) *UmsRoleRouter {
	return &UmsRoleRouter{UmsRoleService: service}
}

// GroupUmsRoleRouter 后台角色管理路由
func (router *UmsRoleRouter) GroupUmsRoleRouter(umsRoleGroup *gin.RouterGroup) {

	{
		// 添加角色
		umsRoleGroup.POST("/create", router.create)
		// 修改角色
		umsRoleGroup.POST("/update/:id", router.update)
		// 批量删除角色
		umsRoleGroup.POST("/delete", router.delete)
		// 获取所有角色
		umsRoleGroup.GET("/listAll", router.listAll)
		// 根据角色名称分页获取角色列表
		umsRoleGroup.GET("/list", router.listPage)
		// 修改角色状态
		umsRoleGroup.POST("/updateStatus/:id", router.updateStatus)
		// 获取角色相关菜单
		umsRoleGroup.GET("/listMenu/:roleId", router.listMenu)
		// 获取角色相关资源
		umsRoleGroup.GET("/listResource/:roleId", router.listResource)
		// 给角色分配菜单
		umsRoleGroup.POST("/allocMenu", router.allocMenu)
		// 给角色分配资源
		umsRoleGroup.POST("/allocResource", router.allocResource)
	}
}

// create 添加角色
//
//	@Description	添加角色
//	@Summary		添加角色
//	@Tags			后台角色管理
//	@Accept			application/json
//	@Produce		application/json
//	@Param			request	body	admin_dto.UmsRoleCreateDTO	true	"添加角色"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/create [post]
func (router *UmsRoleRouter) create(context *gin.Context) {
	var dto admin_dto.UmsRoleCreateDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	m := new(models.UmsRole)
	m.Name = dto.Name
	m.Description = dto.Description
	m.AdminCount = 0
	m.Status = dto.Status
	m.Sort = 0
	rows, err := m.Insert(router.UmsRoleService.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(context, gin_common.DatabaseError)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// update 修改角色
//
//	@Description	修改角色
//	@Summary		修改角色
//	@Tags			后台角色管理
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id		path	int							true	"id"
//	@Param			request	body	admin_dto.UmsRoleCreateDTO	true	"修改角色"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/update/{id} [post]
func (router *UmsRoleRouter) update(context *gin.Context) {
	var dto admin_dto.UmsRoleCreateDTO
	var pathVariable base_dto.PathVariableDTO
	err := context.ShouldBind(&dto)
	err = context.ShouldBindUri(&pathVariable)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsRoleService.UmsRoleUpdate(pathVariable.Id, dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// delete 批量删除角色
//
//	@Description	批量删除角色
//	@Summary		批量删除角色
//	@Tags			后台角色管理
//	@Accept			application/json
//	@Produce		application/json
//	@Param			ids	formData	[]int64	true	"ids"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/delete [post]
func (router *UmsRoleRouter) delete(context *gin.Context) {
	var dto admin_dto.IdsDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsRoleService.UmsRoleDelete(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// listAll 获取所有角色
//
//	@Description	获取所有角色
//	@Summary		获取所有角色
//	@Tags			后台角色管理
//	@Accept			application/json
//	@Produce		application/json
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/listAll [get]
func (router *UmsRoleRouter) listAll(context *gin.Context) {
	var m models.UmsRole
	list, err := m.SelectAll(router.UmsRoleService.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(context, gin_common.DatabaseError)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// listPage 根据角色名称分页获取角色列表
//
//	@Description	根据角色名称分页获取角色列表
//	@Summary		根据角色名称分页获取角色列表
//	@Tags			后台角色管理
//	@Accept			multipart/form-data
//	@Produce		application/json
//	@Param			keyword		query	string	false	"keyword"
//	@Param			pageSize	query	int		true	"pageSize"
//	@Param			pageNum		query	int		true	"pageNum"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/list [get]
func (router *UmsRoleRouter) listPage(context *gin.Context) {
	var dto admin_dto.UmsRoleListPageDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	page, err := router.UmsRoleService.UmsRoleList(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, page)
}

// updateStatus 修改角色状态
//
//	@Description	修改角色状态
//	@Summary		修改角色状态
//	@Tags			后台角色管理
//	@Accept			application/json
//	@Produce		application/json
//	@Param			id		path	int	true	"id"
//	@Param			status	query	int	true	"status"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/updateStatus/{id} [post]
func (router *UmsRoleRouter) updateStatus(context *gin.Context) {
	var pathVariableDTO base_dto.PathVariableDTO
	var dto admin_dto.UmsRoleStatusDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	err = context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsRoleService.UmsRoleUpdateStatus(pathVariableDTO.Id, dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// listMenu 获取角色相关菜单
//
//	@Description	获取角色相关菜单
//	@Summary		获取角色相关菜单
//	@Tags			后台角色管理
//	@Accept			application/json
//	@Produce		application/json
//	@Param			roleId	path	int	true	"roleId"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/listMenu/{roleId} [get]
func (router *UmsRoleRouter) listMenu(context *gin.Context) {
	var dto admin_dto.UmsRolePathVariableDTO
	err := context.ShouldBindUri(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	list, err := router.UmsRoleService.UmsRoleListMenu(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// listResource 获取角色相关资源
//
//	@Description	获取角色相关资源
//	@Summary		获取角色相关资源
//	@Tags			后台角色管理
//	@Accept			application/json
//	@Produce		application/json
//	@Param			roleId	path	int	true	"roleId"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/listResource/{roleId} [get]
func (router *UmsRoleRouter) listResource(context *gin.Context) {
	var dto admin_dto.UmsRolePathVariableDTO
	err := context.ShouldBindUri(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	list, err := router.UmsRoleService.UmsRoleListResource(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// allocMenu 修改角色菜单
//
//	@Description	修改角色菜单
//	@Summary		修改角色菜单
//	@Tags			后台角色管理
//	@Accept			multipart/form-data
//	@Produce		application/json
//	@Param			roleId	query	int		true	"roleId"
//	@Param			menuIds	query	[]int	true	"menuIds"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/allocMenu/{roleId} [post]
func (router *UmsRoleRouter) allocMenu(context *gin.Context) {
	var dto admin_dto.UmsRoleAllocMenuDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsRoleService.UmsRoleAllocMenu(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// allocResource 给角色分配资源
//
//	@Description	给角色分配资源
//	@Summary		给角色分配资源
//	@Tags			后台角色管理
//	@Accept			multipart/form-data
//	@Produce		application/json
//	@Param			roleId		query	int		true	"roleId"
//	@Param			resourceIds	query	[]int	true	"resourceIds"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/role/allocResource/{roleId} [post]
func (router *UmsRoleRouter) allocResource(context *gin.Context) {
	var dto admin_dto.UmsRoleAllocResourceDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	rows, err := router.UmsRoleService.UmsRoleAllocResource(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, rows)
}
