package grpcclients

import (
	"fmt"
	"log"
	"log/slog"

	"google.golang.org/grpc"

	"github.com/chainpulse/backend/api-gateway/pkg/config"
	accountpb "github.com/chainpulse/backend/api-gateway/proto/gen/account-svc"
)

type ProductGrpcClient struct {
	Config        *config.GlobalConfig
	Logger        slog.Logger
	ProGrpcClient accountpb.AccountServiceClient
}

func NewProductGrpcClient(config *config.GlobalConfig, logger slog.Logger) *ProductGrpcClient {
	url := fmt.Sprintf("%s:%d", "localhost", 50051)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to account service: %v", err)
	}
	grpcClient := accountpb.NewAccountServiceClient(conn)
	return &ProductGrpcClient{
		Config:        config,
		Logger:        logger,
		ProGrpcClient: grpcClient,
	}

}
