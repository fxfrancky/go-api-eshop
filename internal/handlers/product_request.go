package handlers

import (
	"github.com/fxfrancky/go-api-eshop/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Product Requests
type createProductRequest struct {
	// UserID      uint   `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Image       string `json:"image"`
	Brand       string `json:"brand" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Description string `json:"description" validate:"required"`
	// Reviews      []Review `json:"reviewList"`
	Rating       float64 `json:"rating"`
	NumReviews   int64   `json:"numReviews"`
	Price        float64 `json:"price" validate:"required"`
	CountInStock int64   `json:"countInStock"`
}

func bindCreateProduct(r *createProductRequest, c *fiber.Ctx, p *models.Product) error {

	// Validate the product
	if err := c.BodyParser(r); err != nil {
		return err
	}
	// Map the product
	// p.UserID = r.UserID
	p.Name = r.Name
	p.Image = r.Image
	p.Brand = r.Brand
	p.Category = r.Category
	p.Description = r.Description
	// p.Reviews = r.Reviews
	p.Rating = r.Rating
	p.NumReviews = r.NumReviews
	p.Price = r.Price
	p.CountInStock = r.CountInStock

	return nil
}

type updateProductRequest struct {
	// ID          uint   `json:"id" validate:"required"`
	// UserID      uint   `json:"user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Image       string `json:"image"`
	Brand       string `json:"brand" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Description string `json:"description" validate:"required"`
	// Reviews      []Review `json:"reviewList"`
	Rating       float64 `json:"rating"`
	NumReviews   int64   `json:"numReviews"`
	Price        float64 `json:"price" validate:"required"`
	CountInStock int64   `json:"countInStock"`
}

func (u *updateProductRequest) bindUpdateProduct(c *fiber.Ctx, p *models.Product) error {

	// Validate the product
	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Map the product
	// p.UserID = u.UserID
	p.Name = u.Name
	p.Image = u.Image
	p.Brand = u.Brand
	p.Category = u.Category
	p.Description = u.Description
	// p.Reviews = r.Reviews
	p.Rating = u.Rating
	p.NumReviews = u.NumReviews
	p.Price = u.Price
	p.CountInStock = u.CountInStock

	return nil
}

func (u *updateProductRequest) populateUpdateProduct(p *models.Product) {
	// u.ID = int(p.ID)
	// u.UserID = p.UserID
	u.Name = p.Name
	u.Image = p.Image
	u.Brand = p.Brand
	u.Category = p.Category
	u.Description = p.Description
	u.Rating = p.Rating
	u.NumReviews = p.NumReviews
	u.Price = p.Price
	u.CountInStock = p.CountInStock
}

// Review Request
type reviewRequest struct {
	Name      string  `json:"name" validate:"required"`
	Rating    float64 `json:"rating" validate:"required"`
	Comment   string  `json:"comment" validate:"required"`
	UserID    uint    `json:"user_id" validate:"required"`
	ProductID uint    `json:"product_id"`
}

func (r *reviewRequest) populateReview(review *models.Review) {
	r.Name = review.Name
	r.Rating = review.Rating
	r.Comment = review.Comment
	r.UserID = review.UserID
	r.ProductID = review.ProductID
}

func bindReviewRequest(rev *reviewRequest, c *fiber.Ctx, r *models.Review) error {

	// Validate the Review
	if err := c.BodyParser(rev); err != nil {
		return err
	}
	// Map the review
	r.UserID = rev.UserID
	r.ProductID = rev.ProductID
	r.Name = rev.Name
	r.Comment = rev.Comment
	r.Rating = rev.Rating
	return nil

}
