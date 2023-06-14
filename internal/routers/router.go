package routers

import (
	"github.com/hongjun500/mall-go/internal/services"
)

type CoreRouter struct {
	*UmsAdminRouter
	*UmsMenuRouter
	*UmsResourceCategoryRouter
	*UmsResourceRouter
	*UmsRoleRouter
	*UmsMemberLevelRouter
}

type CoreRouterInterface interface {
	InitCoreRouter(service *services.CoreService, coreRouter *CoreRouter)
}

func NewCoreRouter(service *services.CoreService) *CoreRouter {
	return &CoreRouter{
		UmsAdminRouter:            NewUmsAdminRouter(service.UmsAdminService),
		UmsMenuRouter:             NewUmsMenuRouter(service.UmsMenuService),
		UmsResourceCategoryRouter: NewUmsResourceCategoryRouter(service.UmsResourceCategoryService),
		UmsResourceRouter:         NewUmsResourceRouter(service.UmsResourceService),
		UmsRoleRouter:             NewUmsRoleRouter(service.UmsRoleService),
		UmsMemberLevelRouter:      NewUmsMemberLevelRouter(service.UmsMemberLevelService),
	}
}
