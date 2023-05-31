package model

import "gorm.io/gorm"

type UmsResourceCategory struct {
	Model
	// 分类名称
	Name string `gorm:"column:name;not null" json:"name"`
	// 排序
	Sort int `gorm:"column:sort;default:0" json:"sort"`
}

func (*UmsResourceCategory) TableName() string {
	return "ums_resource_category"
}

func (umsResourceCategory *UmsResourceCategory) Create(db *gorm.DB) (int64, error) {
	tx := db.Create(umsResourceCategory)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (umsResourceCategory *UmsResourceCategory) Update(db *gorm.DB, id int64) (int64, error) {
	umsResourceCategory.Id = id
	tx := db.Updates(umsResourceCategory)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (umsResourceCategory *UmsResourceCategory) Get(db *gorm.DB, id int64) (*UmsResourceCategory, error) {
	var umsResourceCategoryResult UmsResourceCategory
	tx := db.First(&umsResourceCategoryResult, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &umsResourceCategoryResult, nil
}

func (umsResourceCategory *UmsResourceCategory) Delete(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(umsResourceCategory, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (umsResourceCategory *UmsResourceCategory) ListAll(db *gorm.DB, pageNum, pageSize int) ([]*UmsResourceCategory, error) {
	var umsResourceCategoryList []*UmsResourceCategory
	var err error
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("sort desc").Find(&umsResourceCategoryList).Error
	} else {
		err = db.Order("sort desc").Find(&umsResourceCategoryList).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return umsResourceCategoryList, nil
}
