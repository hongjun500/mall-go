package services

import (
	"crypto/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request_dto/base"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
	"github.com/hongjun500/mall-go/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type UmsAdminService struct {
	DbFactory *database.DbFactory
}

func NewUmsAdminService(dbFactory *database.DbFactory) UmsAdminService {
	return UmsAdminService{DbFactory: dbFactory}
}

// UmsAdminRegister 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param request body ums_admin.UmsAdminRegisterDTO true "用户注册"
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/register [post]
func (s UmsAdminService) UmsAdminRegister(context *gin.Context) {
	var request ums_admin.UmsAdminRegisterDTO
	err := context.ShouldBind(&request)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		// 这些地方需要手动 return
		return
	}
	// 检查用户名是否重复了
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.SelectUmsAdminByUsername(s.DbFactory.GormMySQL, request.Username)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	if umsAdmins != nil && len(umsAdmins) > 0 {
		gin_common.CreateFail(gin_common.UsernameAlreadyExists, context)
		return
	}
	// 密码加密
	hashPassword, _ := HashPassword(request.Password)

	// 创建用户参数
	umsAdmin = &models.UmsAdmin{
		Username: request.Username,
		Password: hashPassword,
		Icon:     request.Icon,
		Email:    request.Email,
		Nickname: request.Nickname,
		Note:     request.Note,
	}
	register, err := umsAdmin.InsertUmsAdmin(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	if register > 0 {
		gin_common.Create(context)
		return
	}
}

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
	// 生成随机盐值
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// 将密码与盐值进行哈希
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// 返回加密后的密码字符串
	return string(hash), nil
}

// VerifyPassword 验证密码
func VerifyPassword(password, hashedPassword string) bool {
	// 进行哈希校验
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// UmsAdminLogin 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param request body ums_admin.UmsAdminLoginDTO true "用户登录"
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/login [post]
func (s UmsAdminService) UmsAdminLogin(context *gin.Context) {
	var umsAdminLogin ums_admin.UmsAdminLoginDTO
	err := context.ShouldBind(&umsAdminLogin)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		// 这些地方需要手动 return
		return
	}
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.SelectUmsAdminByUsername(s.DbFactory.GormMySQL, umsAdminLogin.Username)
	if err != nil {
		gin_common.CreateFail(gin_common.DatabaseError, context)
		return
	}
	if umsAdmins == nil || len(umsAdmins) == 0 {
		gin_common.CreateFail(gin_common.UsernameOrPasswordError, context)
		return
	}
	umsAdmin = umsAdmins[0]
	if !VerifyPassword(umsAdminLogin.Password, umsAdmin.Password) {
		gin_common.CreateFail(gin_common.UsernameOrPasswordError, context)
		return
	}
	if umsAdmin.Status == 0 {
		gin_common.CreateFail(gin_common.AccountLocked, context)
		return
	}
	// 获取当前用户所拥有的资源
	resources := GetResource(s, umsAdmin.Id)
	// 添加策略
	security.AddPolicyFromResource(security.Enforcer, resources)

	token := security.GenerateToken(umsAdmin.Username, umsAdmin.Id, resources)
	if token == "" {
		gin_common.CreateFail(gin_common.TokenGenFail, context)
		return
	}
	now := time.Now()
	umsAdmin.LoginTime = &now
	// 更新登录时间
	_, _ = umsAdmin.UpdateUmsAdminLoginTimeByUserId(s.DbFactory.GormMySQL)

	umsAdminLoginLog := new(models.UmsAdminLoginLog)
	umsAdminLoginLog.AdminId = umsAdmin.Id
	umsAdminLoginLog.Ip = context.ClientIP()
	umsAdminLoginLog.Address = context.Request.Host
	umsAdminLoginLog.UserAgent = context.Request.UserAgent()
	// 记录登录日志
	_, _ = umsAdminLoginLog.SaveLoginLog(s.DbFactory.GormMySQL)

	tokenMap := make(map[string]string)
	tokenMap["token"] = token
	tokenMap["tokenHead"] = conf.GlobalJwtConfigProperties.TokenHead
	gin_common.CreateSuccess(tokenMap, context)

}

func GetResource(s UmsAdminService, adminId int64) []models.UmsResource {
	// 先从缓存中获取
	list, _ := s.GetResourceList(adminId)
	if list != nil && len(list) > 0 {
		return list
	}
	// 从数据库找
	var umsRR models.UmsAdminRoleRelation
	umsResources := umsRR.SelectRoleResourceRelationsByAdminId(s.DbFactory.GormMySQL, adminId)
	if umsResources != nil && len(umsResources) > 0 {
		// 存入缓存
		s.SetResourceList(adminId, umsResources, 0)
	}
	return umsResources
}

// UmsAdminLogout 用户登出
func (s UmsAdminService) UmsAdminLogout(context *gin.Context) {
	gin_common.Create(context)
}

// UmsAdminAuthTest 用户鉴权测试
// @Summary 用户鉴权测试
// @Description 用户鉴权测试
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/authTest [get]
func (s UmsAdminService) UmsAdminAuthTest(context *gin.Context) {
	m := map[string]any{
		"admin": map[string]any{
			"username": "admin",
			"token":    "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJob25nanVuNTAwIiwidXNlcl9pZCI6MTEsImNyZWF0ZWQiOiIyMDIzLTA2LTEwVDE0OjQ1OjA4LjEwMzc3NDYrMDg6MDAiLCJleHAiOjE2ODY5ODQzMDh9.PcLmIhxjenF36OPKmBX5ghPFgrfewSh_OUfT3dS-gUUL8UtyZFrg1gvxMbN8jZpOwJZIP5FQ7A1Yz1cfLl-Exg",
		},
		"hongjun500": map[string]any{
			"username": "hongjun500",
			"token":    "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJob25nanVuNTAwIiwidXNlcklkIjoxMSwiY3JlYXRlZCI6IjIwMjMtMDYtMTNUMTQ6NTg6MzQuNzMxNDA2NyswODowMCIsImV4cCI6MTY4NzI0NDMxNH0.Noxh09smVa6Y81wKdWPZ_II_6uf-mapYgBcie-ZVkM_23VoAoRBvP701Q7XEONDfp-J2HFOoCBbw3JeUuPr5Rw",
		},
	}
	gin_common.CreateSuccess(m, context)
}

// UmsAdminRefreshToken 刷新 token
// @Summary 刷新 token
// @Description 刷新 token
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/refreshToken [post]
func (s UmsAdminService) UmsAdminRefreshToken(context *gin.Context) {
	header := context.GetHeader(conf.GlobalJwtConfigProperties.TokenHeader)
	refreshToken, _ := security.RefreshToken(header)
	if refreshToken == "" {
		gin_common.CreateFail(gin_common.TokenExpired, context)
		return
	}
	tokenMap := make(map[string]string)
	tokenMap["token"] = refreshToken
	tokenMap["tokenHead"] = conf.GlobalJwtConfigProperties.TokenHead
	gin_common.CreateSuccess(tokenMap, context)
}

func (s UmsAdminService) UmsRoleList(adminId int64) []*models.UmsRole {
	var umsRoleRelation models.UmsAdminRoleRelation
	roles, _ := umsRoleRelation.SelectAllByAdminId(s.DbFactory.GormMySQL, adminId)
	return roles
}

// UmsAdminInfo 根据用户 ID 获取用户信息
// @Summary 根据用户 ID 获取用户信息
// @Description 根据用户 ID 获取用户信息
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/info [get]
func (s UmsAdminService) UmsAdminInfo(context *gin.Context) {
	userId := mid.GinJWTGetCurrentUserId(context)
	resultMap := make(map[string]interface{})
	resultMap["username"] = ""
	resultMap["menus"] = nil
	resultMap["icon"] = ""
	resultMap["roles"] = nil
	var umsAdmin models.UmsAdmin
	result, err := umsAdmin.SelectUmsAdminByUserId(s.DbFactory.GormMySQL, userId)
	if err != nil {
		gin_common.CreateSuccess(resultMap, context)
		return
	}
	if result == nil {
		gin_common.CreateSuccess(resultMap, context)
		return
	}

	var umsRole models.UmsRole
	umsMenus, err := umsRole.SelectMenu(s.DbFactory.GormMySQL, userId)
	if err != nil {
		gin_common.CreateSuccess(resultMap, context)
		return
	}
	resultMap["menus"] = umsMenus
	resultMap["username"] = result.Username
	resultMap["icon"] = result.Icon
	roles := s.UmsRoleList(userId)
	roleNames := make([]string, 0)
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}
	resultMap["roles"] = roleNames
	gin_common.CreateSuccess(resultMap, context)
}

// UmsAdminListPage 分页查询用户
// @Summary 分页查询用户
// @Description 分页查询用户
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param keyword query string false "keyword"
// @Param pageSize query int true "pageSize"
// @Param pageNum query int true "pageNum"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/list [get]
func (s UmsAdminService) UmsAdminListPage(context *gin.Context) {
	var request ums_admin.UmsAdminPageDTO
	err := context.ShouldBind(&request)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	var umsAdmin *models.UmsAdmin
	page, err := umsAdmin.SelectUmsAdminPage(s.DbFactory.GormMySQL, request.Username, request.PageNum, request.PageSize)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(page, context)
}

// UmsAdminItem 获取指定用户信息
// @Summary 获取指定用户信息
// @Description 获取指定用户信息
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param id path int true "用户 ID"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/{id} [get]
func (s UmsAdminService) UmsAdminItem(context *gin.Context) {
	var userDTO base.UserDTO
	err := context.ShouldBindUri(&userDTO)
	// 占位符
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	var umsAdmin models.UmsAdmin
	result, err := umsAdmin.SelectUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.Id)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(result, context)
}

// UmsAdminUpdate 修改指定用户信息
// @Summary 修改指定用户信息
// @Description 修改指定用户信息
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param user_id path int true "用户 ID"
// @Param request body ums_admin.UmsAdminUpdateDTO true "修改指定用户信息"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/update/{id} [post]
func (s UmsAdminService) UmsAdminUpdate(context *gin.Context) {
	var umsAdminUpdate ums_admin.UmsAdminUpdateDTO
	var userDTO base.UserDTO
	err := context.ShouldBindUri(&userDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	err = context.ShouldBind(&umsAdminUpdate)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	var umsAdmin models.UmsAdmin
	umsAdmin.Username = umsAdminUpdate.Username
	umsAdmin.Nickname = umsAdminUpdate.Nickname
	umsAdmin.Email = umsAdminUpdate.Email
	umsAdmin.Icon = umsAdminUpdate.Icon
	umsAdmin.Note = umsAdminUpdate.Note
	id, err := umsAdmin.UpdateUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.Id)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	s.DelAdmin(userDTO.UserId)
	gin_common.CreateSuccess(id, context)
}

// UmsAdminDelete 删除指定用户信息
// @Summary 删除指定用户信息
// @Description 删除指定用户信息
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param user_id path int true "用户 ID"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/delete/{id} [post]
func (s UmsAdminService) UmsAdminDelete(context *gin.Context) {
	var userDTO base.UserDTO
	err := context.ShouldBindUri(&userDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	var umsAdmin models.UmsAdmin
	id, err := umsAdmin.DeleteUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.UserId)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	s.DelAdmin(userDTO.UserId)
	s.DelResourceList(userDTO.UserId)
	gin_common.CreateSuccess(id, context)
}

// UmsAdminUpdateStatus 修改指定用户状态
// @Summary 修改指定用户状态
// @Description 修改指定用户状态
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param user_id path int true "用户 ID"
// @Param status formData int true "用户状态"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/updateStatus/{user_id} [post]
func (s UmsAdminService) UmsAdminUpdateStatus(context *gin.Context) {
	var pathVariableDTO base.PathVariableDTO
	err := context.ShouldBindUri(&pathVariableDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	umsAdmin := new(models.UmsAdmin)
	status, _ := strconv.Atoi(context.Query("status"))
	umsAdmin.Status = int64(status)
	umsAdmin.Id = pathVariableDTO.Id
	id, err := umsAdmin.UpdateUmsAdminStatusByUserId(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(gin_common.DatabaseError, context)
		return
	}
	s.DelAdmin(pathVariableDTO.Id)
	gin_common.CreateSuccess(id, context)
}

// UmsAdminUpdatePassword 修改指定用户密码
// @Summary 修改指定用户密码
// @Description 修改指定用户密码
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param request body ums_admin.UmsAdminUpdatePasswordDTO true "修改指定用户密码"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/updatePassword [post]
func (s UmsAdminService) UmsAdminUpdatePassword(context *gin.Context) {
	var umsAdminUpdatePassword ums_admin.UmsAdminUpdatePasswordDTO
	err := context.ShouldBind(&umsAdminUpdatePassword)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}

	var umsAdmin models.UmsAdmin
	umsAdmins, err := umsAdmin.SelectUmsAdminByUsername(s.DbFactory.GormMySQL, umsAdminUpdatePassword.Username)
	if err != nil {
		return
	}
	if umsAdmins == nil || len(umsAdmins) == 0 {
		gin_common.CreateFail("找不到该用户", context)
		return
	}

	getAdmin := umsAdmins[0]

	if !VerifyPassword(umsAdminUpdatePassword.OldPassword, getAdmin.Password) {
		gin_common.CreateFail("旧密码错误", context)
		return
	}
	hashPassword, err := HashPassword(umsAdminUpdatePassword.NewPassword)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	getAdmin.Password = hashPassword
	status, err := getAdmin.UpdateUmsAdminPasswordByUserId(s.DbFactory.GormMySQL)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	// 删除缓存的用户数据
	s.DelAdmin(getAdmin.Id)
	gin_common.CreateSuccess(status, context)
}

// UmsAdminRoleUpdate 修改指定用户角色
// @Summary 修改指定用户角色
// @Description 修改指定用户角色
// @Tags 后台用户管理
// @Accept	multipart/form-data
// @Produce  json
// @Param adminId formData int64 true "用户 ID"
// @Param roleIds formData []int64 true "角色 ID"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/role/update [post]
func (s UmsAdminService) UmsAdminRoleUpdate(context *gin.Context) {
	var umsAdminRoleDTO ums_admin.UmsAdminRoleDTO
	var err error
	count := int64(0)
	err = context.ShouldBind(&umsAdminRoleDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	adminId := umsAdminRoleDTO.AdminId
	roleIds := umsAdminRoleDTO.RoleIds

	if roleIds != nil && len(roleIds) > 0 {
		count = int64(len(roleIds))
	}
	// 先删除原有的绑定关系
	var umsAdminRoleRelation models.UmsAdminRoleRelation
	umsAdminRoleRelation.DelByAdminId(s.DbFactory.GormMySQL, adminId)
	// 建立新的绑定关系
	if count > 0 {
		var umsAdminRoleRelations []*models.UmsAdminRoleRelation
		for _, roleId := range roleIds {
			re := new(models.UmsAdminRoleRelation)
			re.AdminId = adminId
			re.RoleId = roleId
			umsAdminRoleRelations = append(umsAdminRoleRelations, re)
		}
		count, err = umsAdminRoleRelation.InsertList(s.DbFactory.GormMySQL, umsAdminRoleRelations)
		if err != nil {
			gin_common.CreateFail(gin_common.UnknownError, context)
			return
		}
	}
	s.DelResourceList(adminId)
	gin_common.CreateSuccess(count, context)
}

// UmsAdminRoleItem 获取指定用户的角色
// @Summary 获取指定用户的角色
// @Description 获取指定用户的角色
// @Tags 后台用户管理
// @Produce  json
// @Param adminId path int64 true "用户 ID"
// @Security GinJWTMiddleware
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/role/{adminId} [get]
func (s UmsAdminService) UmsAdminRoleItem(context *gin.Context) {
	var umsAdminRoleDTO ums_admin.UmsAdminRoleDTO
	err := context.ShouldBindUri(&umsAdminRoleDTO)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	var umsAdminRoleRelation models.UmsAdminRoleRelation
	umsAdminRoleRelations, err := umsAdminRoleRelation.SelectRoleList(s.DbFactory.GormMySQL, umsAdminRoleDTO.AdminId)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(umsAdminRoleRelations, context)
}
