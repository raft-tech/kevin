/*
Copyright © 2023 Raft LLC
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"kevin/pkg/api"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	streamerReqBody string
)

// streamerCmd represents the streamer command
var streamerCmd = &cobra.Command{
	Use:   "streamer",
	Short: "call the PingPong StreamPong gRPC method",
	Long:  `performs a gRPC client call to the pingpong.PongService's StreamPong method exposed by Kevin running in Server mode`,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("calling Kevin gRPC method pingpong.PongService StreamPong...")
		// dial server
		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", callAddress, callPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}

		// create stream
		client := api.NewPongServiceClient(conn)
		in := &api.Ping{Ping: streamerReqBody}
		stream, err := client.StreamPong(context.Background(), in)
		if err != nil {
			return err
		}

		done := make(chan bool)

		go func() {
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					done <- true //means stream is finished
					return
				}
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp.Pong)
			}
		}()

		<-done //we will wait until all response is received
		return nil
	},
}

func init() {
	callCmd.AddCommand(streamerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// streamerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// streamerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	callCmd.PersistentFlags().StringVarP(&streamerReqBody, "body", "b", "Pong", "body of the request to Streamer service")
}
