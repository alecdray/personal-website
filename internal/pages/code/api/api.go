package api

import (
	"net/http"

	"github.com/CaribouBlue/personal-website/internal/pages/code/templates"
	"github.com/CaribouBlue/personal-website/internal/server/common"
)

type CodeApi struct {
}

var _ common.Api = &CodeApi{}

func NewCodeApi() CodeApi {
	return CodeApi{}
}

func (api *CodeApi) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/code", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates.CodePageComponent(templates.CodePageProps{}).Render(r.Context(), w)
	}))
}
