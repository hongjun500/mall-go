// @author hongjun500
// @date 2023/7/14 13:45
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package r_mall_admin

import "github.com/hongjun500/mall-go/internal/services/s_mall_admin"

type PmsProductCategoryRouter struct {
	s_mall_admin.PmsProductCategoryService
}

func NewPmsProductCategoryRouter(service s_mall_admin.PmsProductCategoryService) *PmsProductCategoryRouter {
	return &PmsProductCategoryRouter{
		PmsProductCategoryService: service,
	}
}
