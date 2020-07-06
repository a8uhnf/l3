package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Add(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "requestId=${id}, remoteIP=${remote_ip}, host=${host}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
}
