// @author hongjun500
// @date 2023/6/14 10:11
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package models

import "gorm.io/gorm"

type UmsMemberLevel struct {
	Model
	// 等级名称
	Name string `gorm:"column:name;" json:"name" `
	// 等级需要的成长值
	GrowthPoint int `gorm:"column:growth_point;" json:"growthPoint" `
	// 是否为默认等级：0->不是；1->是
	DefaultStatus int `gorm:"column:default_status;" json:"defaultStatus" `
	// 免运费标准
	FreeFreightPoint float64 `gorm:"column:free_freight_point;" json:"freeFreightPoint" `
	// 每次评价获取的成长值
	CommentGrowthPoint int `gorm:"column:comment_growth_point;" json:"commentGrowthPoint" `
	// 是否有免邮特权
	PriviledgeFreeFreight int `gorm:"column:priviledge_free_freight;" json:"priviledgeFreeFreight" `
	// 是否有签到特权
	PriviledgeSignIn int `gorm:"column:priviledge_sign_in;" json:"priviledgeSignIn" `
	// 是否有评论获奖励特权
	PriviledgeComment int `gorm:"column:priviledge_comment;" json:"priviledgeComment" `
	// 是否有专享活动特权
	PriviledgePromotion int `gorm:"column:priviledge_promotion;" json:"priviledgePromotion" `
	// 是否有会员价格特权
	PriviledgeMemberPrice int `gorm:"column:priviledge_member_price;" json:"priviledgeMemberPrice" `
	// 是否有生日特权
	PriviledgeBirthday int `gorm:"column:priviledge_birthday;" json:"priviledgeBirthday" `
	// 备注
	Note string `gorm:"column:note;" json:"note" `
}

func (*UmsMemberLevel) TableName() string {
	return "ums_member_level"
}

func (umsMemberLevel *UmsMemberLevel) SelectByDefaultStatus(db *gorm.DB, defaultStatus int) ([]*UmsMemberLevel, error) {
	var umsMemberLevels []*UmsMemberLevel
	err := db.Where("default_status = ?", defaultStatus).Find(&umsMemberLevels).Error
	if err != nil {
		return nil, err
	}
	return umsMemberLevels, nil
}
