package tracing

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func nextRequestId() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func Add(e *echo.Echo) {
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return nextRequestId()
		},
	}))
}
