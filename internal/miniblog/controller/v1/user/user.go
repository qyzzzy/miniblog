// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package user

import (
	"github.com/qylearn/miniblog/internal/miniblog/biz"
	"github.com/qylearn/miniblog/internal/miniblog/store"
	"github.com/qylearn/miniblog/pkg/auth"
	pb "github.com/qylearn/miniblog/pkg/proto/miniblog/v1"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	b biz.IBiz
	a *auth.Authz
	pb.UnimplementedMiniBlogServer
}

// New创建一个user controller。
func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{a: a, b: biz.NewBiz(ds)}
}
