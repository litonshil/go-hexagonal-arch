package db

import (
	"hexagonal-arch/helper/errors"
	"hexagonal-arch/internal/ports/types"
)

// InsertBook creates a new book
func (a Adapter) InsertBook(book types.Book) (types.Book, errors.Error) {
	if err := a.db.Create(&book).Error; err != nil {
		return types.Book{}, errors.InternalDBError(err)
	}

	return book, nil
}
