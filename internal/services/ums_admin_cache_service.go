// @author hongjun500
// @date 2023/6/11
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 用户缓存操作

package services

import (
	"context"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/pkg/constants"
	"github.com/hongjun500/mall-go/pkg/convert"
	"github.com/hongjun500/mall-go/pkg/redis"
	"strconv"
	"time"
)

// DelAdmin 删除用户缓存
func (s UmsAdminService) DelAdmin(adminId int64) {
	m := new(models.UmsAdmin)
	umsAdmin, err := m.SelectUmsAdminByUserId(s.DbFactory.GormMySQL, adminId)
	if err != nil {
		return
	}
	key := constants.RedisDatabase + constants.RedisKeyAdmin + umsAdmin.Username
	redis.Del(s.DbFactory.RedisCli, context.Background(), key)
}

// DelResourceList 删除后台用户资源列表缓存
func (s UmsAdminService) DelResourceList(adminId int64) {
	key := constants.RedisDatabase + constants.RedisKeyResourceList + strconv.FormatInt(adminId, 10)
	redis.Del(s.DbFactory.RedisCli, context.Background(), key)
}

// DelResourceListByRole 当角色相关资源信息改变时删除相关后台用户缓存
func (s UmsAdminService) DelResourceListByRole(roleId int64) {
	re := new(models.UmsAdminRoleRelation)
	roleRelations, err := re.SelectUmsAdminRoleRelationByRoleId(s.DbFactory.GormMySQL, roleId)
	if err != nil {
		return
	}
	if roleRelations != nil && len(roleRelations) > 0 {
		keyPrefix := constants.RedisDatabase + constants.RedisKeyResourceList
		keys := make([]string, 0)
		for _, roleRelation := range roleRelations {
			keys = append(keys, keyPrefix+strconv.FormatInt(roleRelation.AdminId, 10))
		}
		redis.Del(s.DbFactory.RedisCli, context.Background(), keys...)
	}
}

// DelResourceListByRoleIds 当角色相关资源信息改变时删除相关后台用户缓存
func (s UmsAdminService) DelResourceListByRoleIds(roleIds []int64) {
	re := new(models.UmsAdminRoleRelation)
	roleRelations, err := re.SelectUmsAdminRoleRelationInRoleId(s.DbFactory.GormMySQL, roleIds)
	if err != nil {
		return
	}
	if roleRelations != nil && len(roleRelations) > 0 {
		keyPrefix := constants.RedisDatabase + constants.RedisKeyResourceList
		keys := make([]string, 0)
		for _, roleRelation := range roleRelations {
			keys = append(keys, keyPrefix+strconv.FormatInt(roleRelation.AdminId, 10))
		}
		redis.Del(s.DbFactory.RedisCli, context.Background(), keys...)
	}
}

// DelResourceListByResource 当资源信息改变时，删除资源项目后台用户缓存
func (s UmsAdminService) DelResourceListByResource(resourceId int64) {
	rr := new(models.UmsRoleResourceRelation)
	roleResourceRelations, err := rr.SelectAdminIdsByResourceId(s.DbFactory.GormMySQL, resourceId)
	if err != nil {
		return
	}
	if roleResourceRelations != nil && len(roleResourceRelations) > 0 {
		keys := make([]string, 0)
		keyPrefix := constants.RedisDatabase + constants.RedisKeyResourceList
		for _, adminId := range roleResourceRelations {
			keys = append(keys, keyPrefix+strconv.FormatInt(adminId, 10))
		}
		redis.Del(s.DbFactory.RedisCli, context.Background(), keys...)
	}
}

// GetAdminByUsername 获取用户缓存
func (s UmsAdminService) GetAdminByUsername(username string) (models.UmsAdmin, error) {
	key := constants.RedisDatabase + constants.RedisKeyAdmin + username
	umsAdminJsonStr := redis.Get(s.DbFactory.RedisCli, context.Background(), key)
	var umsAdmin models.UmsAdmin
	err := convert.JsonToStruct(umsAdminJsonStr, &umsAdmin)
	if err != nil {
		return umsAdmin, err
	}
	return umsAdmin, nil
}

// SetAdmin 设置用户缓存
func (s UmsAdminService) SetAdmin(umsAdmin models.UmsAdmin, exp time.Duration) {
	key := constants.RedisDatabase + constants.RedisKeyAdmin + umsAdmin.Username
	jsonStr := convert.StructToJson(umsAdmin)
	redis.SetExpiration(s.DbFactory.RedisCli, context.Background(), key, jsonStr, exp)
}

// GetResourceList 获取后台用户资源列表
func (s UmsAdminService) GetResourceList(adminId int64) ([]models.UmsResource, error) {
	key := constants.RedisDatabase + constants.RedisKeyResourceList + strconv.FormatInt(adminId, 10)
	umsResourceJsonStr := redis.LRange(s.DbFactory.RedisCli, context.Background(), key, 0, -1)
	var umsResources []models.UmsResource
	for _, jsonStr := range umsResourceJsonStr {
		var resource models.UmsResource
		err := convert.JsonToStruct(jsonStr, &resource)
		if err != nil {
			return umsResources, err
		}
		umsResources = append(umsResources, resource)
	}
	return umsResources, nil
}

// SetResourceList 设置后台用户资源列表
func (s UmsAdminService) SetResourceList(adminId int64, resourceList []models.UmsResource, exp time.Duration) {
	key := constants.RedisDatabase + constants.RedisKeyResourceList + strconv.FormatInt(adminId, 10)
	sliceStructToJson := convert.SliceStructToJson(resourceList)
	redis.LRPush(s.DbFactory.RedisCli, context.Background(), key, sliceStructToJson, exp)
}
