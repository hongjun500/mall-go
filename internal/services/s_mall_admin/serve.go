// @author hongjun500
// @date 2023/6/28 15:11
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/internal/database"
)

type CoreAdminService struct {
	UmsAdminService
	UmsMenuService
	UmsResourceCategoryService
	UmsResourceService
	UmsRoleService
	UmsMemberLevelService
	PmsProductCategoryService
}

func NewCoreAdminService(dbFactory *database.DbFactory) *CoreAdminService {
	return &CoreAdminService{
		UmsAdminService:            NewUmsAdminService(dbFactory),
		UmsMenuService:             NewUmsMenuService(dbFactory),
		UmsResourceCategoryService: NewUmsResourceCategoryService(dbFactory),
		UmsResourceService:         NewUmsResourceService(dbFactory),
		UmsRoleService:             NewUmsRoleService(dbFactory),
		UmsMemberLevelService:      NewUmsMemberLevelService(dbFactory),
		PmsProductCategoryService:  NewPmsProductCategoryService(dbFactory),
	}
}
