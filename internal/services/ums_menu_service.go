// @author hongjun500
// @date 2023/6/11
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台菜单相关服务

package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request_dto/base"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
	"time"
)

type UmsMenuService struct {
	DbFactory *database.DbFactory
}

func NewUmsMenuService(dbFactory *database.DbFactory) UmsMenuService {
	return UmsMenuService{DbFactory: dbFactory}
}

// UmsMenuCreate 添加后台菜单
// @Description 添加后台菜单
// @Summary 添加后台菜单
// @Description 添加后台菜单
// @Tags 后台菜单管理
// @Accept  json
// @Produce  json
// @Param   request body    ums_admin.UmsMenuCreateDTO   true "添加后台菜单"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /menu/create [post]
func (s UmsMenuService) UmsMenuCreate(context *gin.Context) {
	var umsMenuCreateDTO ums_admin.UmsMenuCreateDTO
	err := context.ShouldBind(&umsMenuCreateDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	umsMenu := new(models.UmsMenu)
	umsMenu.ParentID = umsMenuCreateDTO.ParentId
	now := time.Now()
	umsMenu.CreateTime = &now
	umsMenu.Title = umsMenuCreateDTO.Title
	umsMenu.Level = int64(umsMenuCreateDTO.Level)
	umsMenu.Sort = int64(umsMenuCreateDTO.Sort)
	umsMenu.Name = umsMenuCreateDTO.Name
	umsMenu.Icon = umsMenuCreateDTO.Icon
	umsMenu.Hidden = int64(umsMenuCreateDTO.Hidden)
	// 计算层级
	updateLevel(umsMenu, s)
	menu, err := umsMenu.InsertUmsMenu(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(menu, context)
}

// updateLevel 更新菜单层级
func updateLevel(umsMenu *models.UmsMenu, s UmsMenuService) {
	if umsMenu.ParentID == 0 {
		umsMenu.Level = 0
	} else {
		// 有父级菜单时根据父级菜单 level 设置
		parentMenu, _ := umsMenu.SelectById(s.DbFactory.GormMySQL, umsMenu.ParentID)
		if parentMenu != nil {
			umsMenu.Level = parentMenu.Level + 1
		} else {
			umsMenu.Level = 0
		}
	}
}

// UmsMenuUpdate 修改后台菜单
// @Description 修改后台菜单
// @Summary 修改后台菜单
// @Description 修改后台菜单
// @Tags 后台菜单管理
// @Accept  json
// @Produce  json
// @Param   id path int64 true "菜单ID"
// @Param   request body    ums_admin.UmsMenuCreateDTO   true "修改后台菜单"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /menu/update/{id} [post]
func (s UmsMenuService) UmsMenuUpdate(context *gin.Context) {
	var umsMenuCreateDTO ums_admin.UmsMenuCreateDTO
	err := context.ShouldBind(&umsMenuCreateDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	umsMenu := new(models.UmsMenu)
	umsMenu.Id = umsMenuCreateDTO.Id
	umsMenu.ParentID = umsMenuCreateDTO.ParentId
	umsMenu.Title = umsMenuCreateDTO.Title
	umsMenu.Level = int64(umsMenuCreateDTO.Level)
	umsMenu.Sort = int64(umsMenuCreateDTO.Sort)
	umsMenu.Name = umsMenuCreateDTO.Name
	umsMenu.Icon = umsMenuCreateDTO.Icon
	umsMenu.Hidden = int64(umsMenuCreateDTO.Hidden)
	// 计算层级
	updateLevel(umsMenu, s)
	var menus []*models.UmsMenu
	menus = append(menus, umsMenu)
	err = umsMenu.Update(s.DbFactory.GormMySQL, menus)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(len(menus), context)
}

// UmsMenuDelete 删除后台菜单
// @Description 删除后台菜单
// @Summary 删除后台菜单
// @Description 删除后台菜单
// @Tags 后台菜单管理
// @Accept  json
// @Produce  json
// @Param   id path int64 true "菜单ID"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /menu/delete/{id} [post]
func (s UmsMenuService) UmsMenuDelete(context *gin.Context) {
	var dto ums_admin.UmsMenuCreateDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	umsMenu := new(models.UmsMenu)

	result, err := umsMenu.Delete(s.DbFactory.GormMySQL, dto.Id)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(result, context)
}

// UmsMenuItem 根据ID获取菜单详情
// @Description 根据ID获取菜单详情
// @Summary 根据ID获取菜单详情
// @Description 根据ID获取菜单详情
// @Tags 后台菜单管理
// @Accept  json
// @Produce  json
// @Param   id path int64 true "菜单ID"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /menu/{id} [get]
func (s UmsMenuService) UmsMenuItem(context *gin.Context) {
	var dto ums_admin.UmsMenuCreateDTO
	err := context.ShouldBindUri(&dto)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	umsMenu := new(models.UmsMenu)
	result, err := umsMenu.SelectById(s.DbFactory.GormMySQL, dto.Id)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(result, context)
}

// UmsMenuPageList 分页查询后台菜单
// @Summary 分页查询后台菜单
// @Description 分页查询后台菜单
// @Tags 后台菜单管理
// @Accept  multipart/form-data
// @Produce  json
// @Param   parentId path int64 true "父级菜单ID"
// @Param   pageNum formData int64 true "页码"
// @Param   pageSize formData int64 true "每页数量"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /menu/list/{parentId} [get]
func (s UmsMenuService) UmsMenuPageList(context *gin.Context) {
	// @Param   pageNum formData int64 true "页码"
	// @Param   pageSize formData int64 true "每页数量"
	// 这会导致swagger文档中的参数不正确

	var pageDTO base.PageDTO
	var parentIdDTO ums_admin.UmsMenuListDTO
	err := context.ShouldBind(&pageDTO)
	err = context.ShouldBindUri(&parentIdDTO)
	// parentId, _ := strconv.ParseInt(context.Param("parentId"), 10, 64)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	umsMenu := new(models.UmsMenu)

	page, err := umsMenu.SelectPage(s.DbFactory.GormMySQL, pageDTO.PageNum, pageDTO.PageSize, parentIdDTO.ParentId)

	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(page, context)
}

// UmsMenuUpdateHidden 修改菜单显示状态
// @Description 修改菜单显示状态
// @Summary 修改菜单显示状态
// @Description 修改菜单显示状态
// @Tags 后台菜单管理
// @Accept  json
// @Produce  json
// @Param   id path int64 true "菜单ID"
// @Param   hidden formData int64 true "是否隐藏"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /menu/updateHidden/{id} [post]
func (s UmsMenuService) UmsMenuUpdateHidden(context *gin.Context) {
	var dto ums_admin.UmsMenuHiddenDTO
	err := context.ShouldBind(&dto)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	umsMenu := new(models.UmsMenu)
	result, err := umsMenu.UpdateHidden(s.DbFactory.GormMySQL, dto.Id, dto.Hidden)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(result, context)
}

// UmsMenuTreeList 树形结构返回所有菜单列表
// @Description 树形结构返回所有菜单列表
// @Summary 树形结构返回所有菜单列表
// @Description 树形结构返回所有菜单列表
// @Tags 后台菜单管理
// @Accept  json
// @Produce  json
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /menu/treeList [get]
func (s UmsMenuService) UmsMenuTreeList(context *gin.Context) {
	umsMenu := new(models.UmsMenu)
	result, err := umsMenu.SelectAll(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	// 转换树形结构
	var umsMenuNodes []*models.UmsMenuNode
	for _, menu := range result {
		if menu.ParentID == 0 {
			umsMenuNodes = append(umsMenuNodes, convertMenuTreeNode(menu, result))
		}
	}
	gin_common.CreateSuccess(umsMenuNodes, context)
}

// ConvertMenuTreeNode 转换菜单树形结构
func convertMenuTreeNode(menu *models.UmsMenu, menus []*models.UmsMenu) *models.UmsMenuNode {
	var umsMenuNode = &models.UmsMenuNode{}
	var umsMenuNodeChildren []*models.UmsMenuNode
	umsMenuNode.UmsMenu = *menu
	for _, umsMenu := range menus {
		if umsMenu.ParentID == menu.Id {
			umsMenuNodeChildren = append(umsMenuNodeChildren, convertMenuTreeNode(umsMenu, menus))
		}
	}
	umsMenuNode.Children = umsMenuNodeChildren
	return umsMenuNode
}
