package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheckHandler returns a basic status message to check if the server is running
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Server is up and running!")
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/trending", TrendingHanler).Methods("GET")
	router.HandleFunc("/category", CategoryHandler).Methods("GET")
	router.HandleFunc("/search", SearchHandler).Methods("GET")
	router.HandleFunc("/proxy-video", ProxyVideoHandler).Methods("GET")

	return router

}
