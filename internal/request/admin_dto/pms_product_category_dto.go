// @author hongjun500
// @date 2023/7/14 14:29
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package admin_dto

type PmsProductCategoryDTO struct {
	// 父分类的编号
	ParentId int64 `json:"parentId"`
	// 分类名称
	Name string `json:"name" binding:"required"`
	// 分类单位
	ProductUnit string `json:"productUnit"`
	// 是否显示在导航栏：0->不显示；1->显示
	NavStatus int `json:"navStatus"`
	// 显示状态：0->不显示；1->显示
	ShowStatus int `json:"showStatus"`
	// 排序
	Sort int `json:"sort"`
	// 图标
	Icon string `json:"icon"`
	// 关键字
	Keywords string `json:"keywords"`
	// 描述
	Description string `json:"description"`
	// 产品相关筛选属性集合
	ProductAttributeIdList []int64 `json:"productAttributeIdList"`
}
