// @author hongjun500
// @date 2023/7/26 14:10
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

type PmsBrandService struct {
	DbFactory *database.DbFactory
}

func NewPmsBrandService(dbFactory *database.DbFactory) PmsBrandService {
	return PmsBrandService{DbFactory: dbFactory}
}

// ListAll 获取所有品牌
func (service PmsBrandService) ListAll() ([]*models.PmsBrand, error) {
	var brand models.PmsBrand
	all, err := brand.SelectAll(service.DbFactory.GormMySQL)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return all, nil
}

// Create 创建品牌
func (service PmsBrandService) Create(PmsBrandCreateDto admin_dto.PmsBrandDTO) (int64, error) {
	brand := &models.PmsBrand{
		Name:          PmsBrandCreateDto.Name,
		FirstLetter:   PmsBrandCreateDto.FirstLetter,
		Sort:          PmsBrandCreateDto.Sort,
		FactoryStatus: PmsBrandCreateDto.FactoryStatus,
		ShowStatus:    PmsBrandCreateDto.ShowStatus,
		Logo:          PmsBrandCreateDto.Logo,
		BigPic:        PmsBrandCreateDto.BigPic,
		BrandStory:    PmsBrandCreateDto.BrandStory,
	}
	if brand.FirstLetter == "" {
		brand.FirstLetter = string(brand.Name[0])
	}
	rows, err := brand.Insert(service.DbFactory.GormMySQL)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}

// Update 更新品牌
func (service PmsBrandService) Update(id int64, PmsBrandUpdateDto admin_dto.PmsBrandDTO) (int64, error) {
	brand := &models.PmsBrand{}
	brand.Id = id
	brand.Name = PmsBrandUpdateDto.Name
	brand.FirstLetter = PmsBrandUpdateDto.FirstLetter
	brand.Sort = PmsBrandUpdateDto.Sort
	brand.FactoryStatus = PmsBrandUpdateDto.FactoryStatus
	brand.ShowStatus = PmsBrandUpdateDto.ShowStatus
	brand.Logo = PmsBrandUpdateDto.Logo
	brand.BigPic = PmsBrandUpdateDto.BigPic
	brand.BrandStory = PmsBrandUpdateDto.BrandStory
	if brand.FirstLetter == "" {
		brand.FirstLetter = string(brand.Name[0])
	}
	// 更新品牌时需要更新商品的品牌名称
	product := new(models.PmsProduct)
	product.BrandName = brand.Name
	_, _ = product.UpdateBrandNameByBrandId(service.DbFactory.GormMySQL, id, brand.Name)
	rows, err := brand.Update(service.DbFactory.GormMySQL)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}

// Delete 删除品牌
func (service PmsBrandService) Delete(id int64) (int64, error) {
	brand := &models.PmsBrand{}
	brand.Id = id
	rows, err := brand.Delete(service.DbFactory.GormMySQL)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}

// DeleteBatch 批量删除品牌
func (service PmsBrandService) DeleteBatch(ids []int64) (int64, error) {
	brand := &models.PmsBrand{}
	rows, err := brand.DeleteBatch(service.DbFactory.GormMySQL, ids)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}

// ListBrand 分页查询品牌
func (service PmsBrandService) ListBrand(keyword string, showStatus, pageNum, pageSize int) (*pkg.CommonPage, error) {
	brand := &models.PmsBrand{}
	page, err := brand.SelectPage(service.DbFactory.GormMySQL, keyword, showStatus, pageNum, pageSize)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return page, nil
}

// GetBrand 获取单个品牌
func (service PmsBrandService) GetBrand(id int64) (*models.PmsBrand, error) {
	brand := &models.PmsBrand{}
	brand, err := brand.SelectById(service.DbFactory.GormMySQL, id)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return brand, nil
}

// UpdateShowStatus 批量更新显示状态
func (service PmsBrandService) UpdateShowStatus(ids []int64, showStatus int) (int64, error) {
	brand := &models.PmsBrand{}
	rows, err := brand.UpdateShowStatusOrFactoryStatus(service.DbFactory.GormMySQL, ids, 0, showStatus)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}

// UpdateFactoryStatus 批量更新厂家制造商状态
func (service PmsBrandService) UpdateFactoryStatus(ids []int64, factoryStatus int) (int64, error) {
	brand := &models.PmsBrand{}
	rows, err := brand.UpdateShowStatusOrFactoryStatus(service.DbFactory.GormMySQL, ids, 1, factoryStatus)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return rows, nil
}
