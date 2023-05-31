package testting

import (
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/models"
	"testing"
)

func TestResourceListAll(t *testing.T) {
	var umsResource models.UmsResource
	all, err := umsResource.ListAll(conf.Db)
	if err != nil {
		return
	}
	t.Log("list all success, len = ", len(all))
}

func TestResourceListPage(t *testing.T) {
	var umsResource models.UmsResource
	page, err := umsResource.ListPage(conf.Db, 0, "", "", 1, 10)
	if err != nil {
		return
	}
	t.Log("list page success, len = ", len(page))
}
