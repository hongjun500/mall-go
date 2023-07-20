// @author hongjun500
// @date 2023/7/20 13:26
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

import (
	"testing"

	"github.com/hongjun500/mall-go/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSelectWithAttr(t *testing.T) {
	var pmsProductAttributeCategory models.PmsProductAttributeCategory
	attr, err := pmsProductAttributeCategory.SelectWithAttr(TestModelGormMySQL)
	if err != nil {
		return
	}
	assert.True(t, attr != nil)
}
