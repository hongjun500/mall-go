// @author hongjun500
// @date 2023/6/14 10:08
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package ums_member

type (
	UmsMemberLevelListDTO struct {
		DefaultStatus int `form:"defaultStatus" json:"defaultStatus" binding:"omitempty,oneof=0 1"`
	}
)
