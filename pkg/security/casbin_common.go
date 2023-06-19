// @author hongjun500
// @date 2023/6/16 16:49
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package security

import (
	"sync"

	"github.com/casbin/casbin"
	"github.com/hongjun500/mall-go/internal/models"
)

var (
	rbacModel = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch(r.obj, p.obj) || r.sub == "admin"
	`
	once     sync.Once
	Enforcer *casbin.Enforcer
)

func init() {
	once.Do(func() {
		modelFromString := casbin.NewModel(rbacModel)
		Enforcer = casbin.NewEnforcer(modelFromString)
	})
}

// AddPolicyFromResource 将基于每个不同的用户资源添加到 casbin 中
func AddPolicyFromResource(enforcer *casbin.Enforcer, sub string, resources []models.UmsResource) {
	for _, umsResource := range resources {
		enforcer.AddPolicy(sub, umsResource.Url, "*")
	}
}

// m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act) || r = "admin"
// 这里不用搞那么严格，直接路径匹配即可，不用考虑请求方法
