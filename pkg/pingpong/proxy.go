package pingpong

import (
	context "context"
	"fmt"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"kevin/pkg/api"
	"log"
	"time"
)

type ProxyServer struct {
	api.UnimplementedPongServiceServer
}

var (
	ProxyCallPort    string
	ProxyCallAddress string
)

func (s *ProxyServer) SayPong(context.Context, *emptypb.Empty) (*api.Pong, error) {
	pongResp, err := CallPingPong(ProxyCallPort, ProxyCallAddress)
	if err != nil {
		return nil, err
	}
	log.Printf("Proxying call to PingPong SayPong gRPC method at %s:%s", ProxyCallAddress, ProxyCallPort)
	return pongResp, nil
}

func (s *ProxyServer) StreamPong(in *api.Ping, srv api.PongService_StreamPongServer) error {
	for i := 0; i < 5; i++ {
		//time sleep to simulate server process time
		time.Sleep(time.Second)
		resp := api.Pong{Pong: in.Ping}
		if err := srv.Send(&resp); err != nil {
			fmt.Printf("send error %v", err)
		}
		log.Printf("Proxying call to PingPong StreamPong gRPC method at %s:%s", ProxyCallAddress, ProxyCallPort)
	}
	return nil
}

func (s *ProxyServer) WritePong(context.Context, *emptypb.Empty) (*api.Pong, error) {
	writerResp, err := CallWritePong(ProxyCallPort, ProxyCallAddress)
	if err != nil {
		return nil, err
	}
	log.Printf("Proxying call to PingPong WritePong gRPC method at %s:%s", ProxyCallAddress, ProxyCallPort)
	return writerResp, nil
}
