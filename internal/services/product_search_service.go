//	@author	hongjun500
//	@date	2023/6/19 17:59
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/es_index"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
)

type ProductSearchService struct {
	DbFactory *database.DbFactory
}

func NewProductSearchService(dbFactory *database.DbFactory) ProductSearchService {
	return ProductSearchService{DbFactory: dbFactory}
}

// ImportAll 将数据库中的商品信息导入到 es
//
//	@Summary		将数据库中的商品信息导入到 es
//	@Description	将数据库中的商品信息导入到 es
//	@Tags			搜索商品管理
//	@Accept			application/json
//	@Produce		application/json
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/product/importAll [post]
func (p *ProductSearchService) ImportAll(context *gin.Context) {
	var product models.PmsProduct
	pmsProducts, err := product.GetProductInfoById(p.DbFactory, 0)
	if err != nil {
		gin_common.CreateFail(context, gin_common.DatabaseError)
		return
	}
	var esProduct es_index.EsProduct
	esProducts := es_index.ConvertEsProductFromPmsProduct(pmsProducts)

	// 将 esProducts 导入到 es
	esProduct.PutEsProductsDocument(p.DbFactory, esProducts)
	gin_common.Create(context)
}
