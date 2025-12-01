package app

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/azoma13/computer-assembly-service/payment/config"
	v1 "github.com/azoma13/computer-assembly-service/payment/internal/api/payment/v1"
	"github.com/azoma13/computer-assembly-service/payment/internal/service"
	payment_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal("app - Run - config.NewConfig: %w", err)
	}

	log.Println("Start gRPC server...")
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		log.Fatal("failed to listen: %w", err)
	}
	defer func() {
		if servErr := lis.Close(); servErr != nil {
			log.Println("failed close to listen server")
		}
	}()

	log.Println("Initializing gRPC server...")
	s := grpc.NewServer()

	log.Println("Initializing service...")
	deps := service.ServicesDependencies{}
	service := service.NewService(deps)

	log.Println("Initializing api...")
	api := v1.NewPaymentAPI(*service)

	payment_v1.RegisterPaymentServiceServer(s, api)

	reflection.Register(s)

	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Fatal("failed to server: %w", err)
		}
	}()

	log.Println("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
	log.Println("Shutting down...")
	s.GracefulStop()
}
