/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: 分页请求的数据传输对象
 */

package base

type PageDTO struct {
	// 页码
	PageNum int `json:"pageNum" form:"pageNum" binding:"required" default:"1"`
	// 每页数量
	PageSize int `json:"pageSize" form:"pageSize" binding:"required" default:"10"`
}
