package handlers

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/chainpulse/backend/api-gateway/internal/routes/handlers/dto"
	accountpb "github.com/chainpulse/backend/api-gateway/proto/gen/account-svc"
	"github.com/google/uuid"
	"github.com/mr-tron/base58" // Solana wallet address decoding
	"github.com/patrickmn/go-cache"
)

// In-memory nonce storage (5 min TTL)
var nonceCache = cache.New(5*time.Minute, 10*time.Minute)

// RequestNonceHandler generates a unique nonce for the client to sign
func (s *Handlers) RequestNonceHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		WalletAddress string `json:"walletAddress"`
	}
	s.handlers.Logger.Info("Requesting nonce for wallet address", "walletAddress", req.WalletAddress)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.WalletAddress == "" {
		s.handlers.Logger.Error("Invalid request for nonce", "error", err)
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid wallet address")
		return
	}

	nonce := "Sign this to login: " + uuid.NewString()
	nonceCache.Set(req.WalletAddress, nonce, cache.DefaultExpiration)
	s.handlers.Logger.Info("Generated nonce for wallet address", "walletAddress", req.WalletAddress, "nonce", nonce)
	dto.WriteSuccessResponse(w, map[string]string{
		"nonce": nonce,
	}, "Nonce generated")
}

// VerifySignatureHandler verifies signed nonce and logs in or registers a buyer
func (s *Handlers) VerifySignatureHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		WalletAddress string `json:"walletAddress"`
		Signature     string `json:"signature"` // base64 encoded
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.handlers.Logger.Error("Failed to decode request body", "error", err)
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// 1. Fetch stored nonce
	value, found := nonceCache.Get(req.WalletAddress)
	if !found {
		s.handlers.Logger.Error("Nonce not found or expired", "walletAddress", req.WalletAddress)
		dto.WriteErrorResponse(w, http.StatusUnauthorized, "Nonce expired or missing")
		return
	}
	nonce := value.(string)

	// 2. Decode Solana wallet address (base58) and signature (base64)
	pubKeyBytes, err := base58.Decode(req.WalletAddress)
	if err != nil || len(pubKeyBytes) != ed25519.PublicKeySize {
		s.handlers.Logger.Error("Invalid wallet address", "walletAddress", req.WalletAddress, "error", err)
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid wallet address")
		return
	}

	sigBytes, err := base64.StdEncoding.DecodeString(req.Signature)
	if err != nil || len(sigBytes) != ed25519.SignatureSize {
		s.handlers.Logger.Error("Invalid signature", "signature", req.Signature, "error", err)
		dto.WriteErrorResponse(w, http.StatusBadRequest, "Invalid signature")
		return
	}

	// 3. Verify signature
	if !ed25519.Verify(pubKeyBytes, []byte(nonce), sigBytes) {
		s.handlers.Logger.Error("Signature verification failed", "walletAddress", req.WalletAddress)
		dto.WriteErrorResponse(w, http.StatusUnauthorized, "Signature verification failed")
		return
	}

	// 4. Look up buyer by wallet address
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	buyerResp, err := s.accGrpcClient.GetBuyerByWallet(ctx, &accountpb.GetByWalletRequest{
		WalletAddress: req.WalletAddress,
	})

	if err != nil && strings.Contains(err.Error(), "not found") {
		s.handlers.Logger.Info("Buyer not found, registering new buyer", "walletAddress", req.WalletAddress)
		// 5. Register wallet-only buyer if not found
		_, err := s.accGrpcClient.CreateBuyer(ctx, &accountpb.CreateBuyerRequest{
			WalletAddress: req.WalletAddress,
		})
		if err != nil {
			s.handlers.Logger.Error("Failed to create buyer", "walletAddress", req.WalletAddress, "error", err)
			dto.WriteGrpcErrorResponse(w, err)
			return
		}
		buyerResp, err = s.accGrpcClient.GetBuyerByWallet(ctx, &accountpb.GetByWalletRequest{
			WalletAddress: req.WalletAddress,
		})
		if err != nil {
			s.handlers.Logger.Error("Failed to fetch buyer after creation", "walletAddress", req.WalletAddress, "error", err)
			dto.WriteGrpcErrorResponse(w, err)
			return
		}
	} else if err != nil {
		s.handlers.Logger.Error("Failed to fetch buyer", "walletAddress", req.WalletAddress, "error", err)
		dto.WriteGrpcErrorResponse(w, err)
		return
	}

	// 6. Respond with buyer info (JWT/session token could be added here)
	s.handlers.Logger.Info("Wallet login successful", "walletAddress", req.WalletAddress, "buyerId", buyerResp.BuyerId)
	dto.WriteSuccessResponse(w, map[string]string{
		"buyerId":       buyerResp.BuyerId,
		"walletAddress": buyerResp.WalletAddress,
	}, "Wallet login successful")
}
