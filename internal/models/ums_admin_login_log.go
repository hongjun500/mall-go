//	@author	hongjun500
//	@date	2023/6/10
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 用户登录日志

package models

import "gorm.io/gorm"

type UmsAdminLoginLog struct {
	Model
	// 用户id
	AdminId int64 `gorm:"column:admin_id;not null;" json:"adminId"`
	// 登录 ip
	Ip string `gorm:"column:ip;type:varchar(64);not null;" json:"ip"`
	// 地址
	Address string `gorm:"column:address;type:varchar(255);not null;" json:"address"`
	// 浏览器登录类型
	UserAgent string `gorm:"column:user_agent;type:varchar(255);not null;" json:"userAgent"`
}

// TableName 自定义表名
func (*UmsAdminLoginLog) TableName() string {
	return "ums_admin_login_log"
}

// SaveLoginLog 保存登录日志
func (loginLog *UmsAdminLoginLog) SaveLoginLog(db *gorm.DB) (int64, error) {
	result := db.Create(loginLog)
	return result.RowsAffected, result.Error
}
