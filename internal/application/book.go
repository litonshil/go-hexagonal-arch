package app

import (
	"hexagonal-arch/helper/errors"
	"hexagonal-arch/internal/ports/types"
)

// CreateBook creates a new book
func (a Adapter) CreateBook(bookReq types.BookReqBody) (types.BookResp, errors.Error) {
	// Converting request body to DB model
	book := bookReq.ToBook()
	book, err := a.db.InsertBook(book)
	if err != nil {
		return types.BookResp{}, err
	}

	// Converting DB model to response body
	return book.ToBookResp(), nil
}
