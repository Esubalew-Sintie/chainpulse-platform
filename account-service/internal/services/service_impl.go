package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/chainpulse/backend/account/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrBuyerExistsByEmail  = errors.New("buyer with this email already exists")
	ErrBuyerExistsByPhone  = errors.New("buyer with this phone already exists")
	ErrBuyerExistsByWallet = errors.New("buyer with this wallet address already exists")
	ErrInvalidBuyerData    = errors.New("invalid buyer data")
	ErrBuyerNotFound       = errors.New("buyer not found")
	ErrInvalidCredential   = errors.New("invalid email or password")
)

// CreateBuyer validates buyer info and checks existence before creating
func (s *authService) CreateBuyer(ctx context.Context, input *models.Buyer) error {
	// 🛡️ Validation: required fields
	if input == nil {
		return ErrInvalidBuyerData
	}
	if input.PhoneNumber == nil || *input.PhoneNumber == "" {
		return errors.New("phone number is required")
	}
	if input.Password == "" {
		return errors.New("password is required")
	}
	if len(input.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if input.Email != nil && !s.IsValidEmail(*input.Email) {
		return errors.New("invalid email format")
	}

	// 🔍 Check if buyer exists by email
	if input.Email != nil {
		existingByEmail, err := s.repo.GetBuyerByEmail(ctx, *input.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to check existing buyer by email: %w", err)
		}
		if existingByEmail != nil {
			return ErrBuyerExistsByEmail
		}
	}

	// 🔍 Check if buyer exists by phone
	existingByPhone, err := s.repo.GetBuyerByPhone(ctx, *input.PhoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check existing buyer by phone: %w", err)
	}
	if existingByPhone != nil {
		return ErrBuyerExistsByPhone
	}

	// 🔐 Hash password before saving
	hashedPassword, err := s.hashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	input.Password = hashedPassword

	return s.repo.CreateBuyer(ctx, input)
}
func (s *authService) CreateWalletOnlyBuyer(ctx context.Context, input *models.Buyer) error {
	if input == nil || input.WalletAddress == nil || *input.WalletAddress == "" {
		log.Println("wallet address is required for wallet-only buyer creation")
		return errors.New("wallet address is required")
	}

	// Check if buyer exists by wallet
	existing, err := s.repo.GetBuyerByWallet(ctx, *input.WalletAddress)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("failed to check buyer by wallet: %v", err)
		return fmt.Errorf("failed to check buyer by wallet: %w", err)
	}
	if existing != nil {
		log.Println("buyer with wallet address %s already exists", *input.WalletAddress,"existing: ",*existing)
		return ErrBuyerExistsByWallet
	}

	log.Println("creating wallet-only buyer with wallet address %s", *input.WalletAddress)
	return s.repo.CreateBuyer(ctx, input)
}

func (s *authService) LoginBuyer(ctx context.Context, email, password string) (*models.Buyer, error) {
	buyer, err := s.repo.GetBuyerByPhone(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredential
		}
		return nil, fmt.Errorf("failed to find buyer: %w", err)
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(buyer.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredential
	}

	return buyer, nil
}

func (s *authService) GetBuyerByID(ctx context.Context, id uuid.UUID) (*models.Buyer, error) {
	return s.repo.GetBuyerByID(ctx, id)
}
func (s *authService) GetBuyerByWallet(ctx context.Context, walletAddress string) (*models.Buyer, error) {
	return s.repo.GetBuyerByWallet(ctx, walletAddress)
}

func (s *authService) GetBuyerByEmail(ctx context.Context, email string) (*models.Buyer, error) {
	return s.repo.GetBuyerByEmail(ctx, email)
}

func (s *authService) UpdateBuyer(ctx context.Context, input *models.Buyer) error {
	// Add validation or business logic here if needed
	return s.repo.UpdateBuyer(ctx, input)
}

func (s *authService) DeleteBuyer(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteBuyer(ctx, id)
}

func (s *authService) GetSettings(ctx context.Context) (*models.Settings, error) {
	return s.repo.GetSettings(ctx)
}

func (s *authService) UpdateSettings(ctx context.Context, input *models.Settings) error {
	return s.repo.UpdateSettings(ctx, input)
}

func (s *authService) InsertSettings(ctx context.Context, input *models.Settings) error {
	return s.repo.InsertSettings(ctx, input)
}
