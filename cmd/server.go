package cmd

import (
	"fmt"
	"log"
	grpc_server "user/app/interface/grpc/user"
	"user/app/interface/persistence/consul"
	"user/app/usecase"
	"user/config"
	"user/registry"

	"github.com/yautze/tools/network"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/spf13/cobra"
	"github.com/yautze/tools/logger"
	"github.com/yautze/tools/srv"
	"google.golang.org/grpc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&config.ConsulAdress, "consul", "c", "127.0.0.1:8500", "consul address")
	serverCmd.Flags().StringVarP(&config.PORT, "port", "p", "0", "service port")
	serverCmd.Flags().BoolVarP(&config.ISUSEIP, "ip", "i", false, "host (defalt: false)")
}

func start() {
	config.Setup()

	l := logger.WithField("service", config.APP)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(l),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_logrus.StreamServerInterceptor(l),
				grpc_recovery.StreamServerInterceptor(),
			),
		),
	)

	// create service
	s := srv.NewService(
		srv.SetServer(server),
		srv.SetName(config.APP),
		srv.SetRegistry(consul.GetClient()),
		srv.SetVersion(config.VERSION),
		srv.SetPort(config.PORT),
		srv.SetHost(func() string {
			if config.ISUSEIP {
				return network.HostIP()
			}
			return "127.0.0.1"
		}()),
	)

	// init hook
	s.Init(
		srv.AfterStart(func() {
			fmt.Printf("Start gRPC Server on Port: %v\n", s.Options().Port)
		}),
	)

	r := registry.New()
	grpc_server.Apply(
		server,
		grpc_server.New(
			r["user-usecase"].(usecase.UserUsecase),
		),
	)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
