// @author hongjun500
// @date 2023/6/21 15:52
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package common

import (
	"strconv"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hongjun500/mall-go/internal/es_index"
	"github.com/hongjun500/mall-go/pkg/convert"
	"github.com/hongjun500/mall-go/pkg/elasticsearch"
)

var (
	index = "pms_product"
)

type Product struct {
	Id        int64   `json:"id" es_type:"long"`
	Name      string  `json:"name" es_type:"text" es_analyzer:"ik_max_word"`
	Price     float64 `json:"price" es_type:"float"`
	Count     int64   `json:"count" es_type:"long"`
	BrandName string  `json:"brand_name" es_type:"keyword"`
}

func TestDeleteIndex(t *testing.T) {
	t.Log("delete index success: ", elasticsearch.DeleteIndex(dbFactory, ctx, "pms"))
}

func TestHasIndex(t *testing.T) {
	t.Log("index is exist: ", elasticsearch.HasIndex(dbFactory, ctx, index))
}

func TestCreateIndex(t *testing.T) {
	settings := &types.IndexSettings{NumberOfShards: "1", NumberOfReplicas: "0"}

	t.Log("create index success: ", elasticsearch.CreateIndex(dbFactory, ctx, index, settings))
}

func TestPutMapping(t *testing.T) {
	_ = elasticsearch.GetStructTag(es_index.EsProduct{})
	esProduct := es_index.EsProduct{}
	t.Log("put mapping success: ", elasticsearch.PutMappingByStruct(dbFactory, ctx, index, esProduct))
}

func TestBulkIndex(t *testing.T) {
	product1 := Product{
		Id:        1,
		Name:      "商品1",
		Price:     100.0,
		Count:     100,
		BrandName: "apple",
	}
	product2 := product1
	product2.Id = 2
	product2.Name = "商品2"
	product2.BrandName = "google"

	var ps []Product
	for i := 0; i < 1000; i++ {
		var product Product
		product.Id = int64(i + 2)
		product.Name = "商品" + strconv.Itoa(i+2)
		product.Price = 100.0 + float64(i+2)
		product.Count = 100 + int64(i+2)
		product.BrandName = "brand" + strconv.Itoa(i+2)
		ps = append(ps, product)
	}
	// products := []Product{product1, product2}
	// t.Log("bulk index success: ", elasticsearch.BulkAddDocument(dbFactory, ctx, index, products))
	t.Log("bulk index success: ", elasticsearch.BulkAddDocument(dbFactory, ctx, index, ps))
}

func TestSearchAll(t *testing.T) {
	searchAllReq := &search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
		Size: some.Int(1001),
	}
	t.Logf("search all: %+v", convert.AnyToJson(searchAllReq))
	res, err := elasticsearch.SearchDocument(dbFactory, ctx, "pms", searchAllReq)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestSearch(t *testing.T) {
	searchReq := &search.Request{
		Query: &types.Query{
			Term: map[string]types.TermQuery{
				"name": {
					Value: "品",
				},
			},
		},
	}
	t.Log(searchReq)
	searchAllReq := &search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}
	document, err := elasticsearch.SearchDocument(dbFactory, ctx, index, searchAllReq)
	if err != nil {
		t.Error(err)
	}
	t.Log(document)
}

func TestDeleteDocument(t *testing.T) {
	t.Log("delete document success: ", elasticsearch.DeleteDocument(dbFactory, ctx, index, "1"))
}
