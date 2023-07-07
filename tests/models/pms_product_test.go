// @author hongjun500
// @date 2023/6/26 17:32
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package models

import (
	"testing"

	"github.com/hongjun500/mall-go/internal/models"
)

func TestSelectProductInfoById(t *testing.T) {
	var product models.PmsProduct
	pmsProducts, err := product.SelectProductInfoById(TestModelDbFactory, 0)
	if err != nil {
		t.Errorf("SelectProductInfoById() error = %v", err)
		return
	}
	t.Logf("SelectProductInfoById() got = %v", pmsProducts)
}
