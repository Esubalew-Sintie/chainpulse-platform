package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Badge string

const (
	BadgeBasic     Badge = "basic"
	BadgeExclusive Badge = "exclusive"
)

type Buyer struct {
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"buyer_id"`
	Email             *string   `gorm:"type:varchar(100);unique" json:"email,omitempty"`
	PhoneNumber       *string   `gorm:"type:varchar(20);unique" json:"phone_number,omitempty"`
	Password          string    `gorm:"type:varchar(255)" json:"-"`
	WalletAddress     *string   `gorm:"type:varchar(255);unique" json:"wallet_address,omitempty"`
	FirstName         *string   `gorm:"type:varchar(100)" json:"first_name,omitempty"`
	LastName          *string   `gorm:"type:varchar(100)" json:"last_name,omitempty"`
	Address           *string   `gorm:"type:text" json:"address,omitempty"`
	ProfilePictureURL *string   `gorm:"type:varchar(255)" json:"profile_picture_url,omitempty"`
	Badge             Badge     `gorm:"type:varchar(10);default:'basic'" json:"badge"`

	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`
}

type Settings struct {
	ID                 uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"setting_id"`
	BadgesPrice        int             `json:"badges_price"`
	RegistrationReward string          `gorm:"type:varchar(255)" json:"registration_reward"`
	RatingReward       string          `gorm:"type:varchar(255)" json:"rating_reward"`
	PurchasingReward   string          `gorm:"type:varchar(255)" json:"purchasing_reward"`
	CoinToCurrencyRate float64         `gorm:"type:decimal(10,2)" json:"coin_to_currency_rate"`
	DeliveryPriceRules json.RawMessage `gorm:"type:jsonb" json:"delivery_price_rules"`

	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now()" json:"updated_at"`
}
