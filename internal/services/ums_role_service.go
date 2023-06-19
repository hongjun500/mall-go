// @author hongjun500
// @date 2023/6/13 15:48
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request_dto/base"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
)

type UmsRoleService struct {
	DbFactory *database.DbFactory
}

func NewUmsRoleService(dbFactory *database.DbFactory) UmsRoleService {
	return UmsRoleService{DbFactory: dbFactory}
}

// UmsRoleCreate 添加角色
// @Description 添加角色
// @Summary 添加角色
// @Tags 后台角色管理
// @Accept application/json
// @Produce application/json
// @Param request body ums_admin.UmsRoleCreateDTO true "添加角色"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/create [post]
func (s UmsRoleService) UmsRoleCreate(context *gin.Context) {
	var dto ums_admin.UmsRoleCreateDTO
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
	rows, err := m.Insert(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// UmsRoleUpdate 修改角色
// @Description 修改角色
// @Summary 修改角色
// @Tags 后台角色管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "id"
// @Param request body ums_admin.UmsRoleCreateDTO true "修改角色"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/update/{id} [post]
func (s UmsRoleService) UmsRoleUpdate(context *gin.Context) {
	var dto ums_admin.UmsRoleCreateDTO
	var pathVariable base.PathVariableDTO
	err := context.ShouldBind(&dto)
	err = context.ShouldBindUri(&pathVariable)
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
	rows, err := m.Update(s.DbFactory.GormMySQL, pathVariable.Id)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// UmsRoleDelete 批量删除角色
// @Description 批量删除角色
// @Summary 批量删除角色
// @Tags 后台角色管理
// @Accept application/json
// @Produce application/json
// @Param ids query []int64 true "ids"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/delete [post]
func (s UmsRoleService) UmsRoleDelete(context *gin.Context) {
	var dto ums_admin.IdsDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	m := new(models.UmsRole)
	rows, err := m.Delete(s.DbFactory.GormMySQL, dto.Ids)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	s.DelResourceListByRoleIds(dto.Ids)
	gin_common.CreateSuccess(context, rows)
}

// UmsRoleListAll 获取所有角色
// @Description 获取所有角色
// @Summary 获取所有角色
// @Tags 后台角色管理
// @Accept application/json
// @Produce application/json
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/listAll [get]
func (s UmsRoleService) UmsRoleListAll(context *gin.Context) {
	var m models.UmsRole
	list, err := m.SelectAll(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// UmsRoleList 根据角色名称分页获取角色列表
// @Description 根据角色名称分页获取角色列表
// @Summary 根据角色名称分页获取角色列表
// @Tags 后台角色管理
// @Accept multipart/form-data
// @Produce application/json
// @Param keyword query string false "keyword"
// @Param pageSize query int true "pageSize"
// @Param pageNum query int true "pageNum"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/list [get]
func (s UmsRoleService) UmsRoleList(context *gin.Context) {
	var dto ums_admin.UmsRoleListPageDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	m := new(models.UmsRole)
	page, err := m.SelectPage(s.DbFactory.GormMySQL, dto.Keyword, dto.PageNum, dto.PageSize)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, page)
}

// UmsRoleUpdateStatus 修改角色状态
// @Description 修改角色状态
// @Summary 修改角色状态
// @Tags 后台角色管理
// @Accept application/json
// @Produce application/json
// @Param id path int true "id"
// @Param status query int true "status"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/updateStatus/{id} [post]
func (s UmsRoleService) UmsRoleUpdateStatus(context *gin.Context) {
	var pathVariableDTO base.PathVariableDTO
	var dto ums_admin.UmsRoleStatusDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	err = context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	var m models.UmsRole
	rows, err := m.UpdateStatus(s.DbFactory.GormMySQL, pathVariableDTO.Id, dto.Status)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// UmsRoleListMenu 获取角色相关菜单
// @Description 获取角色相关菜单
// @Summary 获取角色相关菜单
// @Tags 后台角色管理
// @Accept application/json
// @Produce application/json
// @Param roleId path int true "roleId"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/listMenu/{roleId} [get]
func (s UmsRoleService) UmsRoleListMenu(context *gin.Context) {
	var dto ums_admin.UmsRolePathVariableDTO
	err := context.ShouldBindUri(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	var m models.UmsRole
	list, err := m.SelectMenu(s.DbFactory.GormMySQL, dto.RoleId)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// UmsRoleListResource 获取角色相关资源
// @Description 获取角色相关资源
// @Summary 获取角色相关资源
// @Tags 后台角色管理
// @Accept application/json
// @Produce application/json
// @Param roleId path int true "roleId"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/listResource/{roleId} [get]
func (s UmsRoleService) UmsRoleListResource(context *gin.Context) {
	var dto ums_admin.UmsRolePathVariableDTO
	err := context.ShouldBindUri(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	var m models.UmsRole
	list, err := m.SelectResourceByRoleId(s.DbFactory.GormMySQL, dto.RoleId)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, list)
}

// UmsRoleAllocMenu 修改角色菜单
// @Description 修改角色菜单
// @Summary 修改角色菜单
// @Tags 后台角色管理
// @Accept multipart/form-data
// @Produce application/json
// @Param roleId query int true "roleId"
// @Param menuIds query []int true "menuIds"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/allocMenu/{roleId} [post]
func (s UmsRoleService) UmsRoleAllocMenu(context *gin.Context) {
	var dto ums_admin.UmsRoleAllocMenuDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	var m models.UmsRole
	rows, err := m.UpdateRoleFromAllocMenu(s.DbFactory.GormMySQL, dto.RoleId, dto.MenuIds)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// UmsRoleAllocResource 给角色分配资源
// @Description 给角色分配资源
// @Summary 给角色分配资源
// @Tags 后台角色管理
// @Accept multipart/form-data
// @Produce application/json
// @Param roleId query int true "roleId"
// @Param resourceIds query []int true "resourceIds"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /role/allocResource/{roleId} [post]
func (s UmsRoleService) UmsRoleAllocResource(context *gin.Context) {
	var dto ums_admin.UmsRoleAllocResourceDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	var m models.UmsRole
	rows, err := m.UpdateRoleFromAllocResource(s.DbFactory.GormMySQL, dto.RoleId, dto.ResourceIds)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	s.DelResourceListByRole(dto.RoleId)
	gin_common.CreateSuccess(context, rows)
}
