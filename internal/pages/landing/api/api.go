package api

import (
	"net/http"

	"github.com/CaribouBlue/personal-website/internal/pages/landing/templates"
	"github.com/CaribouBlue/personal-website/internal/server/common"
)

type LandingApi struct{}

var _ common.Api = LandingApi{}

func NewLandingApi() LandingApi {
	return LandingApi{}
}

func (api LandingApi) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/landing", http.StatusFound)
	}))

	mux.Handle("/landing", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates.LandingPageComponent(templates.LandingPageProps{}).Render(r.Context(), w)
	}))
}
