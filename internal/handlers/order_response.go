package handlers

import (
	"time"

	"github.com/fxfrancky/go-api-eshop/internal/models"
)

// Order Response
type orderResponse struct {
	ID                uint               `json:"id,omitempty"`
	UserID            uint               `json:"user_id,omitempty"`
	OrderItems        []models.OrderItem `json:"orderItemList,omitempty"`
	ShippingAddressID uint               `json:"shipping_address_id,omitempty"`
	PaymentMethod     string             `json:"payment_method,omitempty"`
	PaymentResult     uint               `json:"payment_result_id,omitempty"`
	TaxPrice          float64            `json:"tax_price,omitempty"`
	ShippingPrice     float64            `json:"shipping_price,omitempty"`
	TotalPrice        float64            `json:"total_price,omitempty"`
	IsPaid            bool               `json:"is_paid,omitempty"`
	PaidAt            time.Time          `json:"paid_at,omitempty"`
	IsDelivered       bool               `json:"is_delivered,omitempty"`
	DeliveredAt       time.Time          `json:"delivered_at"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	DeletedAt         time.Time          `json:"deleted_at"`
}

func newOrderResponse(o *models.Order) *orderResponse {

	orderResp := new(orderResponse)
	orderResp.ID = o.ID
	orderResp.UserID = o.UserID
	orderResp.ShippingAddressID = o.ShippingAddress.ID
	orderResp.PaymentMethod = o.PaymentMethod
	orderResp.PaymentResult = o.PaymentResultID
	orderResp.TaxPrice = o.TaxPrice
	orderResp.ShippingPrice = o.ShippingPrice
	orderResp.TotalPrice = o.TotalPrice
	orderResp.IsPaid = o.IsPaid
	orderResp.PaidAt = o.PaidAt
	orderResp.IsDelivered = o.IsDelivered
	orderResp.DeliveredAt = o.DeliveredAt
	orderResp.CreatedAt = o.CreatedAt
	orderResp.DeletedAt = o.DeliveredAt
	orderResp.UpdatedAt = o.UpdatedAt

	return orderResp
}

type orderListResponse struct {
	Orders      []*orderResponse `json:"orders"`
	OrdersCount int64            `json:"ordersCount"`
}

func newOrderListResponse(orders []models.Order, count int64) *orderListResponse {
	r := new(orderListResponse)
	r.Orders = make([]*orderResponse, 0)
	for _, o := range orders {
		or := newOrderResponse(&o)
		r.Orders = append(r.Orders, or)
	}
	r.OrdersCount = count
	return r
}
