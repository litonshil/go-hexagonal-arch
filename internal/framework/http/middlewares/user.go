// Package middlewares contains all the middlewares used by the API
// nolint: wrapcheck
package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func authorizeUser(config config) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}

	//var binder echo.DefaultBinder

	// var binder echo.DefaultBinder
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			//var user types.User
			//if err := binder.BindHeaders(c, &user); err != nil {
			//	cErr := errors.InvalidRequestParsingError(err)
			//	return c.JSON(cErr.Status(), cErr)
			//}
			//
			//if !user.IsAdmin {
			//	cErr := errors.ForbiddenError("missing authorization headers")
			//	return c.JSON(cErr.Status(), cErr)
			//}
			//
			//c.Set("user", user)

			return next(c)
		}
	}
}
