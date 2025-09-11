package api

import (
	"net/http"

	"github.com/CaribouBlue/personal-website/internal/pages/blog/templates"
	"github.com/CaribouBlue/personal-website/internal/server/common"
)

type BlogApi struct {
}

var _ common.Api = &BlogApi{}

func NewBlogApi() BlogApi {
	return BlogApi{}
}

func (api *BlogApi) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/blog", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates.BlogPageComponent(templates.BlogPageProps{}).Render(r.Context(), w)
	}))
}
