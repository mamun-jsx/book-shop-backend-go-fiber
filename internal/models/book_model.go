package models

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"`
	AuthorName       string    `gorm:"type:varchar(255);not null" json:"author_name"`
	Pages            int       `gorm:"type:integer;not null" json:"pages"`
	Type             string    `gorm:"type:varchar(100);not null" json:"type"`
	Price            float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	ShortDescription string    `gorm:"type:text" json:"short_description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
