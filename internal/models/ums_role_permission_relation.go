package models

type UmsRolePermissionRelation struct {
	Model
	RoleId       int64 `gorm:"column:role_id;not null" json:"roleId"`
	PermissionId int64 `gorm:"column:permission_id;not null" json:"permissionId"`
}
