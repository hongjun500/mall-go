package services

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/gin_common/mid"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_admin"
	"golang.org/x/crypto/bcrypt"
)

type UmsAdminService struct {
	DbFactory *database.DbFactory
}

func NewUmsAdminService(dbFactory *database.DbFactory) *UmsAdminService {
	return &UmsAdminService{DbFactory: dbFactory}
}

// UmsAdminRegister 用户注册
func (s *UmsAdminService) UmsAdminRegister(context *gin.Context) {
	var request ums_admin.UmsAdminRequest
	err := context.ShouldBind(&request)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		// 这些地方需要手动 return
		return
	}
	// 检查用户名是否重复了
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.GetUmsAdminByUsername(s.DbFactory.GormMySQL, request.Username)
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
	register, err := umsAdmin.CreateUmsAdmin(s.DbFactory.GormMySQL)
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

func (s *UmsAdminService) UmsAdminLogin(context *gin.Context) {
	var request ums_admin.UmsAdminRequest
	err := context.ShouldBind(&request)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		// 这些地方需要手动 return
		// return
	}
	/*var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.GetUmsAdminByUsername(s.DbFactory.GormMySQL, request.Username)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	if umsAdmins == nil || len(umsAdmins) == 0 {
		gin_common.CreateFail(gin_common.UsernameOrPasswordError, context)
		return
	}
	umsAdmin = umsAdmins[0]
	if !VerifyPassword(request.Password, umsAdmin.Password) {
		gin_common.CreateFail(gin_common.UsernameOrPasswordError, context)
		return
	}
	if umsAdmin.Status == 0 {
		gin_common.CreateFail(gin_common.AccountLocked, context)
		return
	}*/

	if request.Username != "hongjun500" || request.Password != "123456" {
		gin_common.CreateFail(gin_common.UsernameOrPasswordError, context)
		return
	}
	umsAdmin := &models.UmsAdmin{
		Username: request.Username,
		Password: request.Password,
	}
	token := mid.GenerateToken(umsAdmin.Username)
	// todo 添加登录记录

	if token == "" {
		gin_common.CreateFail(gin_common.TokenGenFail, context)
		return
	}
	tokenMap := make(map[string]string)
	tokenMap["token"] = token
	tokenMap["tokenHead"] = "Bearer "
	gin_common.CreateSuccess(tokenMap, context)
}

func (s *UmsAdminService) UmsAdminRefreshToken(context *gin.Context) {
	header := context.GetHeader("Authorization")
	refreshToken, _ := mid.RefreshToken(header)
	if refreshToken == "" {
		gin_common.CreateFail(gin_common.TokenExpired, context)
		return
	}
	tokenMap := make(map[string]string)
	tokenMap["token"] = refreshToken
	tokenMap["tokenHead"] = "Bearer "
	gin_common.CreateSuccess(tokenMap, context)
}

// UmsAdminInfo 根据用户 ID 获取用户信息
func (s *UmsAdminService) UmsAdminInfo(context *gin.Context) {
	id := context.MustGet("id")
	if id == "" || id == nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
	}
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.GetUmsAdminByUserId(s.DbFactory.GormMySQL, id.(int64))
	if err != nil {
		gin_common.Create(context)
		return
	}
	if umsAdmins == nil {
		gin_common.Create(context)
		return
	}
	gin_common.CreateSuccess(umsAdmins, context)
}

func (s *UmsAdminService) UmsAdminListPage(context *gin.Context) {
	var request ums_admin.UmsAdminRequest
	err := context.ShouldBind(&request)
	if err != nil {
		gin_common.CreateFail(gin_common.ParameterValidationError, context)
		return
	}
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.GetUmsAdminListPage(s.DbFactory.GormMySQL, request.Username, request.PageNum, request.PageSize)
	if err != nil {
		gin_common.CreateFail(gin_common.UnknownError, context)
		return
	}
	gin_common.CreateSuccess(umsAdmins, context)
}
