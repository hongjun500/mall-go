package model

type UmsRoleMenuRelation struct {
	Model
	// 角色ID
	RoleID int64 `gorm:"column:role_id;not null" json:"role_id"`
	// 菜单ID
	MenuID int64 `gorm:"column:menu_id;not null" json:"menu_id"`
}
