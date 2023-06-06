package routers

import (
	"github.com/hongjun500/mall-go/internal/services"
)

type CoreRouter struct {
	*UmsAdminRouter
}

func NewCoreRouter(service *services.CoreService) *CoreRouter {
	return &CoreRouter{
		UmsAdminRouter: CreateUmsAdminRouter(service.UmsAdminService),
	}
}
