package main

import (
	"context"
	"fmt"
	pb "proto"
	"google.golang.org/grpc"
)

func main() {
	// 1. 创建与gRPC服务端的连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}
	defer conn.Close()
	// 2. 实例化gRPC客户端
	client := pb.NewUserInfoServiceClient(conn)
	// 3. 组装参数
	req := new(pb.UserRequest)
	req.Name = "zhangsan"
	// 4. 调用接口
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常：%s\n", err)
	}
	fmt.Printf("响应结果: %v\n", resp)
}
