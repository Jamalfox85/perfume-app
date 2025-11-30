package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jamalfox85/perfume-app/backend/api"
	"github.com/jamalfox85/perfume-app/backend/api/middleware"
)

func main() {
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	// Disable Extra Gin Logs
	gin.SetMode(gin.ReleaseMode)

	// Load Supabase JWKS
	middleware.InitJWKS()

	// Create app instance
	app := api.NewApplication()

	cfg := newConfig()

	srv := &http.Server{
		Addr:	net.JoinHostPort("", cfg.port()),
		Handler: app.Routes(),
	}
	srvErr := make(chan error, 1)
	go func() {
		slog.Info("Server Running on port " + cfg.port())
		srvErr <- srv.ListenAndServe()
	}()

	select {
		case err = <-srvErr:
			if err != nil && err != http.ErrServerClosed {
				slog.Error("Server error: " + err.Error())
			}
			return
		case <-ctx.Done():
			slog.Info("Shutting down server...")
			stop()
	}

	err = srv.Shutdown(context.Background())
	return
}