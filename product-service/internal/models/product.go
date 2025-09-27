package models

import (
	"time"

	"github.com/lib/pq"
)

type Category struct {
	CategoryID       string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"category_id"`
	Name             string    `gorm:"type:varchar(50);not null" json:"name"`
	ParentCategoryID *string   `gorm:"type:uuid" json:"parent_category_id"`
	ImageURL         *string   `gorm:"type:varchar(255)" json:"image_url"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
type Product struct {
	ProductID            string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"product_id"`
	Name                 string         `gorm:"type:varchar(100);not null" json:"name"`
	Description          *string        `gorm:"type:text" json:"description"`
	MainCategoryID       string         `gorm:"type:uuid;not null" json:"main_category_id"`
	SubCategoryID        *string        `gorm:"type:uuid" json:"sub_category_id"`
	BadgeType            *string        `gorm:"type:varchar(9);check:badge_type IN ('exclusive','normal','bonda')" json:"badge_type"`
	IsAcceptOffer        bool           `gorm:"default:false" json:"is_accept_offer"`
	IsForExclusiveAccess bool           `gorm:"default:false" json:"is_for_exclusive_access"`
	SKU                  string         `gorm:"type:varchar(50);unique" json:"sku"`
	ProductImages        pq.StringArray `gorm:"type:text[]" json:"product_images"`
	ThumbnailImages      pq.StringArray `gorm:"type:text[]" json:"thumbnail_images"`
	BasePrice            float64        `gorm:"type:decimal(10,2);not null" json:"base_price"`
	DiscountType         *string        `gorm:"type:varchar(10);check:discount_type IN ('percentage','fixed')" json:"discount_type"`
	DiscountValue        float64        `gorm:"type:decimal(10,2);not null" json:"discount_value"`
	AllowOffers          bool           `gorm:"default:true" json:"allow_offers"`
	Attributes           map[string]any `gorm:"type:jsonb" json:"attributes"`
	CreatedAt            time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
type ProductSize struct {
	SizeID    string                 `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"size_id"`
	ProductID string                 `gorm:"type:uuid;not null" json:"product_id"`
	Size      string                 `gorm:"type:varchar(10);not null" json:"size"`
	Name      map[string]interface{} `gorm:"type:jsonb" json:"name"` // e.g., [{ "color": "red", "quantity": 10 }]
}
type ProductReview struct {
	ReviewID   string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"review_id"`
	ProductID  string    `gorm:"type:uuid;not null" json:"product_id"`
	BuyerID    string    `gorm:"type:uuid;not null" json:"buyer_id"`
	Rating     int       `gorm:"type:int;check:rating BETWEEN 1 AND 5" json:"rating"`
	ReviewText *string   `gorm:"type:text" json:"review_text"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type PriceOffer struct {
	OfferID      string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"offer_id"`
	BuyerID      string    `gorm:"type:uuid;not null" json:"buyer_id"`
	ProductID    string    `gorm:"type:uuid;not null" json:"product_id"`
	OfferedPrice float64   `gorm:"type:decimal(10,2);not null" json:"offered_price"`
	Status       string    `gorm:"type:varchar(10);default:'pending'" json:"status"` // pending, accepted, rejected
	Count        int       `gorm:"type:int;not null" json:"count"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
type Wishlist struct {
	WishlistID string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"wishlist_id"`
	BuyerID    string    `gorm:"type:uuid;not null" json:"buyer_id"`
	ProductID  string    `gorm:"type:uuid;not null" json:"product_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
