// @author hongjun500
// @date 2023/6/21 15:52
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package common

import (
	"testing"

	"github.com/hongjun500/mall-go/pkg/elasticsearch"
)

type Product struct {
	Id        int64   `json:"id" es_type:"long"`
	Name      string  `json:"name" es_type:"text" es_analyzer:"ik_max_word"`
	Price     float64 `json:"price" es_type:"float"`
	Count     int64   `json:"count" es_type:"long"`
	BrandName string  `json:"brand_name" es_type:"keyword"`
}

func TestGetStructTags(t *testing.T) {
	product := Product{
		Id:        1,
		Name:      "测试商品",
		Price:     100.0,
		Count:     100,
		BrandName: "华为",
	}
	tags := elasticsearch.GetStructTag(product)
	t.Log(tags)
}
