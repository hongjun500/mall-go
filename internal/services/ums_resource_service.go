// @author hongjun500
// @date 2023/6/13 11:14
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// @Description:

package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request_dto/base"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
)

type UmsResourceService struct {
	DbFactory *database.DbFactory
}

func NewUmsResourceService(dbFactory *database.DbFactory) UmsResourceService {
	return UmsResourceService{DbFactory: dbFactory}
}

// UmsResourceCreate 添加后台资源
// @Description 添加后台资源
// @Summary 添加后台资源
// @Tags 后台资源管理
// @Accept  json
// @Produce  json
// @Param   request body    ums_admin.UmsResourceCreateDTO   true "添加后台资源"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resource/create [post]
func (s UmsResourceService) UmsResourceCreate(context *gin.Context) {
	var dto ums_admin.UmsResourceCreateDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	// var umsResource *models.UmsResource
	// 上面这种直接赋值有问题，空指针, 如果是以 umsResource := &models.UmsResource{} 的方式就不会有问题, 或者使用 new() 函数
	umsResource := new(models.UmsResource)
	umsResource.Name = dto.Name
	umsResource.Url = dto.Url
	umsResource.Description = dto.Description
	umsResource.CategoryId = dto.CategoryId
	rows, err := umsResource.Insert(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, rows)
}

// UmsResourceUpdate 修改后台资源
// @Description 修改后台资源
// @Summary 修改后台资源
// @Tags 后台资源管理
// @Accept  json
// @Produce  json
// @Param id path int true "资源id"
// @Param   request body    ums_admin.UmsResourceCreateDTO   true "修改后台资源"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resource/update/{id} [post]
func (s UmsResourceService) UmsResourceUpdate(context *gin.Context) {
	var dto ums_admin.UmsResourceCreateDTO
	var pathVariableDTO base.PathVariableDTO
	err := context.ShouldBind(&dto)
	err = context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	m := new(models.UmsResource)
	m.Name = dto.Name
	m.Url = dto.Url
	m.Description = dto.Description
	m.CategoryId = dto.CategoryId
	rows, err := m.Update(s.DbFactory.GormMySQL, pathVariableDTO.Id)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	s.DelResourceListByResource(pathVariableDTO.Id)
	gin_common.CreateSuccess(context, rows)
}

// UmsResourceItem 根据ID获取资源详情
// @Description 根据ID获取资源详情
// @Summary 根据ID获取资源详情
// @Tags 后台资源管理
// @Accept  json
// @Produce  json
// @Param id path int true "资源id"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resource/{id} [get]
func (s UmsResourceService) UmsResourceItem(context *gin.Context) {
	var pathVariableDTO base.PathVariableDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	m := new(models.UmsResource)
	m.Id = pathVariableDTO.Id
	umsResource, err := m.SelectUmsResourceById(s.DbFactory.GormMySQL, m.Id)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, umsResource)
}

// UmsResourceDelete 根据ID删除后台资源
// @Description 根据ID删除后台资源
// @Summary 根据ID删除后台资源
// @Tags 后台资源管理
// @Accept  json
// @Produce  json
// @Param id path int true "资源id"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resource/delete/{id} [post]
func (s UmsResourceService) UmsResourceDelete(context *gin.Context) {
	var pathVariableDTO base.PathVariableDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	m := new(models.UmsResource)
	m.Id = pathVariableDTO.Id
	rows, err := m.Delete(s.DbFactory.GormMySQL, m.Id)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	s.DelResourceListByResource(pathVariableDTO.Id)
	gin_common.CreateSuccess(context, rows)
}

// UmsResourcePageList 分页模糊查询后台资源
// @Description 分页模糊查询后台资源
// @Summary 分页模糊查询后台资源
// @Tags 后台资源管理
// @Accept  multipart/form-data
// @Produce  json
// @Param   categoryId query int64 false "资源分类ID"
// @Param   nameKeyword query string false "资源名称"
// @Param   urlKeyword query string false "资源URL"
// @Param   pageNum query int64  true "页码"
// @Param   pageSize query int64 true "每页数量"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resource/list/resource [get]
func (s UmsResourceService) UmsResourcePageList(context *gin.Context) {
	var dto ums_admin.UmsResourcePageListDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	m := new(models.UmsResource)
	page, err := m.SelectPage(s.DbFactory.GormMySQL, dto.CategoryId, dto.NameKeyword, dto.UrlKeyword, dto.PageNum, dto.PageSize)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, page)
}

// UmsResourceList 查询所有后台资源
// @Description 查询所有后台资源
// @Summary 查询所有后台资源
// @Tags 后台资源管理
// @Accept  json
// @Produce  json
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resource/listAll [get]
func (s UmsResourceService) UmsResourceList(context *gin.Context) {
	m := new(models.UmsResource)
	list, err := m.SelectAll(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, list)
}
