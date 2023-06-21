// @author hongjun500
// @date 2023/6/21 16:11
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description: product 索引 pms

package es_index

type Product struct {
	Id                  int64                    `json:"id" es_type:"long"`
	ProductSn           string                   `json:"product_sn" es_type:"keyword"`
	BrandId             int64                    `json:"brand_id" es_type:"long"`
	BrandName           string                   `json:"brand_name" es_type:"keyword"`
	ProductCategoryId   int64                    `json:"product_category_id" es_type:"long"`
	ProductCategoryName string                   `json:"product_category_name" es_type:"keyword"`
	Pic                 string                   `json:"pic" es_type:"keyword"`
	Name                string                   `json:"name" es_type:"text" es_analyzer:"ik_max_word"`
	SubTitle            string                   `json:"sub_title" es_type:"text" es_analyzer:"ik_max_word"`
	KeyWord             string                   `json:"key_word" es_type:"text" es_analyzer:"ik_max_word"`
	Price               string                   `json:"price" es_type:"float"`
	Sale                int                      `json:"sale" es_type:"integer"`
	NewStatus           int                      `json:"new_status" es_type:"integer"`
	RecommandStatus     int                      `json:"recommand_status" es_type:"integer"`
	Stock               int                      `json:"stock" es_type:"integer"`
	PromotionType       int                      `json:"promotion_type" es_type:"integer"`
	Sort                int                      `json:"sort" es_type:"integer"`
	AttrValueList       []*ProductAttributeValue `json:"attr_value_list" es_type:"nested"`
}

type ProductAttributeValue struct {
	Id                 int64  `json:"id" es_type:"long"`
	ProductAttributeId int64  `json:"product_id" es_type:"long"`
	Value              string `json:"value" es_type:"keyword"`
	Type               int    `json:"type" es_type:"integer"`
	Name               string `json:"name" es_type:"keyword"`
}
