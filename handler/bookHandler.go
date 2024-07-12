package handler

import (
	"log"
	"restful-api-gofiber/database"
	"restful-api-gofiber/models"
	"restful-api-gofiber/models/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {
	body := new(request.BookCreateRequest)
	if err := ctx.BodyParser(body); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": 400,
			"error":  err,
		})
	}
	log.Println(body)

	//Validasi Request

	validate := validator.New()
	errValidate := validate.Struct(body)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": errValidate.Error(),
		})
	}

	//Validation Required Image
	filename := ctx.Locals("filename")
	if filename == nil {
		return ctx.Status(fiber.StatusUnavailableForLegalReasons).JSON(fiber.Map{
			"error":   true,
			"message": "Image Cover Required",
		})
	}
	filenameString := filename.(string)

	newBook := models.Book{
		Title:  body.Title,
		Author: body.Author,
		Cover:  filenameString,
	}

	errCreate := database.DB.Debug().Create(&newBook).Error

	if errCreate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": errCreate.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"message": "Success",
		"data":    newBook,
	})
}

func BookHandlerGetAll(ctx *fiber.Ctx) error {
	return ctx.Next()
}
