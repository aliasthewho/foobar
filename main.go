package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"foobar/config"
	"foobar/internal/handlers"
	"foobar/pkg/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Initialize configuration
	cfg := config.DefaultConfig()

	// Initialize logger
	logLevel := zapcore.InfoLevel // default
	if err := logLevel.Set(cfg.Log.Level); err != nil {
		logLevel = zapcore.InfoLevel // fallback to info level
	}

	if err := logger.Initialize(logLevel); err != nil {
		panic("failed to initialize logger: " + err.Error())
	}

	if err := logger.Sync(); err != nil {
		log.Printf("Failed to sync logger: %v", err)
	}

	// Log a sample message
	logger.Logger.Info("Application starting",
		zap.String("version", "1.0.0"),
		zap.String("environment", "development"),
	)

	// Setup HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handlers.SetupRouter(),
	}

	// Graceful shutdown
	go func() {
		logger.Logger.Info("Starting HTTP server", zap.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Logger.Info("Server exiting")
}
