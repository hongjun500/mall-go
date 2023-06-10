package models

type UmsRoleMenuRelation struct {
	Model
	// 角色ID
	RoleID int64 `gorm:"column:role_id;not null" json:"roleId"`
	// 菜单ID
	MenuID int64 `gorm:"column:menu_id;not null" json:"menuId"`
}
