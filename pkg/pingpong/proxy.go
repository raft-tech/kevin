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
	PongProxy.Inc()
	log.Printf("Proxying call to PingPong SayPong gRPC method at %s:%s", ProxyCallAddress, ProxyCallPort)
	pongResp, err := CallPingPong(ProxyCallPort, ProxyCallAddress)
	if err != nil {
		PongProxyErrors.Inc()
		return nil, err
	}
	return pongResp, nil
}

func (s *ProxyServer) StreamPong(in *api.Ping, srv api.PongService_StreamPongServer) error {
	StreamProxy.Inc()
	for i := 0; i < 5; i++ {
		//time sleep to simulate server process time
		time.Sleep(time.Second)
		log.Printf("Proxying call to PingPong StreamPong gRPC method at %s:%s", ProxyCallAddress, ProxyCallPort)
		resp := api.Pong{Pong: in.Ping}
		if err := srv.Send(&resp); err != nil {
			StreamProxyErrors.Inc()
			fmt.Printf("send error %v", err)
		}
	}
	return nil
}

func (s *ProxyServer) WritePong(context.Context, *emptypb.Empty) (*api.Pong, error) {
	WriterProxy.Inc()
	log.Printf("Proxying call to PingPong WritePong gRPC method at %s:%s", ProxyCallAddress, ProxyCallPort)
	writerResp, err := CallWritePong(ProxyCallPort, ProxyCallAddress)
	if err != nil {
		WriterProxyErrors.Inc()
		return nil, err
	}
	return writerResp, nil
}
