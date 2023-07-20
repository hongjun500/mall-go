//	@author	hongjun500
//	@date	2023/6/11
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 后台用户路由

package r_mall_admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
	"github.com/hongjun500/mall-go/pkg/security"
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
	umsAdminGroup.POST("/refreshToken", router.refreshToken)
	// 根据用户 ID 获取用户信息
	umsAdminGroup.GET("/info", router.info)
	// umsAdminGroup.GET("/info/:user_id", router.UmsAdminService.UmsAdminInfo)
	// 用户列表分页
	umsAdminGroup.GET("/list", router.list)
	// 获取指定用户信息
	umsAdminGroup.GET("/:id", router.detail)
	// 修改指定用户信息
	umsAdminGroup.POST("/update/:id", router.update)
	// 删除指定用户
	umsAdminGroup.POST("/delete/:id", router.delete)
	// 修改指定用户状态
	umsAdminGroup.POST("/updateStatus/:id", router.updateStatus)
	// 给用户分配角色
	umsAdminGroup.POST("/role/update", router.roleUpdate)
	// 获取指定用户的角色
	umsAdminGroup.GET("/role/:adminId", router.role)
	// 修改指定用户密码
	umsAdminGroup.POST("/updatePassword", router.updatePassword)

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
		unAuthGroup.POST("/logout", router.logout)
		unAuthGroup.GET("/authTest", router.authTest)
	}
}

// register 用户注册
//
//	@Summary		用户注册
//	@Description	用户注册
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body		admin_dto.UmsAdminRegisterDTO	true	"用户注册"
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/admin/register [post]
func (router *UmsAdminRouter) register(context *gin.Context) {
	var request admin_dto.UmsAdminRegisterDTO
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
//	@Param			request	body		admin_dto.UmsAdminLoginDTO	true	"用户登录"
//	@Success		200		{object}	gin_common.GinCommonResponse
//	@Router			/admin/login [post]
func (router *UmsAdminRouter) login(context *gin.Context) {
	var umsAdminLogin admin_dto.UmsAdminLoginDTO
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

func (router *UmsAdminRouter) logout(context *gin.Context) {
	// 清除策略
	// security.Enforcer.RemovePolicy("admin")
	gin_common.Create(context)
}

// authTest 用户鉴权测试
//
//	@Summary		用户鉴权测试
//	@Description	用户鉴权测试
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/authTest [get]
func (*UmsAdminRouter) authTest(context *gin.Context) {
	m := map[string]any{
		"admin": map[string]any{
			"username": "admin",
			"token":    "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJob25nanVuNTAwIiwidXNlcl9pZCI6MTEsImNyZWF0ZWQiOiIyMDIzLTA2LTEwVDE0OjQ1OjA4LjEwMzc3NDYrMDg6MDAiLCJleHAiOjE2ODY5ODQzMDh9.PcLmIhxjenF36OPKmBX5ghPFgrfewSh_OUfT3dS-gUUL8UtyZFrg1gvxMbN8jZpOwJZIP5FQ7A1Yz1cfLl-Exg",
		},
		"hongjun500": map[string]any{
			"username": "hongjun500",
			"token":    "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJob25nanVuNTAwIiwidXNlcklkIjoxMSwiY3JlYXRlZCI6IjIwMjMtMDYtMTlUMTQ6Mjg6MjYuODM1MTgxNyswODowMCIsImV4cCI6MTY4Nzc2MDkwNn0.Hgt5qXSE25_zCHiCbtlEVdU2v-qsRG5-PR-Pckf7cThwGqbbOiHe2NAS-Yia8W8ALqIzI9mzSpSLc50dJLLsIw",
		},
	}
	gin_common.CreateSuccess(context, m)
}

// refreshToken 刷新 token
//
//	@Summary		刷新 token
//	@Description	刷新 token
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/refreshToken [post]
func (router *UmsAdminRouter) refreshToken(context *gin.Context) {
	header := context.GetHeader(conf.GlobalJwtConfigProperties.TokenHeader)
	refreshToken, _ := security.RefreshToken(header)
	if refreshToken == "" {
		gin_common.CreateFail(context, gin_common.TokenExpired)
		return
	}
	tokenMap := make(map[string]string)
	tokenMap["token"] = refreshToken
	tokenMap["tokenHead"] = conf.GlobalJwtConfigProperties.TokenHead
	gin_common.CreateSuccess(context, tokenMap)
}

// info 根据用户 ID 获取用户信息
//
//	@Summary		根据用户 ID 获取用户信息
//	@Description	根据用户 ID 获取用户信息
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/info [get]
func (router *UmsAdminRouter) info(context *gin.Context) {
	username := mid.GinJWTGetCurrentUsername(context)
	info, err := router.UmsAdminService.UmsAdminInfo(username)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, info)
}

// list 分页查询用户
//
//	@Summary		分页查询用户
//	@Description	分页查询用户
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			keyword		query	string	false	"keyword"
//	@Param			pageSize	query	int		true	"pageSize"
//	@Param			pageNum		query	int		true	"pageNum"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/list [get]
func (router *UmsAdminRouter) list(context *gin.Context) {
	var request admin_dto.UmsAdminPageDTO
	err := context.ShouldBind(&request)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	page, err := router.UmsAdminService.UmsAdminListPage(request)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, page)
}

// detail 获取指定用户信息
//
//	@Summary		获取指定用户信息
//	@Description	获取指定用户信息
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"用户 ID"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/{id} [get]
func (router *UmsAdminRouter) detail(context *gin.Context) {
	var userDTO base_dto.UserDTO
	err := context.ShouldBindUri(&userDTO)
	// 占位符
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	item, err := router.UmsAdminService.UmsAdminItem(userDTO)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, item)
}

// update 修改指定用户信息
//
//	@Summary		修改指定用户信息
//	@Description	修改指定用户信息
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int							true	"用户 ID"
//	@Param			request	body	admin_dto.UmsAdminUpdateDTO	true	"修改指定用户信息"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/update/{id} [post]
func (router *UmsAdminRouter) update(context *gin.Context) {
	var umsAdminUpdate admin_dto.UmsAdminUpdateDTO
	var userDTO base_dto.UserDTO
	err := context.ShouldBindUri(&userDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	err = context.ShouldBind(&umsAdminUpdate)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	update, err := router.UmsAdminService.UmsAdminUpdate(userDTO, umsAdminUpdate)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, update)
}

// delete 删除指定用户信息
//
//	@Summary		删除指定用户信息
//	@Description	删除指定用户信息
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	int	true	"用户 ID"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/delete/{id} [post]
func (router *UmsAdminRouter) delete(context *gin.Context) {
	var userDTO base_dto.UserDTO
	err := context.ShouldBindUri(&userDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	id, err := router.UmsAdminService.UmsAdminDelete(userDTO)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, id)
}

// updateStatus 修改指定用户状态
//
//	@Summary		修改指定用户状态
//	@Description	修改指定用户状态
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int	true	"用户 ID"
//	@Param			status	formData	int	true	"用户状态"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/updateStatus/{user_id} [post]
func (router *UmsAdminRouter) updateStatus(context *gin.Context) {
	var pathVariableDTO base_dto.PathVariableDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	status, _ := strconv.Atoi(context.Query("status"))
	id, err := router.UmsAdminService.UmsAdminUpdateStatus(pathVariableDTO.Id, int64(status))
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, id)
}

// roleUpdate 修改指定用户角色
//
//	@Summary		修改指定用户角色
//	@Description	修改指定用户角色
//	@Tags			后台用户管理
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			adminId	formData	int64	true	"用户 ID"
//	@Param			roleIds	formData	[]int64	true	"角色 ID"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/role/update [post]
func (router *UmsAdminRouter) roleUpdate(context *gin.Context) {
	var umsAdminRoleDTO admin_dto.UmsAdminRoleDTO
	var err error
	err = context.ShouldBind(&umsAdminRoleDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	update, err := router.UmsAdminService.UmsAdminRoleUpdate(umsAdminRoleDTO)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, update)
}

// UmsAdminRoleItem 获取指定用户的角色
//
//	@Summary		获取指定用户的角色
//	@Description	获取指定用户的角色
//	@Tags			后台用户管理
//	@Produce		json
//	@Param			adminId	path	int64	true	"用户 ID"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/role/{adminId} [get]
func (router *UmsAdminRouter) role(context *gin.Context) {
	var umsAdminRoleDTO admin_dto.UmsAdminRoleDTO
	err := context.ShouldBindUri(&umsAdminRoleDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	item, err := router.UmsAdminService.UmsAdminRoleItem(umsAdminRoleDTO)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, item)
}

// updatePassword 修改指定用户密码
//
//	@Summary		修改指定用户密码
//	@Description	修改指定用户密码
//	@Tags			后台用户管理
//	@Accept			json
//	@Produce		json
//	@Param			request	body	admin_dto.UmsAdminUpdatePasswordDTO	true	"修改指定用户密码"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/admin/updatePassword [post]
func (router *UmsAdminRouter) updatePassword(context *gin.Context) {
	var umsAdminUpdatePasswordDTO admin_dto.UmsAdminUpdatePasswordDTO
	err := context.ShouldBind(&umsAdminUpdatePasswordDTO)
	if err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	update, err := router.UmsAdminService.UmsAdminUpdatePassword(umsAdminUpdatePasswordDTO)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, update)
}
