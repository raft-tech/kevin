package cmd

import (
	"context"
	"fmt"
	"kevin/pkg/pingpong"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
