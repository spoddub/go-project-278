package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := setupRouter()
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, request)

	if w.Code != http.StatusOK {
		t.Fatalf("want %d; got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "pong" {
		t.Fatalf("want %s; got %s", "pong", w.Body.String())
	}
}
