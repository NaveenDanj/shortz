package main

import (
	"config"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"auth/internal/app"
	"auth/internal/platform/betterauth"
	"auth/internal/platform/db"

	"gorm.io/gorm/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := db.Open(cfg.AuthConfig.DatabaseURL, logger.Warn)
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	auth, err := betterauth.New(cfg.AuthConfig, sqlDB)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	stopCleanup := auth.StartCleanup(ctx, time.Hour)
	defer stopCleanup()

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           app.NewRouter(auth),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(shutdownCtx)
	}()

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
