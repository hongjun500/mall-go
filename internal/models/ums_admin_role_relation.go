package models

import (
	"gorm.io/gorm"
)

type UmsAdminRoleRelation struct {
	Model
	AdminId int64 `gorm:"column:admin_id;not null" json:"adminId"`
	RoleId  int64 `gorm:"column:role_id;not null" json:"roleId"`
}

func (*UmsAdminRoleRelation) TableName() string {
	return "ums_admin_role_relation"
}

func (re *UmsAdminRoleRelation) SelectAllByAdminId(db *gorm.DB, adminId int64) ([]*UmsRole, error) {
	var list []*UmsRole
	tx := db.Raw("select r.* from ums_admin_role_relation ar left join ums_role r on ar.role_id = r.id where ar.admin_id = ?", adminId).Scan(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}

func (re *UmsAdminRoleRelation) SelectUmsAdminRoleRelationByRoleId(db *gorm.DB, roleId int64) ([]*UmsAdminRoleRelation, error) {
	var list []*UmsAdminRoleRelation
	tx := db.Where("role_id = ?", roleId).Find(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}

func (re *UmsAdminRoleRelation) SelectUmsAdminRoleRelationInRoleId(db *gorm.DB, roleId []int64) ([]*UmsAdminRoleRelation, error) {
	var list []*UmsAdminRoleRelation
	tx := db.Where("role_id in ?", roleId).Find(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}

func (re *UmsAdminRoleRelation) DelByAdminId(db *gorm.DB, adminId int64) {
	db.Where("admin_id = ?", adminId).Delete(&UmsAdminRoleRelation{})
}

// InsertList 批量插入用户角色关系
func (re *UmsAdminRoleRelation) InsertList(db *gorm.DB, list []*UmsAdminRoleRelation) (int64, error) {
	tx := db.Create(&list)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// SelectRoleList 获取用于所有角色
func (re *UmsAdminRoleRelation) SelectRoleList(db *gorm.DB, adminId int64) ([]*UmsRole, error) {
	var list []*UmsRole
	tx := db.Select("r.*").Table("ums_admin_role_relation ar").Joins("left join ums_role r on ar.role_id = r.id").Where("ar.admin_id = ?", adminId).Scan(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}

func (re *UmsAdminRoleRelation) SelectRoleResourceRelationsByAdminId(db *gorm.DB, adminId int64) []UmsResource {
	// sql := "SELECT ur.id id, ur.create_time createTime, ur.`name` `name`, ur.url url, ur.description description, ur.category_id categoryId FROM ums_admin_role_relation ar LEFT JOIN ums_role r ON ar.role_id = r.id LEFT JOIN ums_role_resource_relation rrr ON r.id = rrr.role_id LEFT JOIN ums_resource ur ON ur.id = rrr.resource_id WHERE ar.admin_id = ? AND ur.id IS NOT NULL GROUP BY ur.id"
	var umsResources []UmsResource
	err := db.Table("ums_admin_role_relation").
		Select("ur.id, ur.name, ur.url").
		Joins("LEFT JOIN ums_role r ON ums_admin_role_relation.role_id = r.id").
		Joins("LEFT JOIN ums_role_resource_relation rrr ON r.id = rrr.role_id").
		Joins("LEFT JOIN ums_resource ur ON ur.id = rrr.resource_id").
		Where("ums_admin_role_relation.admin_id = ? AND ur.id IS NOT NULL", adminId).
		Group("ur.id").
		Find(&umsResources).Error
	if err != nil {
		return nil
	}
	return umsResources
}
