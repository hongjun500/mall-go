//	@author	hongjun500
//	@date	2023/6/13 10:47
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台资源分类服务

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/request/ums_admin_dto"
)

type UmsResourceCategoryService struct {
	DbFactory *database.DbFactory
}

func NewUmsResourceCategoryService(dbFactory *database.DbFactory) UmsResourceCategoryService {
	return UmsResourceCategoryService{DbFactory: dbFactory}
}

// UmsResourceCategoryList 查询所有后台资源分类
func (s UmsResourceCategoryService) UmsResourceCategoryList() ([]*models.UmsResourceCategory, error) {
	var umsResourceCategory models.UmsResourceCategory
	list, err := umsResourceCategory.SelectAll(s.DbFactory.GormMySQL)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// UmsResourceCategoryCreate 添加后台资源分类
func (s UmsResourceCategoryService) UmsResourceCategoryCreate(dto ums_admin_dto.UmsResourceCategoryCreateDTO) (int64, error) {
	umsResourceCategory := new(models.UmsResourceCategory)
	umsResourceCategory.Name = dto.Name
	umsResourceCategory.Sort = dto.Sort
	rows, err := umsResourceCategory.Insert(s.DbFactory.GormMySQL)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// UmsResourceCategoryUpdate 修改后台资源分类
func (s UmsResourceCategoryService) UmsResourceCategoryUpdate(pathVariableDTO base_dto.PathVariableDTO, dto ums_admin_dto.UmsResourceCategoryCreateDTO) (int64, error) {
	umsResourceCategory := new(models.UmsResourceCategory)
	umsResourceCategory.Name = dto.Name
	umsResourceCategory.Sort = dto.Sort
	rows, err := umsResourceCategory.Update(s.DbFactory.GormMySQL, pathVariableDTO.Id)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// UmsResourceCategoryDelete 删除后台资源分类
func (s UmsResourceCategoryService) UmsResourceCategoryDelete(id int64) (int64, error) {
	umsResourceCategory := new(models.UmsResourceCategory)
	rows, err := umsResourceCategory.Delete(s.DbFactory.GormMySQL, id)
	if err != nil {
		return 0, err
	}
	return rows, nil
}
