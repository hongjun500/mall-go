//	@author	hongjun500
//	@date	2023/6/21 16:11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: product 索引 pms

package es_index

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionscoremode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/hongjun500/mall-go/internal"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/pkg"
)

type EsProduct struct {
	Id                  int64                     `json:"id" es_type:"long"`
	ProductSn           string                    `json:"productSn" es_type:"keyword"`
	BrandId             int64                     `json:"brandId" es_type:"long"`
	BrandName           string                    `json:"brandName" es_type:"keyword"`
	ProductCategoryId   int64                     `json:"productCategoryId" es_type:"long"`
	ProductCategoryName string                    `json:"productCategoryName" es_type:"keyword"`
	Pic                 string                    `json:"pic" es_type:"keyword"`
	Name                string                    `json:"name" es_type:"text" es_analyzer:"ik_max_word"`
	SubTitle            string                    `json:"subTitle" es_type:"text" es_analyzer:"ik_max_word"`
	KeyWord             string                    `json:"keywords" es_type:"text" es_analyzer:"ik_max_word"`
	Price               string                    `json:"price" es_type:"float"`
	Sale                int                       `json:"sale" es_type:"integer"`
	NewStatus           int                       `json:"newStatus" es_type:"integer"`
	RecommandStatus     int                       `json:"recommandStatus" es_type:"integer"`
	Stock               int                       `json:"stock" es_type:"integer"`
	PromotionType       int                       `json:"promotionType" es_type:"integer"`
	Sort                int                       `json:"sort" es_type:"integer"`
	AttrValues          []EsProductAttributeValue `json:"attrValueList" es_type:"nested"`
}

type EsProductAttributeValue struct {
	Id                 int64  `json:"id" es_type:"long"`
	ProductAttributeId int64  `json:"productId" es_type:"long"`
	Value              string `json:"value" es_type:"keyword"`
	Type               int    `json:"type" es_type:"integer"`
	Name               string `json:"name" es_type:"keyword"`
}

// IndexName 索引名称
func (*EsProduct) IndexName() string {
	return "pms"
}

// ConvertEsProductFromPmsProduct 将从数据库中查询的 pmsProduct 结果 转换为 esProduct
func ConvertEsProductFromPmsProduct(pmsProducts []*models.PmsProduct) []*EsProduct {
	if pmsProducts == nil {
		return nil
	}
	var esProducts []*EsProduct
	for _, pmsProduct := range pmsProducts {
		esProduct := &EsProduct{
			Id:                  pmsProduct.Id,
			ProductSn:           pmsProduct.ProductSn,
			BrandId:             pmsProduct.BrandId,
			BrandName:           pmsProduct.BrandName,
			ProductCategoryId:   pmsProduct.ProductCategoryId,
			ProductCategoryName: pmsProduct.ProductCategoryName,
			Pic:                 pmsProduct.Pic,
			Name:                pmsProduct.Name,
			SubTitle:            pmsProduct.SubTitle,
			KeyWord:             pmsProduct.KeyWord,
			Price:               pmsProduct.Price,
			Sale:                pmsProduct.Sale,
			NewStatus:           pmsProduct.NewStatus,
			RecommandStatus:     pmsProduct.RecommendStatus,
			Stock:               pmsProduct.Stock,
			PromotionType:       pmsProduct.PromotionType,
			Sort:                pmsProduct.Sort,
		}
		var attrValues []EsProductAttributeValue
		for _, pmsProductAttributeValue := range pmsProduct.ProductAttributeValueList {
			attrValues = append(attrValues, EsProductAttributeValue{
				Id:                 pmsProductAttributeValue.Id,
				ProductAttributeId: pmsProductAttributeValue.ProductAttributeId,
				Value:              pmsProductAttributeValue.Value,
				Type:               pmsProductAttributeValue.ProductAttribute.Type,
				Name:               pmsProductAttributeValue.ProductAttribute.Name,
			})
		}
		esProduct.AttrValues = attrValues
		esProducts = append(esProducts, esProduct)
	}
	return esProducts
}

// PutEsProductsDocument 将 esProducts 存入 es 文档
func (esProduct *EsProduct) PutEsProductsDocument(db *database.DbFactory, esProducts []*EsProduct) {
	// 判断索引是否存在
	if !internal.HasIndex(db, context.Background(), esProduct.IndexName()) {
		// 创建索引
		settings := &types.IndexSettings{NumberOfShards: "1", NumberOfReplicas: "0"}
		internal.CreateIndex(db, context.Background(), esProduct.IndexName(), settings)
	}
	// 创建 mapping
	internal.PutMappingByStruct(db, context.Background(), esProduct.IndexName(), EsProduct{})

	internal.BulkAddDocument(db, context.Background(), esProduct.IndexName(), esProducts)
}

func (esProduct *EsProduct) DelDocument(db *database.DbFactory, id int64) bool {
	ok := internal.DeleteDocument(db, context.Background(), esProduct.IndexName(), id)
	return ok
}

func (esProduct *EsProduct) DelDocuments(db *database.DbFactory, ids []int64) bool {
	ok := internal.BulkDeleteDocument(db, context.Background(), esProduct.IndexName(), ids)
	return ok
}

// SearchByNameOrSubtitleOrKeyword 根据名称或副标题或关键字搜索
func (esProduct *EsProduct) SearchByNameOrSubtitleOrKeyword(db *database.DbFactory, keyword string, pageNum, pageSize int) (*pkg.CommonPage, error) {
	// or 用 bool 的 should 来实现
	query := []types.Query{
		{
			Match: map[string]types.MatchQuery{
				"name": {
					Query: keyword,
				},
			},
		},
		{
			Match: map[string]types.MatchQuery{
				"subTitle": {
					Query: keyword,
				},
			},
		},
		{
			Match: map[string]types.MatchQuery{
				"keyWord": {
					Query: keyword,
				},
			},
		},
	}
	searchRequest := &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Should: query,
			},
		},
	}
	page := internal.NewElasticSearchPage(db.Es, esProduct.IndexName(), pageNum, pageSize)
	page.SearchRequest = searchRequest
	err := page.Paginate()
	if err != nil {
		return nil, err
	}
	return page.CommonPage, nil
}

// SearchByNameOrSubtitle 根据关键字搜索名称或者副标题复合查询
func (esProduct *EsProduct) SearchByNameOrSubtitle(db *database.DbFactory, keyword string, brandId, productCategoryId int64, pageNum, pageSize, sort int) (*pkg.CommonPage, error) {
	page := internal.NewElasticSearchPage(db.Es, esProduct.IndexName(), pageNum, pageSize)
	query := types.Query{}
	filters := make([]types.Query, 0)

	if brandId != 0 || productCategoryId != 0 {
		if brandId != 0 {
			filters = append(filters, types.Query{
				Term: map[string]types.TermQuery{
					"brandId": {Value: brandId},
				},
			})
		}
		if productCategoryId != 0 {
			filters = append(filters, types.Query{
				Term: map[string]types.TermQuery{
					"productCategoryId": {Value: productCategoryId},
				},
			})
		}
	}
	if keyword == "" {
		// 没有关键字，直接查询
		query = types.Query{
			MatchAll: &types.MatchAllQuery{},
		}
	} else {
		scoreQuery := types.FunctionScoreQuery{
			Query: &types.Query{
				MatchAll: types.NewMatchAllQuery(),
			},
			ScoreMode: &functionscoremode.Sum,
			MinScore:  (*types.Float64)(some.Float64(2)),
			Functions: []types.FunctionScore{
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"name": {Query: keyword},
						},
					},
					Weight: (*types.Float64)(some.Float64(10)),
				},
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"subTitle": {Query: keyword},
						},
					},
					Weight: (*types.Float64)(some.Float64(5)),
				},
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"keywords": {Query: keyword},
						},
					},
					Weight: (*types.Float64)(some.Float64(2)),
				},
			},
		}
		queryScore := types.Query{
			FunctionScore: &scoreQuery,
		}
		query = queryScore
	}
	// 相关度排序
	var options []types.SortCombinations
	if sort == 1 {
		// 按新品从新到旧
		options = []types.SortCombinations{
			types.SortOptions{SortOptions: map[string]types.FieldSort{
				"id": {Order: &sortorder.Desc},
			}},
		}
	} else if sort == 2 {
		// 按销量从高到低
		options = []types.SortCombinations{
			types.SortOptions{SortOptions: map[string]types.FieldSort{
				"sale": {Order: &sortorder.Desc},
			}},
		}
	} else if sort == 3 {
		// 按价格从低到高
		options = []types.SortCombinations{
			types.SortOptions{SortOptions: map[string]types.FieldSort{
				"price": {Order: &sortorder.Asc},
			}},
		}
	} else if sort == 4 {
		// 按价格从低到高
		options = []types.SortCombinations{
			types.SortOptions{SortOptions: map[string]types.FieldSort{
				"price": {Order: &sortorder.Desc},
			}},
		}
	} else {
		options = []types.SortCombinations{
			types.SortOptions{Score_: &types.ScoreSort{Order: &sortorder.Desc}},
		}
	}

	page.SearchRequest = &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must:   append([]types.Query{}, query),
				Filter: filters,
			},
		},
		Sort: options,
	}
	err := page.Paginate()
	if err != nil {
		return nil, err
	}
	return page.CommonPage, nil
}
