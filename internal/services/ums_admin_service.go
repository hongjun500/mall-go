package services

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UmsAdminService struct {
	DbFactory *database.DbFactory
}

func NewUmsAdminService(dbFactory *database.DbFactory) *UmsAdminService {
	return &UmsAdminService{DbFactory: dbFactory}
}

// UmsAdminRequest 用户注册请求参数
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

// UmsAdminRegister 用户注册
func (s *UmsAdminService) UmsAdminRegister(context *gin.Context) {
	var request UmsAdminRequest
	err := context.ShouldBind(&request)
	if err != nil {
		// return nil
		return
	}
	// 检查用户名是否重复了
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.GetUmsAdminByUsername(s.DbFactory.GormMySQL, request.Username)
	if err != nil {
		return
	}
	if umsAdmins != nil && len(umsAdmins) > 0 {
		// return umsAdmin
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
		// return nil
		return
	}
	if register > 0 {
		// return umsAdmin
		return
	}
	// return nil
	return
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
