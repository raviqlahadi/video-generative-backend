package api

import (
	"net/http"
	"net/url"

	"github.com/raviqlahadi/video-generative-backend/internal/services"
	"github.com/raviqlahadi/video-generative-backend/pkg/utils"
)

func TrendingHanler(w http.ResponseWriter, r *http.Request) {
	videos, err := services.GetTrendingVideos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, videos)
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	hashtag := r.URL.Query().Get("hashtag")
	// Decode the URL-encoded 'hashtag'
	decodedHashtag, err := url.QueryUnescape(hashtag)
	if err != nil {
		// If URL decoding fails, return a Bad Request error
		http.Error(w, "Invalid hashtag encoding", http.StatusBadRequest)
		return
	}

	// URL-encode the decoded hashtag to make sure it's correctly encoded when sent to the external API
	encodedHashtag := url.QueryEscape(decodedHashtag)

	videos, err := services.GetCategoryVideos(encodedHashtag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, videos)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	videos, err := services.GetSearchVideos(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, videos)
}
