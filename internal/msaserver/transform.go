package msaserver

import (
	"context"
	"log"
	"net"
	"tempotaletl/api/proto"
	"time"

	"google.golang.org/grpc"
)

type TransformService struct {
	proto.UnimplementedTransformServer
}

func (s *TransformService) Work(_ context.Context, req *proto.TransformRequest) (*proto.TransformResponse, error) {
	time.Sleep(2 * time.Second)
	return &proto.TransformResponse{
		Success: true,
		Result:  "transform success",
		Data:    req.Start + 100,
	}, nil
}

func RunTransform() {
	service := &TransformService{}

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("fail to listen port: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterTransformServer(grpcServer, service)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve Transform Server: %v", err)
	}
}
