// Package types holds all the types includes API, database, contracts
package types

import (
	"time"
)

// BookReqBody is the API model for book
type BookReqBody struct {
	// Code is the identification code.
	// +required
	Code string `json:"code"`
	// Title is the book title
	// +required
	Title string `json:"title"`
	// Description is the book description
	// +optional
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BookResp is the API model for the book response
type BookResp struct {
	ID uint `json:"id"`
	// BookReqBody is the API model for book
	// +embedded
	BookReqBody
}
