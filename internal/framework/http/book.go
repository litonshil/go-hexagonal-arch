// Package http implements HTTP server
// nolint: wrapcheck
package http

import (
	"hexagonal-arch/helper/errors"
	"hexagonal-arch/internal/ports/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary	Creates a book
// @Description	Creates a book with its Bindings
// @Tags	book
// @Accept	json
// @Produce	json
// @Param	book	body	types.BookReqBody	true	"Request body of Book"
// @Success	200	{object}	types.BookResp
// @Failure	400	{object}	errors.CustomError
// @Router	/api/v1/book [POST]
func (a Adapter) createBook(c echo.Context) error {
	var request types.BookReqBody
	if err := c.Bind(&request); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSON(cErr.Status(), cErr)
	}

	// Validate the request body
	//if cErr := types.Validate(&request, c.Request().Method); cErr != nil {
	//	return c.JSON(cErr.Status(), cErr)
	//}

	bookResp, cErr := a.api.CreateBook(request)
	if cErr != nil {
		return c.JSON(cErr.Status(), cErr)
	}

	return c.JSON(http.StatusOK, bookResp)
}
