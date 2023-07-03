//	@author	hongjun500
//	@date	2023/6/14 10:01
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package r_mall_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/mall-go/internal/gin_common"
	"github.com/hongjun500/mall-go/internal/request/ums_member_dto"
	"github.com/hongjun500/mall-go/internal/services/s_mall_admin"
)

type UmsMemberLevelRouter struct {
	s_mall_admin.UmsMemberLevelService
}

func NewUmsMemberLevelRouter(service s_mall_admin.UmsMemberLevelService) *UmsMemberLevelRouter {
	return &UmsMemberLevelRouter{UmsMemberLevelService: service}
}

// GroupUmsMemberLevelRouter 会员等级管理路由
func (router *UmsMemberLevelRouter) GroupUmsMemberLevelRouter(umsMemberLevelGroup *gin.RouterGroup) {
	{
		// 查看所有会员等级
		umsMemberLevelGroup.GET("/list", router.list)
	}
}

// list 查看所有会员等级
//
//	@Description	查看所有会员等级
//	@Summary		查看所有会员等级
//	@Tags			会员等级管理
//	@Accept			multipart/form-data
//	@Produce		application/json
//	@Param			defaultStatus	query	int	false	"是否为默认等级：0->不是；1->是"
//	@Security		GinJWTMiddleware
//	@Success		200	{object}	gin_common.GinCommonResponse
//	@Router			/memberLevel/list [get]
func (router *UmsMemberLevelRouter) list(context *gin.Context) {
	var dto ums_member_dto.UmsMemberLevelListDTO
	if err := context.ShouldBind(&dto); err != nil {
		gin_common.CreateFail(context, gin_common.ParameterValidationError)
		return
	}
	list, err := router.UmsMemberLevelService.UmsMemberLevelList(dto)
	if err != nil {
		gin_common.CreateFail(context, err)
		return
	}
	gin_common.CreateSuccess(context, list)
}
