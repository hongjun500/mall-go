package models

import (
	"github.com/hongjun500/mall-go/internal/gorm_common"
	"gorm.io/gorm"
	"time"
)

type UmsAdmin struct {
	Model
	// 用户名
	Username string `gorm:"column:username;not null" json:"username"`
	// 密码
	Password string `gorm:"column:password;not null" json:"password"`
	// 头像
	Icon string `gorm:"column:icon;" json:"icon"`
	// 邮箱
	Email string `gorm:"column:email;" json:"email"`
	// 昵称
	Nickname string `gorm:"column:nick_name;" json:"nickName"`
	// 备注信息
	Note string `gorm:"column:note;" json:"note"`
	// 最后登录时间
	LoginTime *time.Time `gorm:"column:login_time;" json:"loginTime"`
	// 帐号启用状态：0->禁用；1->启用
	Status int64 `gorm:"column:status;default:1" json:"status"`
}

func (*UmsAdmin) TableName() string {
	return "ums_admin"
}

// SelectUmsAdminByUsername 根据用户名获取用户信息
func (umsAdmin *UmsAdmin) SelectUmsAdminByUsername(db *gorm.DB, username string) ([]*UmsAdmin, error) {
	var umsAdmins []*UmsAdmin
	tx := db.Where("username = ?", username).Find(&umsAdmins)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return umsAdmins, nil
}

// InsertUmsAdmin 添加用户
func (umsAdmin *UmsAdmin) InsertUmsAdmin(db *gorm.DB) (int64, error) {
	tx := db.Create(umsAdmin)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// SelectUmsAdminByUserId 根据用户 ID 获取用户信息
func (umsAdmin *UmsAdmin) SelectUmsAdminByUserId(db *gorm.DB, userId int64) (*UmsAdmin, error) {
	// 如果这里的指针接收者是 nil,需要放入的是一个二级指针
	// tx := db.First(&umsAdmin, userId)
	tx := db.First(umsAdmin, userId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return umsAdmin, nil
}

// SelectUmsAdminPage 分页获取用户列表
func (umsAdmin *UmsAdmin) SelectUmsAdminPage(db *gorm.DB, keyword string, pageNum, pageSize int) (gorm_common.CommonPage, error) {
	/*var umsAdmins []*UmsAdmin
	dbQuery := db.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	if keyword != "" {
		dbQuery = dbQuery.Where("username like  ? or nick_name like ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	tx := dbQuery.Find(&umsAdmins)
	if tx.Error != nil {
		return nil, tx.Error
	}*/
	var umsAdmins []*UmsAdmin
	page := gorm_common.NewPage(pageNum, pageSize)
	if err := gorm_common.ExecutePagedQuery(db, page, &umsAdmins, func(query *gorm.DB) *gorm.DB {
		if keyword != "" {
			return query.Where("username like  ? or nick_name like ?", "%"+keyword+"%", "%"+keyword+"%")
		}
		return query
	}); err != nil {
		return nil, err // 处理错误
	}

	return page, nil
}

// UpdateUmsAdminByUserId 更新用户信息
func (umsAdmin *UmsAdmin) UpdateUmsAdminByUserId(db *gorm.DB, userId int64) (int64, error) {
	tx := db.Where("id = ?", userId).Updates(umsAdmin)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
