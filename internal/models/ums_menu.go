package models

import (
	"github.com/hongjun500/mall-go/internal"
	"github.com/hongjun500/mall-go/pkg"
	"gorm.io/gorm"
)

type UmsMenu struct {
	Model
	// 父级ID
	ParentID int64 `gorm:"column:parent_id;" json:"parentId"`
	// 菜单名称
	Title string `gorm:"column:title;not null" json:"title"`
	// 菜单级数
	Level int64 `gorm:"column:level;not null" json:"level"`
	// 菜单排序
	Sort int64 `gorm:"column:sort;default:0" json:"sort"`
	// 前端名称
	Name string `gorm:"column:name;" json:"name"`
	// 前端图标
	Icon string `gorm:"column:icon;" json:"icon"`
	// 前端隐藏
	Hidden int64 `gorm:"column:hidden;default:0" json:"hidden"`
}

type UmsMenuNode struct {
	UmsMenu
	Children []*UmsMenuNode `json:"children"`
}

func (*UmsMenu) TableName() string {
	return "ums_menu"
}

func (umsMenu *UmsMenu) InsertUmsMenu(db *gorm.DB) (int64, error) {
	// umsMenu.UpdateLevel(db)
	tx := db.Create(umsMenu)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// Update 批量更新菜单 传入的是一个切片
func (umsMenu *UmsMenu) Update(db *gorm.DB, menus []*UmsMenu) error {
	if menus == nil || len(menus) == 0 {
		return nil
	}
	var ids []int64
	for _, menu := range menus {
		menu.UpdateLevel(db)
		ids = append(ids, menu.Model.Id)
		db.Save(menu)
	}
	return nil
}

// UpdateLevel 更新菜单级数
func (umsMenu *UmsMenu) UpdateLevel(db *gorm.DB) {
	if umsMenu.ParentID == 0 {
		umsMenu.Level = 0
	} else {
		// 有父级菜单时根据父级菜单 level 设置
		parentMenu, _ := umsMenu.SelectById(db, umsMenu.ParentID)
		if parentMenu != nil {
			umsMenu.Level = parentMenu.Level + 1
		} else {
			umsMenu.Level = 0
		}
	}
}

func (umsMenu *UmsMenu) SelectById(db *gorm.DB, id int64) (*UmsMenu, error) {
	var menu UmsMenu
	tx := db.First(&menu, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &menu, nil
}

func (umsMenu *UmsMenu) Delete(db *gorm.DB, id int64) (int64, error) {
	tx := db.Delete(umsMenu, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// SelectPage 获取菜单分页列表
func (umsMenu *UmsMenu) SelectPage(db *gorm.DB, pageNum, pageSize int, parentId int64) (*pkg.CommonPage, error) {
	var menus []*UmsMenu
	page := internal.NewGormPage(db, pageNum, pageSize)
	page.List = &menus
	page.OrderBy = "sort"
	page.Sort = "desc"
	page.QueryFunc = func(dbQuery *gorm.DB) *gorm.DB {
		if parentId != 0 {
			return dbQuery.Where("parent_id = ?", parentId)
		}
		return dbQuery
	}
	err := page.Paginate()
	if err != nil {
		return page.CommonPage, err
	}
	return page.CommonPage, nil
}

// UpdateHidden 更新菜单显示状态
func (umsMenu *UmsMenu) UpdateHidden(db *gorm.DB, id, hidden int64) (int64, error) {
	tx := db.Model(umsMenu).Where("id = ?", id).Update("hidden", hidden)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

// SelectAll 查询所有菜单
func (umsMenu *UmsMenu) SelectAll(db *gorm.DB) ([]*UmsMenu, error) {
	var menus []*UmsMenu
	tx := db.Order("sort desc").Find(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return menus, nil
}

// ListMenuTree 获取菜单树形列表
func (umsMenu *UmsMenu) ListMenuTree(db *gorm.DB) ([]*UmsMenuNode, error) {
	var menus []*UmsMenu
	tx := db.Order("sort desc").Find(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var umsMenuNodes []*UmsMenuNode
	for _, menu := range menus {
		if menu.ParentID == 0 {
			umsMenuNodes = append(umsMenuNodes, convertMenuTreeNode(menu, menus))
		}
	}
	return umsMenuNodes, nil
}

// ConvertMenuTreeNode 转换菜单树形结构
func convertMenuTreeNode(menu *UmsMenu, menus []*UmsMenu) *UmsMenuNode {
	var umsMenuNode = &UmsMenuNode{}
	var umsMenuNodeChildren []*UmsMenuNode
	umsMenuNode.UmsMenu = *menu
	for _, umsMenu := range menus {
		if umsMenu.ParentID == menu.Id {
			umsMenuNodeChildren = append(umsMenuNodeChildren, convertMenuTreeNode(umsMenu, menus))
		}
	}
	umsMenuNode.Children = umsMenuNodeChildren
	return umsMenuNode
}
