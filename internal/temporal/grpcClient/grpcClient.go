package grpcclient

import (
	"context"
	"log"
	"tempotaletl/api/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	conn, err := grpc.Dial("extract:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldnt connect to Extract Service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto.NewExtractClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.Work(ctx, &proto.ExtractRequest{
		Start: data,
	})

	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.DeadlineExceeded {
			log.Printf("call failed: deadline exceeded")
		} else {
			log.Fatalf("call failed: %v", err)
		}

		return nil, err
	}

	return &Response{
		Result: response.GetResult(),
		Data:   response.GetData(),
	}, nil
}

func CallLoadService(data int32) (*Response, error) {
	conn, err := grpc.Dial("load:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	conn, err := grpc.Dial("transform:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
