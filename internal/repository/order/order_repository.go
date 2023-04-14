package repository

import "github.com/fxfrancky/go-api-eshop/internal/models"

type OrderRepository interface {
	// Users
	CreateOrder(order models.Order) error
	AddOrderItems(order *models.Order, orderItems []models.OrderItem) error
	GetOrderById(orderId int) (*models.Order, error)
	UpdateOrderTopaid(order *models.Order) error
	UpdateOrder(order *models.Order) error
	UpdateOrderToDelivered(order *models.Order) error
	GetUserOrders(userId, offset, limit int) ([]models.Order, int64, error)
	AllOrders(offset, limit int) ([]models.Order, int64, error)
	DeleteOrder(order *models.Order) error
	AddOrderItemToOrder(order *models.Order, orderItem *models.OrderItem) error
}
