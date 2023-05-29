package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

// Model  基础Model
type Model struct {
	Id        int64                 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt  *time.Time            `gorm:"column:created_at;not null" json:"created_at"`
	UpdateAt  *time.Time            `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt *time.Time            `gorm:"column:deleted_at;" json:"deleted_at"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"is_del"`
}

func (*Model) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	tx.Statement.SetColumn("created_at", now)
	tx.Statement.SetColumn("updated_at", now)
	tx.Statement.SetColumn("deleted_at", nil)
	return
}

func (*Model) BeforeUpdate(tx *gorm.DB) (err error) {
	if !tx.Statement.Changed("updated_at") {
		tx.Statement.SetColumn("updated_at", time.Now())
	}
	return
}

func (*Model) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("deleted_at", time.Now())
	return
}
