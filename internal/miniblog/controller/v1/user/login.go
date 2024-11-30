// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/qylearn/miniblog/internal/pkg/core"
	"github.com/qylearn/miniblog/internal/pkg/errno"
	"github.com/qylearn/miniblog/internal/pkg/log"
	v1 "github.com/qylearn/miniblog/pkg/api/miniblog/v1"
)

// 登录miniblog并返回一个Token.
func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")
	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)
}
