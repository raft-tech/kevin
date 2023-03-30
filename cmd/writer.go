/*
Copyright © 2023 Raft LLC
*/
package cmd

import (
	"github.com/spf13/cobra"
	"kevin/internal"
	"kevin/pkg/pingpong"
)

// writerCmd represents the pingpong command
var writerCmd = &cobra.Command{
	Use:   "writer",
	Short: "call the PingPong WritePong gRPC method",
	Long: `performs a gRPC client call to the pingpong.PongService's WritePong method exposed by Kevin running in Server mode
ensure that there is a ./data/kevin-input.txt file for the server to read from -- or you will experience issues`,
	RunE: func(cmd *cobra.Command, args []string) error {
		go internal.Metrics(metricsPort)
		_, err := pingpong.CallWritePong(callPort, callAddress)
		return err
	},
}

func init() {
	callCmd.AddCommand(writerCmd)
}
