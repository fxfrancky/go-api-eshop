package repository

import "github.com/fxfrancky/go-api-eshop/internal/models"

type ProductRepository interface {
	// Product
	CreateProduct(product models.Product) error
	GetProductById(productId int) (*models.Product, error)
	DeleteProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	GetTopProducts(limit int) ([]models.Product, int64, error)
	GetAllProducts(offset, limit int) ([]models.Product, int64, error)
	// Product Review
	AddReviewToProduct(product *models.Product, review *models.Review) error
	GetReviewById(reviewId int) (*models.Review, error)
	DeleteReview(review *models.Review) error
	UpdateReview(review *models.Review) error
	GetAllReviewByProductName(product_name string, offset, limit int) ([]models.Review, int64, error)
}
