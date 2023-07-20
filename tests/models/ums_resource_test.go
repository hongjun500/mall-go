package models

import (
	"testing"

	"github.com/hongjun500/mall-go/internal/models"
)

func TestResourceListAll(t *testing.T) {
	var umsResource models.UmsResource
	all, err := umsResource.SelectAll(TestModelGormMySQL)
	if err != nil {
		return
	}
	t.Log("list all success, len = ", len(all))
}

func TestResourceListPage(t *testing.T) {
	var umsResource models.UmsResource
	page, err := umsResource.SelectPage(TestModelGormMySQL, 0, "", "", 1, 10)
	if err != nil {
		return
	}
	t.Log("list page success, len = ", page.Total)
}
