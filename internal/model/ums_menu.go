package model

type UmsMenu struct {
	Model
	// 父级ID
	ParentID int64 `gorm:"column:parent_id;" json:"parent_id"`
	// 菜单名称
	Title string `gorm:"column:title;not null" json:"title"`
	// 菜单级数
	Level int64 `gorm:"column:level;not null" json:"level"`
	// 菜单排序
	Sort int64 `gorm:"column:sort;default:0" json:"sort"`
	// 前端名称
	Name string `gorm:"column:name;" json:"name"`
	// 前端图标
	Icon string `gorm:"column:icon;" json:"icon"`
	// 前端隐藏
	Hidden int64 `gorm:"column:hidden;default:0" json:"hidden"`
}
