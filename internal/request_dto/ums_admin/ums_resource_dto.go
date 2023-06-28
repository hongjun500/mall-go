//	@author	hongjun500
//	@date	2023/6/13 11:22
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package ums_admin

import "github.com/hongjun500/mall-go/internal/request_dto/base"

type UmsResourceCreateDTO struct {
	// 资源名称
	Name string `json:"name"`
	// 资源URL
	Url string `json:"url"`
	// 描述
	Description string `json:"description"`
	// 资源分类ID
	CategoryId int64 `json:"categoryId"`
}

type UmsResourcePageListDTO struct {
	base.PageDTO
	// 资源分类ID
	CategoryId int64 `form:"categoryId"`
	// 资源名称模糊关键字
	NameKeyword string `form:"nameKeyword"`
	// 资源URL模糊关键字
	UrlKeyword string `form:"urlKeyword"`
}

type ()
