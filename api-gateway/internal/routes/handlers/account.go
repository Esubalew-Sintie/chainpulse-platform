package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/chainpulse/backend/api-gateway/internal/routes/handlers/dto"
	accountpb "github.com/chainpulse/backend/api-gateway/proto/gen/account-svc"
	"github.com/gorilla/mux"
)

func (s *Handlers) CreateBuyerHandler(w http.ResponseWriter, r *http.Request) {
	var req accountpb.CreateBuyerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.accGrpcClient.CreateBuyer(ctx, &req)
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	dto.WriteSuccessResponse(w, nil, "Buyer created successfully")
}

func (s *Handlers) LoginBuyerHandler(w http.ResponseWriter, r *http.Request) {
	var req accountpb.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.accGrpcClient.LoginBuyer(ctx, &req)
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	dto.WriteSuccessResponse(w, res, "Login successful")
}

func (s *Handlers) GetBuyerByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.accGrpcClient.GetBuyerByID(ctx, &accountpb.GetByIDRequest{Id: id})
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	dto.WriteSuccessResponse(w, res, "Buyer retrieved successfully")
}

func (s *Handlers) UpdateBuyerHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req accountpb.UpdateBuyerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	req.BuyerId = id

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.accGrpcClient.UpdateBuyer(ctx, &req)
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	dto.WriteSuccessResponse(w, nil, "Buyer updated successfully")
}

func (s *Handlers) DeleteBuyerHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.accGrpcClient.DeleteBuyer(ctx, &accountpb.GetByIDRequest{Id: id})
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	dto.WriteSuccessResponse(w, nil, "Buyer deleted successfully")
}

func (s *Handlers) GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.accGrpcClient.GetSettings(ctx, &emptypb.Empty{})
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	dto.WriteSuccessResponse(w, res, "Settings retrieved")
}

func (s *Handlers) UpdateSettingsHandler(w http.ResponseWriter, r *http.Request) {
	var req accountpb.SettingsRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.accGrpcClient.UpdateSettings(ctx, &req)
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	dto.WriteSuccessResponse(w, nil, "Settings updated")
}

func (s *Handlers) InsertSettingsHandler(w http.ResponseWriter, r *http.Request) {
	var req accountpb.SettingsRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.accGrpcClient.InsertSettings(ctx, &req)
	if err != nil {
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	dto.WriteSuccessResponse(w, nil, "Settings inserted")
}
