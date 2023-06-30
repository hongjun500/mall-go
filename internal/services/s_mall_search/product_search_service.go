//	@author	hongjun500
//	@date	2023/6/19 17:59
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package s_mall_search

import (
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

	// 将 esProducts 导入到 es
	esProduct.PutEsProductsDocument(p.DbFactory, esProducts)
	return nil
}

// Delete 根据 id 删除商品
func (p ProductSearchService) Delete(id int64) (bool, error) {
	esProduct := new(es_index.EsProduct)
	return esProduct.DelDocument(p.DbFactory, id), nil
}
