package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// readinessProbeHandler it called by kubernetes deployment controller continuously
// in certain time interval to check whether DNS webserver is ready to receive the http request
func readinessProbeHandler(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

// livenessProbeHandler its called by kubernetes deployment controller
// continuously in certain time interval to check whether DNS webserver is alive or not
// if its failing then webserver will be restarted
func livenessProbeHandler(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

// ReadinessProbe add route for readiness probe
func ReadinessProbe(e *echo.Echo) {
	e.GET("/ready", readinessProbeHandler)
}

// LivenessProbe add route for liveness probe
func LivenessProbe(e *echo.Echo) {
	e.GET("/live", livenessProbeHandler)
}
