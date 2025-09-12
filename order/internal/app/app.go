package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/azoma13/computer-assembly-service/order/config"
	orderV1API "github.com/azoma13/computer-assembly-service/order/internal/controller/http/v1"
	"github.com/azoma13/computer-assembly-service/order/internal/repo"
	"github.com/azoma13/computer-assembly-service/order/internal/service"
	"github.com/azoma13/computer-assembly-service/order/pkg/httpserver"
	"github.com/azoma13/computer-assembly-service/order/pkg/postgres"
	"github.com/azoma13/computer-assembly-service/order/pkg/validator"
	"github.com/labstack/echo/v4"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("app - Run - config.NewConfig: %v", err)
		return
	}

	log.Println("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatalf("app - Run - postgres.New: %v", err)
		return
	}
	defer pg.Close()

	log.Println("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	log.Println("Initializing services...")
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	service := service.NewServices(deps)

	log.Println("Initializing handlers and routes...")
	handler := echo.New()
	handler.Validator = validator.NewCustomValidator()

	orderV1API.NewRouter(handler, service)

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
