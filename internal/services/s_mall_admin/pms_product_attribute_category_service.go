// @author hongjun500
// @date 2023/7/20 13:43
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/pkg"
)

type PmsProductAttributeCategoryService struct {
	DbFactory *database.DbFactory
}

func NewPmsProductAttributeCategoryService(DbFactory *database.DbFactory) PmsProductAttributeCategoryService {
	return PmsProductAttributeCategoryService{DbFactory: DbFactory}
}

// Create 创建商品属性分类
func (s PmsProductAttributeCategoryService) Create(name string) (int64, error) {
	pmsProductAttributeCategory := &models.PmsProductAttributeCategory{
		Name: name,
	}
	rows, err := pmsProductAttributeCategory.Insert(s.DbFactory.GormMySQL)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}

// Update 更新商品属性分类
func (s PmsProductAttributeCategoryService) Update(id int64, name string) (int64, error) {
	pmsProductAttributeCategory := &models.PmsProductAttributeCategory{}
	rows, err := pmsProductAttributeCategory.UpdateName(s.DbFactory.GormMySQL, id, name)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}

// Delete 删除商品属性分类
func (s PmsProductAttributeCategoryService) Delete(id int64) (int64, error) {
	pmsProductAttributeCategory := &models.PmsProductAttributeCategory{}
	rows, err := pmsProductAttributeCategory.Delete(s.DbFactory.GormMySQL, id)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}

// GetByID 根据id获取商品属性分类
func (s PmsProductAttributeCategoryService) GetByID(id int64) (*models.PmsProductAttributeCategory, error) {
	pmsProductAttributeCategory := &models.PmsProductAttributeCategory{}
	pmsProductAttributeCategory, err := pmsProductAttributeCategory.SelectById(s.DbFactory.GormMySQL, id)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return pmsProductAttributeCategory, nil
}

// ListPage 分页查询商品属性分类
func (s PmsProductAttributeCategoryService) ListPage(pageNum int, pageSize int) (*pkg.CommonPage, error) {
	pmsProductAttributeCategory := &models.PmsProductAttributeCategory{}
	page, err := pmsProductAttributeCategory.SelectPage(s.DbFactory.GormMySQL, pageNum, pageSize)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return page, nil
}

// ListWithAttr 查询所有商品属性分类及其下属性
func (s PmsProductAttributeCategoryService) ListWithAttr() ([]*models.PmsProductAttributeCategoryItem, error) {
	pmsProductAttributeCategory := &models.PmsProductAttributeCategory{}
	items, err := pmsProductAttributeCategory.SelectWithAttr(s.DbFactory.GormMySQL)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return items, nil
}
