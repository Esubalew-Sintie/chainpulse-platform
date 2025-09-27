package services

import (
	"context"
	"errors"
	"log/slog"

	postgresrepo "github.com/chainpulse/backend/product/internal/infrastructure/postgres_repo"
	"github.com/chainpulse/backend/product/internal/models"
	"github.com/chainpulse/backend/product/pkgs/config"
)

type productSvc struct {
	repo   postgresrepo.IProductRepo
	config *config.GlobalConfig
	logger *slog.Logger
}

func NewProductService(repo postgresrepo.IProductRepo, config *config.GlobalConfig, logger *slog.Logger) IProductSvc {
	return &productSvc{
		repo:   repo,
		config: config,
		logger: logger,
	}
}

// Products

func (s *productSvc) CreateProduct(ctx context.Context, product *models.Product) error {
	if product.Name == "" || product.BasePrice <= 0 {
		return errors.New("invalid product input")
	}
	return s.repo.CreateProduct(ctx, product)
}

func (s *productSvc) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	return s.repo.GetProductByID(ctx, id)
}

func (s *productSvc) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	return s.repo.GetAllProducts(ctx)
}

func (s *productSvc) UpdateProduct(ctx context.Context, product *models.Product) error {
	return s.repo.UpdateProduct(ctx, product)
}

func (s *productSvc) DeleteProduct(ctx context.Context, id string) error {
	return s.repo.DeleteProduct(ctx, id)
}

// Categories

func (s *productSvc) CreateCategory(ctx context.Context, category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}
	return s.repo.CreateCategory(ctx, category)
}

func (s *productSvc) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetAllCategories(ctx)
}

func (s *productSvc) GetCategoryByID(ctx context.Context, id string) (*models.Category, error) {
	return s.repo.GetCategoryByID(ctx, id)
}

// Sizes

func (s *productSvc) AddProductSize(ctx context.Context, size *models.ProductSize) error {
	return s.repo.AddProductSize(ctx, size)
}

func (s *productSvc) GetSizesByProductID(ctx context.Context, productID string) ([]models.ProductSize, error) {
	return s.repo.GetSizesByProductID(ctx, productID)
}

// Reviews

func (s *productSvc) AddReview(ctx context.Context, review *models.ProductReview) error {
	if review.ProductID == "" || review.BuyerID == "" {
		return errors.New("product_id and buyer_id are required")
	}
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return s.repo.AddReview(ctx, review)
}

func (s *productSvc) GetReviewsByProductID(ctx context.Context, productID string) ([]models.ProductReview, error) {
	return s.repo.GetReviewsByProductID(ctx, productID)
}

// Offers

func (s *productSvc) CreatePriceOffer(ctx context.Context, offer *models.PriceOffer) error {
	if offer.BuyerID == "" || offer.ProductID == "" {
		return errors.New("buyer_id and product_id are required")
	}
	return s.repo.CreatePriceOffer(ctx, offer)
}

func (s *productSvc) GetOffersByProductID(ctx context.Context, productID string) ([]models.PriceOffer, error) {
	return s.repo.GetOffersByProductID(ctx, productID)
}

// Wishlist

func (s *productSvc) AddToWishlist(ctx context.Context, wishlist *models.Wishlist) error {
	if wishlist.BuyerID == "" || wishlist.ProductID == "" {
		return errors.New("buyer_id and product_id are required")
	}
	return s.repo.AddToWishlist(ctx, wishlist)
}

func (s *productSvc) RemoveFromWishlist(ctx context.Context, buyerID, productID string) error {
	return s.repo.RemoveFromWishlist(ctx, buyerID, productID)
}

func (s *productSvc) GetWishlistByBuyerID(ctx context.Context, buyerID string) ([]models.Wishlist, error) {
	return s.repo.GetWishlistByBuyerID(ctx, buyerID)
}
