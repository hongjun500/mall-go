// @author hongjun500
// @date 2023/7/20 16:38
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

import (
	"testing"

	"github.com/hongjun500/mall-go/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSelectListFromProductAttrInfo(t *testing.T) {
	item := new(models.PmsProductAttribute)
	info, err := item.SelectListFromProductAttrInfo(TestModelGormMySQL, 26)
	if err != nil {
		return
	}
	assert.NotNil(t, info)
	assert.Lenf(t, info, 1, "info len is not 1")
}
