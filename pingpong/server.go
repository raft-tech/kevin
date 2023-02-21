package pingpong

import (
	context "context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	UnimplementedPongServiceServer
}

func (S *Server) SayHello(context.Context, *emptypb.Empty) (*Pong, error) {
	return &Pong{Pong: "Pong"}, nil

}

/// func (S *Server) SayHello(context.Context, *Pong) (*Pong, error) {
//	return &Pong{Pong: "Pong"}, nil
