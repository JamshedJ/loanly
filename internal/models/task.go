package models

import "time"

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"not null"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Done        *bool     `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
