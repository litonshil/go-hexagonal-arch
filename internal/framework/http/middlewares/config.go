package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
)

type config struct {
	Skipper middleware.Skipper
}
