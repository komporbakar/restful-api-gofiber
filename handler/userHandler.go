package handler

import (
	"fmt"
	"log"
	"restful-api-gofiber/database"
	"restful-api-gofiber/models"
	"restful-api-gofiber/models/request"
	"restful-api-gofiber/models/response"
	"restful-api-gofiber/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user")
	fmt.Println(userId)
	var users []models.User
	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println("error", result.Error)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": 200,
		"data":   users,
	})
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	body := new(request.UserCreateRequest)

	if err := ctx.BodyParser(body); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": 400,
			"error":  err,
		})
	}

	// VALIDATE REQUEST

	validate := validator.New()
	errValidate := validate.Struct(body)

	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": errValidate.Error(),
		})
	}

	hashPassword, _ := utils.HashingPassword(body.Password)

	newUser := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hashPassword,
		Phone:    body.Phone,
		Role:     "user",
		Status:   "active",
	}

	errCreate := database.DB.Create(&newUser).Error
	if errCreate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": errCreate.Error(),
		})
	}

	res := response.UserResponse{
		Id:        newUser.Id,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
		Role:      newUser.Role,
		Status:    newUser.Status,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"data":  res,
	})

}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var user models.User
	err := database.DB.First(&user, "id = ?", id).Error

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	data := response.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	userResponse := response.WebResponse{
		Error:   false,
		Message: "Success",
		Data:    data,
	}

	return ctx.Status(fiber.StatusOK).JSON(userResponse)
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {

	body := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Bad Request	",
		})
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	var user models.User
	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	// VALIDATE REQUEST
	validate := validator.New()
	errValidate := validate.Struct(body)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": errValidate.Error(),
		})
	}

	// UPDATE USER
	if body.Name != "" {
		user.Name = body.Name
	}
	if body.Phone != "" {
		user.Phone = body.Phone
	}
	if body.Password != "" {
		user.Password = body.Password
	}

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": errUpdate,
		})
	}

	data := response.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	res := response.WebResponse{
		Error:   false,
		Message: "Success Update User",
		Data:    response.UserResponse(data),
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {

	body := new(request.EmailUpdateRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Bad Request	",
		})
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	var user models.User
	var isCheckEmail models.User
	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	errCheckEmail := database.DB.Where("email = ?", body.Email).First(&isCheckEmail).Error
	if errCheckEmail == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Email already exist",
		})
	}

	// VALIDATE REQUEST
	validate := validator.New()
	errValidate := validate.Struct(body)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": errValidate.Error(),
		})
	}

	// UPDATE USER
	user.Email = body.Email

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": errUpdate,
		})
	}

	data := response.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	res := response.WebResponse{
		Error:   false,
		Message: "Success Update Email",
		Data:    response.UserResponse(data),
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func UserHandlerDelete(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	var user models.User

	err := database.DB.Debug().First(&user, "id = ?", id).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user, "id = ?", id).Error
	if errDelete != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": errDelete,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Success Delete User",
	})
}
