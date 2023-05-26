package model

type UmsResourceCategory struct {
	Model
	// 分类名称
	Name string `gorm:"column:name;not null" json:"name"`
	// 排序
	Sort int `gorm:"column:sort;default:0" json:"sort"`
}
