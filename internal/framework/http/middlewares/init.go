package middlewares

import (
	"github.com/labstack/echo/v4"
	"strings"

	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	metricsPath := "/metrics"

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `${time_custom} ${remote_ip} ${host} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out} "${user_agent}"` + "\n",
		CustomTimeFormat: "2006-01-02T15:04:05.00",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(context echo.Context) bool {
			return context.Request().RequestURI == metricsPath
		},
		Level: 5,
	}))

	skipList := []string{
		"/swagger",
		"/metrics",
		"/health",
		"/api/v1/internal",
	}

	e.Use(authorizeUser(config{
		Skipper: func(c echo.Context) bool {
			reqURI := c.Request().RequestURI

			for _, skip := range skipList {
				if strings.HasPrefix(reqURI, skip) {
					return true
				}
			}
			return false
		},
	}))

}
