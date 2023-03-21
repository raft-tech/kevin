package pingpong

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"kevin/pkg/api"
	"log"
)

func CallPingPong(port string, address string) (*api.Pong, error) {
	fmt.Println("calling Kevin gRPC method pingpong.PongService SayPong...")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := api.NewPongServiceClient(conn)
	pongResp, err := client.SayPong(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	fmt.Println(pongResp.Pong)
	return pongResp, nil
}

func CallStreamPong(port string, address string, streamerReqBody string) error {
	log.Println("calling Kevin gRPC method pingpong.PongService StreamPong...")
	// dial server
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	// create stream
	client := api.NewPongServiceClient(conn)
	in := &api.Ping{Ping: streamerReqBody}
	stream, err := client.StreamPong(context.Background(), in)
	if err != nil {
		return err
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(resp.Pong)
		}
	}()

	<-done //we will wait until all response is received
	return nil
}
