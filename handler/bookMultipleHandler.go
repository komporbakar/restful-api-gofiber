package handler

import (
	"log"
	"restful-api-gofiber/database"
	"restful-api-gofiber/models"
	"restful-api-gofiber/models/request"
	"restful-api-gofiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookMultipleHandlerCreate(ctx *fiber.Ctx) error {
	body := new(request.BookMultipleRequest)
	log.Println(body, "body")
	if err := ctx.BodyParser(body); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": 400,
			"error":  err,
		})
	}

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
	filenames := ctx.Locals("filenames")
	log.Println(filenames, "filenames")
	if filenames == nil {
		return ctx.Status(fiber.StatusUnavailableForLegalReasons).JSON(fiber.Map{
			"error":   true,
			"message": "Image Covers Required",
		})
	}

	for _, filename := range filenames.([]string) {
		filenameString := filename

		newBook := models.Image{
			CategoryId: body.CategoryId,
			Image:      filenameString,
		}

		errCreate := database.DB.Debug().Create(&newBook).Error
		if errCreate != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  401,
				"message": errCreate.Error(),
			})
		}
	}
	// filenameString := filenames.(string)

	// newBook := models.Image{
	// 	CategoryId: 1,
	// 	Image:      filenameString,
	// }

	// errCreate := database.DB.Debug().Create(&newBook).Error

	// if errCreate != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"status":  401,
	// 		"message": errCreate.Error(),
	// 	})
	// }

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Success",
	})

}

func ImageHandlerDelete(ctx *fiber.Ctx) error {
	getId := ctx.Params("id")
	var image models.Image

	errImage := database.DB.Debug().First(&image, "id = ?", getId).Error
	if errImage != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": errImage.Error(),
		})
	}
	log.Println(image, "image")

	errDelete := utils.HandleRemoveFile(image.Image)

	if errDelete != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": errDelete.Error(),
		})
	}

	err := database.DB.Debug().Delete(&image, "id = ?", getId).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   true,
		"message": "Success",
	})
}
