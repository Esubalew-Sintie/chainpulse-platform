package postgresrepo

import (
	"context"

	"gorm.io/gorm"
	"github.com/chainpulse/backend/product/internal/models"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IProductRepo {
	return &Repo{db: db}
}

// Categories
func (r *Repo) CreateCategory(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *Repo) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (r *Repo) GetCategoryByID(ctx context.Context, id string) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).First(&category, "category_id = ?", id).Error
	return &category, err
}

// Products
func (r *Repo) CreateProduct(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *Repo) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).First(&product, "product_id = ?", id).Error
	return &product, err
}

func (r *Repo) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err
}

func (r *Repo) UpdateProduct(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *Repo) DeleteProduct(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.Product{}, "product_id = ?", id).Error
}

// Sizes
func (r *Repo) AddProductSize(ctx context.Context, size *models.ProductSize) error {
	return r.db.WithContext(ctx).Create(size).Error
}

func (r *Repo) GetSizesByProductID(ctx context.Context, productID string) ([]models.ProductSize, error) {
	var sizes []models.ProductSize
	err := r.db.WithContext(ctx).Where("product_id = ?", productID).Find(&sizes).Error
	return sizes, err
}

// Reviews
func (r *Repo) AddReview(ctx context.Context, review *models.ProductReview) error {
	return r.db.WithContext(ctx).Create(review).Error
}

func (r *Repo) GetReviewsByProductID(ctx context.Context, productID string) ([]models.ProductReview, error) {
	var reviews []models.ProductReview
	err := r.db.WithContext(ctx).Where("product_id = ?", productID).Find(&reviews).Error
	return reviews, err
}

// Offers
func (r *Repo) CreatePriceOffer(ctx context.Context, offer *models.PriceOffer) error {
	return r.db.WithContext(ctx).Create(offer).Error
}

func (r *Repo) GetOffersByProductID(ctx context.Context, productID string) ([]models.PriceOffer, error) {
	var offers []models.PriceOffer
	err := r.db.WithContext(ctx).Where("product_id = ?", productID).Find(&offers).Error
	return offers, err
}

// Wishlist
func (r *Repo) AddToWishlist(ctx context.Context, wishlist *models.Wishlist) error {
	return r.db.WithContext(ctx).Create(wishlist).Error
}

func (r *Repo) RemoveFromWishlist(ctx context.Context, buyerID, productID string) error {
	return r.db.WithContext(ctx).Where("buyer_id = ? AND product_id = ?", buyerID, productID).Delete(&models.Wishlist{}).Error
}

func (r *Repo) GetWishlistByBuyerID(ctx context.Context, buyerID string) ([]models.Wishlist, error) {
	var wishlist []models.Wishlist
	err := r.db.WithContext(ctx).Where("buyer_id = ?", buyerID).Find(&wishlist).Error
	return wishlist, err
}
