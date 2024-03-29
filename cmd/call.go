/*
Copyright © 2023 Raft LLC
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kevin/pkg/pingpong"
)

var (
	callAddress  string
	callPort     string
	repeats      int
	delaySeconds int
)

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Use Kevin as a client to call Kevin gRPC services",
	Long:  `A subcommand which prefaces all 'client mode' interactions with Kevin gRPC services`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Specify a subcommand in order to perform a client call to the desired Kevin gRPC service")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(callCmd)

	callCmd.PersistentFlags().StringVarP(&callAddress, "address", "a", "kevin-server.kevin.svc.cluster.local", "address to dial gRPC services on")
	callCmd.PersistentFlags().StringVarP(&callPort, "port", "p", "9000", "port to dial gRPC services on")
	pingpong.ProxyCallAddress = callAddress
	pingpong.ProxyCallPort = callPort

	callCmd.PersistentFlags().IntVarP(&repeats, "repeats", "r", 1, "number of times to perform a given call operation")
	callCmd.PersistentFlags().IntVarP(&delaySeconds, "delay", "d", 0, "seconds of delay before next client call")
}
