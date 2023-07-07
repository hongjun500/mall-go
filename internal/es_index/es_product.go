//	@author	hongjun500
//	@date	2023/6/21 16:11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: product 索引 pms

package es_index

import (
	"fmt"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/functionscoremode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/pkg/convert"
)

// EsProduct 商品 索引
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
	AttrValues          []EsProductAttributeValue `json:"attrValues" es_type:"nested"`
}

// EsProductAttributeValue 商品属性值 索引
type EsProductAttributeValue struct {
	Id                 int64  `json:"id" es_type:"long"`
	ProductAttributeId int64  `json:"productAttributeId" es_type:"long"`
	Value              string `json:"value" es_type:"keyword"`
	Type               int    `json:"type" es_type:"integer"`
	Name               string `json:"name" es_type:"keyword"`
}

// EsProductRelatedInfo 商品属性相关，用于 Aggregation 查询之后的结果
type EsProductRelatedInfo struct {
	BrandNames           []string      `json:"brandNames"`
	ProductCategoryNames []string      `json:"productCategoryNames"`
	ProductAttrs         []ProductAttr `json:"productAttrs"`
}

type ProductAttr struct {
	AttrId     int64    `json:"attrId"`
	AttrName   string   `json:"attrName"`
	AttrValues []string `json:"attrValues"`
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

// SearchByNameOrSubtitleOrKeyword 根据名称或副标题或关键字搜索 构建DSL语句
func (esProduct *EsProduct) SearchByNameOrSubtitleOrKeyword() (*types.Query, error) {
	// or 用 bool 的 should 来实现
	query := []types.Query{
		{
			Match: map[string]types.MatchQuery{
				"name": {
					Query: esProduct.KeyWord,
				},
			},
		},
		{
			Match: map[string]types.MatchQuery{
				"subTitle": {
					Query: esProduct.KeyWord,
				},
			},
		},
		{
			Match: map[string]types.MatchQuery{
				"keyWord": {
					Query: esProduct.KeyWord,
				},
			},
		},
	}
	DSL := &types.Query{
		Bool: &types.BoolQuery{
			Should: query,
		},
	}
	fmt.Println("DSL", DSL)
	return DSL, nil
	/*

		page := internal.NewElasticSearchPage(es, esProduct.IndexName(), pageNum, pageSize)
		page.SearchRequest = searchRequest
		err := page.Paginate()
		if err != nil {
			return nil, err
		}
		return page.CommonPage, nil*/
}

// SearchByNameOrSubtitle 根据关键字搜索名称或者副标题复合查询并排序 构建DSL语句
func (esProduct *EsProduct) SearchByNameOrSubtitle(sort int) (*types.Query, types.Sort, error) {
	// page := internal.NewElasticSearchPage(es, esProduct.IndexName(), pageNum, pageSize)
	query := types.Query{}
	filters := make([]types.Query, 0)

	if esProduct.BrandId != 0 || esProduct.ProductCategoryId != 0 {
		if esProduct.BrandId != 0 {
			filters = append(filters, types.Query{
				Term: map[string]types.TermQuery{
					"brandId": {Value: esProduct.BrandId},
				},
			})
		}
		if esProduct.ProductCategoryId != 0 {
			filters = append(filters, types.Query{
				Term: map[string]types.TermQuery{
					"productCategoryId": {Value: esProduct.ProductCategoryId},
				},
			})
		}
	}
	if esProduct.KeyWord == "" {
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
							"name": {Query: esProduct.KeyWord},
						},
					},
					Weight: (*types.Float64)(some.Float64(10)),
				},
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"subTitle": {Query: esProduct.KeyWord},
						},
					},
					Weight: (*types.Float64)(some.Float64(5)),
				},
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"keywords": {Query: esProduct.KeyWord},
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

	/*page.SearchRequest = &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must:   append([]types.Query{}, query),
				Filter: filters,
			},
		},
		Sort: options,
	}*/
	DSL := &types.Query{
		Bool: &types.BoolQuery{
			Must:   append([]types.Query{}, query),
			Filter: filters,
		},
	}
	Sort := options
	fmt.Println("DSL", convert.AnyToJson(DSL))
	return DSL, Sort, nil
	/*err := page.Paginate()
	if err != nil {
		return nil, err
	}
	return page.CommonPage, nil*/
}

// SearchById 根据商品id搜索 用于推荐同类商品 构建DSL语句
func (esProduct *EsProduct) SearchById() (*types.Query, error) {
	if esProduct.Id == 0 {
		return nil, nil
	} else {
		keyword := esProduct.Name
		brandId := esProduct.BrandId
		productCategoryId := esProduct.ProductCategoryId
		// 根据商品标题、品牌、分类进行搜索
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
					Weight: (*types.Float64)(some.Float64(8)),
				},
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"subTitle": {Query: keyword},
						},
					},
					Weight: (*types.Float64)(some.Float64(2)),
				},
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"keywords": {Query: keyword},
						},
					},
					Weight: (*types.Float64)(some.Float64(2)),
				},
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"brandId": {Query: strconv.FormatInt(brandId, 10)},
						},
					},
					Weight: (*types.Float64)(some.Float64(5)),
				},
				{
					Filter: &types.Query{
						Match: map[string]types.MatchQuery{
							"productCategoryId": {Query: strconv.FormatInt(productCategoryId, 10)},
						},
					},
					Weight: (*types.Float64)(some.Float64(3)),
				},
			},
		}

		boolQuery := types.NewBoolQuery()
		boolQuery.MustNot = []types.Query{
			{
				Term: map[string]types.TermQuery{
					"id": {Value: strconv.FormatInt(esProduct.Id, 10)},
				},
			},
		}
		DSL := &types.Query{
			Bool: &types.BoolQuery{
				Filter: []types.Query{
					{
						Bool: boolQuery,
					},
				},
				Must: []types.Query{
					{
						FunctionScore: &scoreQuery,
					},
				},
			},
		}
		fmt.Println("DSL", convert.AnyToJson(DSL))
		return DSL, nil
	}
}

// SearchRelatedInfo 根据关键字搜索 (聚合搜索品牌、分类、属性)
func (esProduct *EsProduct) SearchRelatedInfo() (*types.Query, map[string]types.Aggregations, error) {
	query := types.NewQuery()
	if esProduct.KeyWord == "" {
		query = &types.Query{
			MatchAll: types.NewMatchAllQuery(),
		}
	} else {
		query = &types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Fields: []string{"name", "subTitle", "keywords"},
				Query:  esProduct.KeyWord,
			},
		}
	}

	// 聚合搜索品牌、分类、属性

	Aggregations := map[string]types.Aggregations{
		"brandNames": {
			Terms: &types.TermsAggregation{
				Field: some.String("brandName"),
			},
		},
		"productCategoryNames": {
			Terms: &types.TermsAggregation{
				Field: some.String("productCategoryName"),
			},
		},
		"allAttrValues": {
			Nested: &types.NestedAggregation{
				Path: some.String("attrValues"),
			},
			Aggregations: map[string]types.Aggregations{
				"productAttrs": {
					Filter: &types.Query{
						Bool: &types.BoolQuery{
							Filter: []types.Query{
								{
									Term: map[string]types.TermQuery{
										"attrValues.type": {Value: 1},
									},
								},
							},
						},
					},
					Aggregations: map[string]types.Aggregations{
						"attrIds": {
							Terms: &types.TermsAggregation{
								Field: some.String("attrValues.productAttributeId"),
							},
							Aggregations: map[string]types.Aggregations{
								"dds": {
									Terms: &types.TermsAggregation{
										Field: some.String("attrValues.value"),
									},
								},
								"attrNames": {
									Terms: &types.TermsAggregation{
										Field: some.String("attrValues.name"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	DSL := query
	return DSL, Aggregations, nil
}

func ConvertProductRelatedInfo(aggregations map[string]types.Aggregate) (*EsProductRelatedInfo, error) {
	esProductRelatedInfo := new(EsProductRelatedInfo)

	handleBucket := func(bucket []types.StringTermsBucket) []string {
		names := make([]string, 0, len(bucket))
		for _, termsBucket := range bucket {
			names = append(names, termsBucket.Key.(string))
		}
		return names
	}
	brandNames := aggregations["brandNames"].(*types.StringTermsAggregate)
	productCategoryNames := aggregations["productCategoryNames"].(*types.StringTermsAggregate)
	esProductRelatedInfo.BrandNames = handleBucket(brandNames.Buckets.([]types.StringTermsBucket))
	esProductRelatedInfo.ProductCategoryNames = handleBucket(productCategoryNames.Buckets.([]types.StringTermsBucket))

	allAttrValues := aggregations["allAttrValues"].(*types.NestedAggregate)

	log.Printf("aggregations: %v", brandNames)
	log.Printf("productCategoryNames: %v", productCategoryNames)
	log.Printf("allAttrValues: %v", allAttrValues)
	data := make([]any, 0, len(aggregations)) // 预分配足够的容量
	fmt.Println(len(data))
	fmt.Println(aggregations)
	return esProductRelatedInfo, nil
}
