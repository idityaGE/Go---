package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/idityaGE/go-mongo-gofiber/routes"
)

var ReqCount = 0

func main() {
	app := fiber.New()
	api := app.Group("/api")         // /api
	v1 := api.Group("/v1", countReq) // /api/v1
	routes.Handler(v1)               // /api/v1/*

	fmt.Println("Server is running on port 3000")
	log.Fatal(app.Listen(":3000"))
}

func countReq(c *fiber.Ctx) error { // a simple middleware
	c.Set("Version", "v1")
	ReqCount += 1
	return c.Next()
}
