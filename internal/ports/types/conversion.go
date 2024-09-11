package types

// ToBook converts the API model BookBody to DB model Book
func (b BookReqBody) ToBook() Book {
	book := Book{
		Code:        b.Code,
		Title:       b.Title,
		Description: b.Description,
	}

	return book
}

// ToBookResp converts the Book to BookResp
// This is used to convert the DB model to the response model
func (p *Book) ToBookResp() BookResp {
	resp := BookResp{
		ID: p.ID,
		BookReqBody: BookReqBody{
			Code:        p.Code,
			Description: p.Description,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		},
	}

	return resp
}
