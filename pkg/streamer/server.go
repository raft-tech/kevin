package streamer

import (
	"fmt"
	"time"
)

type Server struct {
	UnimplementedStreamServiceServer
}

func (s *Server) FetchResponse(in *Request, srv StreamService_FetchResponseServer) error {

	//use wait group to allow process to be concurrent
	for i := 0; i < 5; i++ {
		//time sleep to simulate server process time
		time.Sleep(time.Second)
		resp := Response{Result: in.Request}
		if err := srv.Send(&resp); err != nil {
			fmt.Printf("send error %v", err)
		}
	}
	return nil
}
