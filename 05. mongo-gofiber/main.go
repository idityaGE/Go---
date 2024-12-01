package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idityaGE/go-mongo-gofiber/controllers"
	"github.com/idityaGE/go-mongo-gofiber/routes"
	"go.uber.org/zap"
)

var (
	logger   *zap.Logger
	reqCount int
)

func main() {
	// Initialize zap logger
	initLogger()
	defer logger.Sync() // Ensure logs are flushed before program exits

	// Pass logger to controllers
	controllers.SetLogger(logger)

	// Initialize Fiber app
	app := fiber.New()

	// Define routes with grouped middleware
	api := app.Group("/api")
	v1 := api.Group("/v1", countRequestsMiddleware)
	routes.Handler(v1)

	logger.Info("Server is starting", zap.String("port", "3000"))
	if err := app.Listen(":3000"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

// initLogger initializes the zap logger
func initLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	logger.Info("Logger initialized")
}

// countRequestsMiddleware increments request count and sets version header
func countRequestsMiddleware(c *fiber.Ctx) error {
	reqCount++
	logger.Info("Request received", zap.Int("request_count", reqCount))
	c.Set("Version", "v1")
	return c.Next()
}
