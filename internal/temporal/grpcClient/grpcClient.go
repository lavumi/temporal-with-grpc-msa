package grpcclient

import (
	"context"
	"log"
	"tempotaletl/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Response struct {
	Result string
	Data   int32
}

// func CallCrawlService(data int32) (*Response, error) {
// 	conn, err := grpc.Dial("localhost:3031", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("Couldnt connect to crawl Service: %v", err)
// 		return nil, err
// 	}
// 	defer conn.Close()

// 	client := proto.NewCrawlClient(conn)
// 	response, err := client.Work(context.Background(), &proto.CrawlRequest{
// 		Start: data,
// 	})
// 	return &Response{
// 		Result: response.GetResult(),
// 		Data:   response.GetData(),
// 	}, err
// }

func CallExtractService(data int32) (*Response, error) {
	conn, err := grpc.Dial("localhost:3032", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldnt connect to Extract Service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto.NewExtractClient(conn)
	response, err := client.Work(context.Background(), &proto.ExtractRequest{
		Start: data,
	})
	return &Response{
		Result: response.GetResult(),
		Data:   response.GetData(),
	}, err
}

func CallLoadService(data int32) (*Response, error) {
	conn, err := grpc.Dial("localhost:3034", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldnt connect to Load Service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto.NewLoadClient(conn)
	response, err := client.Work(context.Background(), &proto.LoadRequest{
		Start: data,
	})
	return &Response{
		Result: response.GetResult(),
		Data:   response.GetData(),
	}, err
}

func CallTransformService(data int32) (*Response, error) {
	conn, err := grpc.Dial("localhost:3033", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldnt connect to Transform Service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto.NewTransformClient(conn)
	response, err := client.Work(context.Background(), &proto.TransformRequest{
		Start: data,
	})
	return &Response{
		Result: response.GetResult(),
		Data:   response.GetData(),
	}, err
}
