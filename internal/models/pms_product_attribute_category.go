//	@author	hongjun500
//	@date	2023/6/26 16:24
//	@tool	ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

import (
	"time"

	"github.com/hongjun500/mall-go/internal"
	"github.com/hongjun500/mall-go/pkg"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type PmsProductAttributeCategory struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt  *time.Time `gorm:"column:created_at;not null" json:"createdAt"`
	UpdateAt  *time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"deletedAt"`
	// 原有表结构字段 用于兼容
	CreateTime *time.Time            `gorm:"-"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"isDel"`
	Name       string                `json:"value" gorm:"value"`
	// 属性数量
	AttributeCount int `json:"attributeCount" gorm:"attribute_count"`
	// 参数数量
	ParamCount int `json:"paramCount" gorm:"param_count"`
}

type PmsProductAttributeCategoryItem struct {
	PmsProductAttributeCategory
	ProductAttributeList []*PmsProductAttribute `json:"productAttributeList" gorm:"foreignKey:ProductAttributeCategoryId"`
}

func (*PmsProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) Insert(db *gorm.DB) (int64, error) {
	tx := db.Create(pmsProductAttributeCategory)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) UpdateName(db *gorm.DB, id int64, name string) (int64, error) {
	tx := db.Model(pmsProductAttributeCategory).Where("id = ?", id).Update("name", name)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) Delete(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(pmsProductAttributeCategory, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) SelectById(db *gorm.DB, id int64) (*PmsProductAttributeCategory, error) {
	tx := db.First(pmsProductAttributeCategory, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return pmsProductAttributeCategory, nil
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) SelectPage(db *gorm.DB, pageNum, pageSize int) (*pkg.CommonPage, error) {
	var pmsProductAttributeCategoryList []*PmsProductAttributeCategory
	page := internal.NewGormPage(db, pageNum, pageSize)
	page.List = &pmsProductAttributeCategoryList
	err := page.Paginate()
	if err != nil {
		return nil, err
	}
	return page.CommonPage, nil
}

func (pmsProductAttributeCategory *PmsProductAttributeCategory) SelectWithAttr(db *gorm.DB) ([]*PmsProductAttributeCategoryItem, error) {
	var pmsProductAttributeCategoryList []*PmsProductAttributeCategoryItem
	tx := db.Preload("ProductAttributeList").
		Joins("left join pms_product_attribute on pms_product_attribute_category.id = pms_product_attribute.product_attribute_category_id").
		Where("pms_product_attribute.type = ?", 0).
		Find(&pmsProductAttributeCategoryList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return pmsProductAttributeCategoryList, nil
}
