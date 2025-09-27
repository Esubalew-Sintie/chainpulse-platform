package service

import (
	"context"

	postgresrepo "github.com/chainpulse/backend/account/internal/infrastructure/postgres_repo"
	"github.com/chainpulse/backend/account/internal/models"
	"github.com/google/uuid"
)

type IAuthSvc interface {
	CreateBuyer(ctx context.Context, input *models.Buyer) error
	CreateWalletOnlyBuyer(ctx context.Context, input *models.Buyer) error 
	LoginBuyer(ctx context.Context, email, password string) (*models.Buyer, error)
	GetBuyerByID(ctx context.Context, id uuid.UUID) (*models.Buyer, error)
	GetBuyerByEmail(ctx context.Context, email string) (*models.Buyer, error)
	GetBuyerByWallet(ctx context.Context, walletAddress string) (*models.Buyer, error)
	UpdateBuyer(ctx context.Context, input *models.Buyer) error
	DeleteBuyer(ctx context.Context, id uuid.UUID) error

	GetSettings(ctx context.Context) (*models.Settings, error)
	UpdateSettings(ctx context.Context, input *models.Settings) error
	InsertSettings(ctx context.Context, input *models.Settings) error
}

type authService struct {
	repo postgresrepo.IAuth
}

func NewAuthService(repo postgresrepo.IAuth) IAuthSvc {
	return &authService{repo: repo}
}
