package main

import (
	order_services "back-usm/internals/order/core/services"
	product_services "back-usm/internals/product/core/services"

	order_handlers "back-usm/internals/order/handlers"
	product_handlers "back-usm/internals/product/handlers"

	order_repository "back-usm/internals/order/repository"
	product_repository "back-usm/internals/product/repository"

	server "back-usm/cmd/server"
	"back-usm/utils"
)

func main() {
	// DSNs databases
	productsDB := utils.GetEnvVar("MYSQL_PRODUCTS_DSN")
	ordersDB := utils.GetEnvVar("MYSQL_ORDERS_DSN")

	// Repositories
	productRepository := product_repository.NewProductRepository(productsDB)
	orderRepository := order_repository.NewOrderRepository(ordersDB)

	// Services
	productService := product_services.NewProductServices(productRepository)
	orderService := order_services.NewOrderServices(orderRepository)

	// Handlers
	productHandlers := product_handlers.NewProductHandlers(productService)
	orderHandlers := order_handlers.NewOrderHandlers(orderService)

	// Server
	server := server.NewServer(productHandlers, orderHandlers)

	// Init
	server.Start()
}
