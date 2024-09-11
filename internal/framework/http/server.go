// Package http implements HTTP server
// nolint: wrapcheck
package http

import (
	"hexagonal-arch/internal/framework/http/middlewares"
	"hexagonal-arch/internal/ports"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Adapter implements HTTP interface
type Adapter struct {
	// Echo is the HTTP server.
	echo *echo.Echo
	// DefaultBinder is used to bind HTTP headers, query & path parameters into a struct.
	binder echo.DefaultBinder
	// BookServicePort is the port to the application's business logic.
	api ports.BookServicePort
}

// NewAdapter creates a new Adapter struct and returns a pointer to it.
func NewAdapter(api ports.BookServicePort) *Adapter {
	return &Adapter{
		echo: echo.New(),
		api:  api,
	}
}

func (a Adapter) Run() {
	middlewares.Init(a.echo)

	a.registerAPI()

	a.echo.GET("/swagger/*", echoSwagger.WrapHandler)
	a.echo.Logger.Fatal(a.echo.Start(":7766"))
}

func (a Adapter) registerAPI() {
	group := a.echo.Group("/api/v1")
	group.POST("/book", a.createBook)
}
