// @author hongjun500
// @date 2023/6/10
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: gorm 分页

package gorm_common

type GormCommonPage struct {
	// 当前页
	PageNum int `json:"page_num" form:"page_num" binding:"required" default:"1"`
	// 每页数量
	PageSize int `json:"page_size" form:"page_size" binding:"required" default:"10"`
	// 总页数
	TotalPage int `json:"total_page" form:"total_page"`
	// 总记录数
	Total int `json:"total" form:"total"`
	// 分页数据
	List any `json:"list" form:"list"`
}
