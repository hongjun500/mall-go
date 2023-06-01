package models

import (
	"github.com/hongjun500/mall-go/internal/initialize"
	"github.com/hongjun500/mall-go/internal/models"
	"testing"
)

func TestResourceListAll(t *testing.T) {
	var umsResource models.UmsResource
	all, err := umsResource.ListAll(initialize.SqlSession.DbMySQL)
	if err != nil {
		return
	}
	t.Log("list all success, len = ", len(all))
}

func TestResourceListPage(t *testing.T) {
	var umsResource models.UmsResource
	page, err := umsResource.ListPage(initialize.SqlSession.DbMySQL, 0, "", "", 1, 10)
	if err != nil {
		return
	}
	t.Log("list page success, len = ", len(page))
}
