// @author hongjun500
// @date 2023/6/20 10:45
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package database

import (
	"log"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/putmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/stretchr/testify/assert"
)

func TestNewEsClient(t *testing.T) {
	assert.NotNil(t, esTypedClient)
}

func TestExistIndex(t *testing.T) {
	success, err := esTypedClient.Indices.Exists("test-index").IsSuccess(ctx)
	assert.Nil(t, err)
	assert.True(t, success)
	success, err = esTypedClient.Indices.Exists("fsdfasdf").IsSuccess(ctx)
	assert.Nil(t, err)
	assert.False(t, success)
}

// 创建mapping
func TestCreateMapping(t *testing.T) {
	/*mapping := types.NewTypeMapping()
	settings := types.NewIndexSettings()
	settings.NumberOfShards = "1"
	settings.NumberOfReplicas = "1"
	mapping.Properties = map[string]types.Property{
		"price": types.NewIntegerNumberProperty(),
		"name":  types.NewTextProperty(),
		"count": types.NewIntegerNumberProperty(),
	}*/
	request := putmapping.NewRequest()
	request.Properties = map[string]types.Property{
		"price": types.NewIntegerNumberProperty(),
		"name":  types.NewTextProperty(),
		"count": types.NewIntegerNumberProperty(),
	}
	_, err := esTypedClient.Indices.PutMapping("test-index1").Request(request).Do(ctx)
	if err != nil {
		return
	}
}

func TestDeleteIndex(t *testing.T) {
	_, err := esTypedClient.Indices.Delete("test-index").Do(ctx)
	assert.Nil(t, err)
}

func TestSelectAllIndex(t *testing.T) {

}

func TestCreateIndex(t *testing.T) {
	// body := `{"name":"hongjun500"}`
	res, err := esTypedClient.Indices.Create("test-index").
		Request(&create.Request{
			Mappings: &types.TypeMapping{
				Properties: map[string]types.Property{
					"price": types.NewIntegerNumberProperty(),
					"name":  types.NewTextProperty(),
				},
			},
		}).
		Do(nil)
	assert.Nil(t, err)
	assert.Equal(t, "test-index", res.Index)
	assert.True(t, res.Acknowledged)

	type doc struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Price int    `json:"price"`
	}
	document := doc{
		Name:  "hongjun500",
		Age:   18,
		Price: 100,
	}
	log.Print(document)
}

func TestSearchIndex(t *testing.T) {

	do, err := esTypedClient.Search().Index("test-index").Do(nil)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), do.Hits.Total.Value)
}
