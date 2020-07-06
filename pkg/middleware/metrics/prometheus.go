package metrics

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"

	metricspkg "github.com/a8uhnf/l3/pkg/metrics"
)

type Stats struct {
	prometheusHistogram *prometheus.HistogramVec
}

func New(name string) *Stats {
	return &Stats{
		prometheusHistogram: metricspkg.RegisterMetrics(name),
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := time.Now()
		if err := next(c); err != nil {
			c.Error(err)
		}
		method := c.Request().Method
		elapsedTime := time.Now().Sub(t).Nanoseconds()
		statusCode := c.Response().Status
		url := c.Request().URL
		fmt.Println(elapsedTime, statusCode)
		s.prometheusHistogram.WithLabelValues(method, url.String()).Observe(float64(elapsedTime))
		return nil
	}
}

// Add prometheus middleware
func (s *Stats) Add(e *echo.Echo) {
	prometheus.MustRegister(s.prometheusHistogram)
	e.Use(s.Process)
}
