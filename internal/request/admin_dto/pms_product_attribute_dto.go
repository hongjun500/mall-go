// @author hongjun500
// @date 2023/7/21 10:36
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package admin_dto

import "github.com/hongjun500/mall-go/internal/request/base_dto"

type (
	PmsProductAttributeDTO struct {
		// 属性分类ID
		ProductAttributeCategoryId int64 `json:"productAttributeCategoryId"`
		// 属性名称
		Name string `json:"name"`
		// 属性选择类型：0->唯一；1->单选；2->多选
		SelectType int `json:"selectType"`
		// 属性录入方式：0->手工录入；1->从列表中选取
		InputType int `json:"inputType"`
		// 可选值列表，以逗号隔开
		InputList string `json:"inputList"`
		// 排序字段
		Sort int `json:"sort"`
		// 分类筛选样式：1->普通；1->颜色
		FilterType int `json:"filterType"`
		// 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
		SearchType int `json:"searchType"`
		// 相同属性产品是否关联；0->不关联；1->关联
		RelatedStatus int `json:"relatedStatus"`
		// 是否支持手动新增；0->不支持；1->支持
		HandAddStatus int `json:"handAddStatus"`
		// 属性的类型；0->规格；1->参数
		Type int `json:"type"`
	}

	PmsProductAttributeListDTO struct {
		// 属性分类ID
		CategoryId int64 `path:"cid" binding:"required"`
		Type       int   `query:"type"`
		base_dto.PageDTO
	}

	PmsProductAttributeInfoDTO struct {
		// 属性ID
		CategoryId int64 `path:"productCategoryId" binding:"required"`
	}
)
