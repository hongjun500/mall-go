package models

import (
	"time"

	"github.com/hongjun500/mall-go/internal"
	"github.com/hongjun500/mall-go/pkg"

	"gorm.io/gorm"
)

type UmsAdmin struct {
	Model
	Username  string     `gorm:"column:username;not null" json:"username"` // 用户名
	Password  string     `gorm:"column:password;not null" json:"password"` // 密码
	Icon      string     `gorm:"column:icon;" json:"icon"`                 // 头像
	Email     string     `gorm:"column:email;" json:"email"`               // 邮箱
	Nickname  string     `gorm:"column:nick_name;" json:"nickName"`        // 昵称
	Note      string     `gorm:"column:note;" json:"note"`                 // 备注信息
	LoginTime *time.Time `gorm:"column:login_time;" json:"loginTime"`      // 最后登录时间
	Status    int64      `gorm:"column:status;default:1" json:"status"`    // 帐号启用状态：0->禁用；1->启用
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
func (umsAdmin *UmsAdmin) SelectUmsAdminPage(db *gorm.DB, keyword string, pageNum, pageSize int) (*pkg.CommonPage, error) {
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
	page := internal.NewGormPage(db, pageNum, pageSize)
	page.List = &umsAdmins
	page.QueryFunc = func(query *gorm.DB) *gorm.DB {
		if keyword != "" {
			return query.Where("username like  ? or nick_name like ?", "%"+keyword+"%", "%"+keyword+"%")
		}
		return query
	}
	err := page.Paginate()
	if err != nil {
		return page.CommonPage, err
	}
	return page.CommonPage, nil
}

// UpdateUmsAdminByUserId 更新用户信息
func (umsAdmin *UmsAdmin) UpdateUmsAdminByUserId(db *gorm.DB, userId int64) (int64, error) {
	tx := db.Where("id = ?", userId).Updates(umsAdmin)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// DeleteUmsAdminByUserId 删除用户信息
func (umsAdmin *UmsAdmin) DeleteUmsAdminByUserId(db *gorm.DB, userId int64) (int64, error) {
	tx := db.Where("id = ?", userId).Delete(umsAdmin)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// UpdateUmsAdminPasswordByUserId 更改用户密码
func (umsAdmin *UmsAdmin) UpdateUmsAdminPasswordByUserId(db *gorm.DB) (int64, error) {
	tx := db.Model(umsAdmin).Select("password").Updates(UmsAdmin{Password: umsAdmin.Password})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// UpdateUmsAdminLoginTimeByUserId 更改用户登录时间
func (umsAdmin *UmsAdmin) UpdateUmsAdminLoginTimeByUserId(db *gorm.DB) (int64, error) {
	tx := db.Model(umsAdmin).Select("login_time").Updates(UmsAdmin{LoginTime: umsAdmin.LoginTime})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// UpdateUmsAdminStatusByUserId 更改用户状态
func (umsAdmin *UmsAdmin) UpdateUmsAdminStatusByUserId(db *gorm.DB) (int64, error) {
	tx := db.Model(umsAdmin).Select("status").Updates(UmsAdmin{Status: umsAdmin.Status})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
