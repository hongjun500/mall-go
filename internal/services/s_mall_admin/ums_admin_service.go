package s_mall_admin

import (
	"crypto/rand"
	"time"

	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/gorm_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/base_dto"
	"github.com/hongjun500/mall-go/internal/request/ums_admin_dto"
	"github.com/hongjun500/mall-go/internal/services"
	"github.com/hongjun500/mall-go/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type UmsAdminService struct {
	DbFactory *database.DbFactory
}

func NewUmsAdminService(dbFactory *database.DbFactory) UmsAdminService {
	return UmsAdminService{DbFactory: dbFactory}
}

func (s UmsAdminService) getAdminByUsername(username string) (models.UmsAdmin, error) {
	// 先从缓存中获取
	cacheAdmin, _ := services.GetAdmin(s.DbFactory, username)
	if cacheAdmin.Id > 0 {
		return cacheAdmin, nil
	}
	// 没有再从数据库中获取
	var umsAdmin models.UmsAdmin
	umsAdmins, err := umsAdmin.SelectUmsAdminByUsername(s.DbFactory.GormMySQL, username)
	if err != nil {
		return umsAdmin, err
	}
	if umsAdmins != nil && len(umsAdmins) > 0 {
		umsAdmin = *umsAdmins[0]
		// 存入缓存
		services.SetAdmin(s.DbFactory, umsAdmin, 0)
		return umsAdmin, nil
	}
	return umsAdmin, nil
}

func (s UmsAdminService) GetResource(adminId int64) []models.UmsResource {
	// 先从缓存中获取
	list, _ := services.GetResourceList(s.DbFactory, adminId)
	if list != nil && len(list) > 0 {
		return list
	}
	// 从数据库找
	var umsRR models.UmsAdminRoleRelation
	umsResources := umsRR.SelectRoleResourceRelationsByAdminId(s.DbFactory.GormMySQL, adminId)
	if umsResources != nil && len(umsResources) > 0 {
		// 存入缓存
		services.SetResourceList(s.DbFactory, adminId, umsResources, 0)
	}
	return umsResources
}

// HashPassword 加密密码
func hashPassword(password string) (string, error) {
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

func (s UmsAdminService) UmsAdminRegister(request ums_admin_dto.UmsAdminRegisterDTO) error {
	// 检查用户名是否重复了
	var umsAdmin *models.UmsAdmin
	umsAdmins, err := umsAdmin.SelectUmsAdminByUsername(s.DbFactory.GormMySQL, request.Username)
	if err != nil {
		return gin_common.NewGinCommonError(err.Error())
	}
	if umsAdmins != nil && len(umsAdmins) > 0 {
		return gin_common.NewGinCommonError(gin_common.UsernameAlreadyExists)
	}
	// 密码加密
	hash, _ := hashPassword(request.Password)

	// 创建用户参数
	umsAdmin = &models.UmsAdmin{
		Username: request.Username,
		Password: hash,
		Icon:     request.Icon,
		Email:    request.Email,
		Nickname: request.Nickname,
		Note:     request.Note,
	}
	registers, err := umsAdmin.InsertUmsAdmin(s.DbFactory.GormMySQL)
	if err != nil || registers <= 0 {
		return gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return nil
}

func (s UmsAdminService) UmsAdminLogin(umsAdminLogin ums_admin_dto.UmsAdminLoginDTO, others ...string) (map[string]string, error) {

	umsAdmin, err := s.getAdminByUsername(umsAdminLogin.Username)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	if umsAdmin.Id == 0 {
		return nil, gin_common.NewGinCommonError(gin_common.UsernameOrPasswordError)
	}
	if !VerifyPassword(umsAdminLogin.Password, umsAdmin.Password) {
		return nil, gin_common.NewGinCommonError(gin_common.UsernameOrPasswordError)
	}
	if umsAdmin.Status == 0 {
		return nil, gin_common.NewGinCommonError(gin_common.AccountLocked)
	}
	if "admin" != umsAdmin.Username {
		// 获取用户资源
		umsResources := s.GetResource(umsAdmin.Id)
		// 添加用户可访问资源策略
		security.AddPolicyFromResource(security.Enforcer, umsAdmin.Username, umsResources)
	}

	token := security.GenerateToken(umsAdmin.Username, umsAdmin.Id)
	if token == "" {
		return nil, gin_common.NewGinCommonError(gin_common.TokenGenFail)
	}
	now := time.Now()
	umsAdmin.LoginTime = &now
	// 更新登录时间
	_, _ = umsAdmin.UpdateUmsAdminLoginTimeByUserId(s.DbFactory.GormMySQL)

	umsAdminLoginLog := new(models.UmsAdminLoginLog)
	umsAdminLoginLog.AdminId = umsAdmin.Id
	umsAdminLoginLog.Ip = others[0]
	umsAdminLoginLog.Address = others[1]
	umsAdminLoginLog.UserAgent = others[2]
	// 记录登录日志
	_, _ = umsAdminLoginLog.SaveLoginLog(s.DbFactory.GormMySQL)

	tokenMap := make(map[string]string)
	tokenMap["token"] = token
	tokenMap["tokenHead"] = conf.GlobalJwtConfigProperties.TokenHead
	return tokenMap, nil
}

func (s UmsAdminService) UmsRoleList(adminId int64) []*models.UmsRole {
	var umsRoleRelation models.UmsAdminRoleRelation
	roles, _ := umsRoleRelation.SelectAllByAdminId(s.DbFactory.GormMySQL, adminId)
	return roles
}

func (s UmsAdminService) UmsAdminInfo(username string) (map[string]any, error) {
	// username := mid.GinJWTGetCurrentUsername(context)
	resultMap := make(map[string]any)
	resultMap["username"] = ""
	resultMap["menus"] = nil
	resultMap["icon"] = ""
	resultMap["roles"] = nil
	// var umsAdmin models.UmsAdmin
	// result, err := umsAdmin.SelectUmsAdminByUserId(s.DbFactory.GormMySQL, userId)
	result, err := s.getAdminByUsername(username)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	if result.Id == 0 {
		return resultMap, nil
	}

	var umsRole models.UmsRole
	umsMenus, err := umsRole.SelectMenu(s.DbFactory.GormMySQL, result.Id)
	if err != nil {
		return resultMap, nil
	}
	resultMap["menus"] = umsMenus
	resultMap["username"] = result.Username
	resultMap["icon"] = result.Icon
	roles := s.UmsRoleList(result.Id)
	roleNames := make([]string, 0)
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}
	resultMap["roles"] = roleNames
	return resultMap, nil
}

// UmsAdminListPage 分页查询用户
func (s UmsAdminService) UmsAdminListPage(request ums_admin_dto.UmsAdminPageDTO) (gorm_common.CommonPage, error) {
	var umsAdmin *models.UmsAdmin
	page, err := umsAdmin.SelectUmsAdminPage(s.DbFactory.GormMySQL, request.Username, request.PageNum, request.PageSize)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return page, nil
}

func (s UmsAdminService) UmsAdminItem(userDTO base_dto.UserDTO) (*models.UmsAdmin, error) {
	var umsAdmin models.UmsAdmin
	getUmsAdmin, err := umsAdmin.SelectUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.Id)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.UnknownError)
	}
	return getUmsAdmin, nil
}

// UmsAdminUpdate 修改指定用户信息
func (s UmsAdminService) UmsAdminUpdate(userDTO base_dto.UserDTO, umsAdminUpdate ums_admin_dto.UmsAdminUpdateDTO) (int64, error) {
	var umsAdmin models.UmsAdmin
	umsAdmin.Username = umsAdminUpdate.Username
	umsAdmin.Nickname = umsAdminUpdate.Nickname
	umsAdmin.Email = umsAdminUpdate.Email
	umsAdmin.Icon = umsAdminUpdate.Icon
	umsAdmin.Note = umsAdminUpdate.Note
	id, err := umsAdmin.UpdateUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.Id)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	services.DelAdmin(s.DbFactory, userDTO.UserId)
	return id, nil
}

// UmsAdminDelete 删除指定用户信息
func (s UmsAdminService) UmsAdminDelete(userDTO base_dto.UserDTO) (int64, error) {
	var umsAdmin models.UmsAdmin
	id, err := umsAdmin.DeleteUmsAdminByUserId(s.DbFactory.GormMySQL, userDTO.UserId)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	services.DelAdmin(s.DbFactory, userDTO.UserId)
	services.DelResourceList(s.DbFactory, userDTO.UserId)
	return id, nil
}

// UmsAdminUpdateStatus 修改指定用户状态
func (s UmsAdminService) UmsAdminUpdateStatus(id, status int64) (int64, error) {
	umsAdmin := new(models.UmsAdmin)
	umsAdmin.Status = status
	umsAdmin.Id = id
	id, err := umsAdmin.UpdateUmsAdminStatusByUserId(s.DbFactory.GormMySQL)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	services.DelAdmin(s.DbFactory, id)
	return id, nil
}

// UmsAdminUpdatePassword 修改指定用户密码
func (s UmsAdminService) UmsAdminUpdatePassword(umsAdminUpdatePasswordDTO ums_admin_dto.UmsAdminUpdatePasswordDTO) (int64, error) {
	var umsAdmin models.UmsAdmin
	umsAdmins, err := umsAdmin.SelectUmsAdminByUsername(s.DbFactory.GormMySQL, umsAdminUpdatePasswordDTO.Username)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	if umsAdmins == nil || len(umsAdmins) == 0 {
		return 0, gin_common.NewGinCommonError("找不到该用户")
	}
	getAdmin := umsAdmins[0]
	if !VerifyPassword(umsAdminUpdatePasswordDTO.OldPassword, getAdmin.Password) {
		return 0, gin_common.NewGinCommonError("旧密码错误")
	}
	hash, err := hashPassword(umsAdminUpdatePasswordDTO.NewPassword)
	if err != nil {
		return 0, gin_common.NewGinCommonError("密码加密失败")
	}
	getAdmin.Password = hash
	status, err := getAdmin.UpdateUmsAdminPasswordByUserId(s.DbFactory.GormMySQL)
	if err != nil {
		return 0, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	// 删除缓存的用户数据
	services.DelAdmin(s.DbFactory, getAdmin.Id)
	return status, nil
}

// UmsAdminRoleUpdate 修改指定用户角色
func (s UmsAdminService) UmsAdminRoleUpdate(umsAdminRoleDTO ums_admin_dto.UmsAdminRoleDTO) (int64, error) {
	adminId := umsAdminRoleDTO.AdminId
	roleIds := umsAdminRoleDTO.RoleIds
	var count int64
	if roleIds != nil && len(roleIds) > 0 {
		count = int64(len(roleIds))
	}
	// 先删除原有的绑定关系
	var umsAdminRoleRelation models.UmsAdminRoleRelation
	umsAdminRoleRelation.DelByAdminId(s.DbFactory.GormMySQL, adminId)
	// 建立新地绑定关系
	if count > 0 {
		var umsAdminRoleRelations []*models.UmsAdminRoleRelation
		for _, roleId := range roleIds {
			re := new(models.UmsAdminRoleRelation)
			re.AdminId = adminId
			re.RoleId = roleId
			umsAdminRoleRelations = append(umsAdminRoleRelations, re)
		}
		rows, err := umsAdminRoleRelation.InsertList(s.DbFactory.GormMySQL, umsAdminRoleRelations)
		if err != nil {
			return rows, gin_common.NewGinCommonError(gin_common.DatabaseError)
		}
	}
	services.DelResourceList(s.DbFactory, adminId)
	return count, nil
}

// UmsAdminRoleItem 获取指定用户的角色
func (s UmsAdminService) UmsAdminRoleItem(umsAdminRoleDTO ums_admin_dto.UmsAdminRoleDTO) ([]*models.UmsRole, error) {
	var umsAdminRoleRelation models.UmsAdminRoleRelation
	umsAdminRoleRelations, err := umsAdminRoleRelation.SelectRoleList(s.DbFactory.GormMySQL, umsAdminRoleDTO.AdminId)
	if err != nil {
		return nil, gin_common.NewGinCommonError(gin_common.DatabaseError)
	}
	return umsAdminRoleRelations, nil
}
