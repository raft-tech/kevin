/*
Copyright © 2023 Raft LLC
*/
package cmd

import (
	"kevin/pkg/api"
	"kevin/pkg/pingpong"
	"time"

	"github.com/spf13/cobra"

	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts serving all Kevin gRPC Services",
	Long: `Starts serving all Kevin gRPC Services

Services currently available:
- pingpong.PongService SayPong
- streamer.StreamService FetchResponse`,

	RunE: func(cmd *cobra.Command, args []string) error {
		port, _ := cmd.Flags().GetString("port")
		log.Printf("starting Kevin gRPC services at localhost:%s", port)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			return err
		}
		grpcServer := grpc.NewServer()
		api.RegisterPongServiceServer(grpcServer, &pingpong.Server{})
		reflection.Register(grpcServer)

		go pingpong.Metrics(metricsPort, metricsEnabled)

		if err := grpcServer.Serve(lis); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().StringP("port", "p", "9000", "port to serve gRPC services on")
}
