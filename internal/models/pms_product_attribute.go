//	@author	hongjun500
//	@date	2023/6/26 16:24
//	@tool	ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

import (
	"github.com/hongjun500/mall-go/internal"
	"github.com/hongjun500/mall-go/pkg"
	"gorm.io/gorm"
)

type PmsProductAttribute struct {
	Model
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId" gorm:"product_attribute_category_id"`
	Name                       string `json:"name" gorm:"attr_name"`
	// 属性选择类型：0->唯一；1->单选；2->多选
	SelectType int `json:"selectType" json:"selectType" gorm:"select_type"`
	// 属性录入方式：0->手工录入；1->从列表中选取
	InputType int `json:"inputType" gorm:"input_type"`
	// 可选值列表，以逗号隔开
	InputList string `json:"inputList" gorm:"input_list"`
	// 排序字段：最高的可以单独上传图片
	Sort int `json:"sort" gorm:"sort"`
	// 分类筛选样式：1->普通；1->颜色
	FilterType int `json:"filterType" gorm:"filter_type"`
	// 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	SearchType int `json:"searchType" gorm:"search_type"`
	// 相同属性产品是否关联；0->不关联；1->关联
	RelatedStatus int `json:"relatedStatus" gorm:"related_status"`
	// 是否支持手动新增；0->不支持；1->支持
	HandAddStatus int `json:"handAddStatus" gorm:"hand_add_status"`
	// 属性的类型；0->规格；1->参数
	Type int `json:"type" gorm:"type"`
}

func (*PmsProductAttribute) TableName() string {
	return "pms_product_attribute"
}

// ProductAttrInfo 用于join查询的结构体，必须指定gorm的column，否则属性将是零值
type ProductAttrInfo struct {
	AttributeId         int64 `json:"attributeId" gorm:"column:attributeId"`
	AttributeCategoryId int64 `json:"attributeCategoryId" gorm:"column:attributeCategoryId"`
}

func (productAttribute *PmsProductAttribute) Insert(db *gorm.DB) (int64, error) {
	tx := db.Create(productAttribute)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (productAttribute *PmsProductAttribute) UpdateColumnById(db *gorm.DB, id int64, columns ...string) (int64, error) {
	if len(columns) == 0 {
		return 0, nil
	}
	tx := db.Model(productAttribute).Where("id = ?", id).Select(columns).Updates(productAttribute)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (productAttribute *PmsProductAttribute) UpdateById(db *gorm.DB, id int64) (int64, error) {
	tx := db.Model(productAttribute).Where("id = ?", id).Updates(productAttribute)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (productAttribute *PmsProductAttribute) DeleteById(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(productAttribute, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (productAttribute *PmsProductAttribute) DeleteByIds(db *gorm.DB, ids []int64) (int64, error) {
	tx := db.Where("id in (?)", ids).Delete(productAttribute)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (productAttribute *PmsProductAttribute) SelectById(db *gorm.DB, id int64) (*PmsProductAttribute, error) {
	tx := db.First(productAttribute, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return productAttribute, nil
}

func (productAttribute *PmsProductAttribute) SelectListByPage(db *gorm.DB, cid int64, ctype, pageNum, pageSize int) (*pkg.CommonPage, error) {
	var productAttributes []*PmsProductAttribute
	page := internal.NewGormPage(db, pageNum, pageSize, "sort", "desc")
	page.List = &productAttributes
	page.QueryFunc = func(db *gorm.DB) *gorm.DB {
		if cid != 0 {
			db = db.Where("product_attribute_category_id = ?", cid)
		}
		if ctype != 0 {
			db = db.Where("type = ?", ctype)
		}
		return db
	}
	err := page.Paginate()
	if err != nil {
		return nil, err
	}
	return page.CommonPage, nil
}

func (productAttribute *PmsProductAttribute) SelectListFromProductAttrInfo(db *gorm.DB, productCategoryId int64) ([]*ProductAttrInfo, error) {
	var productAttrInfos []*ProductAttrInfo
	tx := db.Table("pms_product_category_attribute_relation as pcar").
		Select("pa.id as attributeId, pac.id as attributeCategoryId").
		Joins("left join pms_product_attribute pa on pa.id = pcar.product_attribute_id").
		Joins("left join pms_product_attribute_category  pac on pa.product_attribute_category_id = pac.id").
		Where("pcar.product_category_id = ?", productCategoryId).
		Find(&productAttrInfos)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return productAttrInfos, nil
}
