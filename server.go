package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-demo/proto/proto"
	"log"
	"net"
)

type SearchService struct {
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + "Server"}, nil
}

const PORT = "9001"

func main() {
	// 创建grpc Server对象
	server := grpc.NewServer()

	// 注册服务
	pb.RegisterSearchServiceServer(server, &SearchService{})

	// 创建Listen 监听tcp端口
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}

	server.Serve(lis)
}
