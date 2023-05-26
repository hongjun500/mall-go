package model

type UmsRoleResourceRelation struct {
	Model
	// 角色ID
	RoleID int64 `gorm:"column:role_id;not null" json:"role_id"`
	// 资源ID
	ResourceID int64 `gorm:"column:resource_id;not null" json:"resource_id"`
}
