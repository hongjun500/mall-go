package testting

import (
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/model"
	"testing"
)

func TestCreateMenu(t *testing.T) {
	var menu = model.UmsMenu{
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
	row, err := menu.Create(conf.Db)
	if err != nil {
		t.Errorf("create error: %v", err)
		return
	}
	t.Log("create success", row)
}

func TestUpdateMenu(t *testing.T) {
	var menu model.UmsMenu
	umsMenu1, _ := menu.Get(conf.Db, 31)
	umsMenu2, _ := menu.Get(conf.Db, 36)
	umsMenu1.Title = "一级菜单->修改"
	umsMenu1.Name = "一级菜单->修改"
	umsMenu1.Hidden = 1
	umsMenu2.Title = "一级菜单下的二级菜单->修改"
	umsMenu2.Name = "一级菜单下的二级菜单->修改"
	umsMenu2.Hidden = 1
	var menus []*model.UmsMenu
	menus = append(menus, umsMenu1, umsMenu2)
	err := menu.Update(conf.Db, menus)
	if err != nil {
		return
	}
	t.Log("update success RowsAffected = ", len(menus))
}

func TestDeleteMenu(t *testing.T) {
	var menu model.UmsMenu
	i, err := menu.Delete(conf.Db, 36)
	if err != nil {
		return
	}
	t.Log("delete success RowsAffected = ", i)
}

func TestListPageMenu(t *testing.T) {
	var menu model.UmsMenu
	menus, err := menu.ListPage(conf.Db, 1, 10)
	if err != nil {
		return
	}
	t.Log("list all success, len = ", len(menus))
}

func TestUpdateHidden(t *testing.T) {
	var menu model.UmsMenu
	row, err := menu.UpdateHidden(conf.Db, 36, 1)
	if err != nil {
		return
	}
	t.Log("update success RowsAffected = ", row)
}
