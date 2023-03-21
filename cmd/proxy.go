package cmd

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"kevin/pkg/pingpong"
	"net/http"
)

func pingPongHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", callAddress, callPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	client := pingpong.NewPongServiceClient(conn)
	pongResp, err := client.SayPong(context.Background(), &emptypb.Empty{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(pongResp.Pong)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pongResp.Pong))
}

// metricCmd represents the metric command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Use the proxy method within Kevin",
	Long:  `Use the proxy method within Kevin to communicate with both client and server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		r := mux.NewRouter()
		// Register the proxy server on port 8000.

		// Create the server endpoint. Its endpoint is '/server'.
		r.HandleFunc("/proxy/pingpong", pingPongHandler)
		http.Handle("/pingpong", r)

		r.HandleFunc("/proxy/streamer", pingPongHandler)
		http.Handle("/client", r)

		// Register the proxy server on port 8000.
		return http.ListenAndServe("localhost:8000", r)
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)

	proxyCmd.PersistentFlags().StringVarP(&callAddress, "address", "a", "kevin-server.kevin.svc.cluster.local", "address to dial gRPC services on")
	proxyCmd.PersistentFlags().StringVarP(&callPort, "port", "p", "9000", "port to dial gRPC services on")
}
