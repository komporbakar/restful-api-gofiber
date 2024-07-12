package routes

import (
	"restful-api-gofiber/config"
	"restful-api-gofiber/handler"
	"restful-api-gofiber/middleware"
	"restful-api-gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

func Setup(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	api := r.Group("/api")
	v1 := api.Group("/v1", middleware.Auth)

	//ROUTE AUTH
	api.Post("/login", handler.LoginHandler)

	// ROUTE USER
	v1.Get("/user", handler.UserHandlerGetAll)
	v1.Get("/user/:id", handler.UserHandlerGetById)
	v1.Post("/user", handler.UserHandlerCreate)
	v1.Put("/user/:id", handler.UserHandlerUpdate)
	v1.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	v1.Delete("/user/:id", handler.UserHandlerDelete)

	//Route Book
	v1.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)
	v1.Post("/book-multiple", utils.HandleMultipleFile, handler.BookMultipleHandlerCreate)
	v1.Delete("/image/:id", handler.ImageHandlerDelete)
}
