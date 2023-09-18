// @author hongjun500
// @date 2023/7/26 11:07
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

type PmsBrand struct {
	Id        int64                 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt  *time.Time            `gorm:"column:created_at;not null" json:"createdAt"`
	UpdateAt  *time.Time            `gorm:"column:updated_at;not null" json:"updatedAt"`
	DeletedAt *time.Time            `gorm:"column:deleted_at;" json:"deletedAt"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"isDel"`

	// 品牌名称
	Name string `gorm:"column:name" json:"name"`
	// 首字母
	FirstLetter string `gorm:"column:first_letter" json:"firstLetter"`
	// 排序
	Sort int `gorm:"column:sort" json:"sort"`
	// 是否为品牌制造商：0->不是；1->是
	FactoryStatus int `gorm:"column:factory_status" json:"factoryStatus"`
	// 是否显示
	ShowStatus int `gorm:"column:show_status" json:"showStatus"`
	// 产品数量
	ProductCount int `gorm:"column:product_count" json:"productCount"`
	// 产品评论数量
	ProductCommentCount int `gorm:"column:product_comment_count" json:"productCommentCount"`
	// 品牌logo
	Logo string `gorm:"column:logo" json:"logo"`
	// 专区大图
	BigPic string `gorm:"column:big_pic" json:"bigPic"`
	// 品牌故事
	BrandStory string `gorm:"column:brand_story" json:"brandStory"`
}

func (*PmsBrand) TableName() string {
	return "pms_brand"
}

func (brand *PmsBrand) SelectAll(db *gorm.DB) ([]*PmsBrand, error) {
	var brands []*PmsBrand
	err := db.Find(&brands).Error
	return brands, err
}

func (brand *PmsBrand) SelectById(db *gorm.DB, id int64) (*PmsBrand, error) {
	err := db.First(brand, id).Error
	return brand, err
}

func (brand *PmsBrand) Insert(db *gorm.DB) (int64, error) {
	tx := db.Create(brand)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (brand *PmsBrand) Update(db *gorm.DB) (int64, error) {
	tx := db.Save(brand)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (brand *PmsBrand) Delete(db *gorm.DB) (int64, error) {
	tx := db.Delete(brand)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (brand *PmsBrand) DeleteBatch(db *gorm.DB, ids []int64) (int64, error) {
	tx := db.Where("id IN (?)", ids).Delete(brand)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (brand *PmsBrand) SelectPage(db *gorm.DB, keyword string, showStatus, pageNum, pageSize int) (*pkg.CommonPage, error) {
	var brands []*PmsBrand
	page := internal.NewGormPage(db, pageNum, pageSize, "sort", "desc")
	page.List = &brands
	page.QueryFunc = func(query *gorm.DB) *gorm.DB {
		if keyword != "" {
			query = query.Where("name LIKE ?", "%"+keyword+"%")
		}
		if showStatus != 0 {
			query = query.Where("show_status = ?", showStatus)
		}
		return query
	}
	err := page.Paginate()
	if err != nil {
		return nil, err
	}
	return page.CommonPage, nil
}

// UpdateShowStatusOrFactoryStatus 批量更新显示状态或厂家状态 column: show_status, factory_status
func (brand *PmsBrand) UpdateShowStatusOrFactoryStatus(db *gorm.DB, ids []int64, tColumn int, value any) (int64, error) {
	var tx *gorm.DB
	if tColumn == 0 {
		tx = db.Model(brand).Where("id IN (?)", ids).Update("show_status", value)
	} else {
		tx = db.Model(brand).Where("id IN (?)", ids).Update("factory_status", value)
	}
	return tx.RowsAffected, tx.Error
}
