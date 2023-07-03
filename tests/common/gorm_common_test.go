// @author hongjun500
// @date 2023/6/10
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: gorm_common_test.go

package common

import (
	"github.com/hongjun500/mall-go/internal"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestExecutePagedQuery(t *testing.T) {
	page := internal.NewGormPage(db, 1, 10)
	var admins []*models.UmsAdmin
	page.List = &admins
	keywords := "hongjun500"
	page.QueryFunc = func(db *gorm.DB) *gorm.DB {
		if keywords != "" {
			return db.Where("username LIKE ? OR nick_name LIKE ?", "%"+keywords+"%", "%"+keywords+"%")
		}
		return db
	}
	err := page.Paginate()
	if err != nil {
		return
	}
	assert.Len(t, admins, 1)
	assert.Equal(t, page.Total, int64(1))
	t.Log("ExecutePagedQuery admins: ", admins)
	t.Log("ExecutePagedQuery page: ", page)
}

// 测试原生sql分页
// ignore
func TestExecutePageSqlQuery(t *testing.T) {

	page := internal.NewNativeSqlPage(db, 1, 10)
	var menus []*models.UmsMenu
	page.List = &menus
	page.QueryArgs = []interface{}{1}
	page.CountArgs = []interface{}{1}
	page.CountSQL = "SELECT count(*) FROM ums_admin_role_relation arr LEFT JOIN ums_role r ON arr.role_id = r.id LEFT JOIN ums_role_menu_relation rmr ON r.id = rmr.role_id LEFT JOIN ums_menu m ON rmr.menu_id = m.id WHERE arr.admin_id = ? AND m.id IS NOT NULL GROUP BY m.id"
	page.QuerySQL = "SELECT m.id id, m.parent_id parentId, m.create_time createTime, m.title title, m.LEVEL LEVEL, m.sort sort, m.NAME NAME, m.icon icon, m.hidden hidden FROM ums_admin_role_relation arr LEFT JOIN ums_role r ON arr.role_id = r.id LEFT JOIN ums_role_menu_relation rmr ON r.id = rmr.role_id LEFT JOIN ums_menu m ON rmr.menu_id = m.id WHERE arr.admin_id = ? AND m.id IS NOT NULL GROUP BY m.id"
	err := page.Paginate()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(menus)
	t.Log(page)

	// gorm.ExecutePagedSQLQuery()
}
