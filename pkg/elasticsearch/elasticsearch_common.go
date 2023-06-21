// @author hongjun500
// @date 2023/6/21 15:44
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package elasticsearch

import (
	"context"
	"reflect"

	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/putmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hongjun500/mall-go/internal/database"
)

const (
	EsType     = "es_type"
	EsAnalyzer = "es_analyzer"
	EsField    = "field"
	IkMaxWord  = "ik_max_word"
)

var (
	ikMaxWord = IkMaxWord
	// tags map[string]map[string]string
)

type tags map[string]map[string]string

// GetStructTag 获取结构体的 elasticsearch 标签
func GetStructTag(t any) tags {
	tt := reflect.TypeOf(t)
	esParam := make(map[string]map[string]string, 0)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		tag := field.Tag
		esTags := make(map[string]string)
		if tag != "" {
			esTags[EsType] = tag.Get(EsType)
			esTags[EsAnalyzer] = tag.Get(EsAnalyzer)
		}
		esParam[field.Name] = esTags
	}
	return esParam
}

// HasIndex 索引是否存在
func HasIndex(db *database.DbFactory, ctx context.Context, index string) bool {
	success, err := db.Es.TypedCli.Indices.Exists(index).IsSuccess(ctx)
	if err != nil {
		return false
	}
	return success
}

// CreateIndex 创建索引
func CreateIndex(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	settings := params[1].(*types.IndexSettings)
	mappings := params[2].(*types.TypeMapping)
	creq := create.NewRequest()
	creq.Settings = settings
	creq.Mappings = mappings
	res, err := db.Es.TypedCli.Indices.Create(index).Request(creq).Do(ctx)
	if err != nil {
		return false
	}
	return res.Acknowledged
}

// PutMapping 根据结构体更新 mapping
func PutMapping(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	// mappings := params[1].(*types.TypeMapping)
	putreq := putmapping.NewRequest()

	putreq.Properties = map[string]types.Property{
		"price": types.NewFloatNumberProperty(),
		"name": &types.TextProperty{
			Analyzer: &ikMaxWord,
			Type:     "text",
		},
		"count": types.NewKeywordProperty(),
	}

	_, err := db.Es.TypedCli.Indices.PutMapping(index).Request(putreq).Do(ctx)
	if err != nil {
		return false
	}
	return true
}
