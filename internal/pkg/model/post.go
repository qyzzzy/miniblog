// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package model

import "time"

type PostM struct {
	ID        int64     `gorm:"column:id;primary_key"` //
	Username  string    `gorm:"column:username"`       //
	PostID    string    `gorm:"column:postID"`         //
	Title     string    `gorm:"column:title"`          //
	Content   string    `gorm:"column:content"`        //
	CreatedAt time.Time `gorm:"column:createdAt"`      //
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //
}

// TableName 用来指定映射的 MySQL 表名.
func (p *PostM) TableName() string {
	return "post"
}
