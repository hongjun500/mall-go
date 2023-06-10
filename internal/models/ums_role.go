package models

import (
	"gorm.io/gorm"
)

type UmsRole struct {
	*Model
	// 名称
	Name string `gorm:"column:name;type:varchar(100);" json:"name"`
	// 描述
	Description string `gorm:"column:description;type:varchar(500);" json:"description"`
	// 后台用户数量
	AdminCount int `gorm:"column:admin_count;type:int(11);" json:"admin_count"`
	// 0->禁用；1->启用
	Status int `gorm:"column:status;type:tinyint(1);default:1;" json:"status"`
	// 排序
	Sort int `gorm:"column:sort;type:int(11);default:0;" json:"sort"`
}

func (*UmsRole) TableName() string {
	return "ums_role"
}

func (*UmsRole) CreateAt() string {
	return "create_at"
}

func (role *UmsRole) Insert(db *gorm.DB) (int64, error) {
	tx := db.Create(role)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (role *UmsRole) Update(db *gorm.DB, id int64) (int64, error) {
	role.Id = id
	tx := db.Updates(role)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (role *UmsRole) SelectUmsRoleById(db *gorm.DB, id int64) (*UmsRole, error) {
	var umsRole UmsRole
	tx := db.First(&umsRole, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &umsRole, nil
}

func (role *UmsRole) Delete(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(role, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// SelectAll 获取所有角色信息
func (role *UmsRole) SelectAll(db *gorm.DB) ([]*UmsRole, error) {
	var roles []*UmsRole
	tx := db.Find(&roles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return roles, nil
}

// SelectPage 根据 name 的关键字分页获取角色信息
func (role *UmsRole) SelectPage(db *gorm.DB, keyword string, pageNum int, pageSize int) ([]*UmsRole, error) {
	var roles []*UmsRole
	dbQuery := db.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	if keyword != "" {
		dbQuery = dbQuery.Where("name like ?", "%"+keyword+"%")
	}
	tx := dbQuery.Find(&roles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return roles, nil
}

// SelectMenu 根据管理员 ID 获取角色的菜单
func (role *UmsRole) SelectMenu(db *gorm.DB, adminId int64) ([]*UmsMenu, error) {
	var menus []*UmsMenu
	sql := "SELECT m.id id, m.parent_id parentId, m.create_time createTime, m.title title, m.level level, m.sort sort, m.name name, m.icon icon, m.hidden hidden FROM ums_admin_role_relation arr LEFT JOIN ums_role r ON arr.role_id = r.id LEFT JOIN ums_role_menu_relation rmr ON r.id = rmr.role_id LEFT JOIN ums_menu m ON rmr.menu_id = m.id WHERE arr.admin_id = ? AND m.id IS NOT NULL GROUP BY m.id"
	tx := db.Raw(sql, adminId).Scan(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return menus, nil
}

// SelectMenuByRoleId 根据角色 ID 获取角色的菜单
func (role *UmsRole) SelectMenuByRoleId(db *gorm.DB, roleId int64) ([]*UmsRole, error) {
	var menus []*UmsRole
	sql := "SELECT m.id id, m.parent_id parentId, m.create_time createTime, m.title title, m.level level, m.sort sort, m.name name, m.icon icon, m.hidden hidden FROM ums_role_menu_relation rmr LEFT JOIN ums_menu m ON rmr.menu_id = m.id WHERE rmr.role_id = ? AND m.id IS NOT NULL GROUP BY m.id"
	tx := db.Raw(sql, roleId).Scan(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return menus, nil
}

// SelectResourceByRoleId  根据角色 ID 获取角色的资源
func (role *UmsRole) SelectResourceByRoleId(db *gorm.DB, roleId int64) ([]*UmsResource, error) {
	var resources []*UmsResource
	sql := "SELECT r.id id, r.create_time createTime, r.`name` `name`, r.url url, r.description description, r.category_id categoryId FROM ums_role_resource_relation rrr LEFT JOIN ums_resource r ON rrr.resource_id = r.id WHERE rrr.role_id = ? AND r.id IS NOT NULL GROUP BY r.id"
	tx := db.Raw(sql, roleId).Scan(&resources)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return resources, nil
}

// AllocMenu 给角色分配菜单
func (role *UmsRole) AllocMenu(db *gorm.DB, roleId int64, menuIdList []int64) (int64, error) {
	// 先删除原有的关系
	tx := db.Delete(&UmsRoleMenuRelation{}, "role_id = ?", roleId)
	if tx.Error != nil {
		return 0, tx.Error
	}
	// 再添加新的关系
	var roleMenuRelations []*UmsRoleMenuRelation
	for _, menuId := range menuIdList {
		// 追加到 roleMenuRelations 后面
		roleMenuRelations = append(roleMenuRelations, &UmsRoleMenuRelation{
			RoleID: roleId,
			MenuID: menuId,
		})
	}
	tx = db.Create(&roleMenuRelations)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// AllocResource 给角色分配资源
func (role *UmsRole) AllocResource(db *gorm.DB, roleId int64, resourceIdList []int64) (int64, error) {
	// 先删除原有的关系
	tx := db.Delete(&UmsRoleResourceRelation{}, "role_id = ?", roleId)
	if tx.Error != nil {
		return 0, tx.Error
	}
	// 再添加新的关系
	var roleResourceRelations []*UmsRoleResourceRelation
	for _, resourceId := range resourceIdList {
		// 追加到 roleResourceRelations 后面
		roleResourceRelations = append(roleResourceRelations, &UmsRoleResourceRelation{
			RoleID:     roleId,
			ResourceID: resourceId,
		})
	}
	tx = db.Create(&roleResourceRelations)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
