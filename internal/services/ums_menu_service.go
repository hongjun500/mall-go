// @author hongjun500
// @date 2023/6/11
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台菜单相关服务

package services

import "github.com/hongjun500/mall-go/internal/database"

type UmsMenuService struct {
	DbFactory *database.DbFactory
}

func NewUmsMenuService(dbFactory *database.DbFactory) UmsMenuService {
	return UmsMenuService{DbFactory: dbFactory}
}
