// @author hongjun500
// @date 2023/6/15 11:25
// @tool ThinkPadX1隐士
// Created with GoLand 2022.2
// Description: 自定义动态权限数据源

package security

import (
	"regexp"
	"strconv"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"

	"github.com/hongjun500/mall-go/internal/database"
	"github.com/hongjun500/mall-go/internal/models"
)

var configAttributeMap = make(map[string]string, 0)
var dbFactory *database.DbFactory

func SetDbFactory(factory *database.DbFactory) {
	dbFactory = factory
}

func NewEnforcerSecurity() *casbin.Enforcer {
	// 使用Gorm适配器创建Casbin的Enforcer
	adapter, _ := gormadapter.NewAdapterByDB(dbFactory.GormMySQL)

	e, _ := casbin.NewEnforcer("path/to/model.conf", adapter)
	// 将权限添加到 Casbin 模型
	// e.AddPolicy(role.Name, permission.Name, permission.Action)

	return e
}

func loadDataSource() map[string]string {
	// 查询数据库，获取所有的权限配置
	umsResource := &models.UmsResource{}
	umsResources, _ := umsResource.SelectAll(dbFactory.GormMySQL)
	for _, resource := range umsResources {
		configAttributeMap[resource.Url] = strconv.Itoa(int(resource.Id)) + ":" + resource.Name
	}
	return configAttributeMap
}

func clearDataSource() {
	configAttributeMap = make(map[string]string, 0)
}

func GetConfigAttributes(context *gin.Context) []string {
	if len(configAttributeMap) == 0 {
		configAttributeMap = loadDataSource()
	}
	// 获取当前访问的路径
	url := context.Request.URL
	path := url.Path
	configAttributes := make([]string, 0)
	for key, _ := range configAttributeMap {
		if matchPath(key, path) {
			configAttributes = append(configAttributes, configAttributeMap[key])
		}
	}
	return configAttributes
}

// 使用正则表达式匹配路径
func matchPath(pattern, path string) bool {
	regexpPattern := "^" + pattern + "$"
	match, err := regexp.MatchString(regexpPattern, path)
	if err != nil {
		return false
	}
	return match
}

// 定义Casbin模型的字符串
const casbinModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && (r.act == p.act || p.act == "*")
`
