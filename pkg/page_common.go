// @author hongjun500
// @date 2023/7/3
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 通用分页数据

package pkg

type CommonPage struct {
	// 当前页
	PageNum int `json:"pageNum" form:"page_num" binding:"required" default:"1"`
	// 每页数量
	PageSize int `json:"pageSize" form:"page_size" binding:"required" default:"10"`
	// 总页数
	TotalPage int64 `json:"totalPage" form:"total_page"`
	// 总记录数
	Total int64 `json:"total" form:"total"`
	// 分页数据
	List any `json:"list" form:"list"`
	// 排序
	OrderBy string `json:"orderBy" form:"order_by"`
	// 排序方式
	Sort string `json:"sort" form:"sort"`
}

type Paginator interface {
	Paginate() error
}
