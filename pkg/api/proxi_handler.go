package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ProxyVideoHandler(w http.ResponseWriter, r *http.Request) {
	// Get the video URL from the query parameters
	videoURL := r.URL.Query().Get("url")
	if videoURL == "" {
		http.Error(w, "Missing video URL", http.StatusBadRequest)
		return
	}

	// Decode the URL-encoded 'video_url' parameter
	decodedURL, err := url.QueryUnescape(videoURL)
	if err != nil {
		http.Error(w, "Failed to decode video_url", http.StatusBadRequest)
		return
	}

	// Create a new request to fetch the video
	req, err := http.NewRequest("GET", decodedURL, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Add headers to the request to mimic a browser request
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	req.Header.Set("Accept", "video/mp4")
	req.Header.Set("Referer", "https://www.tiktok.com/")
	req.Header.Set("Origin", "https://www.tiktok.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	// Send the request to TikTok
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error fetching video", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// If TikTok returns an error, log the response and send back an error
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error fetching video: %s", resp.Status), resp.StatusCode)
		return
	}

	// Set the content type to the same as TikTok's response (video content)
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	// Set a cache-control header if desired
	w.Header().Set("Cache-Control", "no-cache")

	// Stream the video content to the frontend
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Error streaming video", http.StatusInternalServerError)
	}
}
