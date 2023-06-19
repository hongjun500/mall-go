package services

import "github.com/hongjun500/mall-go/internal/database"

type CoreAdminService struct {
	UmsAdminService
	UmsMenuService
	UmsResourceCategoryService
	UmsResourceService
	UmsRoleService
	UmsMemberLevelService
}

type CoreSearchService struct {
	ProductSearchService
}

func NewCoreAdminService(dbFactory *database.DbFactory) *CoreAdminService {
	return &CoreAdminService{
		UmsAdminService:            NewUmsAdminService(dbFactory),
		UmsMenuService:             NewUmsMenuService(dbFactory),
		UmsResourceCategoryService: NewUmsResourceCategoryService(dbFactory),
		UmsResourceService:         NewUmsResourceService(dbFactory),
		UmsRoleService:             NewUmsRoleService(dbFactory),
		UmsMemberLevelService:      NewUmsMemberLevelService(dbFactory),
	}
}

func NewCoreSearchService(dbFactory *database.DbFactory) *CoreSearchService {
	return &CoreSearchService{
		ProductSearchService: NewProductSearchService(dbFactory),
	}
}
