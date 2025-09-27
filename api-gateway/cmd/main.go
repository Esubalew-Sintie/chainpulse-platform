package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/chainpulse/backend/api-gateway/grpcclients"
	"github.com/chainpulse/backend/api-gateway/internal/routes"
	"github.com/chainpulse/backend/api-gateway/internal/routes/handlers"
	"github.com/chainpulse/backend/api-gateway/pkg/config"
	"github.com/chainpulse/backend/api-gateway/pkg/logger"
	"github.com/gorilla/mux"
)

const (
	version = "1.0.0"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("failed to load config:", err)
		slog.Error(err.Error())
		os.Exit(1)
	}

	lvl := logger.ParseLogLevel(cfg.LOGGER.LOG_LEVEL)
	logger := logger.NewLogger(string(cfg.ENVIRONMENT), lvl, version)
	if logger == nil {
		slog.Error("failed to initialize logger")
		os.Exit(1)
	}

	accGrpcClient := grpcclients.NewAccountGrpcClient(cfg, *logger)
	if accGrpcClient == nil {
		logger.Error("failed to initialize account gRPC client")
		os.Exit(1)
	}
	proGrpcClient := grpcclients.NewProductGrpcClient(cfg, *logger)
	if proGrpcClient == nil {
		logger.Error("failed to initialize product gRPC client")
		os.Exit(1)
	}
	payGrpcClient := grpcclients.NewPaymentGrpcClient(cfg, *logger)
	if payGrpcClient == nil {
		logger.Error("failed to initialize payment gRPC client")
		os.Exit(1)
	}

	r := mux.NewRouter()

	rts := routes.NewServer(r, cfg, *logger)

	hand := handlers.NewHandlers(rts, accGrpcClient.AccGrpcClient, *proGrpcClient, *payGrpcClient, cfg)
	err = handlers.SetupServer(hand)
	if err != nil {
		logger.Error("failed to setup server", slog.String("error", err.Error()))
		os.Exit(1)
	}
	logger.Info("API Gateway is running", slog.String("version", version), slog.String("environment", string(cfg.ENVIRONMENT)))

}
