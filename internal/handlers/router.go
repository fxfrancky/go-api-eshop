package handlers

import (
	"context"
	"fmt"

	"github.com/fxfrancky/go-api-eshop/config"
	"github.com/fxfrancky/go-api-eshop/internal/initializers"
	"github.com/fxfrancky/go-api-eshop/internal/middleware"
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func (h *Handler) NewRoutes(config *config.Config, swagg swagger.Config) *fiber.App {

	apiPath := "/api/" + config.APIVersion
	app := fiber.New()

	app = middleware.EnableCors(app)

	// Main /api/v1 route Group
	v1 := app.Group(apiPath)
	ctx := context.TODO()
	value, err := initializers.RedisClient.Get(ctx, "test").Result()

	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		panic(err)
	}
	// Check that the api is healthy /api/v1/healthchecker
	v1.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": value,
		})
	})
	// add docs /api/v1/swagger
	v1.Get("/swagger/*", swagger.HandlerDefault)
	v1.Get("/swagger/*", swagger.New(swagg))

	// Authentification and user management
	grpUsers := v1.Group("/auth")                  // /api/v1/auth
	grpUsers.Post("/register", h.SignUpUser)       // /api/v1/auth/register
	grpUsers.Post("/login", h.SignInUser)          // /api/v1/auth/login
	grpUsers.Get("/refresh", h.RefreshAccessToken) // /api/v1/auth/refresh
	// Protected users routes
	grpUsers.Get("/logout", middleware.DeserializeUser, h.LogoutUser) // /api/v1/auth/logout
	//grpUsers.Get("/users/me", middleware.DeserializeUser, h.GetMe)    // /api/v1/auth/users/me

	// user management secure
	usersGrp := v1.Group("/users")                                                // /api/v1/users
	usersGrp.Get("/me", middleware.DeserializeUser, h.GetMe)                      // /api/v1/users/me
	usersGrp.Put("/:id", middleware.DeserializeUser, h.UpdateUser)                // /api/v1/users/:id [put]
	usersGrp.Delete("/:id", middleware.DeserializeUser, h.DeleteUser)             // /api/v1/users/:id [delete]
	usersGrp.Get("/:email", middleware.DeserializeUser, h.GetUserByEmail)         // /api/v1/users/:email [get]
	usersGrp.Get("/all/:limit?/:offset?", middleware.DeserializeUser, h.AllUsers) // /api/v1/users/top/:limit?/:offset? [get]

	// products management protected
	products := v1.Group("/products")                                                // /api/v1/products
	products.Post("", middleware.DeserializeUser, h.CreateProduct)                   // /api/v1/products [post]
	products.Put("/:id", middleware.DeserializeUser, h.UpdateProduct)                // /api/v1/products/:id [put]
	products.Delete("/:id", middleware.DeserializeUser, h.DeleteProduct)             // /api/v1/products/:id [delete]
	products.Get("/:id", middleware.DeserializeUser, h.GetProduct)                   // /api/v1/products/:id [get]
	products.Get("/all/:limit?/:offset?", middleware.DeserializeUser, h.AllProducts) // /api/v1/products/all/:limit?/:offset? [get]
	products.Get("/top/:limit?", middleware.DeserializeUser, h.TopProducts)          // /api/v1/products/top/:limit?/:offset? [get]

	// products review protected
	reviews := v1.Group("/reviews")
	reviews.Post("/:product_id", middleware.DeserializeUser, h.AddReviewToProduct)                         // /api/v1/reviews [post]
	reviews.Put("/:id", middleware.DeserializeUser, h.UpdateReview)                                        // /api/v1/reviews/:id [put]
	reviews.Delete("/:id", middleware.DeserializeUser, h.DeleteReview)                                     // /api/v1/reviews/:id [delete]
	reviews.Get("/:id", middleware.DeserializeUser, h.GetReview)                                           // /api/v1/reviews/:id [get]
	reviews.Get("/:productName/product/:limit?/:offset?", middleware.DeserializeUser, h.AllProductsReview) // /api/v1/reviews/{productName}/product/:limit?/:offset? [get]

	// orders management protected
	orders := v1.Group("/orders")                                                             // /api/v1/orders
	orders.Post("", middleware.DeserializeUser, h.CreateOrder)                                // /api/v1/orders [post]
	orders.Put("/:id", middleware.DeserializeUser, h.UpdateOrder)                             // /api/v1/orders/:id [put]
	orders.Put("/paid/:id", middleware.DeserializeUser, h.UpdateOrderTopaid)                  // /api/v1/orders/paid/:id [put]
	orders.Put("/delivered/:id", middleware.DeserializeUser, h.UpdateOrderToDelivered)        // /api/v1/orders/delivered/:id [put]
	orders.Delete("/:id", middleware.DeserializeUser, h.DeleteOrder)                          // /api/v1/orders/:id [delete]
	orders.Get("/:id", middleware.DeserializeUser, h.GetOrderById)                            // /api/v1/orders/:id [get]
	orders.Get("/all/:limit?/:offset?", middleware.DeserializeUser, h.AllOrders)              // /api/v1/orders/all/:limit?/:offset? [get]
	orders.Get("/:userId/user/:limit?/:offset?", middleware.DeserializeUser, h.AllUserOrders) // /api/v1/orders/:userId/user/:limit?/:offset? [get]

	// Close Other Routes
	v1.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	// 	router.Post("/products", middleware.DeserializeUser, h.CreateProduct)
	// 	router.Put("/products", middleware.DeserializeUser, h.UpdateProduct)

	// 	router.Post("/register", h.SignUpUser)
	// 	router.Post("/login", h.SignInUser)
	// 	router.Get("/logout", middleware.DeserializeUser, h.LogoutUser)
	// 	router.Get("/refresh", h.RefreshAccessToken)

	// app := fiber.New()
	// micro := fiber.New()
	// app.Use(recover.New())
	// app.Mount(apiPath, micro)
	// app.Use(logger.New())

	// // Enable Cors
	// app = middleware.EnableCors(app)
	// // add docs
	// app.Get("/swagger/*", swagger.HandlerDefault)
	// app.Get("/swagger/*", swagger.New(swagg))
	// // micro.Use(recover.New())
	// // micro.Use(csrf.New())
	// // /api/auth/register
	// micro.Route("/auth", func(router fiber.Router) {
	// 	// User routes
	// 	router.Post("/register", h.SignUpUser)
	// 	router.Post("/login", h.SignInUser)
	// 	router.Get("/logout", middleware.DeserializeUser, h.LogoutUser)
	// 	router.Get("/refresh", h.RefreshAccessToken)
	// 	// Product routes
	// 	router.Post("/products", middleware.DeserializeUser, h.CreateProduct)
	// 	router.Put("/products", middleware.DeserializeUser, h.UpdateProduct)
	// })

	// micro.Get("/users/me", middleware.DeserializeUser, h.GetMe)

	// ctx := context.TODO()
	// value, err := initializers.RedisClient.Get(ctx, "test").Result()

	// if err == redis.Nil {
	// 	fmt.Println("key: test does not exist")
	// } else if err != nil {
	// 	panic(err)
	// }

	// micro.Get("/healthchecker", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 		"status":  "success",
	// 		"message": value,
	// 	})
	// })

	// micro.All("*", func(c *fiber.Ctx) error {
	// 	path := c.Path()
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"status":  "fail",
	// 		"message": fmt.Sprintf("Path: %v does not exists on this server", path),
	// 	})
	// })

	return app

}
