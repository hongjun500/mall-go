package models

import (
	"gorm.io/gorm"
)

type UmsAdminRoleRelation struct {
	Model
	AdminId int64 `gorm:"column:admin_id;not null" json:"admin_id"`
	RoleId  int64 `gorm:"column:role_id;not null" json:"role_id"`
}

func (*UmsAdminRoleRelation) TableName() string {
	return "ums_admin_role_relation"
}

func (m *UmsAdminRoleRelation) SelectAllByAdminId(db *gorm.DB, adminId int64) ([]*UmsRole, error) {
	var list []*UmsRole
	tx := db.Raw("select r.* from ums_admin_role_relation ar left join ums_role r on ar.role_id = r.id where ar.admin_id = ?", adminId).Scan(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}
