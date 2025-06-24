package postgresrepo

import (
	"context"
	"errors"

	"github.com/chainpulse/backend/account/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IAuth {
	return &Repo{db: db}
}

func (r *Repo) CreateBuyer(ctx context.Context, buyer *models.Buyer) error {
	return r.db.WithContext(ctx).Create(buyer).Error
}

func (r *Repo) GetBuyerByID(ctx context.Context, id uuid.UUID) (*models.Buyer, error) {
	var buyer models.Buyer
	if err := r.db.WithContext(ctx).First(&buyer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &buyer, nil
}

func (r *Repo) GetBuyerByEmail(ctx context.Context, email string) (*models.Buyer, error) {
	var buyer models.Buyer
	if err := r.db.WithContext(ctx).First(&buyer, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &buyer, nil
}

func (r *Repo) GetBuyerByPhone(ctx context.Context, phone string) (*models.Buyer, error) {
	var buyer models.Buyer
	if err := r.db.WithContext(ctx).First(&buyer, "phone_number = ?", phone).Error; err != nil {
		return nil, err
	}
	return &buyer, nil
}

func (r *Repo) UpdateBuyer(ctx context.Context, buyer *models.Buyer) error {
	return r.db.WithContext(ctx).Save(buyer).Error
}

func (r *Repo) DeleteBuyer(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Buyer{}, "id = ?", id).Error
}

// --- SettingsRepository implementation ---

func (r *Repo) GetSettings(ctx context.Context) (*models.Settings, error) {
	var settings models.Settings
	err := r.db.WithContext(ctx).First(&settings).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &settings, err
}

func (r *Repo) UpdateSettings(ctx context.Context, settings *models.Settings) error {
	return r.db.WithContext(ctx).Save(settings).Error
}

func (r *Repo) InsertSettings(ctx context.Context, settings *models.Settings) error {
	return r.db.WithContext(ctx).Create(settings).Error
}
