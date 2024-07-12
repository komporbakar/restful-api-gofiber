package handler

import (
	"restful-api-gofiber/database"
	"restful-api-gofiber/models"
	"restful-api-gofiber/models/request"
	"restful-api-gofiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {

	body := new(request.Login)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Bad Request",
		})
	}

	validate := validator.New()
	loginValidate := validate.Struct(body)

	if loginValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": loginValidate.Error(),
		})
	}

	// CHECK EMAIL
	user := models.User{}
	isCheckEmail := database.DB.Where("email = ?", body.Email).First(&user).Error
	if isCheckEmail != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Credential Not Match",
		})
	}

	// CHECK PASSWORD
	isValid := utils.CheckPasswordHash(body.Password, user.Password)
	// isCheckPassword := database.DB.Where("password = ?", body.Password).First(&user).Error
	if !isValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Credential Not Match",
		})
	}

	//Generate Token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Login Success",
		"token":   token,
	})
}
