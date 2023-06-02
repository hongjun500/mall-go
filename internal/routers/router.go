package routers

import (
	"github.com/hongjun500/mall-go/internal/services"
)

type CoreRouter struct {
	*UmsAdminRouter
}

func InitCoreRouter(service *services.CoreService) *CoreRouter {
	return &CoreRouter{
		UmsAdminRouter: NewUmsAdminRouter(service.UmsAdminService),
	}
}
