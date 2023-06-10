package models

import (
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
}

func (*UmsResource) TableName() string {
	return "ums_resource"
}

func (usmResource *UmsResource) SelectAll(db *gorm.DB) ([]*UmsResource, error) {
	var umsResources []*UmsResource
	tx := db.Find(&umsResources)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return umsResources, nil
}

func (usmResource *UmsResource) SelectPage(db *gorm.DB, categoryId int64, NameKeyword string, UrlKeyword string, pageNum, pageSize int) ([]*UmsResource, error) {
	var umsResources []*UmsResource
	dbQuery := db.Offset(pageNum).Limit(pageSize)
	if categoryId != 0 {
		dbQuery = dbQuery.Where("category_id = ?", categoryId)
	}
	if NameKeyword != "" {
		dbQuery = dbQuery.Where("name like ?", "%"+NameKeyword+"%")
	}
	if UrlKeyword != "" {
		dbQuery = dbQuery.Where("url like ?", "%"+UrlKeyword+"%")
	}
	tx := dbQuery.Find(&umsResources)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return umsResources, nil
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
	// todo: 更新缓存
	tx := db.Updates(usmResource)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (usmResource *UmsResource) GetUmsResourceById(db *gorm.DB, id int64) (*UmsResource, error) {
	var umsResourceModel UmsResource
	tx := db.First(&umsResourceModel, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &umsResourceModel, nil
}

func (usmResource *UmsResource) Delete(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(usmResource, id)
	// todo: 清除缓存
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
