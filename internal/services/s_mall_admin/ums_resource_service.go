//	@author	hongjun500
//	@date	2023/6/13 11:14
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
//	@Description:

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/ums_admin_dto"
	"github.com/hongjun500/mall-go/internal/services"
	"github.com/hongjun500/mall-go/pkg"
)

type UmsResourceService struct {
	DbFactory *database.DbFactory
}

func NewUmsResourceService(dbFactory *database.DbFactory) UmsResourceService {
	return UmsResourceService{DbFactory: dbFactory}
}

// UmsResourceCreate 添加后台资源
func (s UmsResourceService) UmsResourceCreate(dto ums_admin_dto.UmsResourceCreateDTO) (int64, error) {
	// var umsResource *models.UmsResource
	// 上面这种直接赋值有问题，空指针, 如果是以 umsResource := &models.UmsResource{} 的方式就不会有问题, 或者使用 new() 函数
	umsResource := new(models.UmsResource)
	umsResource.Name = dto.Name
	umsResource.Url = dto.Url
	umsResource.Description = dto.Description
	umsResource.CategoryId = dto.CategoryId
	rows, err := umsResource.Insert(s.DbFactory.GormMySQL)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// UmsResourceUpdate 修改后台资源
func (s UmsResourceService) UmsResourceUpdate(id int64, dto ums_admin_dto.UmsResourceCreateDTO) (int64, error) {
	m := new(models.UmsResource)
	m.Name = dto.Name
	m.Url = dto.Url
	m.Description = dto.Description
	m.CategoryId = dto.CategoryId
	rows, err := m.Update(s.DbFactory.GormMySQL, id)
	if err != nil {
		return 0, err
	}
	services.DelResourceListByResource(s.DbFactory, id)
	return rows, nil
}

// UmsResourceItem 根据ID获取资源详情
func (s UmsResourceService) UmsResourceItem(id int64) (*models.UmsResource, error) {
	m := new(models.UmsResource)
	m.Id = id
	umsResource, err := m.SelectUmsResourceById(s.DbFactory.GormMySQL, m.Id)
	if err != nil {
		return nil, err
	}
	return umsResource, nil
}

// UmsResourceDelete 根据ID删除后台资源
func (s UmsResourceService) UmsResourceDelete(id int64) (int64, error) {
	m := new(models.UmsResource)
	m.Id = id
	rows, err := m.Delete(s.DbFactory.GormMySQL, m.Id)
	if err != nil {
		return 0, err
	}
	services.DelResourceListByResource(s.DbFactory, id)
	return rows, nil
}

// UmsResourcePageList 分页模糊查询后台资源
func (s UmsResourceService) UmsResourcePageList(dto ums_admin_dto.UmsResourcePageListDTO) (*pkg.CommonPage, error) {
	m := new(models.UmsResource)
	page, err := m.SelectPage(s.DbFactory.GormMySQL, dto.CategoryId, dto.NameKeyword, dto.UrlKeyword, dto.PageNum, dto.PageSize)
	if err != nil {
		return nil, err
	}
	return page, nil
}

// UmsResourceList 查询所有后台资源
/*func (s UmsResourceService) UmsResourceList(context *gin.Context) {
	m := new(models.UmsResource)
	list, err := m.SelectAll(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(context, gin_common.UnknownError)
		return
	}
	gin_common.CreateSuccess(context, list)
}*/
