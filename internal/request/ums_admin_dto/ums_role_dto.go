//	@author	hongjun500
//	@date	2023/6/13 15:54
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 角色请求传输数据对象

package ums_admin_dto

import "github.com/hongjun500/mall-go/internal/request/base_dto"

type (
	UmsRolePathVariableDTO struct {
		// 角色ID
		RoleId int64 `uri:"roleId"`
	}
	UmsRoleCreateDTO struct {
		// 角色名称
		Name string `json:"name" binding:"required"`
		// 角色描述
		Description string `json:"description"`
		// 后台用户数量
		AdminCount int `json:"adminCount"`
		// 启用状态：0->禁用；1->启用
		Status int `json:"status"`
		// 排序
		Sort int `json:"sort"`
	}
	IdsDTO struct {
		Ids []int64 `json:"ids" form:"ids" binding:"required"`
	}
	UmsRoleListPageDTO struct {
		base_dto.PageDTO
		Keyword string `json:"keyword" form:"keyword"`
	}
	UmsRoleStatusDTO struct {
		// 启用状态：0->禁用；1->启用
		Status int `form:"status" binding:"required"`
	}
	UmsRoleAllocMenuDTO struct {
		// 角色ID
		RoleId int64 `json:"roleId" form:"roleId" binding:"required"`
		// 菜单ID列表
		MenuIds []int64 `json:"menuIds" form:"menuIds" binding:"required"`
	}
	UmsRoleAllocResourceDTO struct {
		// 角色ID
		RoleId int64 `json:"roleId" form:"roleId" binding:"required"`
		// 资源ID列表
		ResourceIds []int64 `json:"resourceIds" form:"resourceIds" binding:"required"`
	}
)
