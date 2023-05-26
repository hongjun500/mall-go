package model

type UmsPermission struct {
	Model
	// 父级ID
	ParentID int64 `gorm:"column:parent_id;" json:"parent_id"`
	// 名称
	Name string `gorm:"column:name;not null" json:"name"`
	// 权限值
	Value string `gorm:"column:value;not null" json:"value"`
	// 图标
	Icon string `gorm:"column:icon;" json:"icon"`
	// 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）
	Type int64 `gorm:"column:type;not null" json:"type"`
	// 前端资源路径
	Uri string `gorm:"column:uri;" json:"uri"`
	// 启用状态；0->禁用；1->启用
	Status int64 `gorm:"column:status;default:1" json:"status"`
	// 排序值
	Sort int64 `gorm:"column:sort;default:0" json:"sort"`
}
