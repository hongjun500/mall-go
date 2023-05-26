package model

import "gorm.io/gorm"

type UmsRole struct {
	Model
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

func Create(db *gorm.DB, role *UmsRole) (int, error) {
	tx := db.Create(&role)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return 1, nil
}

func Update(db *gorm.DB, id int64, role *UmsRole) (int64, error) {
	role.Id = id
	tx := db.Updates(&role)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func Delete(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(&UmsRole{}, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// ListAll 获取所有角色信息
func ListAll(db *gorm.DB) (roles []*UmsRole, err error) {
	tx := db.Find(&roles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return roles, nil
}

// ListPage 根据 name 的关键字分页获取角色信息
func ListPage(db *gorm.DB, keyword string, page int, pageSize int) (roles []*UmsRole, err error) {
	tx := db.Where("name like ", "%"+keyword+"%").Offset((page - 1) * pageSize).Limit(pageSize).Find(&roles)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return roles, nil
}

// ListMenu 根据管理员 ID 获取角色的菜单
func ListMenu(db *gorm.DB, adminId int64) (menus []*UmsMenu, err error) {
	sql := "SELECT m.id id, m.parent_id parentId, m.create_time createTime, m.title title, m.level level, m.sort sort, m.name name, m.icon icon, m.hidden hidden FROM ums_admin_role_relation arr LEFT JOIN ums_role r ON arr.role_id = r.id LEFT JOIN ums_role_menu_relation rmr ON r.id = rmr.role_id LEFT JOIN ums_menu m ON rmr.menu_id = m.id WHERE arr.admin_id = ? AND m.id IS NOT NULL GROUP BY m.id"
	tx := db.Raw(sql, adminId).Scan(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return menus, nil
}

// ListMenuByRoleId 根据角色 ID 获取角色的菜单
func ListMenuByRoleId(db *gorm.DB, roleId int64) (menus []*UmsRole, err error) {
	sql := "SELECT m.id id, m.parent_id parentId, m.create_time createTime, m.title title, m.level level, m.sort sort, m.name name, m.icon icon, m.hidden hidden FROM ums_role_menu_relation rmr LEFT JOIN ums_menu m ON rmr.menu_id = m.id WHERE rmr.role_id = ? AND m.id IS NOT NULL GROUP BY m.id"
	tx := db.Raw(sql, roleId).Scan(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return menus, nil
}

// ListResourceByRoleId 根据角色 ID 获取角色的资源
func ListResourceByRoleId(db *gorm.DB, roleId int64) (resources []*UmsResource, err error) {
	sql := "SELECT r.id id, r.create_time createTime, r.`name` `name`, r.url url, r.description description, r.category_id categoryId FROM ums_role_resource_relation rrr LEFT JOIN ums_resource r ON rrr.resource_id = r.id WHERE rrr.role_id = ? AND r.id IS NOT NULL GROUP BY r.id"
	tx := db.Raw(sql, roleId).Scan(&resources)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return resources, nil
}

// AllocMenu 给角色分配菜单
func AllocMenu(db *gorm.DB, roleId int64, menuIdList []int64) (int64, error) {
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
func AllocResource(db *gorm.DB, roleId int64, resourceIdList []int64) (int64, error) {
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
