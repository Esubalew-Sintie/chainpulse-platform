package postgresrepo
import (
	"context"
	"github.com/chainpulse/backend/product/internal/models"
)

type IProductRepo interface {
	// Categories
	CreateCategory(ctx context.Context, category *models.Category) error
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*models.Category, error)

	// Products
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, id string) error

	// Sizes
	AddProductSize(ctx context.Context, size *models.ProductSize) error
	GetSizesByProductID(ctx context.Context, productID string) ([]models.ProductSize, error)

	// Reviews
	AddReview(ctx context.Context, review *models.ProductReview) error
	GetReviewsByProductID(ctx context.Context, productID string) ([]models.ProductReview, error)

	// Offers
	CreatePriceOffer(ctx context.Context, offer *models.PriceOffer) error
	GetOffersByProductID(ctx context.Context, productID string) ([]models.PriceOffer, error)

	// Wishlist
	AddToWishlist(ctx context.Context, wishlist *models.Wishlist) error
	RemoveFromWishlist(ctx context.Context, buyerID, productID string) error
	GetWishlistByBuyerID(ctx context.Context, buyerID string) ([]models.Wishlist, error)
}
