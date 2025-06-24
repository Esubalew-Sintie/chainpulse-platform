package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"log"

	"github.com/chainpulse/backend/api-gateway/dto"
	accountpb "github.com/chainpulse/backend/api-gateway/proto/gen"
	"google.golang.org/grpc"
)

var grpcClient accountpb.AccountServiceClient

func init() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // TODO: secure for production
	if err != nil {
		log.Fatalf("Could not connect to account service: %v", err)
	}
	grpcClient = accountpb.NewAccountServiceClient(conn)
}

func CreateBuyerHandler(w http.ResponseWriter, r *http.Request) {
	var req accountpb.CreateBuyerRequest

	// Parse JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON input: "+err.Error())
		return
	}

	// Context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call gRPC
	_, err := grpcClient.CreateBuyer(ctx, &req)
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	// Return structured success response
	dto.WriteSuccessResponse(w, nil, "Buyer created successfully")
}
