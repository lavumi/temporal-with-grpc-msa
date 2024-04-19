package msaserver

import (
	"context"
	"log"
	"net"
	"tempotaletl/api/proto"
	"time"

	"google.golang.org/grpc"
)

type ExtractService struct {
	proto.UnimplementedExtractServer
}

func (s *ExtractService) Work(_ context.Context, req *proto.ExtractRequest) (*proto.ExtractResponse, error) {
	time.Sleep(2 * time.Second)
	return &proto.ExtractResponse{
		Success: true,
		Result:  "extract success",
		Data:    req.Start + 10,
	}, nil
}

func RunExtract() {
	service := &ExtractService{}

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("fail to listen port: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterExtractServer(grpcServer, service)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve extract Server: %v", err)
	}
}
