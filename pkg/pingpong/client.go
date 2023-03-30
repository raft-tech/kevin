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
	"time"
)

func CallPingPong(port string, address string) (*api.Pong, error) {
	callstart := time.Now()
	log.Println("calling Kevin gRPC method pingpong.PongService SayPong...")
	PongClientCalled.Inc()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		PongClientErrors.Inc()
		return nil, err
	}
	client := api.NewPongServiceClient(conn)
	pongResp, err := client.SayPong(context.Background(), &emptypb.Empty{})
	if err != nil {
		PongClientErrors.Inc()
		return nil, err
	}
	fmt.Println(pongResp.Pong)
	PongClientLastDurationSeconds.Set(float64(time.Now().Sub(callstart).Milliseconds()))
	return pongResp, nil
}

func CallStreamPong(port string, address string, streamerReqBody string) error {
	callstart := time.Now()
	PongStreamClientCalled.Inc()
	log.Println("calling Kevin gRPC method pingpong.PongService StreamPong...")
	// dial server
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		PongStreamClientErrors.Inc()
		return err
	}

	// create stream
	client := api.NewPongServiceClient(conn)
	in := &api.Ping{Ping: streamerReqBody}
	stream, err := client.StreamPong(context.Background(), in)
	if err != nil {
		PongStreamClientErrors.Inc()
		return err
	}

	// buffer the channel
	done := make(chan bool, 1)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				PongStreamClientErrors.Inc()
				log.Println(err)
			}
			fmt.Println(resp.Pong)
		}
	}()

	<-done //we will wait until all response is received
	PongStreamClientLastDurationSeconds.Set(float64(time.Now().Sub(callstart).Milliseconds()))
	return nil
}

func CallWritePong(port string, address string) (*api.Pong, error) {
	callstart := time.Now()
	WriterClientCalled.Inc()
	fmt.Println("calling Kevin gRPC method pingpong.PongService WritePong...")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		PongWriterClientErrors.Inc()
		return nil, err
	}
	client := api.NewPongServiceClient(conn)
	pongResp, err := client.WritePong(context.Background(), &emptypb.Empty{})
	if err != nil {
		PongWriterClientErrors.Inc()
		return nil, err
	}
	fmt.Println(pongResp.Pong)
	WriterClientLastDurationSeconds.Set(float64(time.Now().Sub(callstart).Milliseconds()))
	return pongResp, nil
}
