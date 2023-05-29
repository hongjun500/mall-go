package testting

import (
	"github.com/hongjun500/mall-go/internal/conf"
	"github.com/hongjun500/mall-go/internal/model"
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	// 在测试之前，初始化数据库连接
	_, _ = conf.InitMySQLConn()
	code := t.Run()

	// 退出测试
	os.Exit(code)

}

func TestCreate(t *testing.T) {
	t.Log("test create")
	var umsRole model.UmsRole
	umsRole.Name = "二手管理员"
	umsRole.Description = "二手"
	umsRole.AdminCount = 1
	// umsRole.CreateAt = time.Now()
	umsRole.Status = 1
	umsRole.Sort = 0
	create, err := umsRole.Create(conf.Db)
	if err != nil {
		return // 测试失败
	}
	t.Logf("create success, RowsAffected = %v", create)
}

func TestUpdate(t *testing.T) {
	t.Log("test update")
	var umsRole model.UmsRole
	role, err := umsRole.Get(conf.Db, 7)
	if err != nil {
		return
	}
	role.Name = "二手管理员->修改"
	role.Description = "二手->修改"
	role.DeletedAt = nil
	role.Update(conf.Db, 7)

}

func TestDelete(t *testing.T) {
	t.Log("test delete")
	var umsRole model.UmsRole
	i, err := umsRole.Delete(conf.Db, 7)
	if err != nil {
		return // 测试失败
	}
	t.Logf("delete success, RowsAffected = %v", i)
}

func TestListAll(t *testing.T) {

	t.Log("test list all")
	var umsRole model.UmsRole
	all, err := umsRole.ListAll(conf.Db)
	if err != nil {
		return
	}
	t.Logf("all = %v", all)
}

func TestListPage(t *testing.T) {
	t.Log("test list page")
	var umsRole model.UmsRole
	page, err := umsRole.ListPage(conf.Db, "管理员", 1, 10)
	if err != nil {
		return
	}
	t.Logf("page = %v", page)
}
