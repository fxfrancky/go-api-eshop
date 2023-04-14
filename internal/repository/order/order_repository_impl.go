package repository

import (
	"errors"

	"github.com/fxfrancky/go-api-eshop/internal/models"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

// AddOrderItems implements OrderRepository
func (o *OrderRepositoryImpl) AddOrderItems(order *models.Order, orderItems []models.OrderItem) error {
	for _, orderItem := range orderItems {
		err := o.DB.Model(order).Association("OrderItems").Append(orderItem)
		if err != nil {
			return err
		}
		err = o.DB.Where(orderItem.ID).Preload("Product").First(orderItem).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// AllOrders implements OrderRepository
func (o *OrderRepositoryImpl) AllOrders(offset int, limit int) ([]models.Order, int64, error) {
	var (
		orders []models.Order
		count  int64
	)
	o.DB.Model(&orders).Count(&count)
	o.DB.Offset(offset).Limit(limit).Find(&orders)

	return orders, count, nil
}

// CreateOrder implements OrderRepository
func (o *OrderRepositoryImpl) CreateOrder(order models.Order) error {
	// orderItems := order.OrderItems
	// order.OrderItems = make([]models.OrderItem, 0)
	// log.Println("***** Order Items ", orderItems)

	// shippingAddress := order.ShippingAddress
	// order.ShippingAddress = models.ShippingAddress{}
	// log.Println("***** Shipping address ", shippingAddress)

	// paymentResult := order.PaymentResult
	// order.PaymentResult = models.PaymentResult{}
	// log.Println("***** Payment Result ", paymentResult)

	// result := o.DB.Create(&order)
	// if result.Error != nil {
	// 	return result.Error
	// }

	// // Associates Order Items
	// for _, oItems := range orderItems {
	// 	o.DB.Model(order).Association("OrderItems").Append(oItems)
	// }

	// // Associates Shipping address
	// o.DB.Model(order).Association("ShippingAddress").Append(shippingAddress)

	// // Associates Payment Result
	// o.DB.Model(order).Association("PaymentResult").Append(paymentResult)

	tx := o.DB.Begin() // Begin the transaction
	defer tx.Commit()  // Commit the transaction After the execution

	// Create the shipping address
	// tx.Create(&order.ShippingAddress)

	// Create the payment result
	// tx.Create(&order.PaymentResult)

	// for _, &oItems := range order.OrderItems {
	// 	tx.Create(&oItems)
	// }
	// Create Order Items
	// tx.Create(&order.OrderItems)
	tx.Create(&order)

	// Associates Order Items
	// for _, oItems := range order.OrderItems {
	// 	o.DB.Model(order).Association("OrderItems").Append(oItems)
	// }
	// o.DB.Model(&order).Association("OrderItems").Append(&order.OrderItems)

	// // // Associates Shipping address
	// o.DB.Model(&order).Association("ShippingAddress").Append(&order.ShippingAddress)

	// // // Associates Payment Result
	// o.DB.Model(&order).Association("PaymentResult").Append(&order.PaymentResult)

	return nil
}

// GetOrderById implements OrderRepository
func (o *OrderRepositoryImpl) GetOrderById(orderId int) (*models.Order, error) {
	var order models.Order
	err := o.DB.First(&order, orderId).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &order, err
}

// GetUserOrders implements OrderRepository
func (o *OrderRepositoryImpl) GetUserOrders(userId int, offset int, limit int) ([]models.Order, int64, error) {
	var (
		orders []models.Order
		count  int64
	)
	o.DB.Model(&orders).Count(&count)
	o.DB.Offset(offset).Limit(limit).Preload("OrderItems").Preload("ShippingAddress").Preload("PaymentResult").Find(&orders, "user_id = ?", userId)
	return orders, count, nil
}

// UpdateOrderToDelivered implements OrderRepository
func (o *OrderRepositoryImpl) UpdateOrderToDelivered(order *models.Order) error {
	order.IsDelivered = true
	result := o.DB.Model(&order).Where("id = ?", order.ID).Updates(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateOrderTopaid implements OrderRepository
func (o *OrderRepositoryImpl) UpdateOrderTopaid(order *models.Order) error {
	order.IsPaid = true
	result := o.DB.Model(&order).Where("id = ?", order.ID).Updates(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateOrder implements OrderRepository
func (o *OrderRepositoryImpl) UpdateOrder(order *models.Order) error {
	result := o.DB.Model(&order).Where("id = ?", order.ID).Updates(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteOrder implements OrderRepository
func (o *OrderRepositoryImpl) DeleteOrder(order *models.Order) error {
	return o.DB.Delete(order).Error
}

// AddOrderItemToOrder implements OrderRepository
func (o *OrderRepositoryImpl) AddOrderItemToOrder(order *models.Order, orderItem *models.OrderItem) error {
	err := o.DB.Model(o).Association("OrderItems").Append(orderItem)
	if err != nil {
		return err
	}
	return o.DB.Where(orderItem.ID).Preload("Product").First(orderItem).Error
}

func NewOrderRepositoryImpl(DB *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{DB: DB}
}
