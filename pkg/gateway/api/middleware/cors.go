package middleware

import (
	"net/http"
)

func AddCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, X-Requested-With, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0") // TODO CHECK
		w.Header().Set("Vary", "Accept-Encoding") // TODO CHECK
		next.ServeHTTP(w, r)
	})
}

