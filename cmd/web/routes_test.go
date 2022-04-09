package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/sw90lee/go-web/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app *config.AppConfig

	mux := routes(app)
	switch v := mux.(type) {
	case *chi.Mux:
	// 아무것도 없음
	default:
		t.Error(fmt.Sprintf("Type is not *chi.Mux, type is %t", v))
	}
}
