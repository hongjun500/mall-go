package models

import "gorm.io/gorm"

type UmsRoleResourceRelation struct {
	Model
	// 角色ID
	RoleID int64 `gorm:"column:role_id;not null" json:"roleId"`
	// 资源ID
	ResourceID int64 `gorm:"column:resource_id;not null" json:"resourceId"`
}

func (*UmsRoleResourceRelation) TableName() string {
	return "ums_role_resource_relation"
}

func (umsRoleResourceRelation *UmsRoleResourceRelation) SelectAdminIdsByResourceId(db *gorm.DB, resourceId int64) ([]int64, error) {
	var list []int64
	rows, err := db.Table("ums_role_resource_relation").Select("ums_admin_role_relation.admin_id").Distinct("ums_admin_role_relation.admin_id").Joins("left join ums_admin_role_relation on ums_role_resource_relation.role_id = ums_admin_role_relation.role_id").Where("ums_role_resource_relation.resource_id = ?", resourceId).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var adminId int64
		if err := rows.Scan(&adminId); err != nil {
			return nil, err
		}
		list = append(list, adminId)
	}
	return list, nil
}
