package casbin

import (
	"github.com/zeromicro/go-zero/core/logx"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

var e *casbin.Enforcer

func InitCasbin() {

	// 使用MySQL数据库初始化一个gorm适配器
	a, err := gormadapter.NewAdapter("mysql", "root:wzy20040525@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err = casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
}

func GetEnforcer() *casbin.Enforcer {
	return e
}

func CheckPermission(user, action, resource string) bool {
	enforcer := GetEnforcer()

	// 进行权限检验
	result, err := enforcer.Enforce(user, resource, action)
	if err != nil {
		logx.Error(err)
		return false
	}

	return result
}

func AddPermission(user, action, resource string) bool {
	enforcer := GetEnforcer()

	result, err := enforcer.AddPolicy(user, resource, action)
	if err != nil {
		logx.Error(err)
		return false
	}
	return result
}

func RemovePermission(user, action, resource string) bool {
	enforcer := GetEnforcer()

	result, err := enforcer.RemovePolicy(user, resource, action)
	if err != nil {
		logx.Error(err)
		return false
	}
	return result
}
