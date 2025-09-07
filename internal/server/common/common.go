package common

import "net/http"

type Api interface {
	RegisterRoutes(mux *http.ServeMux)
}
