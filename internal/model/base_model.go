package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Model  基础Model
type Model struct {
	Id       int64                 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt time.Time             `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt time.Time             `gorm:"column:update_at;not null" json:"update_at"`
	DeleteAt time.Time             `gorm:"column:delete_at;not null" json:"delete_at"`
	IsDel    soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"is_del"`
}
