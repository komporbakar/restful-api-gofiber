package database

import (
	"fmt"
	"restful-api-gofiber/models"
)

func AutoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Book{},
		&models.Image{},
	)

	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("success")
}
