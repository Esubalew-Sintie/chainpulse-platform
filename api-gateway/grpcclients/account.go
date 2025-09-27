package grpcclients

import (
	"fmt"
	"log"
	"log/slog"

	"google.golang.org/grpc"

	"github.com/chainpulse/backend/api-gateway/pkg/config"
	accountpb "github.com/chainpulse/backend/api-gateway/proto/gen/account-svc"
)

type AccountGrpcClient struct {
	Config        *config.GlobalConfig
	Logger        slog.Logger
	AccGrpcClient accountpb.AccountServiceClient
}

func NewAccountGrpcClient(config *config.GlobalConfig, logger slog.Logger) *AccountGrpcClient {
	url := fmt.Sprintf("%s:%d", "localhost", 50051)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to account service: %v", err)
	}
	grpcClient := accountpb.NewAccountServiceClient(conn)
	return &AccountGrpcClient{
		Config:        config,
		Logger:        logger,
		AccGrpcClient: grpcClient,
	}

}
