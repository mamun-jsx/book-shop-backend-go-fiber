package repository

import "gorm.io/gorm"

type BookRepository struct {
	db *gorm.DB
}

