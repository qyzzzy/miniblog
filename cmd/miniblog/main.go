// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package main

import (
	"os"

	"github.com/qylearn/miniblog/internal/miniblog"
)

func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
