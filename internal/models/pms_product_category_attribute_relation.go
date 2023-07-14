// @author hongjun500
// @date 2023/7/14 15:42
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

import "gorm.io/gorm"

type PmsProductCategoryAttributeRelation struct {
	Model
	// 商品分类id
	ProductCategoryId int64 `json:"productCategoryId"`
	// 商品属性id
	ProductAttributeId int64 `json:"productAttributeId"`
}

func (*PmsProductCategoryAttributeRelation) TableName() string {
	return "pms_product_category_attribute_relation"
}

// CreateBatch 批量创建
func (p *PmsProductCategoryAttributeRelation) CreateBatch(db *gorm.DB, relations []PmsProductCategoryAttributeRelation) (int64, error) {
	tx := db.Create(&relations)
	return tx.RowsAffected, tx.Error
}
