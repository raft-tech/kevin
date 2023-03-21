/*
Copyright © 2023 Raft LLC
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	streamer2 "kevin/pkg/streamer"
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("calling Kevin gRPC method streamer.StreamService FetchResponse...")
		// dial server
		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", callAddress, callPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}

		// create stream
		client := streamer2.NewStreamServiceClient(conn)
		in := &streamer2.Request{Request: streamerReqBody}
		stream, err := client.FetchResponse(context.Background(), in)
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
				fmt.Println(resp.Result)
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
	callCmd.PersistentFlags().StringVarP(&streamerReqBody, "body", "b", "Hello World", "body of the request to Streamer service")
}
