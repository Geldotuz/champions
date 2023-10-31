package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var users = map[string]string{"user1": "password1", "user2": "password2"}

func requestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		for key, value := range headers {
			fmt.Println(key, value)
		}

		requestID := r.Header.Get("X-Request-ID")

		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}

		w.Header().Set("X-Resquest-ID", requestID)

		next.ServeHTTP(w, r)
	})
}

func authMiddlewate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("Autorizado")

		if users[user] == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func yearMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		ctx = context.WithValue(ctx, "year", "1992")
		next.ServeHTTP(w, request.WithContext(ctx))
	})

}
