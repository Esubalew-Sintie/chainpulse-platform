package postgresrepo

import (
	"context"

	"github.com/chainpulse/backend/account/internal/models"
	"github.com/google/uuid"
)

type IAuth interface {
	CreateBuyer(ctx context.Context, buyer *models.Buyer) error
	GetBuyerByID(ctx context.Context, id uuid.UUID) (*models.Buyer, error)
	GetBuyerByEmail(ctx context.Context, email string) (*models.Buyer, error)
	GetBuyerByPhone(ctx context.Context, phone string) (*models.Buyer, error)
	GetBuyerByWallet(ctx context.Context, walletAddress string) (*models.Buyer, error)

	UpdateBuyer(ctx context.Context, buyer *models.Buyer) error
	DeleteBuyer(ctx context.Context, id uuid.UUID) error
	GetSettings(ctx context.Context) (*models.Settings, error)
	UpdateSettings(ctx context.Context, settings *models.Settings) error
	InsertSettings(ctx context.Context, settings *models.Settings) error
}
