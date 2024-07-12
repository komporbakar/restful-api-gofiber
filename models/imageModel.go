package models

import "time"

type Image struct {
	Id          uint   `json:"id" gorm:"primaryKey;AutoIncrement"`
	Image       string `json:"image"`
	CategoryId  uint
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
