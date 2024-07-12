package main

import (
	"restful-api-gofiber/database"
	routes "restful-api-gofiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()

	// RUN MIGRATION
	database.AutoMigrate()

	app := fiber.New()
	app.Use(logger.New())

	// INITIAL ROUTE
	routes.Setup(app)

	app.Listen(":4000")
}
