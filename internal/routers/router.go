package routers

import (
	"github.com/hongjun500/mall-go/internal/services"
)

type CoreRouter struct {
	*UmsAdminRouter
	*UmsMenuRouter
}

type CoreRouterInterface interface {
	InitCoreRouter(service *services.CoreService, coreRouter *CoreRouter)
}

func NewCoreRouter(service *services.CoreService) *CoreRouter {
	return &CoreRouter{
		UmsAdminRouter: NewUmsAdminRouter(service.UmsAdminService),
		UmsMenuRouter:  NewUmsMenuRouter(service.UmsMenuService),
	}
}
