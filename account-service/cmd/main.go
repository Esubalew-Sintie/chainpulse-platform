package main

import (
	"log"
	"log/slog"
	"net"
	"os"

	postgresrepo "github.com/chainpulse/backend/account/internal/infrastructure/postgres_repo"
	"github.com/chainpulse/backend/account/internal/pkgs/config"
	"github.com/chainpulse/backend/account/internal/pkgs/db"
	"github.com/chainpulse/backend/account/internal/pkgs/logger"
	"github.com/chainpulse/backend/account/internal/routes/handlers"
	service "github.com/chainpulse/backend/account/internal/services"
	accountpb "github.com/chainpulse/backend/account/proto"
	"google.golang.org/grpc"
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

	dsn := db.GetDSN(cfg)
	db, cleanup, err := db.OpenDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	if db == nil {
		logger.Error("failed to connect to database")
		os.Exit(1)
	}
	repo := postgresrepo.NewRepo(db) // your actual repo constructor
	svc := service.NewAuthService(repo)
	handler := handlers.NewAccountGrpcHandler(svc)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	accountpb.RegisterAccountServiceServer(server, handler)
	log.Println("✅ gRPC server is running on :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
