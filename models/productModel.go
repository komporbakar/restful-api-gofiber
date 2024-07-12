package models

import "time"

type Product struct {
	Id uint `json:"id" gorm:"primaryKey;AutoIncrement"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price int `json:"price"`
	Stock int `json:"stock"`
	CategoryId int `json:"category_id"`
	Category Category `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}