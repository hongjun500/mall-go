// @author hongjun500
// @date 2023/6/21 15:44
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package elasticsearch

import (
	"bytes"
	"context"
	"github.com/hongjun500/mall-go/pkg/convert"
	"log"
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
		esParam[field.Name] = tags
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
	settings := params[1].(*types.IndexSettings)
	mappings := params[2].(*types.TypeMapping)
	creq := create.NewRequest()
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
					Analyzer: &esAnalyzer,
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
			property[field] = types.NewNestedProperty()
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

// UpdateDocument 根据文档 id 更新文档
func UpdateDocument(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	id := params[1].(string)
	body := params[2].(any)

	toJson := convert.StructToJson(body)
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

// AddDocuments AddDocument 添加多个文档
func AddDocuments(db *database.DbFactory, ctx context.Context, params ...any) bool {
	index := params[0].(string)
	// id := params[1].(string)
	body := params[1].(any)
	res, err := db.Es.TypedCli.Index(index).Request(body).Do(ctx)
	if err != nil {
		log.Printf("add document error: %v", err.Error())
		return false
	}
	return res.Result.Name == "created"
}
