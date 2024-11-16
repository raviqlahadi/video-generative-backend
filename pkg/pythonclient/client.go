package pythonclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/raviqlahadi/video-generative-backend/config"
)

type VideoResponse struct {
	URL            string   `json:"playable_video_url"`
	Author         string   `json:"author"`           // author.uniqueId
	CoverImage     string   `json:"cover_image"`      // video.cover
	DateCreated    string   `json:"date_created"`     // createTime
	Description    string   `json:"description"`      // contents.desc
	Tags           []string `json:"tags"`             // textExtra.hashtagName[]
	OnVideoCaption []string `json:"on_video_caption"` // stickersOnItem.stickerText
}

func fetchFromPython(endpoint string) ([]VideoResponse, error) {
	url := fmt.Sprintf("%s/%s", config.PythonServerUrl, endpoint)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	log.Printf("Requesting URL: %s", url)

	defer resp.Body.Close()

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var videoList []VideoResponse
	for _, item := range result {
		var videoResponse VideoResponse

		// Extract video URL (for 720p resolution)
		if video, ok := item["video"].(map[string]interface{}); ok {
			// Extract cover image
			if videoCover, ok := video["cover"].(string); ok {
				videoResponse.CoverImage = videoCover
			}
			if bitrateInfo, ok := video["bitrateInfo"].([]interface{}); ok {
				for _, bitrateItem := range bitrateInfo {
					if bitrate, ok := bitrateItem.(map[string]interface{}); ok {
						if playAddr, ok := bitrate["PlayAddr"].(map[string]interface{}); ok {
							maxWidht := 0
							if width, ok := playAddr["Width"].(float64); ok && width > float64(maxWidht) {
								if urlList, ok := playAddr["UrlList"].([]interface{}); ok {
									for _, urlItem := range urlList {
										if urlStr, ok := urlItem.(string); ok {
											if strings.HasPrefix(urlStr, "https://www.tiktok.com/aweme/v1/play/") {
												videoResponse.URL = urlStr
												break
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		// Extract author (uniqueId only)
		if author, ok := item["author"].(map[string]interface{}); ok {
			if uniqueId, ok := author["uniqueId"].(string); ok {
				videoResponse.Author = uniqueId
			}
		}

		// Extract creation time
		if createTime, ok := item["createTime"].(float64); ok {
			// Convert Unix timestamp to time.Time
			timestamp := time.Unix(int64(createTime), 0)
			// Format it to "yyyy-mm-dd"
			videoResponse.DateCreated = timestamp.Format("2006-01-02")
		}

		// Extract description
		if desc, ok := item["desc"].(string); ok {
			videoResponse.Description = desc
		}

		// Extract tags
		var tags []string
		if textExtra, ok := item["textExtra"].([]interface{}); ok {
			for _, tag := range textExtra {
				if tagMap, ok := tag.(map[string]interface{}); ok {
					if hashtag, ok := tagMap["hashtagName"].(string); ok && hashtag != "" {
						tags = append(tags, hashtag)
					}
				}
			}
		}
		videoResponse.Tags = tags

		// Extract on-video captions
		var captions []string
		if stickersOnItem, ok := item["stickersOnItem"].([]interface{}); ok {
			for _, sticker := range stickersOnItem {
				if stickerMap, ok := sticker.(map[string]interface{}); ok {
					if text, ok := stickerMap["stickerText"].(string); ok && text != "" {
						captions = append(captions, text)
					}
				}
			}
		}
		videoResponse.OnVideoCaption = captions

		// Add video response to the list
		videoList = append(videoList, videoResponse)
	}

	return videoList, nil
}

func FetchTrendingVideos() ([]VideoResponse, error) {
	return fetchFromPython("trending")
}

func FetchCategoryVideos(hashtag string) ([]VideoResponse, error) {
	return fetchFromPython("category?hashtag=" + hashtag)
}

func FetchSearchVideos(query string) ([]VideoResponse, error) {
	return fetchFromPython("search?query=" + query)
}
