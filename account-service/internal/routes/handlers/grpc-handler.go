package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/chainpulse/backend/account/internal/models"
	service "github.com/chainpulse/backend/account/internal/services"
	accountpb "github.com/chainpulse/backend/account/proto/gen"
	"github.com/gogo/status"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AccountGrpcHandler struct {
	accountpb.UnimplementedAccountServiceServer
	authSvc service.IAuthSvc
}

func NewAccountGrpcHandler(authSvc service.IAuthSvc) *AccountGrpcHandler {
	return &AccountGrpcHandler{authSvc: authSvc}
}

func (h *AccountGrpcHandler) CreateBuyer(ctx context.Context, req *accountpb.CreateBuyerRequest) (*emptypb.Empty, error) {
	// WALLET LOGIN ONLY
	if req.WalletAddress != "" && req.Password == "" && req.PhoneNumber == "" {
		input := &models.Buyer{
			WalletAddress: &req.WalletAddress,
		}
		if err := h.authSvc.CreateWalletOnlyBuyer(ctx, input); err != nil {
			// log the error
			// you can use your preferred logger, here using standard log for example
			// log.Printf("CreateWalletOnlyBuyer error: %v", err)
			// or if you have a logger in the handler, use that
			// h.logger.Errorf("CreateWalletOnlyBuyer error: %v", err)
			println("CreateWalletOnlyBuyer error:", err.Error())
			return nil, err
		}
		return &emptypb.Empty{}, nil
	}

	// PHONE & PASSWORD LOGIN FLOW
	if req.PhoneNumber == "" {
		err := status.Errorf(codes.InvalidArgument, "phone number is required")
		println("CreateBuyer error:", err.Error())
		return nil, err
	}
	if req.Password == "" {
		err := status.Errorf(codes.InvalidArgument, "password is required")
		println("CreateBuyer error:", err.Error())
		return nil, err
	}
	if req.Email == "" {
		err := status.Errorf(codes.InvalidArgument, "email is required")
		println("CreateBuyer error:", err.Error())
		return nil, err
	}

	input := &models.Buyer{
		Email:             &req.Email,
		Password:          req.Password,
		PhoneNumber:       &req.PhoneNumber,
		WalletAddress:     &req.WalletAddress,
		FirstName:         &req.FirstName,
		LastName:          &req.LastName,
		Address:           &req.Address,
		ProfilePictureURL: &req.ProfilePictureUrl,
	}

	if err := h.authSvc.CreateBuyer(ctx, input); err != nil {
		println("CreateBuyer error:", err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *AccountGrpcHandler) LoginBuyer(ctx context.Context, req *accountpb.LoginRequest) (*accountpb.LoginResponse, error) {
	buyer, err := h.authSvc.LoginBuyer(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &accountpb.LoginResponse{
		BuyerId:     buyer.ID.String(),
		Email:       *buyer.Email,
		PhoneNumber: *buyer.PhoneNumber,
	}, nil
}

func (h *AccountGrpcHandler) GetBuyerByID(ctx context.Context, req *accountpb.GetByIDRequest) (*accountpb.BuyerResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	buyer, err := h.authSvc.GetBuyerByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &accountpb.BuyerResponse{
		BuyerId:           buyer.ID.String(),
		Email:             ptrToStr(buyer.Email),
		PhoneNumber:       *buyer.PhoneNumber,
		WalletAddress:     ptrToStr(buyer.WalletAddress),
		FirstName:         ptrToStr(buyer.FirstName),
		LastName:          ptrToStr(buyer.LastName),
		Address:           ptrToStr(buyer.Address),
		ProfilePictureUrl: ptrToStr(buyer.ProfilePictureURL),
		Badge:             string(buyer.Badge),
		CreatedAt:         buyer.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         buyer.UpdatedAt.Format(time.RFC3339),
	}, nil
}
func (h *AccountGrpcHandler) GetBuyerByWallet(ctx context.Context, req *accountpb.GetByWalletRequest) (*accountpb.BuyerResponse, error) {
	if req.WalletAddress == "" {
		return nil, status.Errorf(codes.InvalidArgument, "wallet address is required")
	}

	buyer, err := h.authSvc.GetBuyerByWallet(ctx, req.WalletAddress)
	if err != nil {
		return nil, err
	}
	if buyer == nil {
		return nil, status.Errorf(codes.NotFound, "buyer with wallet address %s not found", req.WalletAddress)
	}

	return &accountpb.BuyerResponse{
		BuyerId:           buyer.ID.String(),
		Email:             ptrToStr(buyer.Email),
		PhoneNumber:       ptrToStr(buyer.PhoneNumber),
		WalletAddress:     ptrToStr(buyer.WalletAddress),
		FirstName:         ptrToStr(buyer.FirstName),
		LastName:          ptrToStr(buyer.LastName),
		Address:           ptrToStr(buyer.Address),
		ProfilePictureUrl: ptrToStr(buyer.ProfilePictureURL),
		Badge:             string(buyer.Badge),
		CreatedAt:         buyer.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         buyer.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (h *AccountGrpcHandler) UpdateBuyer(ctx context.Context, req *accountpb.UpdateBuyerRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(req.BuyerId)
	if err != nil {
		return nil, err
	}
	input := &models.Buyer{
		ID:                id,
		Email:             &req.Email,
		PhoneNumber:       &req.PhoneNumber,
		WalletAddress:     &req.WalletAddress,
		FirstName:         &req.FirstName,
		LastName:          &req.LastName,
		Address:           &req.Address,
		ProfilePictureURL: &req.ProfilePictureUrl,
	}
	if err := h.authSvc.UpdateBuyer(ctx, input); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *AccountGrpcHandler) DeleteBuyer(ctx context.Context, req *accountpb.GetByIDRequest) (*emptypb.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, h.authSvc.DeleteBuyer(ctx, id)
}

func (h *AccountGrpcHandler) GetSettings(ctx context.Context, _ *emptypb.Empty) (*accountpb.SettingsResponse, error) {
	settings, err := h.authSvc.GetSettings(ctx)
	if err != nil {
		return nil, err
	}

	return &accountpb.SettingsResponse{
		SettingId:          settings.ID.String(),
		BadgesPrice:        int32(settings.BadgesPrice),
		RegistrationReward: settings.RegistrationReward,
		RatingReward:       settings.RatingReward,
		PurchasingReward:   settings.PurchasingReward,
		CoinToCurrencyRate: settings.CoinToCurrencyRate,
		DeliveryPriceRules: string(settings.DeliveryPriceRules),
		CreatedAt:          settings.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          settings.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (h *AccountGrpcHandler) UpdateSettings(ctx context.Context, req *accountpb.SettingsRequest) (*emptypb.Empty, error) {
	jsonData := json.RawMessage(req.DeliveryPriceRules)
	input := &models.Settings{
		BadgesPrice:        int(req.BadgesPrice),
		RegistrationReward: req.RegistrationReward,
		RatingReward:       req.RatingReward,
		PurchasingReward:   req.PurchasingReward,
		CoinToCurrencyRate: req.CoinToCurrencyRate,
		DeliveryPriceRules: jsonData,
	}
	return &emptypb.Empty{}, h.authSvc.UpdateSettings(ctx, input)
}

func (h *AccountGrpcHandler) InsertSettings(ctx context.Context, req *accountpb.SettingsRequest) (*emptypb.Empty, error) {
	jsonData := json.RawMessage(req.DeliveryPriceRules)
	input := &models.Settings{
		BadgesPrice:        int(req.BadgesPrice),
		RegistrationReward: req.RegistrationReward,
		RatingReward:       req.RatingReward,
		PurchasingReward:   req.PurchasingReward,
		CoinToCurrencyRate: req.CoinToCurrencyRate,
		DeliveryPriceRules: jsonData,
	}
	return &emptypb.Empty{}, h.authSvc.InsertSettings(ctx, input)
}

func ptrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
