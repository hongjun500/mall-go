// @author hongjun500
// @date 2023/6/20 10:45
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package database

import (
	"log"
	"strings"
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
func TestPutMapping(t *testing.T) {
	/*mapping := types.NewTypeMapping()
	settings := types.NewIndexSettings()
	settings.NumberOfShards = "1"
	settings.NumberOfReplicas = "1"
	mapping.Properties = map[string]types.Property{
		"price": types.NewIntegerNumberProperty(),
		"name":  types.NewTextProperty(),
		"count": types.NewIntegerNumberProperty(),
	}*/
	ik_max_word := "ik_max_word"
	request := putmapping.NewRequest()
	request.Properties = map[string]types.Property{
		"price": types.NewFloatNumberProperty(),
		"name": &types.TextProperty{
			Analyzer: &ik_max_word,
			Type:     "text",
		},
		"count": types.NewKeywordProperty(),
	}

	_, err := esTypedClient.Indices.PutMapping("pms").Request(request).Do(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func TestGetMapping(t *testing.T) {

	mapping, err := esTypedClient.Indices.GetMapping().Index("pms").Do(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, record := range mapping {
		for _, property := range record.Mappings.Properties {
			log.Println(property)
		}
	}
}

func TestDeleteIndex(t *testing.T) {
	_, err := esTypedClient.Indices.Delete("pms").Do(ctx)
	assert.Nil(t, err)
}

func TestSelectAllIndex(t *testing.T) {

	indices, err := esTypedClient.Cat.Indices().Do(ctx)
	if err != nil {
		return
	}
	for _, index := range indices {

		indexName := strings.Trim(*index.Index, `"`)
		log.Println("start del", indexName)
		esTypedClient.Indices.Delete(indexName).Do(ctx)
	}
}

func TestCreateIndex(t *testing.T) {
	indexReq := create.NewRequest()
	settings := types.NewIndexSettings()
	settings.NumberOfShards = "1"
	settings.NumberOfReplicas = "0"

	mapping := types.NewTypeMapping()

	mapping.Properties = map[string]types.Property{
		"id": map[string]interface{}{
			"type": "long",
		},
		"name": map[string]interface{}{
			"type":            "text",
			"analyzer":        "ik_max_word",
			"search_analyzer": "ik_max_word",
		},
	}

	indexReq.Settings = settings
	// indexReq.Mappings = mapping

	res, err := esTypedClient.Indices.Create("pms").Request(indexReq).Do(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	assert.Nil(t, err)
	assert.Equal(t, "pms", res.Index)
	assert.True(t, res.Acknowledged)
}

func TestSearchIndex(t *testing.T) {

	do, err := esTypedClient.Search().Index("test-index").Do(nil)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), do.Hits.Total.Value)
}
