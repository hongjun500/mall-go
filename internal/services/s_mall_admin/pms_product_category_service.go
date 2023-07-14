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
	"github.com/hongjun500/mall-go/internal/request/ums_admin_dto"
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
