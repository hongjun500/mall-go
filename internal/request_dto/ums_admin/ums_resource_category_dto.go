// @author hongjun500
// @date 2023/6/13 11:00
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package ums_admin

type UmsResourceCategoryCreateDTO struct {
	// 资源分类名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
}
