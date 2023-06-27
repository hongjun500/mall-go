// @author hongjun500
// @date 2023/6/26 16:24
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

type PmsProductAttribute struct {
	Model
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId" gorm:"product_attribute_category_id"`
	Name                       string `json:"name" gorm:"attr_name"`
	// 属性选择类型：0->唯一；1->单选；2->多选
	SelectType int `json:"selectType" json:"selectType" gorm:"select_type"`
	// 属性录入方式：0->手工录入；1->从列表中选取
	InputType int `json:"inputType" gorm:"input_type"`
	// 可选值列表，以逗号隔开
	InputList string `json:"inputList" gorm:"input_list"`
	// 排序字段：最高的可以单独上传图片
	Sort int `json:"sort" gorm:"sort"`
	// 分类筛选样式：1->普通；1->颜色
	FilterType int `json:"filterType" gorm:"filter_type"`
	// 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	SearchType int `json:"searchType" gorm:"search_type"`
	// 相同属性产品是否关联；0->不关联；1->关联
	RelatedStatus int `json:"relatedStatus" gorm:"related_status"`
	// 是否支持手动新增；0->不支持；1->支持
	HandAddStatus int `json:"handAddStatus" gorm:"hand_add_status"`
	// 属性的类型；0->规格；1->参数
	Type int `json:"type" gorm:"type"`
}

func (*PmsProductAttribute) TableName() string {
	return "pms_product_attribute"
}
