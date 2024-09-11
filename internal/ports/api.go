// Package ports declares all PORT interface
package ports

import (
	"hexagonal-arch/helper/errors"
	"hexagonal-arch/internal/ports/types"
)

// BookServicePort is the interface for all Book API ports
type BookServicePort interface {
	// CreateBook creates a new book
	CreateBook(body types.BookReqBody) (types.BookResp, errors.Error)
}
