package main

import (
	"log"
	"net/http"

	"github.com/raviqlahadi/video-generative-backend/config"
	"github.com/raviqlahadi/video-generative-backend/pkg/api"
	"github.com/raviqlahadi/video-generative-backend/pkg/middleware"
)

func main() {
	config.LoadConfig()

	router := api.NewRouter()

	corsRouter := middleware.EnableCors(router)

	log.Printf("Server running on 0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", corsRouter); err != nil {
		log.Fatalf("Server failed: %s", err)
	}

}
