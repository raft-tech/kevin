package pingpong

import (
	context "context"
	"fmt"
	"kevin/pkg/api"
	"time"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	api.UnimplementedPongServiceServer
}

func (s *Server) SayPong(context.Context, *emptypb.Empty) (*api.Pong, error) {
	return &api.Pong{Pong: "Pong"}, nil
}

func (s *Server) StreamPong(in *api.Ping, srv api.PongService_StreamPongServer) error {
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
