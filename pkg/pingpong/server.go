package pingpong

import (
	context "context"
	"fmt"
	"kevin/internal"
	"kevin/pkg/api"
	"log"
	"os"
	"time"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	api.UnimplementedPongServiceServer
}

func (s *Server) SayPong(context.Context, *emptypb.Empty) (*api.Pong, error) {
	internal.PongCalled.Inc() // increment prom metric
	log.Println("Saying Pong")
	return &api.Pong{Pong: "Pong"}, nil
}

func (s *Server) StreamPong(in *api.Ping, srv api.PongService_StreamPongServer) error {
	for i := 0; i < 5; i++ {
		//time sleep to simulate server process time
		time.Sleep(time.Second)
		resp := api.Pong{Pong: in.Ping}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Println("Streaming Pong")
	}
	internal.PongStreamed.Inc() // increment prom metric
	return nil
}

func (s *Server) WritePong(context.Context, *emptypb.Empty) (*api.Pong, error) {
	filepath := "./data/kevin-%s.txt"
	inputFile := fmt.Sprintf(filepath, "input")
	internal.PongWriter.Inc()
	internal.WriterBytesRead.Set(0.0)
	internal.WriterBytesWritten.Set(0.0)

	_, err := os.Stat(inputFile)
	if err != nil {
		log.Printf("No input file found at %s", inputFile)
		return nil, err
	}

	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Printf("Error opening input file %s: %s", inputFile, err)
		return nil, err
	}
	internal.WriterBytesRead.Set(float64(len(fileBytes)))

	outputFile := fmt.Sprintf(filepath, "output")
	err = os.WriteFile(outputFile, fileBytes, os.FileMode(644))
	if err != nil {
		log.Printf("Error writing to output file %s: %s", outputFile, err)
		return nil, err
	}
	log.Printf("Writing to %s", outputFile)
	internal.WriterBytesWritten.Set(float64(len(fileBytes)))

	return &api.Pong{Pong: string(fileBytes)}, nil
}
