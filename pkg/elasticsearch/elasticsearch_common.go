//	@author	hongjun500
//	@date	2023/6/21 15:44
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/putmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/some"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/pkg/convert"
)

const (
	EsType     = "es_type"
	EsAnalyzer = "es_analyzer"
	Json       = "json"
	IkMaxWord  = "ik_max_word"
)

var (
	ikMaxWord = IkMaxWord
	// tags map[string]map[string]string
)

type esTags map[string]map[string]string

// GetStructTag 获取结构体的 elasticsearch 标签
func GetStructTag(t any) esTags {
	tt := reflect.TypeOf(t)
	esParam := make(map[string]map[string]string, 0)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		tag := field.Tag
		tags := make(map[string]string, 0)
		if tag != "" {
			tags[EsType] = tag.Get(EsType)
			tags[EsAnalyzer] = tag.Get(EsAnalyzer)
		}
		esParam[tag.Get(Json)] = tags
	}
	return esParam
}

// HasIndex 索引是否存在
func HasIndex(db *database.DbFactory, ctx context.Context, index string) bool {
	success, err := db.Es.TypedCli.Indices.Exists(index).IsSuccess(ctx)
	if err != nil {
		log.Printf("has index error: %v", err.Error())
		return false
	}
	return success
}

// CreateIndex 创建索引
func CreateIndex(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	creq := create.NewRequest()
	var settings *types.IndexSettings
	var mappings *types.TypeMapping

	if len(params) >= 2 {
		settings = params[1].(*types.IndexSettings)
	} else if len(params) >= 3 {
		settings = params[1].(*types.IndexSettings)
		mappings = params[2].(*types.TypeMapping)
	}
	creq.Settings = settings
	creq.Mappings = mappings
	res, err := db.Es.TypedCli.Indices.Create(index).Request(creq).Do(ctx)
	if err != nil {
		log.Printf("create index error: %v", err.Error())
		return false
	}
	return res.Acknowledged
}

// PutMappingByStruct 根据结构体更新 mapping
func PutMappingByStruct(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	// mappings := params[1].(*types.TypeMapping)
	tags := params[1].(esTags)

	property := make(map[string]types.Property, 0)

	for field, tag := range tags {
		esType := tag[EsType]
		esAnalyzer := tag[EsAnalyzer]
		if tag[EsType] == "" {
			continue
		}
		switch esType {
		case "text":
			if esAnalyzer == "" {
				property[field] = types.NewTextProperty()
			} else {
				property[field] = &types.TextProperty{
					Analyzer: some.String(esAnalyzer),
					Type:     "text",
				}
			}
		case "keyword":
			property[field] = types.NewKeywordProperty()
		case "long":
			property[field] = types.NewLongNumberProperty()
		case "float":
			property[field] = types.NewFloatNumberProperty()
		case "date":
			property[field] = types.NewDateProperty()
		case "boolean":
			property[field] = types.NewBooleanProperty()
		case "nested":
			// todo 这里需要处理嵌套结构体
			property[field] = types.NewNestedProperty()
			processNestedStruct(property, field, reflect.ValueOf(field))
		default:
			continue
		}

	}

	putreq := putmapping.NewRequest()
	putreq.Properties = property
	_, err := db.Es.TypedCli.Indices.PutMapping(index).Request(putreq).Do(ctx)
	if err != nil {
		log.Printf("put mapping error: %v", err.Error())
		return false
	}
	return true
}

// 嵌套结构体的 mapping 处理
func processNestedStruct(property map[string]types.Property, field string, nestedStruct reflect.Value) {
	nestedType := nestedStruct.Type()

	// 创建一个嵌套属性
	nestedProperty := types.NewNestedProperty()

	// 递归处理嵌套结构体的字段
	for i := 0; i < nestedStruct.NumField(); i++ {
		nestedField := nestedStruct.Field(i)
		nestedFieldName := nestedType.Field(i).Name

		// 处理嵌套结构体的字段类型
		switch nestedField.Kind() {
		case reflect.Struct:
			// 递归处理嵌套结构体
			processNestedStruct(nestedProperty.Properties, nestedFieldName, nestedField)
		case reflect.String:

			nestedProperty.Properties[nestedFieldName] = types.NewTextProperty()
		case reflect.Int, reflect.Int64:
			nestedProperty.Properties[nestedFieldName] = types.NewLongNumberProperty()
		// 处理其他字段类型
		// ...
		default:
			// 忽略不支持的字段类型
			continue
		}
	}
	// 将处理好的嵌套属性添加到主属性中
	property[field] = nestedProperty
}

// DeleteIndex 删除索引
func DeleteIndex(db *database.DbFactory, ctx context.Context, index string) bool {
	_, err := db.Es.TypedCli.Indices.Delete(index).Do(ctx)
	if err != nil {
		log.Printf("delete index error: %v", err.Error())
		return false
	}
	return true
}

// AddDocument 添加单个文档
func AddDocument(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	id := params[1].(string)
	body := params[2].(any)
	res, err := db.Es.TypedCli.Index(index).Id(id).Request(body).Do(ctx)
	if err != nil {
		log.Printf("add document error: %v", err.Error())
		return false
	}
	return res.Result.Name == "created"
}

// BulkAddDocument 批量添加文档
func BulkAddDocument(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	body := params[1].(any)

	start := time.Now() // 记录开始时间

	bs := convert.AnyToBytes(body)
	var list []any
	err := convert.BytesToAny(bs, &list)
	if err != nil {
		return false
	}
	bodyStr := ""
	for _, data := range list {
		m := data.(map[string]any)
		if m["id"] == nil {
			continue
		}
		by := convert.AnyToBytes(data)
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, m["id"], "\n"))
		by = append(meta, by...)
		by = append(by, []byte("\n")...)
		bodyStr += string(by)
	}
	res, err := db.Es.Cli.Bulk(bytes.NewReader([]byte(bodyStr)), db.Es.Cli.Bulk.WithIndex(index))
	if err != nil {
		log.Printf("bulk add document error: %v", err.Error())
		return false
	}
	elapsed := time.Since(start) // 计算耗时
	fmt.Printf("耗时：%s\n", elapsed)
	return !res.IsError()
}

// UpdateDocument 根据文档 id 更新文档
func UpdateDocument(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	id := params[1].(string)
	body := params[2].(any)

	toJson := convert.AnyToJson(body)
	res, err := db.Es.TypedCli.Update(index, id).Raw(bytes.NewReader([]byte(toJson))).Do(ctx)
	if err != nil {
		log.Printf("update document error: %v", err.Error())
		return false
	}
	return res.Get.Found
}

// DeleteDocument 根据文档 id 删除文档
func DeleteDocument(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	id := params[1].(string)
	res, err := db.Es.TypedCli.Delete(index, id).Do(ctx)
	if err != nil {
		log.Printf("delete document error: %v", err.Error())
		return false
	}
	return res.Result.Name == "deleted"
}

// SearchDocument 根据索引名 index 和 search.Request 条件查询文档
func SearchDocument(db *database.DbFactory, ctx context.Context, params ...any) (any, error) {
	index := params[0].(string)
	body := params[1].(*search.Request)
	start := time.Now()
	res, err := db.Es.TypedCli.Search().Index(index).Request(body).Do(ctx)
	if err != nil {
		log.Printf("search document error: %v", err.Error())
		return nil, err
	}
	hits := res.Hits.Hits
	data := make([]any, 0, len(hits)) // 预分配足够的容量
	for _, hit := range hits {
		if hit.Source_ != nil {
			var result map[string]interface{}
			// var result map[string]interface{}
			err := json.Unmarshal(hit.Source_, &result)
			// result, err = convert.JsonToMap(hit.Source_)
			if err != nil {
				log.Printf("failed to unmarshal search result: %v", err)
				continue
			}
			data = append(data, result)
		} else {
			log.Println("empty source in hit")
		}
	}
	elapsed := time.Since(start) // 计算耗时
	fmt.Printf("耗时：%s\n", elapsed)
	return data, nil
}
