package stream

import (
	"fmt"
	"net/http"
)

func StartSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	fmt.Fprintf(w, "data: %s\n\n", "Starting SSE...")
}
