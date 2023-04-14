package handlers

import (
	"time"

	"github.com/fxfrancky/go-api-eshop/internal/models"
)

// Product Response
type productResponse struct {
	ID uint `json:"id,omitempty"`
	// UserID      uint   `json:"user_id" validate:"required"`
	Name        string `json:"name,omitempty"`
	Image       string `json:"image"`
	Brand       string `json:"brand,omitempty"`
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	// Reviews      []Review
	Rating       float64   `json:"rating"`
	NumReviews   int64     `json:"numReviews"`
	Price        float64   `json:"price,omitempty"`
	CountInStock int64     `json:"countInStock"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

func newProductResponse(p *models.Product) *productResponse {

	productResp := new(productResponse)
	productResp.ID = p.ID
	// productResp.UserID = p.UserID
	productResp.Name = p.Name
	productResp.Image = p.Image
	productResp.Brand = p.Brand
	productResp.Category = p.Category
	productResp.Description = p.Description
	productResp.Rating = p.Rating
	productResp.NumReviews = p.NumReviews
	productResp.Price = p.Price
	productResp.CountInStock = p.CountInStock
	productResp.CreatedAt = p.CreatedAt
	productResp.UpdatedAt = p.UpdatedAt

	return productResp
}

type productListResponse struct {
	Products      []*productResponse `json:"products"`
	ProductsCount int64              `json:"productsCount"`
}

func newProductListResponse(products []models.Product, count int64) *productListResponse {
	r := new(productListResponse)
	r.Products = make([]*productResponse, 0)
	for _, p := range products {
		pr := newProductResponse(&p)
		r.Products = append(r.Products, pr)
	}
	r.ProductsCount = count
	return r
}

// Review Response
type reviewResponse struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Rating      float64   `json:"rating"`
	Comment     string    `json:"comment,omitempty"`
	UserID      uint      `json:"user_id,omitempty"`
	ProductID   uint      `json:"product_id,omitempty"`
	ProductName string    `json:"product_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func newReviewResponse(r *models.Review) *reviewResponse {
	rev := new(reviewResponse)
	rev.ID = r.ID
	rev.Name = r.Name
	rev.Name = r.Name
	rev.Comment = r.Comment
	rev.Rating = r.Rating
	rev.UserID = r.UserID
	rev.ProductID = r.ProductID
	rev.Name = r.Product.Name
	rev.CreatedAt = r.CreatedAt
	rev.UpdatedAt = r.UpdatedAt
	return rev
}

type reviewListResponse struct {
	Reviews      []*reviewResponse `json:"reviews"`
	ReviewsCount int64             `json:"reviewsCount"`
}

func newReviewListResponse(reviews []models.Review, count int64) *reviewListResponse {
	r := new(reviewListResponse)
	r.Reviews = make([]*reviewResponse, 0)
	for _, rev := range reviews {
		rr := newReviewResponse(&rev)
		r.Reviews = append(r.Reviews, rr)
	}
	r.ReviewsCount = count
	return r
}
