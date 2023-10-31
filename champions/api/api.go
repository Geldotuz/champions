package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type API struct {
}

var books = []string{"book1", "book2", "book3"}

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {

	limitParam := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if limit < 0 || limit > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(r.Context().Value("year"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books[:limit])
}
