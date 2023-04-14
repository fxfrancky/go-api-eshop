package handlers

import (
	orderRepository "github.com/fxfrancky/go-api-eshop/internal/repository/order"
	productRepository "github.com/fxfrancky/go-api-eshop/internal/repository/product"
	userRepository "github.com/fxfrancky/go-api-eshop/internal/repository/user"
)

type Handler struct {
	productRepository productRepository.ProductRepository
	orderRepository   orderRepository.OrderRepository
	userRepository    userRepository.UserRepository
}

func NewHandler(productRepo productRepository.ProductRepository, orderRepo orderRepository.OrderRepository, userRepo userRepository.UserRepository) *Handler {
	return &Handler{
		productRepository: productRepo,
		orderRepository:   orderRepo,
		userRepository:    userRepo,
	}
}
