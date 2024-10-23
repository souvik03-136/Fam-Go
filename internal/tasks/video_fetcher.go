package tasks

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/souvik03-136/Fam-Go/internal/database"
)

type VideoFetcher struct {
	DB         *database.Queries
	APIKeys    []string
	CurrentKey int
	Interval   time.Duration
	LastFetch  time.Time
}

func NewVideoFetcher(db *database.Queries, interval time.Duration) *VideoFetcher {
	apiKeys := strings.Split(os.Getenv("YOUTUBE_API_KEYS"), ",")
	return &VideoFetcher{
		DB:         db,
		APIKeys:    apiKeys,
		CurrentKey: 0,
		Interval:   interval,
		LastFetch:  time.Now().Add(-10 * time.Second),
	}
}

func (vf *VideoFetcher) Start(ctx context.Context) {
	ticker := time.NewTicker(vf.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			vf.FetchLatestVideos(ctx)
		}
	}
}

func (vf *VideoFetcher) FetchLatestVideos(ctx context.Context) {
	client := resty.New()
	publishedAfter := vf.LastFetch.Format(time.RFC3339)

	response, err := client.R().
		SetQueryParam("part", "snippet").
		SetQueryParam("order", "date").
		SetQueryParam("type", "video").
		SetQueryParam("publishedAfter", publishedAfter).
		SetQueryParam("key", vf.APIKeys[vf.CurrentKey]).
		Get("https://www.googleapis.com/youtube/v3/search")

	if err != nil || response.StatusCode() != http.StatusOK {
		vf.rotateAPIKey()
		log.Printf("Failed to fetch videos: %v. Switching to next API key.", err)
		return
	}

	var result struct {
		Items []struct {
			Id struct {
				VideoId string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title       string    `json:"title"`
				Description string    `json:"description"`
				PublishedAt time.Time `json:"publishedAt"`
				Thumbnails  struct {
					Default struct {
						Url string `json:"url"`
					} `json:"default"`
				} `json:"thumbnails"`
			} `json:"snippet"`
		} `json:"items"`
	}

	if err := json.Unmarshal(response.Body(), &result); err != nil {
		log.Printf("Failed to unmarshal response: %v", err)
		return
	}

	if len(result.Items) > 0 {
		vf.LastFetch = result.Items[0].Snippet.PublishedAt
	}

	for _, item := range result.Items {
		videoParams := database.CreateVideoParams{
			Title:        item.Snippet.Title,
			Description:  sql.NullString{String: item.Snippet.Description, Valid: true},
			PublishedAt:  item.Snippet.PublishedAt,
			ThumbnailUrl: sql.NullString{String: item.Snippet.Thumbnails.Default.Url, Valid: true},
			VideoUrl:     "https://www.youtube.com/watch?v=" + item.Id.VideoId,
		}
		if err := vf.DB.CreateVideo(ctx, videoParams); err != nil {
			log.Printf("Failed to store video: %v", err)
		}
	}
}

func (vf *VideoFetcher) rotateAPIKey() {
	vf.CurrentKey = (vf.CurrentKey + 1) % len(vf.APIKeys)
}
