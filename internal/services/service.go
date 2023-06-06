package services

import "github.com/hongjun500/mall-go/internal/database"

type CoreService struct {
	*UmsAdminService
}

func NewCoreService(dbFactory *database.DbFactory) *CoreService {
	return &CoreService{
		UmsAdminService: NewUmsAdminService(dbFactory),
	}
}
