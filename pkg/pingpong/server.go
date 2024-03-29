package pingpong

import (
	context "context"
	"fmt"
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
	PongCalled.Inc() // increment prom metric
	log.Println("Saying Pong")
	return &api.Pong{Pong: "Pong"}, nil
}

func (s *Server) StreamPong(in *api.Ping, srv api.PongService_StreamPongServer) error {
	PongStreamed.Inc() // increment prom metric
	for i := 0; i < 5; i++ {
		//time sleep to simulate server process time
		time.Sleep(time.Second)
		resp := api.Pong{Pong: in.Ping}
		if err := srv.Send(&resp); err != nil {
			PongStreamedErrors.Inc()
			log.Printf("send error %v", err)
		}
		log.Println("Streaming Pong")
	}
	return nil
}

func (s *Server) WritePong(context.Context, *emptypb.Empty) (*api.Pong, error) {
	filepath := "./data/kevin-%s.txt"
	inputFile := fmt.Sprintf(filepath, "input")
	PongWriter.Inc()
	WriterBytesRead.Set(0.0)
	WriterBytesWritten.Set(0.0)

	_, err := os.Stat(inputFile)
	if err != nil {
		PongWriterErrors.Inc()
		log.Printf("No input file found at %s", inputFile)
		return nil, err
	}

	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		PongWriterErrors.Inc()
		log.Printf("Error opening input file %s: %s", inputFile, err)
		return nil, err
	}
	WriterBytesRead.Set(float64(len(fileBytes)))

	outputFile := fmt.Sprintf(filepath, "output")
	err = os.WriteFile(outputFile, fileBytes, os.FileMode(644))
	if err != nil {
		PongWriterErrors.Inc()
		log.Printf("Error writing to output file %s: %s", outputFile, err)
		return nil, err
	}
	log.Printf("Writing to %s", outputFile)
	WriterBytesWritten.Set(float64(len(fileBytes)))

	return &api.Pong{Pong: string(fileBytes)}, nil
}
