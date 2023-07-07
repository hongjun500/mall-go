// @author hongjun500
// @date 2023/7/6 10:47
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package es_index_dto

import "github.com/hongjun500/mall-go/internal/request/base_dto"

type (
	// ProductSearchDTO 商品搜索的数据传输对象
	ProductSearchDTO struct {
		// 关键字
		Keyword string `json:"keyword" form:"keyword" query:"keyword"`
		// 分类id
		BrandId int64 `json:"brandId" form:"brandId" query:"brandId"`
		// 分类id
		ProductCategoryId int64 `json:"productCategoryId" form:"productCategoryId" query:"productCategoryId"`
		// 排序方式 0->按相关度；1->按新品；2->按销量；3->价格从低到高；4->价格从高到低
		Sort int `json:"sort" form:"sort" query:"sort" default:"0"`
		base_dto.PageDTO
	}
)
