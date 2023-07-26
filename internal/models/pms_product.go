//	@author	hongjun500
//	@date	2023/6/26 15:35
//	@tool	ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

import (
	"github.com/hongjun500/mall-go/internal/database"
	"gorm.io/gorm"
)

type PmsProduct struct {
	Model
	ProductSn                 string                      `json:"product_sn" gorm:"product_sn"`
	BrandId                   int64                       `json:"brand_id" gorm:"brand_id"`
	BrandName                 string                      `json:"brand_name" gorm:"brand_name"`
	ProductCategoryId         int64                       `json:"product_category_id" gorm:"product_category_id"`
	ProductCategoryName       string                      `json:"product_category_name" gorm:"product_category_name"`
	Pic                       string                      `json:"pic" gorm:"pic"`
	Name                      string                      `json:"name" gorm:"name"`
	SubTitle                  string                      `json:"sub_title" gorm:"sub_title"`
	KeyWord                   string                      `json:"key_word" gorm:"key_word"`
	Price                     string                      `json:"price" gorm:"price"`
	Sale                      int                         `json:"sale" gorm:"sale"`
	NewStatus                 int                         `json:"new_status" gorm:"new_status"`
	RecommendStatus           int                         `json:"recommand_status" gorm:"recommand_status"`
	Stock                     int                         `json:"stock" gorm:"stock"`
	PromotionType             int                         `json:"promotion_type" gorm:"promotion_type"`
	Sort                      int                         `json:"sort" gorm:"sort"`
	ProductAttributeValueList []*PmsProductAttributeValue `json:"attr_value_list" gorm:"foreignKey:ProductId"`
}

func (*PmsProduct) TableName() string {
	return "pms_product"
}

// SelectProductInfoById 根据商品 id 获取商品相关信息
func (pmsProduct *PmsProduct) SelectProductInfoById(db *database.DbFactory, id int64) ([]*PmsProduct, error) {
	var pmsProducts []*PmsProduct
	query := db.GormMySQL.Preload("ProductAttributeValueList").
		Preload("ProductAttributeValueList.ProductAttribute").
		Where("delete_status = ? AND publish_status = ?", 0, 1)
	if id != 0 {
		query = query.Where("id = ?", id)
	}
	query.Find(&pmsProducts)
	return pmsProducts, nil
}

// UpdateProductNameById 根据商品 id 更新商品名称
func (pmsProduct *PmsProduct) UpdateProductNameById(db *database.DbFactory, id int64, name string) (int64, error) {
	tx := db.GormMySQL.Model(pmsProduct).Where("id = ?", id).Update("name", name)
	return tx.RowsAffected, tx.Error
}

// UpdateBrandNameByBrandId 根据品牌 id 更改品牌名称
func (pmsProduct *PmsProduct) UpdateBrandNameByBrandId(db *gorm.DB, brandId int64, name string) (int64, error) {
	tx := db.Model(pmsProduct).Where("brand_id = ?", brandId).Update("brand_name", name)
	return tx.RowsAffected, tx.Error
}
