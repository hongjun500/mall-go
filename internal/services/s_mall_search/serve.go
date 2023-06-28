// @author hongjun500
// @date 2023/6/28 15:19
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package s_mall_search

import "github.com/hongjun500/mall-go/internal/database"

type CoreSearchService struct {
	ProductSearchService
}

func NewCoreSearchService(dbFactory *database.DbFactory) *CoreSearchService {
	return &CoreSearchService{
		ProductSearchService: NewProductSearchService(dbFactory),
	}
}
