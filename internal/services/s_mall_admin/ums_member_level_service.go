//	@author	hongjun500
//	@date	2023/6/14 9:59
//	@tool	ThinkPadX1隐士
// Created with GoLand 2022.2
// Description:

package s_mall_admin

import (
	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
	"github.com/hongjun500/mall-go/internal/request/admin_dto"
)

type UmsMemberLevelService struct {
	DbFactory *database.DbFactory
}

func NewUmsMemberLevelService(dbFactory *database.DbFactory) UmsMemberLevelService {
	return UmsMemberLevelService{DbFactory: dbFactory}
}

// UmsMemberLevelList 查看所有会员等级
func (s UmsMemberLevelService) UmsMemberLevelList(dto admin_dto.UmsMemberLevelListDTO) ([]*models.UmsMemberLevel, error) {
	var umsMemberLevel models.UmsMemberLevel
	list, err := umsMemberLevel.SelectByDefaultStatus(s.DbFactory.GormMySQL, dto.DefaultStatus)
	if err != nil {
		return nil, err
	}
	return list, nil
}
