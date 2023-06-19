// @author hongjun500
// @date 2023/6/19 17:59
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package services

import "github.com/hongjun500/mall-go/internal/database"

type ProductSearchService struct {
	DbFactory *database.DbFactory
}

func NewProductSearchService(dbFactory *database.DbFactory) ProductSearchService {
	return ProductSearchService{DbFactory: dbFactory}
}
