/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: 分页请求的数据传输对象
 */

package base

type PageDTO struct {
	PageNum  int `json:"page_num" form:"page_num" binding:"required" default:"1"`
	PageSize int `json:"page_size" form:"page_size" binding:"required" default:"10"`
}
