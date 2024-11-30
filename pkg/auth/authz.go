// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package auth

import (
	"time"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter2 "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

const (
	// casbin 访问控制模型.
	aclModel = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)`
)

// Authz定义一个授权器.
type Authz struct {
	*casbin.SyncedEnforcer
}

// NewAuthz定义一个使用casbin授权的授权器.
func NewAuthz(db *gorm.DB) (*Authz, error) {
	// 初始化Gorm adapter
	adapter, err := adapter2.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	m, _ := model.NewModelFromString(aclModel)
	// 初始化enforcer
	enforcer, err := casbin.NewSyncedEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	// 加载policy form DB
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	enforcer.StartAutoLoadPolicy(5 * time.Second)
	a := &Authz{enforcer}
	return a, nil
}

// 使用Authorize进行授权.
func (a *Authz) Authorize(sub, obj, act string) (bool, error) {
	return a.Enforce(sub, obj, act)
}
