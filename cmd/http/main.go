package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fxfrancky/go-api-eshop/config"
	_ "github.com/fxfrancky/go-api-eshop/docs"
	"github.com/fxfrancky/go-api-eshop/internal/handlers"
	"github.com/fxfrancky/go-api-eshop/internal/initializers"
	orderRepo "github.com/fxfrancky/go-api-eshop/internal/repository/order"
	productRepo "github.com/fxfrancky/go-api-eshop/internal/repository/product"
	userRepo "github.com/fxfrancky/go-api-eshop/internal/repository/user"
	"github.com/fxfrancky/go-api-eshop/pkg/shutdown"
	"github.com/gofiber/swagger"
)

// @title GO API ESHOP
// @version 1.0
// @description GO API ESHOP Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name ESHOP API Support
// @contact.email contact@owonafx.com
// @license.name api-eshop 2.0
// @schemes http https
// @produce application/json
// @consumes application/json
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {

	// Configure Swagger API
	swagg := swagger.Config{
		Title:        "Swagger API",
		DeepLinking:  false,
		DocExpansion: "none",
	}

	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()
	rootPath := "."

	// Init All Databases
	initializers.LoadDatabases(rootPath)
	// load config
	env, err := config.LoadConfig(rootPath)
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// run the server
	cleanup, err := run(env, swagg)
	// run the cleanup after the server is terminated
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}
	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()

}

func run(env config.Config, swagg swagger.Config) (func(), error) {

	h := initApp(env)
	app := h.NewRoutes(&env, swagg)

	// start the server
	go func() {
		// create the fiber app
		log.Fatal(app.Listen(":8000"))
	}()

	// return a function to close the server and database
	return func() {
		app.Shutdown()
	}, nil
}

func initApp(env config.Config) *handlers.Handler {
	db := initializers.ConnectDB(&env)
	productRepo := productRepo.NewProductRepositoryImpl(db)
	orderRepo := orderRepo.NewOrderRepositoryImpl(db)
	userRepo := userRepo.NewUserRepositoryImpl(db)
	h := handlers.NewHandler(productRepo, orderRepo, userRepo)
	return h
}
