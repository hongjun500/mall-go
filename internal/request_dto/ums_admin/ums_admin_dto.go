/**
 * @author hongjun500
 * @date 2023/6/4
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: ums_admin_dto.go
 */

package ums_admin

import "github.com/hongjun500/mall-go/internal/request_dto/base"

// UmsAdminRegisterDTO 用户注册请求参数
type UmsAdminRegisterDTO struct {
	// base.PageDTO

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

// UmsAdminLoginDTO 用户登录请求参数
type UmsAdminLoginDTO struct {
	// 用户名
	Username string `json:"username" form:"username" binding:"required"`
	// 密文密码
	Password string `json:"password" form:"password" binding:"required"`
}

// UmsAdminPageDTO 用户分页查询请求参数
type UmsAdminPageDTO struct {
	base.PageDTO
	// 用户名
	Username string `json:"username" form:"username"`
}

// UmsAdminUpdateDTO 用户更新请求参数
type UmsAdminUpdateDTO struct {
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

// UmsAdminUpdatePasswordDTO 用户更新密码请求参数
type UmsAdminUpdatePasswordDTO struct {
	// 用户名
	Username string `json:"username" form:"username" binding:"required"`
	// 旧密码
	OldPassword string `json:"oldPassword" form:"old_password" binding:"required"`
	// 新密码
	NewPassword string `json:"newPassword" form:"new_password" binding:"required"`
}
