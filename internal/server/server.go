package server

import (
	"net/http"

	landing "github.com/CaribouBlue/personal-website/internal/pages/landing/api"
)

func NewServer() *http.Server {
	rootMux := http.NewServeMux()

	rootMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/public"))))

	landingApi := landing.NewLandingApi()
	landingApi.RegisterRoutes(rootMux)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: rootMux,
	}

	return server
}
