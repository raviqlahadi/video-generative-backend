package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/trending", TrendingHanler).Methods("GET")
	router.HandleFunc("/category", CategoryHandler).Methods("GET")
	router.HandleFunc("/search", SearchHandler).Methods("GET")
	router.HandleFunc("/proxy-video", ProxyVideoHandler).Methods("GET")

	return router

}
