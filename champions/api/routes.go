package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.Use(requestIDHandler)
	r.Use(authMiddlewate)
	r.Use(yearMiddleware)
	r.HandleFunc("/books", a.getBooks).Methods(http.MethodGet)
}
