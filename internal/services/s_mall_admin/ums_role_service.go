//	@author	hongjun500
//	@date	2023/6/13 15:48
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
	"github.com/hongjun500/mall-go/internal/services"
	"github.com/hongjun500/mall-go/pkg"
)

type UmsRoleService struct {
	DbFactory *database.DbFactory
}

func NewUmsRoleService(dbFactory *database.DbFactory) UmsRoleService {
	return UmsRoleService{DbFactory: dbFactory}
}

// UmsRoleUpdate 修改角色
func (s UmsRoleService) UmsRoleUpdate(id int64, dto admin_dto.UmsRoleCreateDTO) (int64, error) {
	m := new(models.UmsRole)
	m.Name = dto.Name
	m.Description = dto.Description
	m.AdminCount = 0
	m.Status = dto.Status
	m.Sort = 0
	rows, err := m.Update(s.DbFactory.GormMySQL, id)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// UmsRoleDelete 批量删除角色
func (s UmsRoleService) UmsRoleDelete(dto admin_dto.IdsDTO) (int64, error) {
	m := new(models.UmsRole)
	rows, err := m.Delete(s.DbFactory.GormMySQL, dto.Ids)
	if err != nil {
		return 0, err
	}
	services.DelResourceListByRoleIds(s.DbFactory, dto.Ids)
	return rows, nil
}

// UmsRoleList 根据角色名称分页获取角色列表
func (s UmsRoleService) UmsRoleList(dto admin_dto.UmsRoleListPageDTO) (*pkg.CommonPage, error) {
	m := new(models.UmsRole)
	page, err := m.SelectPage(s.DbFactory.GormMySQL, dto.Keyword, dto.PageNum, dto.PageSize)
	if err != nil {
		return nil, err
	}
	return page, nil
}

// UmsRoleUpdateStatus 修改角色状态
func (s UmsRoleService) UmsRoleUpdateStatus(id int64, dto admin_dto.UmsRoleStatusDTO) (int64, error) {
	var m models.UmsRole
	rows, err := m.UpdateStatus(s.DbFactory.GormMySQL, id, dto.Status)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// UmsRoleListMenu 获取角色相关菜单
func (s UmsRoleService) UmsRoleListMenu(dto admin_dto.UmsRolePathVariableDTO) ([]*models.UmsMenu, error) {
	var m models.UmsRole
	list, err := m.SelectMenu(s.DbFactory.GormMySQL, dto.RoleId)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// UmsRoleListResource 获取角色相关资源
func (s UmsRoleService) UmsRoleListResource(dto admin_dto.UmsRolePathVariableDTO) ([]*models.UmsResource, error) {
	var m models.UmsRole
	list, err := m.SelectResourceByRoleId(s.DbFactory.GormMySQL, dto.RoleId)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// UmsRoleAllocMenu 修改角色菜单
func (s UmsRoleService) UmsRoleAllocMenu(dto admin_dto.UmsRoleAllocMenuDTO) (int64, error) {
	var m models.UmsRole
	rows, err := m.UpdateRoleFromAllocMenu(s.DbFactory.GormMySQL, dto.RoleId, dto.MenuIds)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// UmsRoleAllocResource 给角色分配资源
func (s UmsRoleService) UmsRoleAllocResource(dto admin_dto.UmsRoleAllocResourceDTO) (int64, error) {
	var m models.UmsRole
	rows, err := m.UpdateRoleFromAllocResource(s.DbFactory.GormMySQL, dto.RoleId, dto.ResourceIds)
	if err != nil {
		return 0, err
	}
	services.DelResourceListByRole(s.DbFactory, dto.RoleId)
	return rows, nil
}
