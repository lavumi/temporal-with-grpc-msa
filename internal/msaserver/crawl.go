package msaserver

import (
	"context"
	"log"
	"net"
	"tempotaletl/api/proto"

	"google.golang.org/grpc"
)

type CrawlService struct {
	proto.UnimplementedCrawlServer
}

func (s *CrawlService) Work(_ context.Context, req *proto.CrawlRequest) (*proto.CrawlResponse, error) {

	return &proto.CrawlResponse{
		Success: true,
		Result:  "crawling success",
	}, nil
}

func RunCrawl() {
	service := &CrawlService{}

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("fail to listen port: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterCrawlServer(grpcServer, service)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve crawl Server: %v", err)
	}
}
