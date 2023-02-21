/*
Copyright © 2023 Raft LLC

*/
package cmd

import (
    "context"
	"fmt"
    "kevin/pingpong"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// pingpongCmd represents the pingpong command
var pingpongCmd = &cobra.Command{
	Use:   "pingpong",
	Short: "call the PingPong SayPong gRPC method",
	Long: `performs a gRPC client call to the pingpong.PongService's SayPong method exposed by Kevin running in Server mode`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("calling Kevin gRPC method pingpong.PongService SayPong...")
    	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", callAddress, callPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
        client := pingpong.NewPongServiceClient(conn)
        pongResp, err := client.SayPong(context.Background(), &emptypb.Empty{})
        if err != nil {
            return err
        }
        fmt.Println(pongResp.Pong)
		return nil
	},
}

func init() {
	callCmd.AddCommand(pingpongCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingpongCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingpongCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
