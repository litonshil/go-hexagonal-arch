// Package app implements the application layer
package app

import (
	"hexagonal-arch/internal/ports"
)

// Adapter implements the application layer
type Adapter struct {
	db ports.BookRepoPort
}

// This validates that Adapter implements the ports.BookServicePort interface
var _ ports.BookServicePort = Adapter{}

// NewApplication returns a new adapter with the given db port
func NewApplication(db ports.BookRepoPort) *Adapter {
	return &Adapter{
		db: db,
	}
}
