//	@author	hongjun500
//	@date	2023/6/21 16:11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: product 索引 pms

package es_index

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/pkg/elasticsearch"
)

type EsProduct struct {
	Id                  int64                     `json:"id" es_type:"long"`
	ProductSn           string                    `json:"product_sn" es_type:"keyword"`
	BrandId             int64                     `json:"brand_id" es_type:"long"`
	BrandName           string                    `json:"brand_name" es_type:"keyword"`
	ProductCategoryId   int64                     `json:"product_category_id" es_type:"long"`
	ProductCategoryName string                    `json:"product_category_name" es_type:"keyword"`
	Pic                 string                    `json:"pic" es_type:"keyword"`
	Name                string                    `json:"name" es_type:"text" es_analyzer:"ik_max_word"`
	SubTitle            string                    `json:"sub_title" es_type:"text" es_analyzer:"ik_max_word"`
	KeyWord             string                    `json:"key_word" es_type:"text" es_analyzer:"ik_max_word"`
	Price               string                    `json:"price" es_type:"float"`
	Sale                int                       `json:"sale" es_type:"integer"`
	NewStatus           int                       `json:"new_status" es_type:"integer"`
	RecommendStatus     int                       `json:"recommand_status" es_type:"integer"`
	Stock               int                       `json:"stock" es_type:"integer"`
	PromotionType       int                       `json:"promotion_type" es_type:"integer"`
	Sort                int                       `json:"sort" es_type:"integer"`
	AttrValues          []EsProductAttributeValue `json:"attr_value_list" es_type:"nested"`
}

type EsProductAttributeValue struct {
	Id                 int64  `json:"id" es_type:"long"`
	ProductAttributeId int64  `json:"product_id" es_type:"long"`
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
			RecommendStatus:     pmsProduct.RecommendStatus,
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
	if !elasticsearch.HasIndex(db, context.Background(), esProduct.IndexName()) {
		// 创建索引
		settings := &types.IndexSettings{NumberOfShards: "1", NumberOfReplicas: "0"}
		elasticsearch.CreateIndex(db, context.Background(), esProduct.IndexName(), settings)
	}
	esTags := elasticsearch.GetStructTag(EsProduct{})
	// 创建 mapping
	elasticsearch.PutMappingByStruct(db, context.Background(), esProduct.IndexName(), esTags)

	elasticsearch.BulkAddDocument(db, context.Background(), esProduct.IndexName(), esProducts)
}
