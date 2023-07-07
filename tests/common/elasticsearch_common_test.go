// @author hongjun500
// @date 2023/6/21 15:52
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package common

import (
	"strconv"
	"testing"

	"github.com/hongjun500/mall-go/internal"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hongjun500/mall-go/internal/es_index"
	"github.com/hongjun500/mall-go/pkg/convert"
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
	t.Log("delete index success: ", internal.DeleteIndex(dbFactory.Es, ctx, "pms"))
}

func TestHasIndex(t *testing.T) {
	t.Log("index is exist: ", internal.HasIndex(dbFactory.Es, ctx, index))
}

func TestCreateIndex(t *testing.T) {
	settings := &types.IndexSettings{NumberOfShards: "1", NumberOfReplicas: "0"}

	t.Log("create index success: ", internal.CreateIndex(dbFactory.Es, ctx, index, settings))
}

func TestPutMapping(t *testing.T) {
	_ = internal.GetStructTag(es_index.EsProduct{})
	esProduct := es_index.EsProduct{}
	t.Log("put mapping success: ", internal.PutMappingByStruct(dbFactory.Es, ctx, index, esProduct))
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
	t.Log("bulk index success: ", internal.BulkAddDocument(dbFactory.Es, ctx, index, ps))
}

func TestSearchAll(t *testing.T) {
	searchAllReq := &search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
		Size: some.Int(1001),
	}
	t.Logf("search all: %+v", convert.AnyToJson(searchAllReq))
	res, err := internal.SearchDocument(dbFactory.Es, ctx, "pms", searchAllReq)
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
	t.Log(convert.AnyToJson(searchReq))
	searchAllReq := &search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}
	t.Log(convert.AnyToJson(searchAllReq))
	searchMultiMatchReq := &search.Request{
		Query: &types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Query:  "小米",
				Fields: []string{"name", "subTitle", "keywords"},
			},
		},
	}
	t.Log(convert.AnyToJson(searchMultiMatchReq))
	document, err := internal.SearchDocument(dbFactory.Es, ctx, "pms", searchMultiMatchReq)
	if err != nil {
		t.Error(err)
	}
	t.Log(document)
}

func TestDeleteDocument(t *testing.T) {
	t.Log("delete document success: ", internal.DeleteDocument(dbFactory.Es, ctx, index, "1"))
}

func TestElasticSearchPage(t *testing.T) {
	searchReq := &search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}
	t.Log(searchReq)
	page := internal.NewElasticSearchPage(dbFactory.Es, index, 2, 500)
	page.SearchRequest = searchReq
	err := page.Paginate()
	if err != nil {
		t.Error(err)
	}
	t.Log(page.List)
}

func TestElasticSearchPage2(t *testing.T) {

	params := map[string]string{
		"name":     "6",
		"subTitle": "6",
		"keyWord":  "6",
	}

	queries := make([]types.Query, 0, len(params))

	for key, value := range params {
		query := types.NewQuery()
		m := map[string]types.MatchQuery{
			key: {Query: value},
		}
		query.Match = m
		queries = append(queries, *query)
	}

	searchReq := &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Should: /*[]types.Query{
					{
						Match: map[string]types.MatchQuery{
							"name": {Query: "66"},
						},
					},
					{
						Match: map[string]types.MatchQuery{
							"subTitle": {Query: "your_search_term"},
						},
					},
					{
						Match: map[string]types.MatchQuery{
							"keyWord": {Query: "your_search_term"},
						},
					},
				},*/
				queries,
			},
		},
	}
	t.Log(searchReq)
	page := internal.NewElasticSearchPage(dbFactory.Es, "pms", 1, 2000)
	page.SearchRequest = searchReq
	err := page.Paginate()
	if err != nil {
		t.Error(err)
	}

	t.Log(page.List)
}
func TestSearchByNameOrSubtitle(t *testing.T) {
	esProduct := new(es_index.EsProduct)
	dsl, sort, err := esProduct.SearchByNameOrSubtitle(1)
	if err != nil {
		return
	}
	t.Log(dsl)
	t.Log(sort)
}
