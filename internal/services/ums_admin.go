package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/init"
	"github.com/hongjun500/mall-go/internal/models"
)

type UmsAdminRequest struct {
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

func UmsAdminRegister(context *gin.Context, request *UmsAdminRequest) *models.UmsAdmin {
	// 检查用户名是否重复了
	var umsAdmin models.UmsAdmin
	umsAdmins, err := umsAdmin.GetUmsAdminByUsername(init.SqlSession.DbMySQL, "")
	if err != nil {
		return nil
	}
	if umsAdmins != nil && len(umsAdmins) > 0 {
		// return err.Error("f")
	}
	return nil
}
