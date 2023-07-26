// @author hongjun500
// @date 2023/7/26 14:22
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package admin_dto

import "github.com/hongjun500/mall-go/internal/request/base_dto"

type (
	PmsBrandDTO struct {
		// 品牌名称
		Name string `json:"name" binding:"required"`
		// 首字母
		FirstLetter string `json:"first_letter"`
		// 排序
		Sort int `json:"sort" min:"0"`
		// 是否为品牌制造商：0->不是；1->是
		FactoryStatus int `json:"factory_status" min:"0" max:"1"`
		// 是否显示
		ShowStatus int `json:"show_status" min:"0" max:"1"`
		// logo
		Logo string `json:"logo" binding:"required"`
		// 专区大图
		BigPic string `json:"big_pic"`
		// 品牌故事
		BrandStory string `json:"brand_story"`
	}
	PsmBrandPageDTO struct {
		base_dto.PageDTO
		Keyword    string `json:"keyword" query:"keyword"`
		ShowStatus int    `json:"show_status" query:"show_status"`
	}
)
