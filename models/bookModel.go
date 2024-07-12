package models

import "time"

type Book struct {
	Id        uint      `json:"id" gorm:"primaryKey;AutoIncrement"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Cover     string    `json:"cover"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-" gorm:"index,column:deleted_at"`
}
