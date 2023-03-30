package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"kevin/pkg/api"
	"kevin/pkg/pingpong"
	"log"
	"net"
	"time"
)

var (
	proxyAddress   string
	proxyPort      string
	proxyServePort string
)

// metricCmd represents the metric command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Use the proxy method within Kevin",
	Long:  `Use the proxy method within Kevin to communicate with both client and server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("starting Kevin gRPC proxy on port %s...", proxyServePort)
		pingpong.ProxyCallAddress = proxyAddress
		pingpong.ProxyCallPort = proxyPort
		log.Printf("the proxy address: %s\nthe proxy port: %s\n", pingpong.ProxyCallAddress, pingpong.ProxyCallPort)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", proxyServePort))
		if err != nil {
			return err
		}
		grpcServer := grpc.NewServer()
		api.RegisterPongServiceServer(grpcServer, &pingpong.ProxyServer{})
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
	rootCmd.AddCommand(proxyCmd)

	proxyCmd.PersistentFlags().StringVarP(&proxyAddress, "proxy-address", "A", "kevin-server.kevin.svc.cluster.local", "address to dial gRPC services on")
	proxyCmd.PersistentFlags().StringVarP(&proxyPort, "proxy-port", "P", "9000", "port to dial gRPC services on")
	proxyCmd.PersistentFlags().StringVarP(&proxyServePort, "port", "p", "9001", "port to serve proxy on")
}
