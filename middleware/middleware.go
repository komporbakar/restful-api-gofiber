package middleware

import (
	"restful-api-gofiber/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	authorization := ctx.Get("x-token")
	ip := ctx.IP()

	//check ip
	if !utils.CheckIp(ip) {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "Forbidden",
		})
	}

	//check token
	if authorization == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Unauthorized",
		})
	}

	// _, err := utils.VerifyJWT(authorization)
	claims, err := utils.DecodeToken(authorization)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "Unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "Forbidden yakk",
		})
	}

	ctx.Locals("user", claims["id"])

	return ctx.Next()
}

func PermissionCreate(ctx *fiber.Ctx) error {

	return ctx.Next()
}
