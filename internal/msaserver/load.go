package msaserver

import (
	"context"
	"log"
	"net"
	"tempotaletl/api/proto"
	"time"

	"google.golang.org/grpc"
)

type LoadService struct {
	proto.UnimplementedLoadServer
}

func (s *LoadService) Work(_ context.Context, req *proto.LoadRequest) (*proto.LoadResponse, error) {
	time.Sleep(2 * time.Second)
	return &proto.LoadResponse{
		Success: true,
		Result:  "load success",
		Data:    req.Start + 4,
	}, nil
}

func RunLoad() {
	service := &LoadService{}

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("fail to listen port: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterLoadServer(grpcServer, service)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve Load Server: %v", err)
	}
}
