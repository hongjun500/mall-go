package models

import (
	"github.com/hongjun500/mall-go/internal/gorm_common"
	"gorm.io/gorm"
)

type UmsResource struct {
	Model
	// 资源名称
	Name string `gorm:"column:name;not null" json:"name"`
	// 资源URL
	Url string `gorm:"column:url;not null" json:"url"`
	// 描述
	Description string `gorm:"column:description;" json:"description"`
	// 资源分类ID
	CategoryId int64 `gorm:"column:category_id;not null" json:"categoryId"`
	// 对应的拥有该资源的角色，在数据库忽略该字段
	RoleId int64 `gorm:"-" json:"roleId"`
}

func (*UmsResource) TableName() string {
	return "ums_resource"
}

func (usmResource *UmsResource) SelectUmsResourceById(db *gorm.DB, id int64) (*UmsResource, error) {
	tx := db.First(&usmResource, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return usmResource, nil
}

func (usmResource *UmsResource) SelectAll(db *gorm.DB) ([]*UmsResource, error) {
	var umsResources []*UmsResource
	tx := db.Find(&umsResources)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return umsResources, nil
}

func (usmResource *UmsResource) SelectPage(db *gorm.DB, categoryId int64, NameKeyword string, UrlKeyword string, pageNum, pageSize int) (gorm_common.CommonPage, error) {
	var umsResources []*UmsResource
	page := gorm_common.NewPage(pageNum, pageSize)

	err := gorm_common.ExecutePagedQuery(db, page, &umsResources, func(dbQuery *gorm.DB) *gorm.DB {
		if categoryId != 0 {
			dbQuery = dbQuery.Where("category_id = ?", categoryId)
		}
		if NameKeyword != "" {
			dbQuery = dbQuery.Where("name like ?", "%"+NameKeyword+"%")
		}
		if UrlKeyword != "" {
			dbQuery = dbQuery.Where("url like ?", "%"+UrlKeyword+"%")
		}
		return dbQuery
	})
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (usmResource *UmsResource) Insert(db *gorm.DB) (int64, error) {
	tx := db.Create(usmResource)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (usmResource *UmsResource) Update(db *gorm.DB, id int64) (int64, error) {
	usmResource.Id = id
	tx := db.Updates(usmResource)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (usmResource *UmsResource) Delete(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(usmResource, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
