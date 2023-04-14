package handlers

import (
	"net/http"
	"strconv"

	"github.com/fxfrancky/go-api-eshop/internal/models"
	"github.com/fxfrancky/go-api-eshop/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetOrderById func to get an order
// @Summary Get an order
// @Description Get an order. Auth required
// @ID get-order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the order to get"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 200 {object} orderResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Security ApiKeyAuth
// @Router /api/v1/orders/{id} [get]
func (h *Handler) GetOrderById(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	o, err := h.orderRepository.GetOrderById(id)
	if o == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Order"))
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newOrderResponse(o))
}

// CreateOrder func to create a new Order
// @Summary create a new Order
// @Description CreateOrder create a new Order
// @ID create-order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param order body orderRequest true "orderRequest"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success 201 {object} orderResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/orders [post]
func (h *Handler) CreateOrder(c *fiber.Ctx) error {

	var o models.Order
	payload := &orderRequest{}
	if err := bindOrderRequest(payload, c, &o); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}
	err := h.orderRepository.CreateOrder(o)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(newOrderResponse(&o))
}

// UpdateOrder func to update a new Order
// @Summary Update a new Order
// @Description Update a Order. Auth is required
// @ID update-order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the order to update"
// @Param order body orderRequest true "Order to update"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} orderResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/orders/{id} [put]
func (h *Handler) UpdateOrder(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	o, err := h.orderRepository.GetOrderById(id)
	if o == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Order"))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	payload := &orderRequest{}
	payload.populateOrder(o)
	if err := bindOrderRequest(payload, c, o); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	}
	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if err = h.orderRepository.UpdateOrder(o); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newOrderResponse(o))
}

// UpdateOrderTopaid func to update a new Order to paid
// @Summary Update a new Order to Paid
// @Description Update an Order To Paid. Auth is required
// @ID update-order-to-paid
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the order to update to paid"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} orderResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/orders/paid/{id} [put]
func (h *Handler) UpdateOrderTopaid(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	o, err := h.orderRepository.GetOrderById(id)
	if o == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Order"))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	// payload := &orderRequest{}
	// payload.populateOrder(o)
	// if err := bindOrderRequest(payload, c, o); err != nil {
	// 	return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	// }
	// errors := utils.ValidateStruct(payload)
	// if errors != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	// }

	if err = h.orderRepository.UpdateOrderTopaid(o); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newOrderResponse(o))
}

// UpdateOrderToDelivered func to update a new Order to delivered
// @Summary Update a new Order to Delivered
// @Description Update an Order To Delivered. Auth is required
// @ID update-order-to-delivered
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the order to update to delivered"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} orderResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/orders/delivered/{id} [put]
func (h *Handler) UpdateOrderToDelivered(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	o, err := h.orderRepository.GetOrderById(id)
	if o == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Order"))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	// payload := &orderRequest{}
	// payload.populateOrder(o)
	// if err := bindOrderRequest(payload, c, o); err != nil {
	// 	return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
	// }
	// errors := utils.ValidateStruct(payload)
	// if errors != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	// }

	if err = h.orderRepository.UpdateOrderToDelivered(o); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}

	return c.Status(http.StatusOK).JSON(newOrderResponse(o))
}

// DeleteOrder func to delete an Order
// @Summary Delete an Order
// @Description Delete an Order. Auth is required
// @ID delete-order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Id of the order to delete"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} orderResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		403				{string}	string	"Status Forbidden"
// @Failure		404				{string}	string	"Status Not Found"
// @Failure		409				{string}	string	"Status Conflict"
// @Failure		422				{string}	string	"Status UnprocessableEntity"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/orders/{id} [delete]
func (h *Handler) DeleteOrder(c *fiber.Ctx) error {

	id, err := utils.StringToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	o, err := h.orderRepository.GetOrderById(id)
	if o == nil && err == nil {
		return c.Status(http.StatusNotFound).JSON(utils.NotFound("Order"))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	err = h.orderRepository.DeleteOrder(o)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(map[string]interface{}{"result": "order deleted!"})
}

// AllOrders func to display all Orders
// @Summary Display all Orders
// @Description Display all Orders. Auth is required
// @ID all-orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit number of orders returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of orders (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} orderListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/orders/all/{limit}/{offset} [get]
func (h *Handler) AllOrders(c *fiber.Ctx) error {
	var (
		orders []models.Order
		count  int64
	)
	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
	}

	orders, count, err = h.orderRepository.AllOrders(offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newOrderListResponse(orders, count))
}

// AllUserOrders func to display all user Orders
// @Summary Display all user Orders
// @Description Display all user Orders. Auth is required
// @ID all-user-orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Param userId path string true "userId of the order to get"
// @Param limit query integer false "Limit number of orders returned (default is 20)"
// @Param offset query integer false "Offset/Skip number of orders (default is 0)"
// @Param	Authorization	header		string	true	"Authentication header"
// @Success   200       {object} orderListResponse
// @Failure		400				{string}	string	"Status BadRequest"
// @Failure		500				{string}	string	"Status Internal Server Error"
// @Failure		502				{string}	string	"Status BadGateway"
// @Security ApiKeyAuth
// @Router /api/v1/orders/{userId}/user/{limit}/{offset} [get]
func (h *Handler) AllUserOrders(c *fiber.Ctx) error {
	var (
		orders []models.Order
		count  int64
	)
	userId, err := utils.StringToInt(c.Params("userId"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		limit = 20
	}

	orders, count, err = h.orderRepository.GetUserOrders(userId, offset, limit)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(newOrderListResponse(orders, count))
}
