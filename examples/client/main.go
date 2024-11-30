// Copyright 2023 qyzzzy(秦云) <qinyun_77@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/qylearn/miniblog.

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/qylearn/miniblog/internal/pkg/log"
	pb "github.com/qylearn/miniblog/pkg/proto/miniblog/v1"
)

var (
	addr  = flag.String("addr", "localhost:9090", "The adderss to connect to ")
	limit = flag.Int64("limit", 10, "limit to list user.")
)

func main() {
	flag.Parse()
	//建立服务器的连接
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalw("Failed to connect", "err", err)
	}
	defer conn.Close()
	//建立一个客户端
	c := pb.NewMiniBlogClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//请求Lister接口
	r, err := c.ListUser(ctx, &pb.ListUserRequest{Offset: 0, Limit: *limit})
	if err != nil {
		log.Fatalw("could not greet: %v", err)
	}
	// 打印请求结果
	fmt.Println("TotalCount:", r.TotalCount)
	for _, u := range r.Users {
		d, _ := json.Marshal(u)
		fmt.Println(string(d))
	}
}
