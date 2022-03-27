package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf adds CSRF protection to all Post request
func Nosruf(next http.Handler) http.Handler {
	csfHandler := nosurf.New(next)

	csfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csfHandler
}

// SessionLoad loads and saves the session on every request

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
