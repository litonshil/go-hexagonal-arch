package types

import (
	"time"

	"gorm.io/gorm"
)

// Book represents a book in the database.
type Book struct {
	// ID is the primary key of the book.
	ID uint `gorm:"primaryKey"`
	// Code is the book code.
	Code string `gorm:"type:varchar(32);primaryKey"`
	// Title is the title of the book.
	Title string `gorm:"type:varchar(128)"`
	// Description is the description of the book.
	// +optional
	Description *string `gorm:"type:varchar(128)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (*Book) TableName() string {
	return "book"
}
