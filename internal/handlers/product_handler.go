package handlers

import (
	"net/http"
	"strconv"

	"github.com/fxfrancky/go-api-eshop/internal/models"
	"github.com/fxfrancky/go-api-eshop/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetProduct func to get a product
// @Summary Get a product
// @Description Get a product. Auth required
// @ID get-product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the product to get"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 200 {object} productResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Security ApiKeyAuth
// @Router /api/v1/auth/products/{id} [get]
func (h *Handler) GetProduct(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	p, err := h.productRepository.GetProductById(id)
	if p == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Product"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newProductResponse(p))
}

// CreateProduct func to create a new Product
// @Summary create a new Product
// @Description CreateProduct create a new Product
// @ID create-product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param ProductRequest body createProductRequest true "ProductRequest"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 201 {object} productResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/products [post]
func (h *Handler) CreateProduct(c *fiber.Ctx) error {

	var p models.Product
	payload := &createProductRequest{}
	if err := bindCreateProduct(payload, c, &p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}
	err := h.productRepository.CreateProduct(p)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusCreated).JSON(newProductResponse(&p))
}

// UpdateProduct func to update a new Product
// @Summary Update a new Product
// @Description Update a Product. Auth is required
// @ID update-product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the product to update"
// @Param product body updateProductRequest true "Product to update"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} productResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/products/{id} [put]
func (h *Handler) UpdateProduct(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	p, err := h.productRepository.GetProductById(id)
	if p == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Product"))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	payload := &updateProductRequest{}
	payload.populateUpdateProduct(p)
	if err := payload.bindUpdateProduct(c, p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if err = h.productRepository.UpdateProduct(p); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newProductResponse(p))
}

// DeleteProduct func to delete a Product
// @Summary Delete a Product
// @Description Delete a Product. Auth is required
// @ID delete-product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the product to delete"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} productResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/products/{id} [delete]
func (h *Handler) DeleteProduct(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	p, err := h.productRepository.GetProductById(id)
	if p == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Product"))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	err = h.productRepository.DeleteProduct(p)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{"result": "product deleted!"})
}

// AllProducts func to display all Products
// @Summary Display all Products
// @Description Display all Products. Auth is required
// @ID all-products
// @Tags Product
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit number of products returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of products (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} productListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/products/all/{limit}/{offset} [get]
func (h *Handler) AllProducts(c *fiber.Ctx) error {
	var (
		products []models.Product
		count    int64
	)
	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
	}

	products, count, err = h.productRepository.GetAllProducts(offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newProductListResponse(products, count))
}

// TopProducts func to display top Products
// @Summary Display top Products
// @Description Display top Products. Auth is required
// @ID top-products
// @Tags Product
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit number of products returned (default is 20)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} productListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/products/top/{limit} [get]
func (h *Handler) TopProducts(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
	}

	products, count, err := h.productRepository.GetTopProducts(limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newProductListResponse(products, count))
}

// GetReview func to get a review
// @Summary Get a review
// @Description Get a review. Auth required
// @ID get-review
// @Tags Review
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the review to get"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 200 {object} reviewResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Security ApiKeyAuth
// @Router /api/v1/auth/reviews/{id} [get]
func (h *Handler) GetReview(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	r, err := h.productRepository.GetReviewById(id)
	if r == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Review"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newReviewResponse(r))
}

// CreateReview func to create a new Product Review
// @Summary create a new Product Review
// @Description CreateReview create a new Product Review
// @ID create-review
// @Tags Review
// @Accept  json
// @Produce  json
// @Param product_id path string true "Id of the product to get"
// @Param ReviewRequest body reviewRequest true "Review Request"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 201 {object} reviewResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/reviews/{product_id} [post]
func (h *Handler) AddReviewToProduct(c *fiber.Ctx) error {

	// Get The product
	productId, err := utils.StringToInt(c.Params("product_id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	prod, err := h.productRepository.GetProductById(productId)
	if prod == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Product"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	var r models.Review
	payload := &reviewRequest{}
	if err := bindReviewRequest(payload, c, &r); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	// err := h.productRepository.CreateProductReview(&r)
	err = h.productRepository.AddReviewToProduct(prod, &r)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newReviewResponse(&r))
}

// UpdateProduct func to update a new Product
// @Summary Update a review
// @Description Update a review. Auth is required
// @ID update-review
// @Tags Review
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the review to update"
// @Param product body reviewRequest true "Review to update"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} reviewResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/reviews/{id} [put]
func (h *Handler) UpdateReview(c *fiber.Ctx) error {
	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	r, err := h.productRepository.GetReviewById(id)
	if r == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Review"))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	payload := &reviewRequest{}
	payload.populateReview(r)
	if err := bindReviewRequest(payload, c, r); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if err = h.productRepository.UpdateReview(r); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newReviewResponse(r))
}

// DeleteReview func to delete a Review
// @Summary Delete a Review
// @Description Delete a Review. Auth is required
// @ID delete-review
// @Tags Review
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the review to delete"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {string}  string  "product deleted"
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/reviews/{id} [delete]
func (h *Handler) DeleteReview(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	r, err := h.productRepository.GetReviewById(id)
	if r == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Review"))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	err = h.productRepository.DeleteReview(r)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{"result": "review deleted!"})
}

// AllReviews func to display all Reviews By Product Id
// @Summary Display all Reviews By Product Id
// @Description Display all Reviews By Product Id. Auth is required
// @ID all-products-reviews
// @Tags Review
// @Accept  json
// @Produce  json
// @Param productName path string true "Name of the product for reviews"
// @Param limit query integer false "Limit number of products returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of products (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} reviewListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/auth/reviews/{productName}/product/{limit}/{offset} [get]
func (h *Handler) AllProductsReview(c *fiber.Ctx) error {
	var (
		reviews []models.Review
		count   int64
	)

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
		err = nil
	}

	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		offset = 0
		err = nil
	}

	productName := c.Params("productName")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	reviews, count, err = h.productRepository.GetAllReviewByProductName(productName, offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newReviewListResponse(reviews, count))
}
