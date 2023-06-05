/**
* @author hongjun500
* @date 2023/6/4
* @tool ThinkPadX1隐士
* Created with GoLand 2022.2
* Description: 用户基础数据传输对象
 todo 中间件 可用于将请求头中的 token 解析出来的用户信息传递到 请求体中
*/

package base

type UserDTO struct {
	UserId int64 `json:"user_id" uri:"user_id" binding:"required"`
}
