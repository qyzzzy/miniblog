// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQLOptions定义了MySql选项.
type MySQLOptions struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

// DSN 从MySQLOptions返回DSN.
func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local")
}

// NewMySQL使用给定的选项创建一个新的gorm数据库实例。
func NewMySQL(opts *MySQLOptions) (*gorm.DB, error) {
	loglevel := logger.Silent
	if opts.LogLevel != 0 {
		loglevel = logger.LogLevel(opts.LogLevel)
	}
	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxOpenConns设置到数据库的最大打开连接数
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	// SetConnMaxLifetime 设置连接可重用的最长时间
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns 设置空闲连接池的最大连接数
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
