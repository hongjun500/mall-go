// @author hongjun500
// @date 2023/6/14 9:59
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request_dto/ums_member"
)

type UmsMemberLevelService struct {
	DbFactory *database.DbFactory
}

func NewUmsMemberLevelService(dbFactory *database.DbFactory) UmsMemberLevelService {
	return UmsMemberLevelService{DbFactory: dbFactory}
}

// UmsMemberLevelList 查看所有会员等级
// @Description 查看所有会员等级
// @Summary 查看所有会员等级
// @Tags 会员等级管理
// @Accept multipart/form-data
// @Produce application/json
// @Param defaultStatus query int false "是否为默认等级：0->不是；1->是"
// @Security GinJWTMiddleware
// @Success 200 {object} gin_common.GinCommonResponse
// @Router /memberLevel/list [get]
func (s UmsMemberLevelService) UmsMemberLevelList(context *gin.Context) {
	var dto ums_member.UmsMemberLevelListDTO
	if err := context.ShouldBind(&dto); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	var umsMemberLevel models.UmsMemberLevel
	list, err := umsMemberLevel.SelectByDefaultStatus(s.DbFactory.GormMySQL, dto.DefaultStatus)
	if err != nil {
		gin_common.CreateFail(context, gin_common.DatabaseError)
		return
	}
	gin_common.CreateSuccess(context, list)
}
