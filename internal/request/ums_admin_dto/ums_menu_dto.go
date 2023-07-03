//	@author	hongjun500
//	@date	2023/6/12 13:41
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package ums_admin_dto

// UmsMenuCreateDTO 添加后台菜单
type UmsMenuCreateDTO struct {
	// 父级ID
	ParentId int64 `json:"parentId"`
	// 菜单名称
	Title string `json:"title"`
	// 菜单级数
	Level int `json:"level"`
	// 菜单排序
	Sort int `json:"sort"`
	// 前端名称
	Name string `json:"name"`
	// 前端图标
	Icon string `json:"icon"`
	// 前端隐藏
	Hidden int `json:"hidden"`

	Id int64 `json:"id" uri:"id"`
}

type UmsMenuListDTO struct {
	ParentId int64 `json:"parentId" uri:"parentId"`
}

type UmsMenuHiddenDTO struct {
	Id     int64 `json:"id" uri:"id"`
	Hidden int64 `json:"hidden"`
}
