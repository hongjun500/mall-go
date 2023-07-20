//	@author	hongjun500
//	@date	2023/6/26 16:24
//	@tool	ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

type PmsProductAttributeValue struct {
	Model
	ProductId          int64 `json:"productId" gorm:"product_id"`
	ProductAttributeId int64 `json:"productAttributeId" gorm:"attr_product_attribute_id"`
	// 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
	Value string `json:"value" gorm:"value"`

	ProductAttribute *PmsProductAttribute `json:"productAttribute" gorm:"foreignKey:ProductAttributeId"`
}

func (*PmsProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}
