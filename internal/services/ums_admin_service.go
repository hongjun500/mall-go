package services

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request_dto/base"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
	"github.com/hongjun500/mall-go/pkg/jwt"
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
// @Param request body ums_admin.UmsAdminRegisterRequest true "用户注册"
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/register [post]
func (s UmsAdminService) UmsAdminRegister(context *gin.Context) {
	var request ums_admin.UmsAdminRegisterRequest
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
// @Param request body ums_admin.UmsAdminLogin true "用户登录"
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/login [post]
func (s UmsAdminService) UmsAdminLogin(context *gin.Context) {
	var umsAdminLogin ums_admin.UmsAdminLogin
	err := context.ShouldBind(&umsAdminLogin)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		// 这些地方需要手动 return
		return
	}
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.SelectUmsAdminByUsername(s.DbFactory.GormMySQL, umsAdminLogin.Username)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
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

	token := jwt.GenerateToken(umsAdmin.Username)
	// todo 添加登录记录

	if token == "" {
		gin_common.CreateFail(gin_common.TokenGenFail, context)
		return
	}
	tokenMap := make(map[string]string)
	tokenMap["token"] = token
	tokenMap["tokenHead"] = conf.GlobalJwtConfigProperties.TokenHead
	gin_common.CreateSuccess(tokenMap, context)
}

// UmsAdminLogout 用户登出
func (s UmsAdminService) UmsAdminLogout(context *gin.Context) {
	gin_common.Create(context)
}

// UmsAdminRefreshToken 刷新 token
// @Summary 刷新 token
// @Description 刷新 token
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJob25nanVuNTAwIiwiY3JlYXRlZCI6IjIwMjMtMDYtMDVUMTQ6Mjk6NDIuODg5MjMzNCswODowMCIsImV4cCI6MTY4NjU1MTM4Mn0.O1KpbIsAXWkyFgUXXN3isTXUSRq9202cNDjiOz4hDOITNA9Scmrmw_3_T1Bk53hKORpm8cbzL4F6_y0eGqs2nw"
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/refreshToken [post]
func (s UmsAdminService) UmsAdminRefreshToken(context *gin.Context) {
	header := context.GetHeader(conf.GlobalJwtConfigProperties.TokenHeader)
	refreshToken, _ := jwt.RefreshToken(header)
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
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/info [get]
func (s UmsAdminService) UmsAdminInfo(context *gin.Context) {
	var userDTO base.UserDTO
	err := context.ShouldBindUri(&userDTO)
	// 占位符
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	resultMap := make(map[string]interface{})
	resultMap["username"] = ""
	resultMap["menus"] = nil
	resultMap["icon"] = ""
	resultMap["roles"] = nil
	var umsAdmin models.UmsAdmin
	result, err := umsAdmin.SelectUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.UserId)
	if err != nil {
		gin_common.CreateSuccess(resultMap, context)
		return
	}
	if result == nil {
		gin_common.CreateSuccess(resultMap, context)
		return
	}

	var umsRole models.UmsRole
	umsMenus, err := umsRole.SelectMenu(s.DbFactory.GormMySQL, userDTO.UserId)
	if err != nil {
		gin_common.CreateSuccess(resultMap, context)
		return
	}
	resultMap["menus"] = umsMenus
	resultMap["username"] = result.Username
	resultMap["icon"] = result.Icon
	roles := s.UmsRoleList(userDTO.UserId)
	var roleNames []string
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
// @Param request body ums_admin.UmsAdminPage true "分页查询用户"
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/list [post]
func (s UmsAdminService) UmsAdminListPage(context *gin.Context) {
	var request ums_admin.UmsAdminPage
	err := context.ShouldBind(&request)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.SelectUmsAdminPage(s.DbFactory.GormMySQL, request.Username, request.PageNum, request.PageSize)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(umsAdmins, context)
}

// UmsAdminItem 根据用户 ID 获取用户信息
// @Summary 根据用户 ID 获取用户信息
// @Description 根据用户 ID 获取用户信息
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param id path int true "用户 ID"
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/:user_id [get]
func (s UmsAdminService) UmsAdminItem(context *gin.Context) {
	var userDTO base.UserDTO
	err := context.ShouldBindUri(&userDTO)
	// 占位符
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	var umsAdmin models.UmsAdmin
	result, err := umsAdmin.SelectUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.UserId)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(result, context)
}

// UmsAdminUpdate 更新用户信息
// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags 后台用户管理
// @Accept  json
// @Produce  json
// @Param id path int true "用户 ID"
// @Param request body ums_admin.UmsAdminUpdate true "更新用户信息"
// @Success 200 {object}  gin_common.GinCommonResponse
// @Router /admin/{user_id} [post]
func (s UmsAdminService) UmsAdminUpdate(context *gin.Context) {
	umsAdminUpdate := new(ums_admin.UmsAdminUpdate)
	userDTO := new(base.UserDTO)
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
	id, err := umsAdmin.UpdateUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.UserId)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(id, context)
}
