package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/azoma13/computer-assembly-service/order/config"
	v1 "github.com/azoma13/computer-assembly-service/order/internal/api/order/v1"
	hardwareClient "github.com/azoma13/computer-assembly-service/order/internal/client/grpc/hardware/v1"
	paymentClient "github.com/azoma13/computer-assembly-service/order/internal/client/grpc/payment/v1"
	"github.com/azoma13/computer-assembly-service/order/internal/repo"
	repoModels "github.com/azoma13/computer-assembly-service/order/internal/repo/models"
	"github.com/azoma13/computer-assembly-service/order/internal/service"
	"github.com/azoma13/computer-assembly-service/shared/pkg/httpserver"
	order_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/openapi/order/v1"
	"github.com/azoma13/computer-assembly-service/shared/pkg/postgres"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
	payment_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/payment/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	err = pg.AutoMigrate(&repoModels.Order{})
	if err != nil {
		log.Fatal("app - Run - AutoMigrate: %w", err)
	}

	log.Println("Initializing hardware client...")
	hardwareConnect, err := grpc.NewClient(cfg.Grpc.Hardware, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("app - Run - grpc.NewClient(Hardware): %w", err)
	}
	hardwareClient := hardwareClient.NewHardwareClient(hardware_v1.NewHardwareServiceClient(hardwareConnect))
	defer func() {
		if err := hardwareConnect.Close(); err != nil {
			log.Printf("failed to close hardware connection: %v\n", err)
		}
	}()

	log.Println("Initializing payment client...")
	paymentConnect, err := grpc.NewClient(cfg.Grpc.Payment, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("app - Run - grpc.NewClient(Payment): %w", err)
	}
	paymentClient := paymentClient.NewPaymentClient(payment_v1.NewPaymentServiceClient(paymentConnect))
	defer func() {
		if err := paymentConnect.Close(); err != nil {
			log.Printf("failed to close payment connection: %v\n", err)
		}
	}()

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
	api := v1.NewOrderAPI(service.Order)
	orderServer, err := order_v1.NewServer(api)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Mount("/api/v1", orderServer)

	log.Println("Start http server...")
	httpServer := httpserver.New(r, httpserver.Port(cfg.HTTP.Port))

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
