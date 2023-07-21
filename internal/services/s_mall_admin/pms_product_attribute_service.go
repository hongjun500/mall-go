// @author hongjun500
// @date 2023/7/21 10:24
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
	"github.com/hongjun500/mall-go/pkg"
)

type PmsProductAttributeService struct {
	DbFactory *database.DbFactory
}

func NewPmsProductAttributeService(dbFactory *database.DbFactory) PmsProductAttributeService {
	return PmsProductAttributeService{
		DbFactory: dbFactory,
	}
}

// ListPage 根据分类分页获取商品属性
func (s PmsProductAttributeService) ListPage(categoryId int64, categoryType, pageNum, pageSize int) (*pkg.CommonPage, error) {
	var productAttribute models.PmsProductAttribute
	page, err := productAttribute.SelectListByPage(s.DbFactory.GormMySQL, categoryId, categoryType, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return page, nil
}

// Create 创建商品属性
func (s PmsProductAttributeService) Create(productAttributeParam admin_dto.PmsProductAttributeDTO) (int64, error) {
	// todo 这里需要用事务
	productAttribute := new(models.PmsProductAttribute)
	productAttribute.ProductAttributeCategoryId = productAttributeParam.ProductAttributeCategoryId
	productAttribute.Name = productAttributeParam.Name
	productAttribute.SelectType = productAttributeParam.SelectType
	productAttribute.InputType = productAttributeParam.InputType
	productAttribute.InputList = productAttributeParam.InputList
	productAttribute.Sort = productAttributeParam.Sort
	productAttribute.FilterType = productAttributeParam.FilterType
	productAttribute.SearchType = productAttributeParam.SearchType
	productAttribute.RelatedStatus = productAttributeParam.RelatedStatus
	productAttribute.HandAddStatus = productAttributeParam.HandAddStatus
	productAttribute.Type = productAttributeParam.Type
	count, _ := productAttribute.Insert(s.DbFactory.GormMySQL)
	attributeCategory := new(models.PmsProductAttributeCategory)
	attributeCategory, _ = attributeCategory.SelectById(s.DbFactory.GormMySQL, productAttribute.ProductAttributeCategoryId)
	if productAttribute.Type == 0 {
		attributeCategory.AttributeCount = attributeCategory.AttributeCount + int(count)
		s.DbFactory.GormMySQL.Model(attributeCategory).Update("attribute_count", attributeCategory.AttributeCount)
	} else if productAttribute.Type == 1 {
		attributeCategory.ParamCount = attributeCategory.ParamCount + int(count)
		s.DbFactory.GormMySQL.Model(attributeCategory).Update("param_count", attributeCategory.ParamCount)
	}
	return count, nil
}

// Update 修改商品属性信息
func (s PmsProductAttributeService) Update(id int64, productAttributeParam admin_dto.PmsProductAttributeDTO) (int64, error) {
	attr := new(models.PmsProductAttribute)
	attr.Name = productAttributeParam.Name
	attr.SelectType = productAttributeParam.SelectType
	attr.InputType = productAttributeParam.InputType
	attr.InputList = productAttributeParam.InputList
	attr.Sort = productAttributeParam.Sort
	attr.FilterType = productAttributeParam.FilterType
	attr.SearchType = productAttributeParam.SearchType
	attr.RelatedStatus = productAttributeParam.RelatedStatus
	attr.HandAddStatus = productAttributeParam.HandAddStatus
	attr.Type = productAttributeParam.Type
	count, err := attr.UpdateById(s.DbFactory.GormMySQL, id)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return count, nil
}

// GetItem 根据id获取商品属性信息
func (s PmsProductAttributeService) GetItem(id int64) (*models.PmsProductAttribute, error) {
	attr := new(models.PmsProductAttribute)
	attr, err := attr.SelectById(s.DbFactory.GormMySQL, id)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return attr, nil
}

// Delete 批量删除商品属性
func (s PmsProductAttributeService) Delete(ids []int64) (int64, error) {
	// todo 这里需要用事务
	if len(ids) == 0 {
		return 0, nil
	}
	attr := new(models.PmsProductAttribute)
	attr, _ = attr.SelectById(s.DbFactory.GormMySQL, ids[0])
	aType := attr.Type
	attrCategory := new(models.PmsProductAttributeCategory)
	attrCategory, _ = attrCategory.SelectById(s.DbFactory.GormMySQL, attr.ProductAttributeCategoryId)
	count, _ := attr.DeleteByIds(s.DbFactory.GormMySQL, ids)
	if aType == 0 {
		if attrCategory.AttributeCount >= int(count) {
			attrCategory.AttributeCount = attrCategory.AttributeCount - int(count)
		} else {
			attrCategory.AttributeCount = 0
		}
		s.DbFactory.GormMySQL.Model(attrCategory).Update("attribute_count", attrCategory.AttributeCount)
	} else if aType == 1 {
		if attrCategory.ParamCount >= int(count) {
			attrCategory.ParamCount = attrCategory.ParamCount - int(count)
		} else {
			attrCategory.ParamCount = 0
		}
		s.DbFactory.GormMySQL.Model(attrCategory).Update("param_count", attrCategory.ParamCount)
	}
	return count, nil
}

// ListFromProductAttrInfo 获取商品分类对应属性列表
func (s PmsProductAttributeService) ListFromProductAttrInfo(productCategoryId int64) ([]*models.ProductAttrInfo, error) {
	var productAttribute models.PmsProductAttribute
	list, err := productAttribute.SelectListFromProductAttrInfo(s.DbFactory.GormMySQL, productCategoryId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
