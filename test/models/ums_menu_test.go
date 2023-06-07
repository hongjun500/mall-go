package models

import (
	"github.com/hongjun500/mall-go/internal/models"
	"testing"
)

func TestCreateMenu(t *testing.T) {
	var menu = models.UmsMenu{
		ParentID: 31,
		Title:    "一级菜单下的二级菜单",
		Name:     "一级菜单下的二级菜单",

		/*ParentID: 0,
		Title:    "一级菜单",
		Name:     "一级菜单",*/
		Hidden: 0,
		Icon:   "el-icon-s-home",
		Sort:   0,
	}
	row, err := menu.InsertUmsMenu(TestModelGormMySQL)
	if err != nil {
		t.Errorf("create error: %v", err)
		return
	}
	t.Log("create success", row)
}

func TestUpdateMenu(t *testing.T) {
	var menu models.UmsMenu
	umsMenu1, _ := menu.SelectById(TestModelGormMySQL, 31)
	umsMenu2, _ := menu.SelectById(TestModelGormMySQL, 37)
	umsMenu1.Title = "一级菜单->修改"
	umsMenu1.Name = "一级菜单->修改"
	umsMenu1.Hidden = 1
	umsMenu2.Title = "一级菜单下的二级菜单->修改"
	umsMenu2.Name = "一级菜单下的二级菜单->修改"
	umsMenu2.Hidden = 1
	var menus []*models.UmsMenu
	menus = append(menus, umsMenu1, umsMenu2)
	err := menu.Update(TestModelGormMySQL, menus)
	if err != nil {
		return
	}
	t.Log("update success RowsAffected = ", len(menus))
}

func TestDeleteMenu(t *testing.T) {
	var menu models.UmsMenu
	i, err := menu.Delete(TestModelGormMySQL, 37)
	if err != nil {
		return
	}
	t.Log("delete success RowsAffected = ", i)
}

func TestListPageMenu(t *testing.T) {
	var menu models.UmsMenu
	menus, err := menu.SelectPage(TestModelGormMySQL, 1, 10)
	if err != nil {
		return
	}
	t.Log("list all success, len = ", len(menus))
}

func TestUpdateHidden(t *testing.T) {
	var menu models.UmsMenu
	row, err := menu.UpdateHidden(TestModelGormMySQL, 37, 0)
	if err != nil {
		return
	}
	t.Log("update success RowsAffected = ", row)
}

func TestListMenuTree(t *testing.T) {
	var menu models.UmsMenu
	menus, err := menu.ListMenuTree(TestModelGormMySQL)
	if err != nil {
		return
	}
	t.Log("list all success, len = ", len(menus))
}
