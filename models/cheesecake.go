package models

import (
	"time"
)

type CheeseCake struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}
