//	@author	hongjun500
//	@date	2023/6/19 17:59
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package s_mall_search

import (
	"context"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hongjun500/mall-go/internal"

	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/es_index"
	"github.com/hongjun500/mall-go/internal/models"
)

type ProductSearchService struct {
	DbFactory *database.DbFactory
}

func NewProductSearchService(dbFactory *database.DbFactory) ProductSearchService {
	return ProductSearchService{DbFactory: dbFactory}
}

// ImportAll 将所有商品导入到 es
func (p ProductSearchService) ImportAll() error {
	var product models.PmsProduct
	pmsProducts, err := product.GetProductInfoById(p.DbFactory, 0)
	if err != nil {
		return err
	}
	var esProduct es_index.EsProduct
	esProducts := es_index.ConvertEsProductFromPmsProduct(pmsProducts)
	// 判断索引是否存在
	if !internal.HasIndex(p.DbFactory.Es, context.Background(), esProduct.IndexName()) {
		// 创建索引
		settings := &types.IndexSettings{NumberOfShards: "1", NumberOfReplicas: "0"}
		internal.CreateIndex(p.DbFactory.Es, context.Background(), esProduct.IndexName(), settings)
	}
	// 创建 mapping
	internal.PutMappingByStruct(p.DbFactory.Es, context.Background(), esProduct.IndexName(), es_index.EsProduct{})

	internal.BulkAddDocument(p.DbFactory.Es, context.Background(), esProduct.IndexName(), esProducts)

	return nil
}

// Delete 根据 id 删除商品
func (p ProductSearchService) Delete(id int64) (bool, error) {
	esProduct := new(es_index.EsProduct)

	ok := internal.DeleteDocument(p.DbFactory.Es, context.Background(), esProduct.IndexName(), id)
	return ok, nil
}

// DeleteBatch 根据 id 批量删除商品
func (p ProductSearchService) DeleteBatch(ids []int64) (bool, error) {
	esProduct := new(es_index.EsProduct)
	ok := internal.BulkDeleteDocument(p.DbFactory.Es, context.Background(), esProduct.IndexName(), ids)
	return ok, nil
}

// Create 根据id创建商品
func (p ProductSearchService) Create(id int64) (*es_index.EsProduct, error) {
	esProduct := new(es_index.EsProduct)
	var product models.PmsProduct
	pmsProducts, err := product.GetProductInfoById(p.DbFactory, id)
	if err != nil {
		return esProduct, err
	}
	esProducts := es_index.ConvertEsProductFromPmsProduct(pmsProducts)
	esProduct = esProducts[0]
	internal.CreateDocument(p.DbFactory.Es, context.Background(), esProduct.IndexName(), strconv.Itoa(int(esProduct.Id)), esProduct)
	return esProduct, nil
}

// PageSearchByName 根据关键字搜索名称或者副标题
func (p ProductSearchService) PageSearchByName(keyword string, pageNum int, pageSize int, sort int) ([]*es_index.EsProduct, int64, error) {

	return nil, 0, nil
}
