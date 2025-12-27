package main

import (
	"context"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/jackc/pgx/v5/pgxpool"

	"shorty/internal/config"
	db "shorty/internal/db/sqlc"
	httpapi "shorty/internal/http"
)

func initSentry(dsn string) {
	if dsn == "" {
		log.Println("SENTRY_DSN is empty, sentry disabled")
		return
	}

	if err := sentry.Init(sentry.ClientOptions{Dsn: dsn}); err != nil {
		log.Printf("sentry init failed: %v", err)
	}
}

func main() {
	cfg := config.Load()

	initSentry(cfg.SentryDSN)
	defer sentry.Flush(2 * time.Second)

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	defer pool.Close()

	q := db.New(pool)
	router := httpapi.NewRouter(q, cfg.BaseURL)

	_ = router.Run(":" + cfg.AppPort)
}
