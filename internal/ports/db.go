package ports

import (
	"hexagonal-arch/helper/errors"
	"hexagonal-arch/internal/ports/types"
)

// BookRepoPort is the interface for all Book DB ports
type BookRepoPort interface {
	// InsertBook creates a new book
	InsertBook(types.Book) (types.Book, errors.Error)
	// BeginTx starts a new transaction
	BeginTx() BookRepoPort
	// RollbackTx rolls back the transaction
	RollbackTx()
	// CommitTx commits the transaction
	CommitTx() errors.Error
}
