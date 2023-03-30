/*
Copyright © 2023 Raft LLC
*/
package cmd

import (
	"github.com/spf13/cobra"
	"kevin/pkg/pingpong"
	"log"
)

// watchCmd represents the call command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Use Kevin as a client to endlessly watch a Kevin server",
	Long:  `Watches a Kevin server and performs client calls ENDLESSLY to ensure that gRPC services are functional`,
	RunE: func(cmd *cobra.Command, args []string) error {
		go pingpong.Metrics(metricsPort, metricsEnabled)
		log.Println("starting Kevin watch mode...")
		pingpong.WatchKevin(callPort, callAddress, streamerReqBody, delaySeconds)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.PersistentFlags().StringVarP(&callAddress, "address", "a", "kevin-server.kevin.svc.cluster.local", "address to dial gRPC services on")
	watchCmd.PersistentFlags().StringVarP(&callPort, "port", "p", "9000", "port to dial gRPC services on")
	pingpong.ProxyCallAddress = callAddress
	pingpong.ProxyCallPort = callPort

	watchCmd.PersistentFlags().StringVarP(&streamerReqBody, "body", "b", "Pong", "body of the request to Streamer service")

	watchCmd.PersistentFlags().IntVarP(&delaySeconds, "delay", "d", 10, "seconds of delay before next client call")
}
