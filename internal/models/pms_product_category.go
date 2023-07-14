// @author hongjun500
// @date 2023/7/14 15:12
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

import "gorm.io/gorm"

type PmsProductCategory struct {
	Model
	// 父分类的编号
	ParentId int64 `json:"parentId"`
	// 分类名称
	Name string `json:"name"`
	// 分类级别
	Level int `json:"level"`
	// 分类单位
	ProductUnit string `json:"productUnit"`
	// 分类数量
	ProductCount int `json:"productCount"`
	// 是否显示在导航栏
	NavStatus int `json:"navStatus"`
	// 显示状态
	ShowStatus int `json:"showStatus"`
	// 排序
	Sort int `json:"sort"`
	// 图标
	Icon string `json:"icon"`
	// 关键字
	Keywords string `json:"keywords"`
	// 描述
	Description string `json:"description"`
}

func (*PmsProductCategory) TableName() string {
	return "pms_product_category"
}

func (p *PmsProductCategory) Create(db *gorm.DB) (int64, error) {
	tx := db.Create(p)
	return tx.RowsAffected, tx.Error
}

func (p *PmsProductCategory) SelectById(db *gorm.DB, id int64) (*PmsProductCategory, error) {
	pmsProductCategory := new(PmsProductCategory)
	tx := db.Where("id = ?", id).First(&pmsProductCategory)
	return pmsProductCategory, tx.Error
}
