package middleware

import (
	"net/http"
)

func EnableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")    // Allow specific methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Allow headers
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
