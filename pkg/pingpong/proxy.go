package pingpong

import (
	context "context"
	"fmt"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"kevin/pkg/api"
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
	return pongResp, nil
}

// TODO needs some work -- nonfunctional
func (s *ProxyServer) StreamPong(in *api.Ping, srv api.PongService_StreamPongServer) error {
	for i := 0; i < 5; i++ {
		//time sleep to simulate server process time
		time.Sleep(time.Second)
		resp := api.Pong{Pong: in.Ping}
		if err := srv.Send(&resp); err != nil {
			fmt.Printf("send error %v", err)
		}
	}
	return nil
}

func (s *ProxyServer) WritePong(context.Context, *emptypb.Empty) (*api.Pong, error) {
	writerResp, err := CallWritePong(ProxyCallPort, ProxyCallAddress)
	if err != nil {
		return nil, err
	}
	return writerResp, nil
}
