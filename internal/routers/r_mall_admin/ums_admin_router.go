//	@author	hongjun500
//	@date	2023/6/11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台用户路由

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsAdminRouter struct {
	s_mall_admin.UmsAdminService
}

func NewUmsAdminRouter(service s_mall_admin.UmsAdminService) *UmsAdminRouter {
	return &UmsAdminRouter{UmsAdminService: service}
}

// GroupUmsAdminRouter 用户管理路由
func (router *UmsAdminRouter) GroupUmsAdminRouter(umsAdminGroup *gin.RouterGroup) {
	// umsAdminGroup := routerEngine.Group("/admin")

	// 刷新 token
	umsAdminGroup.POST("/refreshToken", router.UmsAdminService.UmsAdminRefreshToken)
	// 根据用户 ID 获取用户信息
	umsAdminGroup.GET("/info", router.UmsAdminService.UmsAdminInfo)
	// umsAdminGroup.GET("/info/:user_id", router.UmsAdminService.UmsAdminInfo)
	// 用户列表分页
	umsAdminGroup.GET("/list", router.UmsAdminService.UmsAdminListPage)
	// 获取指定用户信息
	umsAdminGroup.GET("/:id", router.UmsAdminService.UmsAdminItem)
	// 修改指定用户信息
	umsAdminGroup.POST("/update/:id", router.UmsAdminService.UmsAdminUpdate)
	// 删除指定用户
	umsAdminGroup.POST("/delete/:id", router.UmsAdminService.UmsAdminDelete)
	// 修改指定用户状态
	umsAdminGroup.POST("/updateStatus/:id", router.UmsAdminService.UmsAdminUpdateStatus)
	// 给用户分配角色
	umsAdminGroup.POST("/role/update", router.UmsAdminService.UmsAdminRoleUpdate)
	// 获取指定用户的角色
	umsAdminGroup.GET("/role/:adminId", router.UmsAdminService.UmsAdminRoleItem)
	// 修改指定用户密码
	umsAdminGroup.POST("/updatePassword", router.UmsAdminService.UmsAdminUpdatePassword)

}

// UnauthorizedGroupRouter  未授权路由
func (router *UmsAdminRouter) UnauthorizedGroupRouter(routerEngine *gin.Engine) {
	unAuthGroup := routerEngine.Group("/admin")
	{
		// 用户注册
		unAuthGroup.POST("/register", router.register)

		// 用户登录
		unAuthGroup.POST("/login", router.login)
		// 用户登出
		unAuthGroup.POST("/logout", router.UmsAdminService.UmsAdminLogout)
		unAuthGroup.GET("/authTest", router.UmsAdminService.UmsAdminAuthTest)
	}
}

// register 用户注册
//
//	@Summary		用户注册
//	@Description	用户注册
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ums_admin.UmsAdminRegisterDTO	true	"用户注册"
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/admin/register [post]
func (router *UmsAdminRouter) register(context *gin.Context) {
	var request ums_admin.UmsAdminRegisterDTO
	err := context.ShouldBind(&request)
	if err != nil {
		gin_common.CreateFail(context, err.Error())
		// 这些地方需要手动 return
		return
	}
	err = router.UmsAdminService.UmsAdminRegister(request)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.Create(context)
}

// login 用户登录
//
//	@Summary		用户登录
//	@Description	用户登录
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ums_admin.UmsAdminLoginDTO	true	"用户登录"
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/admin/login [post]
func (router *UmsAdminRouter) login(context *gin.Context) {
	var umsAdminLogin ums_admin.UmsAdminLoginDTO
	err := context.ShouldBind(&umsAdminLogin)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		// 这些地方需要手动 return
		return
	}

	tokenMap, err := router.UmsAdminService.UmsAdminLogin(umsAdminLogin, context.ClientIP(), context.Request.Host, context.Request.UserAgent())
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, tokenMap)

}
