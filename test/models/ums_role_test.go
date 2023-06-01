package models

import (
	"github.com/hongjun500/mall-go/internal/initialize"
	"github.com/hongjun500/mall-go/internal/models"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Log("test create")
	var umsRole models.UmsRole
	umsRole.Name = "二手管理员"
	umsRole.Description = "二手"
	umsRole.AdminCount = 1
	// umsRole.CreateAt = time.Now()
	umsRole.Status = 1
	umsRole.Sort = 0
	create, err := umsRole.Create(initialize.SqlSession.DbMySQL)
	if err != nil {
		return // 测试失败
	}
	t.Logf("create success, RowsAffected = %v", create)
}

func TestUpdate(t *testing.T) {
	t.Log("test update")
	var umsRole models.UmsRole
	role, err := umsRole.Get(initialize.SqlSession.DbMySQL, 7)
	if err != nil {
		return
	}
	role.Name = "二手管理员->修改"
	role.Description = "二手->修改"
	role.DeletedAt = nil
	role.Update(initialize.SqlSession.DbMySQL, 7)

}

func TestDelete(t *testing.T) {
	t.Log("test delete")
	var umsRole models.UmsRole
	i, err := umsRole.Delete(initialize.SqlSession.DbMySQL, 8)
	if err != nil {
		return // 测试失败
	}
	t.Logf("delete success, RowsAffected = %v", i)
}

func TestListAll(t *testing.T) {

	t.Log("test list all")
	var umsRole models.UmsRole
	all, err := umsRole.ListAll(initialize.SqlSession.DbMySQL)
	if err != nil {
		return
	}
	t.Logf("all = %v", all)
}

func TestListPage(t *testing.T) {
	t.Log("test list page")
	var umsRole models.UmsRole
	page, err := umsRole.ListPage(initialize.SqlSession.DbMySQL, "管理员", 1, 10)
	if err != nil {
		return
	}
	t.Logf("page = %v", page)
}

func TestListMenu(t *testing.T) {
	t.Log("test list menu")
	var umsRole models.UmsRole
	menu, err := umsRole.ListMenu(initialize.SqlSession.DbMySQL, 1)
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}
	if len(menu) == 0 {
		t.Errorf("menu is empty")
	}
	for i, umsMenu := range menu {
		t.Logf("menu[%d] = %v", i, umsMenu)
	}

}
