// @author hongjun500
// @date 2023/6/13 10:47
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description: 后台资源分类服务

package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request_dto/base"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
)

type UmsResourceCategoryService struct {
	DbFactory *database.DbFactory
}

func NewUmsResourceCategoryService(dbFactory *database.DbFactory) UmsResourceCategoryService {
	return UmsResourceCategoryService{DbFactory: dbFactory}
}

// UmsResourceCategoryList 查询所有后台资源分类
// @Description 查询所有后台资源分类
// @Summary 查询所有后台资源分类
// @Tags 后台资源分类管理
// @Accept  multipart/form-data
// @Produce  json
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resourceCategory/listAll [get]
func (s UmsResourceCategoryService) UmsResourceCategoryList(context *gin.Context) {
	var umsResourceCategory models.UmsResourceCategory
	list, err := umsResourceCategory.SelectAll(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	gin_common.CreateSuccess(list, context)
}

// UmsResourceCategoryCreate 添加后台资源分类
// @Description 添加后台资源分类
// @Summary 添加后台资源分类
// @Tags 后台资源分类管理
// @Accept  json
// @Produce  json
// @Param   request body    ums_admin.UmsResourceCategoryCreateDTO   true "添加后台资源分类"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resourceCategory/create [post]
func (s UmsResourceCategoryService) UmsResourceCategoryCreate(context *gin.Context) {
	var dto ums_admin.UmsResourceCategoryCreateDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	umsResourceCategory := new(models.UmsResourceCategory)
	umsResourceCategory.Name = dto.Name
	umsResourceCategory.Sort = dto.Sort
	rows, err := umsResourceCategory.Insert(s.DbFactory.GormMySQL)
	if err != nil {
		return
	}
	gin_common.CreateSuccess(rows, context)
}

// UmsResourceCategoryUpdate 修改后台资源分类
// @Description 修改后台资源分类
// @Summary 修改后台资源分类
// @Tags 后台资源分类管理
// @Accept  json
// @Produce  json
// @Param   id path int true "id"
// @Param   request body    ums_admin.UmsResourceCategoryCreateDTO   true "修改后台资源分类"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resourceCategory/update/{id} [post]
func (s UmsResourceCategoryService) UmsResourceCategoryUpdate(context *gin.Context) {
	var dto ums_admin.UmsResourceCategoryCreateDTO
	var pathVariableDTO base.PathVariableDTO
	err := context.ShouldBind(&dto)
	err = context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	umsResourceCategory := new(models.UmsResourceCategory)
	umsResourceCategory.Name = dto.Name
	umsResourceCategory.Sort = dto.Sort
	rows, err := umsResourceCategory.Update(s.DbFactory.GormMySQL, pathVariableDTO.Id)
	if err != nil {
		return
	}
	gin_common.CreateSuccess(rows, context)
}

// UmsResourceCategoryDelete 删除后台资源分类
// @Description 删除后台资源分类
// @Summary 删除后台资源分类
// @Tags 后台资源分类管理
// @Accept  json
// @Produce  json
// @Param   id path int true "id"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /resourceCategory/delete/{id} [post]
func (s UmsResourceCategoryService) UmsResourceCategoryDelete(context *gin.Context) {
	var pathVariableDTO base.PathVariableDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	umsResourceCategory := new(models.UmsResourceCategory)
	rows, err := umsResourceCategory.Delete(s.DbFactory.GormMySQL, pathVariableDTO.Id)
	if err != nil {
		return
	}
	gin_common.CreateSuccess(rows, context)
}
