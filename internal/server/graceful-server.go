package server

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/mrdyuke/shorty/config"
)

func RunServer(cfg *config.Config, router http.Handler) error {
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,

		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen", "error", err)
		}
	}()

	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-sigCtx.Done()
	slog.Info("signal received, shutting down")

	shutdownCtx, stop := context.WithTimeout(context.Background(), cfg.ServerShutdownTimeout)
	defer stop()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}
	slog.Info("server gracefully stopped")

	return nil
}
