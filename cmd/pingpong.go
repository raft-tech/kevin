/*
Copyright © 2023 Raft LLC
*/
package cmd

import (
	"github.com/spf13/cobra"
	"kevin/pkg/pingpong"
	"log"
	"time"
)

// pingpongCmd represents the pingpong command
var pingpongCmd = &cobra.Command{
	Use:   "pingpong",
	Short: "call the PingPong SayPong gRPC method",
	Long:  `performs a gRPC client call to the pingpong.PongService's SayPong method exposed by Kevin running in Server mode`,
	RunE: func(cmd *cobra.Command, args []string) error {
		go pingpong.Metrics(metricsPort, metricsEnabled)
		for i := 0; i < repeats; i++ {
			go func() {
				_, err := pingpong.CallPingPong(callPort, callAddress)
				if err != nil {
					log.Println(err)
				}
			}()
			time.Sleep(time.Duration(delaySeconds) * time.Second)
		}
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
