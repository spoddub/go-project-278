package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         2 * time.Second,
	}))

	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/debug/sentry", func(c *gin.Context) {
		err := errors.New("test error from /debug/sentry")

		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		} else {
			sentry.CaptureException(err)
		}

		c.String(http.StatusInternalServerError, "sent to sentry")
	})

	return router
}

func initSentry() {
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		log.Println("SENTRY_DSN is empty, sentry disabled")
		return
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	}); err != nil {
		log.Printf("sentry init failed: %v", err)
	}
}

func main() {
	initSentry()
	defer sentry.Flush(2 * time.Second)

	router := setupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	_ = router.Run(":" + port)
}
