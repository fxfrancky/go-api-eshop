package handlers

import (
	"time"

	"github.com/fxfrancky/go-api-eshop/internal/models"
	"github.com/gofiber/fiber/v2"
)

type shippingAddressRequest struct {
	Address    string `json:"address" validate:"required"`
	City       string `json:"city" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	Country    string `json:"country" validate:"required"`
}

type orderItemRequest struct {
	Name      string  `json:"name" validate:"required"`
	Quantity  int64   `json:"quantity" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	Image     string  `json:"image" validate:"required"`
	ProductID uint    `json:"product_id" validate:"required"`
}

type paymentResultRequest struct {
	PaymentResultID string `json:"payment_result_id" validate:"required"`
	Status          string `json:"status" validate:"required"`
	UpdateTime      string `json:"update_time" validate:"required"`
	EmailAddress    string `json:"email_address" validate:"required"`
}

// Order Requests
type orderRequest struct {
	UserID        uint               `json:"user_id" validate:"required"`
	OrderItems    []orderItemRequest `json:"orderItemList"`
	PaymentMethod string             `json:"payment_method" validate:"required"`
	// PaymentResultID   uint
	PaymentResult paymentResultRequest `json:"payment_result"`
	TaxPrice      float64              `json:"tax_price" validate:"required"`
	// ShippingAddressID uint
	ShippingAddress shippingAddressRequest
	ShippingPrice   float64   `json:"shipping_price" validate:"required"`
	TotalPrice      float64   `json:"total_price" validate:"required"`
	IsPaid          bool      `json:"is_paid"`
	PaidAt          time.Time `json:"paid_at"`
	IsDelivered     bool      `json:"is_delivered" default:"false"`
	DeliveredAt     time.Time `json:"delivered_at"`
}

// Bind Order
func bindOrderRequest(r *orderRequest, c *fiber.Ctx, o *models.Order) error {
	// Validate the order
	if err := c.BodyParser(r); err != nil {
		return err
	}
	// Map the order
	o.UserID = r.UserID
	// Binding Payment Result
	o.PaymentResult.PaymentResID = r.PaymentResult.PaymentResultID
	o.PaymentResult.Status = r.PaymentResult.Status
	o.PaymentResult.EmailAddress = r.PaymentResult.EmailAddress
	o.PaymentResult.UpdateTime = r.PaymentResult.UpdateTime

	o.PaymentMethod = r.PaymentMethod

	o.TaxPrice = r.TaxPrice
	// Binding Shipping address
	o.ShippingAddress.Address = r.ShippingAddress.Address
	o.ShippingAddress.PostalCode = r.ShippingAddress.PostalCode
	o.ShippingAddress.City = r.ShippingAddress.City
	o.ShippingAddress.Country = r.ShippingAddress.Country
	// o.ShippingAddressID = r.ShippingAddressID

	// Binding Order Items
	for _, item := range r.OrderItems {
		var orderItemMod models.OrderItem
		orderItemMod.Name = item.Name
		// orderItemMod.OrderID = item.OrderID
		orderItemMod.ProductID = item.ProductID
		orderItemMod.Price = item.Price
		orderItemMod.Quantity = item.Quantity
		orderItemMod.Image = item.Image
		o.OrderItems = append(o.OrderItems, orderItemMod)
	}
	o.ShippingPrice = r.ShippingPrice
	o.TotalPrice = r.TotalPrice
	o.IsPaid = r.IsPaid
	o.PaidAt = r.PaidAt
	o.IsDelivered = r.IsDelivered
	o.DeliveredAt = r.DeliveredAt

	return nil
}

func (o *orderRequest) populateOrder(order *models.Order) {
	o.UserID = order.UserID
	o.PaymentMethod = order.PaymentMethod
	// o.PaymentResultID = order.PaymentResultID
	o.TaxPrice = order.TaxPrice
	// o.ShippingAddressID = order.ShippingAddressID
	// o.ShippingAddress.OrderID = order.ShippingAddress.OrderID
	// o.ShippingAddress.Address = order.ShippingAddress.Address
	// o.ShippingAddress.PostalCode = order.ShippingAddress.PostalCode
	// o.ShippingAddress.City = order.ShippingAddress.City
	// o.ShippingAddress.Country = order.ShippingAddress.Country

	// o.ShippingAddress.OrderID = r.ShippingAddress.OrderID
	// o.ShippingAddress.Address = r.ShippingAddress.Address
	// o.ShippingAddress.PostalCode = r.ShippingAddress.PostalCode
	// o.ShippingAddress.City = r.ShippingAddress.City
	o.ShippingPrice = order.ShippingPrice
	o.TotalPrice = order.TotalPrice
	o.IsPaid = order.IsPaid
	o.PaidAt = order.PaidAt
	o.IsDelivered = order.IsDelivered
	o.DeliveredAt = order.DeliveredAt
}

// Bind OrderItems Requests
func bindOrderItemsRequests(orderItemsReq []orderItemRequest, orderItemsModels []models.OrderItem, orderId uint) ([]models.OrderItem, error) {

	// Map the order Items
	// if len(orderItems) > 0 {

	// 	for
	// }

	for _, oItem := range orderItemsReq {
		var orderItemsModel models.OrderItem
		orderItemsModel.Name = oItem.Name
		orderItemsModel.Quantity = oItem.Quantity
		orderItemsModel.Price = oItem.Price
		orderItemsModel.Image = oItem.Image
		// orderItemsModel.OrderID = orderId
		orderItemsModel.ProductID = oItem.ProductID
		orderItemsModels = append(orderItemsModels, orderItemsModel)

	}
	return orderItemsModels, nil
	// o.UserID = r.UserID
	// o.PaymentMethod = r.PaymentMethod
	// o.PaymentResultID = r.PaymentResultID
	// o.TaxPrice = r.TaxPrice

	// 	Name      string  `json:"name" validate:"required"`
	// Quantity  int64   `json:"quantity" validate:"required"`
	// Price     float64 `json:"price" validate:"required"`
	// ProductID uint    `json:"product_id" validate:"required"`
	// OrderID   uint    `json:"order_id" validate:"required"`

}
