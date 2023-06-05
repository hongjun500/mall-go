/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: ums_admin_dto.go
 */

package ums_admin

import "github.com/hongjun500/mall-go/internal/request_dto/base"

// UmsAdminRequest 用户注册请求参数
type UmsAdminRequest struct {
	base.PageDTO

	// 用户名
	Username string `json:"username" form:"username" binding:"required"`
	// 密文密码
	Password string `json:"password" form:"password" binding:"required"`
	// 用户头像
	Icon string `json:"icon" form:"icon"`
	// 邮箱
	Email string `json:"email" form:"email"`
	// 用户昵称
	Nickname string `json:"nickname" form:"nickname"`
	// 备注
	Note string `json:"note" form:"note"`
}

// UmsAdminLogin 用户登录请求参数
type UmsAdminLogin struct {
	// 用户名
	Username string `json:"username" form:"username" binding:"required"`
	// 密文密码
	Password string `json:"password" form:"password" binding:"required"`
}

type UmsAdminPage struct {
	base.PageDTO
	// 用户名
	Username string `json:"username" form:"username"`
}
