package services

import (
	"github.com/raviqlahadi/video-generative-backend/pkg/cache"
	"github.com/raviqlahadi/video-generative-backend/pkg/pythonclient"
)

func GetTrendingVideos() ([]pythonclient.VideoResponse, error) {
	if cached, ok := cache.Get("trending"); ok {
		return cached.([]pythonclient.VideoResponse), nil
	}
	videos, err := pythonclient.FetchTrendingVideos()
	if err == nil {
		cache.Set("trending", videos)
	}
	return videos, err
}

func GetCategoryVideos(hashtag string) ([]pythonclient.VideoResponse, error) {
	cacheKey := "category:" + hashtag
	if cached, ok := cache.Get(cacheKey); ok {
		return cached.([]pythonclient.VideoResponse), nil
	}

	videos, err := pythonclient.FetchCategoryVideos(hashtag)
	if err == nil {
		cache.Set(cacheKey, videos)
	}
	return videos, err
}

func GetSearchVideos(query string) ([]pythonclient.VideoResponse, error) {
	cacheKey := "search:" + query
	if cached, ok := cache.Get(cacheKey); ok {
		return cached.([]pythonclient.VideoResponse), nil
	}
	videos, err := pythonclient.FetchSearchVideos(query)
	if err == nil {
		cache.Set(cacheKey, videos)
	}

	return videos, err
}
