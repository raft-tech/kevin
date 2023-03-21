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
	callAddress string
	callPort    string
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	callCmd.PersistentFlags().StringVarP(&callAddress, "address", "a", "kevin-server.kevin.svc.cluster.local", "address to dial gRPC services on")
	callCmd.PersistentFlags().StringVarP(&callPort, "port", "p", "9000", "port to dial gRPC services on")
	pingpong.ProxyCallAddress = callAddress
	pingpong.ProxyCallPort = callPort
	//os.Setenv("PROXY_CALL_ADDRESS", callAddress)
	//os.Setenv("PROXY_CALL_PORT", callPort)
}
