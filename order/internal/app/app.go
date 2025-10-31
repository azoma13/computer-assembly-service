package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/azoma13/computer-assembly-service/order/config"
	hardwareClient "github.com/azoma13/computer-assembly-service/order/internal/client/grpc/hardware/v1"
	paymentClient "github.com/azoma13/computer-assembly-service/order/internal/client/grpc/payment/v1"
	v1 "github.com/azoma13/computer-assembly-service/order/internal/controller/order/v1"
	"github.com/azoma13/computer-assembly-service/order/internal/models"
	"github.com/azoma13/computer-assembly-service/order/internal/repo"
	"github.com/azoma13/computer-assembly-service/order/internal/service"
	"github.com/azoma13/computer-assembly-service/order/pkg/httpserver"
	"github.com/azoma13/computer-assembly-service/order/pkg/postgres"
	"github.com/azoma13/computer-assembly-service/order/pkg/validator"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
	payment_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/payment/v1"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal("app - Run - config.NewConfig: %w", err)
	}

	log.Println("Initializing postgres...")
	pg, err := postgres.New(cfg.Pg.URL, postgres.MaxPoolSize(cfg.Pg.MaxPoolSize))
	if err != nil {
		log.Fatal("app - Run - postgres.New: %w", err)
	}
	defer pg.Close()

	err = pg.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal("app - Run - AutoMigrate: %w", err)
	}

	log.Println("Initializing hardware client...")
	hardwareConnect, err := grpc.NewClient(cfg.Grpc.Hardware, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("app - Run - grpc.NewClient(Hardware): %w", err)
	}
	hardwareClient := hardwareClient.NewHardwareClient(hardware_v1.NewHardwareServiceClient(hardwareConnect))

	log.Println("Initializing payment client...")
	paymentConnect, err := grpc.NewClient(cfg.Grpc.Payment, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("app - Run - grpc.NewClient(Payment): %w", err)
	}
	paymentClient := paymentClient.NewPaymentClient(payment_v1.NewPaymentServiceClient(paymentConnect))

	log.Println("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	log.Println("Initializing service...")
	deps := service.ServicesDependencies{
		Repos:          repositories,
		HardwareClient: hardwareClient,
		PaymentClient:  paymentClient,
	}
	service := service.NewService(deps)

	log.Println("Initializing routes...")
	handler := echo.New()
	handler.Validator = validator.NewCustomValidator()
	v1.NewRouter(handler, service)

	log.Println("Start http server...")
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	log.Println("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	log.Println("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
