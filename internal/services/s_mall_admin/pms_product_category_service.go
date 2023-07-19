// @author hongjun500
// @date 2023/7/14 14:13
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/request/ums_admin_dto"
	"github.com/hongjun500/mall-go/pkg"
	"gorm.io/gorm"
)

type PmsProductCategoryService struct {
	DbFactory *database.DbFactory
}

func NewPmsProductCategoryService(dbFactory *database.DbFactory) PmsProductCategoryService {
	return PmsProductCategoryService{DbFactory: dbFactory}
}

// CreateProductCategory 添加商品分类
func (s PmsProductCategoryService) CreateProductCategory(dto ums_admin_dto.PmsProductCategoryDTO) (int64, error) {
	pmsProductCategory := new(models.PmsProductCategory)
	pmsProductCategory.ProductCount = 0
	pmsProductCategory.Name = dto.Name
	pmsProductCategory.ParentId = dto.ParentId
	pmsProductCategory.ProductUnit = dto.ProductUnit
	pmsProductCategory.NavStatus = dto.NavStatus
	pmsProductCategory.ShowStatus = dto.ShowStatus
	pmsProductCategory.Sort = dto.Sort
	pmsProductCategory.Icon = dto.Icon
	pmsProductCategory.Keywords = dto.Keywords
	pmsProductCategory.Description = dto.Description

	setCategoryLevel(s.DbFactory.GormMySQL, pmsProductCategory)

	count, err := pmsProductCategory.Create(s.DbFactory.GormMySQL)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	// 创建筛选属性关联
	if len(dto.ProductAttributeIdList) > 0 {
		relations := make([]models.PmsProductCategoryAttributeRelation, len(dto.ProductAttributeIdList))
		for _, attrId := range dto.ProductAttributeIdList {
			item := models.PmsProductCategoryAttributeRelation{ProductAttributeId: attrId, ProductCategoryId: pmsProductCategory.Id}
			relations = append(relations, item)
		}
		m := new(models.PmsProductCategoryAttributeRelation)
		// 分类属性关联关系表的插入
		_, _ = m.CreateBatch(s.DbFactory.GormMySQL, relations)
	}
	return count, nil
}

func setCategoryLevel(db *gorm.DB, category *models.PmsProductCategory) {
	if category.ParentId == 0 {
		category.Level = 0
	} else {
		productCategory, _ := category.SelectById(db, category.ParentId)
		if productCategory != nil {
			category.Level = productCategory.Level + 1
		} else {
			category.Level = 0
		}
	}
}

// UpdateProductCategory 更新商品分类
func (s PmsProductCategoryService) UpdateProductCategory(id int64, dto ums_admin_dto.PmsProductCategoryDTO) (int64, error) {
	pmsProductCategory := new(models.PmsProductCategory)
	pmsProductCategory.Id = id
	pmsProductCategory.ProductCount = 0
	pmsProductCategory.Name = dto.Name
	pmsProductCategory.ParentId = dto.ParentId
	pmsProductCategory.ProductUnit = dto.ProductUnit
	pmsProductCategory.NavStatus = dto.NavStatus
	pmsProductCategory.ShowStatus = dto.ShowStatus
	pmsProductCategory.Sort = dto.Sort
	pmsProductCategory.Icon = dto.Icon
	pmsProductCategory.Keywords = dto.Keywords
	pmsProductCategory.Description = dto.Description

	setCategoryLevel(s.DbFactory.GormMySQL, pmsProductCategory)
	// 更新商品的名称
	product := new(models.PmsProduct)
	_, _ = product.UpdateProductNameById(s.DbFactory, id, pmsProductCategory.Name)

	// 更新筛选属性关联
	if len(dto.ProductAttributeIdList) > 0 {
		// 删除原有关联关系
		_, _ = new(models.PmsProductCategoryAttributeRelation).DeleteByCategoryId(s.DbFactory.GormMySQL, id)
		relations := make([]models.PmsProductCategoryAttributeRelation, len(dto.ProductAttributeIdList))
		for _, attrId := range dto.ProductAttributeIdList {
			item := models.PmsProductCategoryAttributeRelation{ProductAttributeId: attrId, ProductCategoryId: pmsProductCategory.Id}
			relations = append(relations, item)
		}
		m := new(models.PmsProductCategoryAttributeRelation)
		// 分类属性关联关系表的插入
		_, _ = m.CreateBatch(s.DbFactory.GormMySQL, relations)
	} else {
		_, _ = new(models.PmsProductCategoryAttributeRelation).DeleteByCategoryId(s.DbFactory.GormMySQL, id)
	}
	count, err := pmsProductCategory.Update(s.DbFactory.GormMySQL)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return count, nil
}

// ListProductCategory 分页查询商品分类
func (s PmsProductCategoryService) ListProductCategory(parentId int64, pageDTO base_dto.PageDTO) (*pkg.CommonPage, error) {
	pmsProductCategory := new(models.PmsProductCategory)
	pmsProductCategory.ParentId = parentId
	page, err := pmsProductCategory.SelectListPage(s.DbFactory.GormMySQL, pageDTO.PageNum, pageDTO.PageSize)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return page, nil
}

// UpdateNavStatus 批量修改导航状态
func (s PmsProductCategoryService) UpdateNavStatus(ids []int64, navStatus int) (int64, error) {
	pmsProductCategory := new(models.PmsProductCategory)
	count, err := pmsProductCategory.UpdateNavStatusByIds(s.DbFactory.GormMySQL, ids, navStatus)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return count, nil
}

// UpdateShowStatus 批量修改显示状态
func (s PmsProductCategoryService) UpdateShowStatus(ids []int64, showStatus int) (int64, error) {
	pmsProductCategory := new(models.PmsProductCategory)
	count, err := pmsProductCategory.UpdateShowStatusByIds(s.DbFactory.GormMySQL, ids, showStatus)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return count, nil
}

// ListWithChildren 查询所有一级分类及子分类
func (s PmsProductCategoryService) ListWithChildren() ([]*models.PmsProductCategory, error) {
	pmsProductCategory := new(models.PmsProductCategory)
	list, err := pmsProductCategory.SelectList(s.DbFactory.GormMySQL)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return list, nil
}
