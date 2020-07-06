package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"

	"github.com/a8uhnf/l3/pkg/middleware/logger"
	"github.com/a8uhnf/l3/pkg/middleware/metrics"
	"github.com/a8uhnf/l3/pkg/middleware/tracing"
	"github.com/a8uhnf/l3/route"
)

type Server struct {
	Port          string
	SectorID      float64
	ServerRunning chan bool
}

// New initialize location web server
func New(port string, sectorID float64) *Server {
	return &Server{
		Port:          port,
		SectorID:      sectorID,
		ServerRunning: make(chan bool),
	}
}

// Spawn spawn location web server
func (s *Server) Spawn() error {
	e := echo.New()

	// tracing middleware
	tracing.Add(e)
	// logger middleware
	logger.Add(e)
	// prometheus middleware
	stat := metrics.New("l3")
	stat.Add(e)

	// Route for drone navigation system
	route.AddDNS(e, s.SectorID)
	// Add route for prometheus metrics
	route.Prometheus(e)
	// Route for liveness probe
	route.LivenessProbe(e)
	// Route for readiness probe
	route.ReadinessProbe(e)

	go func() {
		<-s.ServerRunning
		e.Close()
	}()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", s.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Error(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	return nil
}
