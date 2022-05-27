package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/proto/proto"
	"log"
)

const PORT = "9001"

func main() {
	//grpc.Dial(":"+PORT, grpc.WithInsecure())
	// 创建与服务端的连接
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	// 创建SearchService的客户端对象
	client := pb.NewSearchServiceClient(conn)

	// 发送rpc请求，等待响应
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	// 输出响应结果
	log.Printf("resp: %s", resp.GetResponse())
}
