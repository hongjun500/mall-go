// @author hongjun500
// @date 2023/7/14 15:12
// @tool ThinkPadX1隐士
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

type PmsProductCategory struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt  *time.Time `gorm:"column:created_at;not null" json:"createdAt"`
	UpdateAt  *time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"deletedAt"`
	// 原有表结构字段 用于兼容
	CreateTime *time.Time            `gorm:"-"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"isDel"`
	// 父分类的编号
	ParentId int64 `json:"parentId" gorm:"column:parent_id"`
	// 分类名称
	Name string `json:"name" gorm:"column:name"`
	// 分类级别
	Level int `json:"level" gor:"column:level"`
	// 分类单位
	ProductUnit string `json:"productUnit" gorm:"column:product_unit"`
	// 分类数量
	ProductCount int `json:"productCount" gorm:"column:product_count"`
	// 是否显示在导航栏
	NavStatus int `json:"navStatus" gorm:"column:nav_status"`
	// 显示状态
	ShowStatus int `json:"showStatus" gorm:"column:show_status"`
	// 排序
	Sort int `json:"sort" gorm:"column:sort"`
	// 图标
	Icon string `json:"icon" gorm:"column:icon"`
	// 关键字
	Keywords string `json:"keywords" gorm:"column:keywords"`
	// 描述
	Description string `json:"description" gorm:"column:description"`
}

type PmsProductCategoryWithChildrenItem struct {
	PmsProductCategory
	Children []*PmsProductCategory `json:"children" gorm:"foreignKey:ParentId"`
}

func (*PmsProductCategory) TableName() string {
	return "pms_product_category"
}

func (p *PmsProductCategory) Create(db *gorm.DB) (int64, error) {
	tx := db.Create(p)
	return tx.RowsAffected, tx.Error
}

func (p *PmsProductCategory) Update(db *gorm.DB) (int64, error) {
	tx := db.Where("id = ?", p.Id).Updates(p)
	return tx.RowsAffected, tx.Error
}

func (p *PmsProductCategory) SelectById(db *gorm.DB, id int64) (*PmsProductCategory, error) {
	pmsProductCategory := new(PmsProductCategory)
	tx := db.Where("id = ?", id).First(&pmsProductCategory)
	return pmsProductCategory, tx.Error
}

func (p *PmsProductCategory) SelectListPage(db *gorm.DB, pageNum, pageSize int) (*pkg.CommonPage, error) {
	var pmsProductCategories []*PmsProductCategory
	page := internal.NewGormPage(db, pageNum, pageSize)
	page.List = &pmsProductCategories
	err := page.Paginate()
	if err != nil {
		return page.CommonPage, err
	}
	return page.CommonPage, nil
}

func (p *PmsProductCategory) DeleteById(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(p, id)
	return tx.RowsAffected, tx.Error
}

func (p *PmsProductCategory) UpdateNavStatusByIds(db *gorm.DB, ids []int64, navStatus int) (int64, error) {
	tx := db.Model(p).Where("id in (?)", ids).Update("nav_status", navStatus)
	return tx.RowsAffected, tx.Error
}

func (p *PmsProductCategory) UpdateShowStatusByIds(db *gorm.DB, ids []int64, showStatus int) (int64, error) {
	tx := db.Model(p).Where("id in (?)", ids).Update("show_status", showStatus)
	return tx.RowsAffected, tx.Error
}

func (p *PmsProductCategory) SelectList(db *gorm.DB) ([]*PmsProductCategoryWithChildrenItem, error) {
	var pmsProductCategoryWithChildrenItem []*PmsProductCategoryWithChildrenItem
	tx := db.Preload("Children").
		Where("parent_id = 0").
		Find(&pmsProductCategoryWithChildrenItem)
	return pmsProductCategoryWithChildrenItem, tx.Error
}
