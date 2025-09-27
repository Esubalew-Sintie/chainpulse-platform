package handlers

import (
	"net/http"

	"github.com/chainpulse/backend/api-gateway/grpcclients"
	"github.com/chainpulse/backend/api-gateway/internal/routes"
	"github.com/chainpulse/backend/api-gateway/internal/routes/handlers/dto"
	"github.com/chainpulse/backend/api-gateway/pkg/config"
	accountpb "github.com/chainpulse/backend/api-gateway/proto/gen/account-svc"
)

type apiClaimsCtxKey string

const (
	maxUploadSize                       = 20 * 1024 * 1024
	userApiClaimsCtxKey apiClaimsCtxKey = "userApiClaimsCtxKey"
)

type Handlers struct {
	handlers          *routes.Routes
	accGrpcClient     accountpb.AccountServiceClient
	productGrpcClient grpcclients.ProductGrpcClient
	paymentGrpcClient grpcclients.PaymGrpcClient
	Config            *config.GlobalConfig
}

func NewHandlers(routes *routes.Routes, AccGrpcClient accountpb.AccountServiceClient, proGrpcClient grpcclients.ProductGrpcClient, payGrpcClient grpcclients.PaymGrpcClient, config *config.GlobalConfig) *Handlers {
	return &Handlers{
		handlers:          routes,
		accGrpcClient:     AccGrpcClient,
		productGrpcClient: proGrpcClient,
		paymentGrpcClient: payGrpcClient,
		Config:            config,
	}
}

// @BasePath /account
func (h *Handlers) SetupRoutes() {
	h.setupAccountRoutes()
	h.setupProductRoutes()
	h.setupPaymentRoutes()
}

func (h *Handlers) setupPaymentRoutes() {
	route := h.handlers.Route.PathPrefix("/payment").Subrouter()

	route.HandleFunc("/create-order", func(w http.ResponseWriter, r *http.Request) {
		h.handlers.Logger.Info("Creating order")
		dto.WriteSuccessResponse(w, nil, "Order created successfully")
	}).Methods("POST")
}

func (h *Handlers) setupProductRoutes() {
	route := h.handlers.Route.PathPrefix("/product").Subrouter()

	route.HandleFunc("/create-product", func(w http.ResponseWriter, r *http.Request) {
		h.handlers.Logger.Info("Creating product")
		dto.WriteSuccessResponse(w, nil, "product created successfully")
	}).Methods("POST")
}

func (h *Handlers) setupAccountRoutes() {
	route := h.handlers.Route.PathPrefix("/account").Subrouter()
	route.HandleFunc("/register", h.CreateBuyerHandler).Methods("POST")
	route.HandleFunc("/buyers", h.CreateBuyerHandler).Methods("POST")
	route.HandleFunc("/request-nonce", h.RequestNonceHandler).Methods("POST")
	route.HandleFunc("/verify-signature", h.VerifySignatureHandler).Methods("POST")
	route.HandleFunc("/buyers/login", h.LoginBuyerHandler).Methods("POST")
	route.HandleFunc("/buyers/{id}", h.GetBuyerByIDHandler).Methods("GET")
	route.HandleFunc("/buyers/{id}", h.UpdateBuyerHandler).Methods("PUT")
	route.HandleFunc("/buyers/{id}", h.DeleteBuyerHandler).Methods("DELETE")

	route.HandleFunc("/settings", h.GetSettingsHandler).Methods("GET")
	route.HandleFunc("/settings", h.UpdateSettingsHandler).Methods("PUT")
	route.HandleFunc("/settings", h.InsertSettingsHandler).Methods("POST")
}
