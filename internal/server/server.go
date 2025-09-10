package server

import (
	"net/http"

	"github.com/CaribouBlue/personal-website/internal/pages/home"
	"github.com/CaribouBlue/personal-website/internal/pages/landing"
)

func NewServer() *http.Server {
	rootMux := http.NewServeMux()

	rootMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/public"))))

	landingPage := landing.NewLandingPage()
	landingPage.Api.RegisterRoutes(rootMux)

	homePage := home.NewHomePage()
	homePage.Api.RegisterRoutes(rootMux)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: rootMux,
	}

	return server
}
