package repository

import (
	"errors"

	"github.com/fxfrancky/go-api-eshop/internal/models"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

// CreateProduct implements ProductRepository
func (p *ProductRepositoryImpl) CreateProduct(product models.Product) error {
	result := p.DB.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteProduct implements ProductRepository
func (p *ProductRepositoryImpl) DeleteProduct(product *models.Product) error {
	return p.DB.Delete(product).Error
}

// GetProductById implements ProductRepository
func (p *ProductRepositoryImpl) GetProductById(productId int) (*models.Product, error) {
	var product models.Product
	err := p.DB.First(&product, productId).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &product, err
}

// GetAllProducts implements ProductRepository
func (p *ProductRepositoryImpl) GetAllProducts(offset, limit int) ([]models.Product, int64, error) {
	var (
		products []models.Product
		count    int64
	)

	p.DB.Model(&products).Count(&count)
	p.DB.Offset(offset).Limit(limit).Find(&products)

	return products, count, nil
}

// GetTopProducts implements ProductRepository
func (p *ProductRepositoryImpl) GetTopProducts(limit int) ([]models.Product, int64, error) {
	var (
		products []models.Product
		count    int64
	)
	p.DB.Model(&products).Count(&count)
	p.DB.Limit(limit).Order("rating desc").Find(&products)

	return products, count, nil
}

// UpdateProduct implements ProductRepository
func (p *ProductRepositoryImpl) UpdateProduct(product *models.Product) error {
	result := p.DB.Model(&product).Where("id = ?", product.ID).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateProductReviews implements ProductRepository
func (p *ProductRepositoryImpl) AddReviewToProduct(product *models.Product, review *models.Review) error {
	err := p.DB.Model(product).Association("Reviews").Append(review)
	if err != nil {
		return err
	}
	return p.DB.Where(review.ID).Preload("User").First(review).Error
}

// DeleteReview implements ProductRepository
func (p *ProductRepositoryImpl) DeleteReview(review *models.Review) error {
	return p.DB.Delete(review).Error
}

// GetReviewById implements ProductRepository
func (p *ProductRepositoryImpl) GetReviewById(reviewId int) (*models.Review, error) {
	var review models.Review
	err := p.DB.Preload("Product").First(&review, reviewId).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &review, err
}

// UpdateReview implements ProductRepository
func (p *ProductRepositoryImpl) UpdateReview(review *models.Review) error {
	result := p.DB.Model(&review).Where("id = ?", review.ID).Updates(review)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllReviewByProductId implements ProductRepository
func (p *ProductRepositoryImpl) GetAllReviewByProductName(productName string, offset int, limit int) ([]models.Review, int64, error) {
	var (
		product models.Product
		count   int64
	)
	// var m model.Article
	err := p.DB.Where(&models.Product{Name: productName}).Preload("Reviews").Preload("Reviews.User").Offset(offset).Limit(limit).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	count = int64(len(product.Reviews))
	return product.Reviews, count, nil
}

func NewProductRepositoryImpl(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: DB}
}
