package api

import (
	"net/http"

	"github.com/CaribouBlue/personal-website/internal/pages/home/templates"
	"github.com/CaribouBlue/personal-website/internal/server/common"
)

type HomeApi struct{}

var _ common.Api = HomeApi{}

func NewHomeApi() HomeApi {
	return HomeApi{}
}

func (api HomeApi) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/home", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates.HomePageComponent(templates.HomePageProps{}).Render(r.Context(), w)
	}))
}
