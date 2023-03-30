/*
Copyright © 2023 Raft LLC
*/
package cmd

import (
	"github.com/spf13/cobra"
	"kevin/internal"
	"kevin/pkg/pingpong"
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
		go internal.Metrics(metricsPort)
		return pingpong.CallStreamPong(callPort, callAddress, streamerReqBody)
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
