package models

import "time"

type Category struct {
	Id          uint      `json:"id" gorm:"primaryKey;AutoIncrement"`
	Name        string    `json:"name"`
	Images      []Image   `json:"images"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
