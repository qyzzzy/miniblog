// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")

		if requestID == "" {
			requestID = uuid.New().String()
		}
		// 将Request-ID保存在gin.Context中，方便后续使用
		c.Set("X-Request-ID", requestID)

		// 将RequestID 保存在HTTP返回头中，header的键位“X-Request-ID”
		c.Writer.Header().Set("X-Request-ID", requestID)
	}
}
