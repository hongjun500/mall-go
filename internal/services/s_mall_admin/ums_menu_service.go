//	@author	hongjun500
//	@date	2023/6/11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台菜单相关服务

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/pkg"
	"time"

	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
)

type UmsMenuService struct {
	DbFactory *database.DbFactory
}

func NewUmsMenuService(dbFactory *database.DbFactory) UmsMenuService {
	return UmsMenuService{DbFactory: dbFactory}
}

// UmsMenuCreate 添加后台菜单
func (s UmsMenuService) UmsMenuCreate(umsMenuCreateDTO admin_dto.UmsMenuCreateDTO) (int64, error) {
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
	rows, err := umsMenu.InsertUmsMenu(s.DbFactory.GormMySQL)
	if err != nil {
		return 0, err
	}
	return rows, nil
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
func (s UmsMenuService) UmsMenuUpdate(umsMenuCreateDTO admin_dto.UmsMenuCreateDTO) (int64, error) {
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
	err := umsMenu.Update(s.DbFactory.GormMySQL, menus)
	if err != nil {
		return 0, err
	}
	return int64(len(menus)), nil
}

// UmsMenuDelete 删除后台菜单
func (s UmsMenuService) UmsMenuDelete(dto admin_dto.UmsMenuCreateDTO) (int64, error) {
	umsMenu := new(models.UmsMenu)
	rows, err := umsMenu.Delete(s.DbFactory.GormMySQL, dto.Id)
	if err != nil {
		return 0, err
	}
	return rows, err
}

// UmsMenuItem 根据ID获取菜单详情
func (s UmsMenuService) UmsMenuItem(dto admin_dto.UmsMenuCreateDTO) (*models.UmsMenu, error) {
	umsMenu := new(models.UmsMenu)
	result, err := umsMenu.SelectById(s.DbFactory.GormMySQL, dto.Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UmsMenuPageList 分页查询后台菜单
func (s UmsMenuService) UmsMenuPageList(pageDTO base_dto.PageDTO, parentIdDTO admin_dto.UmsMenuListDTO) (*pkg.CommonPage, error) {
	umsMenu := new(models.UmsMenu)
	page, err := umsMenu.SelectPage(s.DbFactory.GormMySQL, pageDTO.PageNum, pageDTO.PageSize, parentIdDTO.ParentId)
	if err != nil {
		return nil, err
	}
	return page, nil
}

// UmsMenuUpdateHidden 修改菜单显示状态
func (s UmsMenuService) UmsMenuUpdateHidden(dto admin_dto.UmsMenuHiddenDTO) (int64, error) {
	umsMenu := new(models.UmsMenu)
	rows, err := umsMenu.UpdateHidden(s.DbFactory.GormMySQL, dto.Id, dto.Hidden)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// UmsMenuTreeList 树形结构返回所有菜单列表
func (s UmsMenuService) UmsMenuTreeList() ([]*models.UmsMenuNode, error) {
	umsMenu := new(models.UmsMenu)
	result, err := umsMenu.SelectAll(s.DbFactory.GormMySQL)
	if err != nil {
		return nil, err
	}
	// 转换树形结构
	var umsMenuNodes []*models.UmsMenuNode
	for _, menu := range result {
		if menu.ParentID == 0 {
			umsMenuNodes = append(umsMenuNodes, convertMenuTreeNode(menu, result))
		}
	}
	return umsMenuNodes, nil
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
