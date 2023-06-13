package services

import "github.com/hongjun500/mall-go/internal/database"

type CoreService struct {
	UmsAdminService
	UmsMenuService
	UmsResourceCategoryService
	UmsResourceService
	UmsRoleService
}

func NewCoreService(dbFactory *database.DbFactory) *CoreService {
	return &CoreService{
		UmsAdminService:            NewUmsAdminService(dbFactory),
		UmsMenuService:             NewUmsMenuService(dbFactory),
		UmsResourceCategoryService: NewUmsResourceCategoryService(dbFactory),
		UmsResourceService:         NewUmsResourceService(dbFactory),
		UmsRoleService:             NewUmsRoleService(dbFactory),
	}
}
