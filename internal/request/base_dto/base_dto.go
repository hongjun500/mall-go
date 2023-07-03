//	@author	hongjun500
//	@date	2023/6/13 16:54
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 通用的请求数据对象

package base_dto

type (
	// PageDTO 分页请求的数据传输对象
	PageDTO struct {
		// 页码
		PageNum int `json:"pageNum" form:"pageNum" binding:"required" default:"1"`
		// 每页数量
		PageSize int `json:"pageSize" form:"pageSize" binding:"required" default:"10"`
	}
	// IdsDTO 批量操作的数据传输对象
	IdsDTO struct {
		// ID列表
		Ids []int64 `json:"ids" form:"ids" binding:"required"`
	}
	// PathVariableDTO 路径参数的数据传输对象
	PathVariableDTO struct {
		// ID
		Id int64 `uri:"id" binding:"required"`
	}
	// StatusDTO 状态的数据传输对象
	StatusDTO struct {
		// 状态
		Status int `json:"status" form:"status" binding:"required"`
	}
)
