// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package miniblog

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/qylearn/miniblog/internal/miniblog/controller/v1/user"
	"github.com/qylearn/miniblog/internal/miniblog/store"
	"github.com/qylearn/miniblog/internal/pkg/core"
	"github.com/qylearn/miniblog/internal/pkg/errno"
	"github.com/qylearn/miniblog/internal/pkg/log"
	mw "github.com/qylearn/miniblog/internal/pkg/middleware"
	"github.com/qylearn/miniblog/pkg/auth"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	// 注册 pprof 路由
	pprof.Register(g)

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)

	g.POST("/login", uc.Login)
	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)                             // 创建用户
			userv1.PUT(":name/change-password", uc.ChangePassword) // 修改用户密码
			userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.GET(":name", uc.Get)       // 获取用户信息
			userv1.PUT(":name", uc.Update)    // 更新用户
			userv1.GET("", uc.List)           // 列出用户列表，只有 root 用户才能访问
			userv1.DELETE(":name", uc.Delete) // 删除用户
		}
	}
	return nil
}
